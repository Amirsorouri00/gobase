package auth

import (
	"fmt"
	"portfolio/app"
	api "portfolio/app/auth/controller"
	"portfolio/services/infrastructure/config"
	conf "portfolio/services/infrastructure/config/auth"
	"portfolio/services/infrastructure/log"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.Init(config.Env("portfolio"))
}

// @title Portfolio - Auth
// @version 1.0
// @description This document describes API provided by Portfolio backend.
// @termsOfService http://swagger.io/terms/
// @contact.name Portfolio Developers
// @contact.email support@portfolio.dev
// @BasePath /api/stocks
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Run() {
	cfg := conf.LoadConfig()
	api.Init(cfg.Services.Policy.Host, cfg.Services.Policy.Port)
	// purchase.Init(cfg.Services.Auth.Host, cfg.Services.Auth.Port)
	// i18nLocale, err := locale.LoadLocale()
	// if nil != err {
	// 	log.Fatalf("error loading locale .po files: %s", err.Error())
	// }

	// logFile, err := os.OpenFile(
	// 	cfg.Core.LogRoute, os.O_CREATE|os.O_APPEND|os.O_RDWR,
	// 	os.ModeAppend,
	// )
	// if err != nil {
	// 	log.Fatalf("failed to open log file %s", log.F("err", err))
	// }
	// defer logFile.Close()
	// log.Init(logFile, config.Bool("debug"))

	// authRpcClient, err := rpc.Connect(cfg.Services.Auth.Host, cfg.Services.Auth.Port)
	// if nil != err {
	// 	log.Fatalf("failed to initialize auth service rpc connection: %v", err)
	// }

	db, err := app.InitDB(cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// minioClient, err := storage.NewMinioClient(
	// 	cfg.Minio.Host,
	// 	cfg.Minio.AccessKey,
	// 	cfg.Minio.SecretKey,
	// 	cfg.Minio.Token,
	// 	cfg.Minio.Secure,
	// )
	// fileManagerService := fileManager.NewFileManager(cfg.FileManagerServeUrl)
	// if err != nil {
	// 	log.Fatalf("failed initializing minio client: %v", err)
	// }
	// image.Init(minioClient, cfg.Minio.Bucket)
	// storageClient := storage.NewStorage(minioClient)
	// address.Init()
	// worker.NewMessageGeneratorWorker(db, cfg.IsActiveMessageGeneratorWorker, cfg.PeriodCronMessageGenerator).Start(context.Background())
	r := gin.New()

	r.Use(gin.Recovery())
	// r.Use(locale.BuildLocalesInjector(i18nLocale))
	r.Use(sentrygin.New(sentrygin.Options{Repanic: true, WaitForDelivery: false}))

	api.Populate(r, db)

	addr := fmt.Sprintf(":%s", cfg.Http.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.InstanceName("auth"),
	))

	if err = r.Run(addr); err != nil {
		log.Fatalf("http server failed:%s", log.F("err", err))
	}
}
