package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"senao/pkg/apps"
	"senao/pkg/apps/docs"
	"senao/pkg/core/database"
	"senao/pkg/repository"
	"senao/pkg/usecase"
)

var (
	Cmd = &cobra.Command{
		Use:   "server",
		Short: "start senao server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
	VERSION string
)

func ServerEnv() {
	os.Setenv("API_HOST", "localhost:8080")
	os.Setenv("API_PORT", "8080")
}

func start() {
	log.Info().Msgf("Start RestAPI...")

	//Init logger
	zerolog.SetGlobalLevel(zerolog.Level(0))

	//Database Connection
	db, err := database.GetDB(database.GetDSN())
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer db.Close()

	accountRepo := repository.NewAccountRepo()
	accountUseCase := usecase.NewAccountUsecase(db, accountRepo)

	commonHandler := apps.NewCommonHandler(VERSION)
	accountHandler := apps.NewAccountHandler(accountUseCase)

	// Gin Server
	ServerEnv()
	gin.SetMode(gin.ReleaseMode)
	basePath := "/v1"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = os.Getenv("API_HOST")
	PORT, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", commonHandler.Healthz)
	apiGroup := r.Group(basePath)
	{
		apiGroup.GET("/ping", commonHandler.Ping)
		apiGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		apiGroup.POST("/accounts", accountHandler.CreateAccount)
		apiGroup.POST("/accounts/verify", accountHandler.VerifyAccount)
	}
	r.Run(fmt.Sprintf(":%d", PORT))
}
