package auth

import (
	"context"
	"fmt"
	"os"
	"portfolio/app"
	"portfolio/services/infrastructure/config"
	"portfolio/services/infrastructure/log"

	"git.eways.dev/eways/service/image"
	"git.eways.dev/eways/service/storage"
	fileManager "git.eways.dev/eways/stocks/services/infrastructure/file_manager"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.Init(config.Env("stocks"))
}

// @title E-ways - Stocks
// @version 1.1
// @description This document describes API provided by E-ways backend.
// @termsOfService http://swagger.io/terms/
// @contact.name Eways Developers
// @contact.email support@eways.dev
// @BasePath /api/stocks
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func Run() {
	cfg := conf.LoadConfig()
	api.Init(cfg.Services.Policy.Host, cfg.Services.Policy.Port)
	purchase.Init(cfg.Services.Auth.Host, cfg.Services.Auth.Port)
	i18nLocale, err := locale.LoadLocale()
	if nil != err {
		log.Fatalf("error loading locale .po files: %s", err.Error())
	}

	logFile, err := os.OpenFile(
		cfg.Core.LogRoute, os.O_CREATE|os.O_APPEND|os.O_RDWR,
		os.ModeAppend,
	)
	if err != nil {
		log.Fatalf("failed to open log file %s", log.F("err", err))
	}
	defer logFile.Close()
	log.Init(logFile, config.Bool("debug"))

	authRpcClient, err := rpc.Connect(cfg.Services.Auth.Host, cfg.Services.Auth.Port)
	if nil != err {
		log.Fatalf("failed to initialize auth service rpc connection: %v", err)
	}

	db, err := app.InitDB(cfg.Db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	minioClient, err := storage.NewMinioClient(
		cfg.Minio.Host,
		cfg.Minio.AccessKey,
		cfg.Minio.SecretKey,
		cfg.Minio.Token,
		cfg.Minio.Secure,
	)
	fileManagerService := fileManager.NewFileManager(cfg.FileManagerServeUrl)
	if err != nil {
		log.Fatalf("failed initializing minio client: %v", err)
	}
	image.Init(minioClient, cfg.Minio.Bucket)
	storageClient := storage.NewStorage(minioClient)
	address.Init()
	worker.NewMessageGeneratorWorker(db, cfg.IsActiveMessageGeneratorWorker, cfg.PeriodCronMessageGenerator).Start(context.Background())
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(locale.BuildLocalesInjector(i18nLocale))
	r.Use(sentrygin.New(sentrygin.Options{Repanic: true, WaitForDelivery: false}))

	api.Populate(r, db, authRpcClient.RPCClient, storageClient, fileManagerService)

	addr := fmt.Sprintf(":%s", cfg.Http.Port)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.InstanceName("admin"),
	))

	if err = r.Run(addr); err != nil {
		log.Fatalf("http server failed:%s", log.F("err", err))
	}
}
