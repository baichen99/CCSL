package main

import (
	"ccsl/configs"
	"ccsl/controllers"
	_ "ccsl/docs"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

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
		app.Register(services.NewUserService(pg))
		app.Handle(new(controllers.UserController))
	})
	mvc.Configure(app.Party("/signs"), func(app *mvc.Application) {
		app.Register(services.NewSignService(pg))
		app.Handle(new(controllers.SignController))
	})
	mvc.Configure(app.Party("/performers"), func(app *mvc.Application) {
		app.Register(services.NewPerformerService(pg))
		app.Handle(new(controllers.PerformerController))
	})
	mvc.Configure(app.Party("/lexicons"), func(app *mvc.Application) {
		app.Register(services.NewLexiconService(pg))
		app.Handle(new(controllers.LexiconController))
	})
	// Lexical Database for Chinese National Sign Language
	mvc.Configure(app.Party("/lexical/videos"), func(app *mvc.Application) {
		app.Register(services.NewLexicalVideoService(pg))
		app.Handle(new(controllers.VideoController))
	})
	mvc.Configure(app.Party("/notifications"), func(app *mvc.Application) {
		app.Register(services.NewNotificationService(pg))
		app.Handle(new(controllers.NotificationController))
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
	go shutdown(app)
	host := configs.Conf.Listener.Server + ":" + strconv.Itoa(configs.Conf.Listener.Port)
	app.Run(iris.Addr(host), iris.WithOptimizations, iris.WithoutStartupLog, iris.WithoutInterruptHandler)
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
		docHost := fmt.Sprintf("%s/swagger/doc.json", host)
		config := &swagger.Config{
			URL: docHost,
		}
		app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
		app.Logger().Debugf("Doc server is running on: %s/swagger/index.html", host)
	}
}

func initDB(app *iris.Application) *gorm.DB {
	pg := utils.ConnectPostgres(app)
	pg.SetLogger(configs.GetPostgresLogger())
	// AutoMigrate will create missing tables and missing index keys
	pg.AutoMigrate(
		&models.User{},
		&models.Lexicon{},
		&models.LexicalVideo{},
		&models.Sign{},
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
	)

	// Don't use UNIQUE to declare gorm models because you can't create a alreay deleted object with the same value, manually Add UNIQUE key for table columns

	pg.Exec(
		"CREATE UNIQUE INDEX users_username_key ON users(username) WHERE deleted_at IS NULL",
	)
	pg.Exec(
		"CREATE UNIQUE INDEX signs_name_key ON signs(name) WHERE deleted_at IS NULL",
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
		AddForeignKey("province_code", "provinces(code)", "RESTRICT", "RESTRICT").
		// For other data models, set delete mode to RESTRICT and update mode to CASCADE
		Model(&models.LexicalVideo{}).
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
		AddForeignKey("user_id", "users(id)", "RESTRICT", "CASCADE")
	// utils.InitTestUser(pg)
	return pg
}

func shutdown(app *iris.Application) {
	ch := make(chan os.Signal, 1)
	// Catch exit sign and shutdown server gracefully
	// kill -SIGINT XXXX Or Ctrl+c
	signal.Notify(ch,
		os.Kill,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	)
	select {
	case <-ch:
		app.Logger().Info("SHUTDOWN: receive shutdown sign")
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	}
}
