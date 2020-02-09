package main

import (
	"ccsl/configs"
	"ccsl/controllers"
	_ "ccsl/docs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := initApp()
	pg := initDB(app)
	defer pg.Close()
	initDoc(app)
	// App Handler
	mvc.New(app).Handle(new(controllers.RootController))
	mvc.Configure(app.Party("/files"), func(app *mvc.Application) {
		app.Handle(new(controllers.FileController))
	})
	mvc.Configure(app.Party("/systems"), func(app *mvc.Application) {
		app.Register(services.NewSystemService(pg), services.NewUserService(pg))
		app.Handle(new(controllers.SystemController))
	})
	mvc.Configure(app.Party("/carousels"), func(app *mvc.Application) {
		app.Register(services.NewCarouselService(pg))
		app.Handle(new(controllers.CarouselController))
	})
	mvc.Configure(app.Party("/news"), func(app *mvc.Application) {
		app.Register(services.NewNewsService(pg))
		app.Handle(new(controllers.NewsController))
	})
	mvc.Configure(app.Party("/members"), func(app *mvc.Application) {
		app.Register(services.NewMemberService(pg))
		app.Handle(new(controllers.MemberController))
	})
	mvc.Configure(app.Party("/users"), func(app *mvc.Application) {
		app.Register(services.NewUserService(pg), services.NewNotificationService(pg))
		app.Handle(new(controllers.UserController))
	})
	mvc.Configure(app.Party("/handshapes"), func(app *mvc.Application) {
		app.Register(services.NewHandshapeService(pg))
		app.Handle(new(controllers.HandshapeController))
	})
	mvc.Configure(app.Party("/performers"), func(app *mvc.Application) {
		app.Register(services.NewPerformerService(pg))
		app.Handle(new(controllers.PerformerController))
	})
	mvc.Configure(app.Party("/lexicons"), func(app *mvc.Application) {
		app.Register(services.NewLexiconService(pg))
		app.Handle(new(controllers.LexiconController))
	})
	mvc.Configure(app.Party("/notifications"), func(app *mvc.Application) {
		app.Register(services.NewNotificationService(pg))
		app.Handle(new(controllers.NotificationController))
	})
	// Lexical Database for Chinese National Sign Language
	mvc.Configure(app.Party("/lexical/videos"), func(app *mvc.Application) {
		app.Register(services.NewLexicalVideoService(pg))
		app.Handle(new(controllers.LexicalVideoController))
	})
	// // Corpus for Shanghai Sign Language Verb
	// mvc.Configure(app.Party("/verbs"), func(app *mvc.Application) {
	// 	// TODO
	// })
	// // Corpus for Proper Nouns in CSL
	// mvc.Configure(app.Party("/nouns"), func(app *mvc.Application) {
	// 	// TODO
	// })
	// // Chinese Sign Language Corpus for Sign Texts
	mvc.Configure(app.Party("/texts/videos"), func(app *mvc.Application) {
		// TODO
	})
	// // Literature Database for Sign Language Research
	// mvc.Configure(app.Party("/literature"), func(app *mvc.Application) {
	// 	// TODO
	// })
	// // Database for Technical Terms in Sign Linguistics
	// mvc.Configure(app.Party("/terms"), func(app *mvc.Application) {
	// 	// TODO
	// })

	mvc.Configure(app.Party("/class"), func(app *mvc.Application) {
		app.Register(services.NewClassService(pg))
		app.Handle(new(controllers.ClassController))
	})
	mvc.Configure(app.Party("/course"), func(app *mvc.Application) {
		app.Register(services.NewCourseService(pg))
		app.Handle(new(controllers.CourseController))
	})
	mvc.Configure(app.Party("/assignment"), func(app *mvc.Application) {
		app.Register(services.NewAssignmentService(pg))
		app.Handle(new(controllers.AssignmentController))
	})
	mvc.Configure(app.Party("/submitted_assignment"), func(app *mvc.Application) {
		app.Register(services.NewSubmittedAssignmentService(pg))
		app.Register(services.NewAssignmentService(pg))
		app.Register(services.NewCourseService(pg))
		app.Register(services.NewClassService(pg))
		app.Handle(new(controllers.SubmittedAssignmentController))
	})
	host := fmt.Sprintf("%s:%d", configs.Conf.Listener.Server, configs.Conf.Listener.Port)
	app.Run(iris.Addr(host), iris.WithOptimizations, iris.WithoutStartupLog)
}

func initApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Status:             true,
		IP:                 true,
		Method:             true,
		Path:               true,
		Query:              true,
		MessageContextKeys: []string{"RequestID"},
		MessageHeaderKeys:  []string{"User-Agent"},
	}))
	app.I18n.Load("./locales/*/*.ini", "zh-CN", "en-US")
	app.OnAnyErrorCode(middlewares.ErrorHandler)
	app.AllowMethods(iris.MethodOptions)
	app.Use(cors.New(middlewares.CorsConf))
	app.Use(middlewares.BeforeHandleRequest)
	app.Done(middlewares.AfterHandleRequest)
	app.Configure(iris.WithConfiguration(iris.YAML(configs.AppConf)))
	app.Logger().SetLevel(configs.Conf.Log.Level).SetOutput(configs.GetLogger("api"))
	return app
}

func initDoc(app *iris.Application) {
	// >>>>> DOCS  <<<<<
	// =================
	// @Title CCSL API
	// @Version 1.0
	// @Host localhost:8888
	// =================
	if os.Getenv(configs.EnvName) == configs.EnvDevelopment {
		host := fmt.Sprintf("http://%s:%s", configs.Conf.Listener.Server, strconv.Itoa(configs.Conf.Listener.Port))
		docFile := fmt.Sprintf("%s/swagger/doc.json", host)
		config := &swagger.Config{
			URL: docFile,
		}
		app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
		app.Logger().Debugf("Doc server is running on: %s/swagger/index.html", host)
	}
}

func initDB(app *iris.Application) *gorm.DB {
	pg := utils.ConnectPostgres(app)
	pg.LogMode(true)
	pg.SetLogger(configs.GetPostgresLogger())
	// AutoMigrate will create missing tables and missing index keys
	pg.AutoMigrate(
		&models.User{},
		&models.Lexicon{},
		&models.LexicalVideo{},
		&models.Handshape{},
		&models.Performer{},
		&models.Carousel{},
		&models.News{},
		&models.Member{},
		&models.District{},
		&models.City{},
		&models.Province{},
		&models.JsError{},
		&models.Info{},
		&models.LoginHistory{},
		&models.Notification{},
		&models.Class{},
		&models.Course{},
		&models.Assignment{},
		&models.SubmittedAssignments{},
	)

	// Don't use UNIQUE to declare gorm models because you can't create a already deleted object with the same value, manually Add UNIQUE key for table columns

	pg.Exec(
		"CREATE UNIQUE INDEX users_username_key ON users(username) WHERE deleted_at IS NULL",
	)
	pg.Exec(
		"CREATE UNIQUE INDEX handshapes_name_key ON handshapes(name) WHERE deleted_at IS NULL",
	)
	pg.Exec(
		"CREATE UNIQUE INDEX handshapes_glyph_key ON handshapes(glyph) WHERE deleted_at IS NULL",
	)
	pg.Exec(
		"CREATE UNIQUE INDEX lexicons_chinese_key ON lexicons(chinese) WHERE deleted_at IS NULL",
	)

	// Manually Add foreign key for tables, because gorm won't create foreign keys, to make sure data is clean we need manually add these keys
	// Data of cities are from https://github.com/modood/Administrative-divisions-of-China
	// You cannot neither change nor update them, so set to RESTRICT
	pg.
		Model(&models.District{}).
		AddForeignKey("city_code", "cities(code)", "RESTRICT", "RESTRICT").
		AddForeignKey("province_code", "provinces(code)", "RESTRICT", "RESTRICT").
		Model(&models.City{}).
		AddForeignKey("province_code", "provinces(code)", "RESTRICT", "RESTRICT")

	// For other data models, set delete mode to RESTRICT and update mode to CASCADE
	pg.Model(&models.LexicalVideo{}).
		AddForeignKey("lexicon_id", "lexicons(id)", "RESTRICT", "CASCADE").
		AddForeignKey("performer_id", "performers(id)", "RESTRICT", "CASCADE").
		Model(&models.Performer{}).
		AddForeignKey("region", "districts(code)", "RESTRICT", "CASCADE").
		Model(&models.News{}).
		AddForeignKey("creator_id", "users(id)", "RESTRICT", "CASCADE").
		Model(&models.Carousel{}).
		AddForeignKey("creator_id", "users(id)", "RESTRICT", "CASCADE").
		Model(&models.LoginHistory{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE").
		Model(&models.Notification{}).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE").
		Model(&models.Course{}).
		AddForeignKey("class_id", "classes(id)", "RESTRICT", "CASCADE").
		Model(&models.Assignment{}).
		AddForeignKey("course_id", "courses(id)", "RESTRICT", "CASCADE").
		Model(&models.SubmittedAssignments{}).
		AddForeignKey("assignment_id", "assignments(id)", "RESTRICT", "CASCADE")

	// Create test users for development environment
	if os.Getenv(configs.EnvName) == configs.EnvDevelopment {
		utils.InitTestUser(pg)
	}

	return pg
}
