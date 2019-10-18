package main

import (
	"ccsl/configs"
	"ccsl/controllers"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := initApp()
	pg := initDB(app)
	defer pg.Close()
	mvc.New(app).Handle(new(controllers.RootController))
	mvc.Configure(app.Party("/files"), func(app *mvc.Application) {
		app.Handle(new(controllers.FileController))
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
	// sign language basic elements
	mvc.Configure(app.Party("/signs"), func(app *mvc.Application) {
		app.Register(services.NewSignService(pg))
		app.Handle(new(controllers.SignController))
	})
	mvc.Configure(app.Party("/performers"), func(app *mvc.Application) {
		app.Register(services.NewPerformerService(pg))
		app.Handle(new(controllers.PerformerController))
	})
	// Lexical Database for Chinese National Sign Language
	mvc.Configure(app.Party("/lexical/words"), func(app *mvc.Application) {
		app.Register(services.NewLexicalWordService(pg))
		app.Handle(new(controllers.WordController))
	})
	mvc.Configure(app.Party("/lexical/videos"), func(app *mvc.Application) {
		app.Register(services.NewLexicalVideoService(pg))
		app.Handle(new(controllers.VideoController))
	})
	// Corpus for Shanghai Sign Language Verb
	mvc.Configure(app.Party("/verbs"), func(app *mvc.Application) {
		// TODO
	})
	// Corpus for Proper Nouns in CSL
	mvc.Configure(app.Party("/nouns"), func(app *mvc.Application) {
		// TODO
	})
	// Chinese Sign Language Corpus for Sign Texts
	mvc.Configure(app.Party("/texts"), func(app *mvc.Application) {
		// TODO
	})
	// Literature Database for Sign Language Research
	mvc.Configure(app.Party("/literature"), func(app *mvc.Application) {
		// TODO
	})
	// Database for Technical Terms in Sign Linguistics
	mvc.Configure(app.Party("/terms"), func(app *mvc.Application) {
		// TODO
	})
	go gracefulShutdown(app)
	host := configs.Conf.Listener.Server + ":" + strconv.Itoa(configs.Conf.Listener.Port)
	app.Run(iris.Addr(host), iris.WithOptimizations, iris.WithoutStartupLog, iris.WithoutInterruptHandler)
}

func initApp() *iris.Application {
	app := iris.New()
	app.OnAnyErrorCode(middlewares.ErrorHandler)
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middlewares.NewI18nMidddleware(middlewares.I18nConf))
	app.Use(middlewares.BeforeHandleRequest)
	app.Done(middlewares.AfterHandleRequest)
	app.Configure(iris.WithConfiguration(iris.YAML(configs.AppConf)))
	app.Logger().SetLevel(configs.Conf.Log.Level).SetOutput(configs.GetLogger("api"))
	return app
}

func initDB(app *iris.Application) *gorm.DB {
	pg := utils.ConnectPostgres(app)
	pg.SetLogger(configs.GetPostgresLogger())
	pg.LogMode(true)
	// AutoMigrate will create missing tables and missing index keys
	pg.AutoMigrate(&models.User{}, &models.LexicalWord{}, &models.LexicalVideo{}, &models.Sign{}, &models.Performer{}, &models.Carousel{}, &models.News{}, &models.Member{}, &models.District{}, &models.City{}, &models.Province{})
	// Don't use UNIQUE to declare gorm models because you can't create a alreay deleted object with the same value
	// Manually Add UNIQUE key for table columns
	pg.Exec("CREATE UNIQUE INDEX users_username_key ON users(username) WHERE deleted_at IS NULL")
	pg.Exec("CREATE UNIQUE INDEX signs_name_key ON signs(name) WHERE deleted_at IS NULL")
	pg.Exec("CREATE UNIQUE INDEX lexical_words_chinese_key ON lexical_words(chinese) WHERE deleted_at IS NULL")
	// Manually Add foreign key for tables, because gorm won't create foreign keys, to make sure data is clean we need manually add these keys
	// Data of cities are from https://github.com/modood/Administrative-divisions-of-China
	// You cannot neither change nor update them, so set to RESTRICT
	pg.Model(&models.District{}).AddForeignKey("city_code", "cities(code)", "RESTRICT", "RESTRICT").AddForeignKey("province_code", "provinces(code)", "RESTRICT", "RESTRICT")
	pg.Model(&models.City{}).AddForeignKey("province_code", "provinces(code)", "RESTRICT", "RESTRICT")
	// For other data models, set delete mode to RESTRICT and update mode to CASCADE
	pg.Model(&models.LexicalVideo{}).AddForeignKey("lexical_word_id", "lexical_words(id)", "RESTRICT", "CASCADE")
	pg.Model(&models.LexicalVideo{}).AddForeignKey("performer_id", "performers(id)", "RESTRICT", "CASCADE")
	pg.Model(&models.News{}).AddForeignKey("creator_id", "users(id)", "RESTRICT", "CASCADE")
	pg.Model(&models.Carousel{}).AddForeignKey("creator_id", "users(id)", "RESTRICT", "CASCADE")
	pg.Model(&models.Performer{}).AddForeignKey("region", "districts(code)", "RESTRICT", "CASCADE")
	// These tables are many to many connections table, also need to add foreign keys manually
	pg.Table("lexical_left_sign").AddForeignKey("lexical_video_id", "lexical_videos(id)", "RESTRICT", "CASCADE").AddForeignKey("sign_id", "signs(id)", "RESTRICT", "CASCADE")
	pg.Table("lexical_right_sign").AddForeignKey("lexical_video_id", "lexical_videos(id)", "RESTRICT", "CASCADE").AddForeignKey("sign_id", "signs(id)", "RESTRICT", "CASCADE")
	// utils.InitTestUser(pg)
	return pg
}

func gracefulShutdown(app *iris.Application) {
	ch := make(chan os.Signal, 1)
	// Catch exit sign and shutdown server gracefully
	// kill -SIGINT XXXX Or Ctrl+c
	signal.Notify(ch,
		os.Interrupt,
		syscall.SIGINT,
		os.Kill,
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
