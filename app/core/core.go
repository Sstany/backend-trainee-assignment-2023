package core

import (
	"segmenty/app/core/users"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service struct {
	host   string
	port   string
	logger *zap.Logger
}

func NewService(host string, port string) *Service {
	c := Service{
		host: host,
		port: port,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	c.logger = logger
	return &c
}

func (r *Service) Start() {
	api := r.newApi()

	api.Run(r.host + ":" + r.port)
}

func (r *Service) newApi() *gin.Engine {

	router := gin.New()

	router.Use(CreateLoggerMiddleware(r.logger))

	api := router.Group("/api/v1")

	users.AttachToGroup(api.Group("/users"))

	return router

}

func CreateLoggerMiddleware(logger *zap.Logger) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		logger.Info(ctx.Request.Host, zap.String("Method", ctx.Request.Method), zap.String("URI", ctx.Request.RequestURI))
	}
}