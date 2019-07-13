package main

import (
	"ccsl/configs"
	"ccsl/controllers"
	"ccsl/middlewares"
	"ccsl/models"
	"ccsl/services"
	"ccsl/utils"
	"strconv"

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
	mvc.Configure(app.Party("/performers"), func(app *mvc.Application) {
		app.Register(services.NewPerformerServices(pg))
		app.Handle(new(controllers.PerformerController))
	})
	mvc.Configure(app.Party("/words"), func(app *mvc.Application) {
		app.Register(services.NewWordService(pg))
		app.Handle(new(controllers.WordController))
	})
	mvc.Configure(app.Party("/videos"), func(app *mvc.Application) {
		app.Register(services.NewVideoService(pg))
		app.Handle(new(controllers.VideoController))
	})
	mvc.Configure(app.Party("/users"), func(app *mvc.Application) {
		app.Register(services.NewUserService(pg))
		app.Handle(new(controllers.UserController))
	})
	host := configs.Conf.Listener.Server + ":" + strconv.Itoa(configs.Conf.Listener.Port)
	app.Run(iris.Addr(host), iris.WithOptimizations, iris.WithoutStartupLog)
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
	pg.AutoMigrate(&models.User{}, &models.Word{}, &models.Video{}, &models.Performer{}, &models.Carousel{}, &models.News{})
	pg.Model(&models.Video{}).AddForeignKey("word_id", "words(id)", "RESTRICT", "CASCADE")
	pg.Model(&models.Video{}).AddForeignKey("performer_id", "performers(id)", "RESTRICT", "CASCADE")
	utils.InitProdUser(pg)
	// utils.InitSuper(pg)
	// utils.InitTestUser(pg)
	return pg
}
