package segments

import (
	"segmenty/app/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Segments struct {
	db     *db.Database
	logger *zap.Logger
}

func AttacheToGroup(group *gin.RouterGroup) {
	segmentsEnv := newSegmentsEnv()

	group.GET("/", segmentsEnv.list)
	group.GET("/:slug", segmentsEnv.fetch)
	group.POST("/:slug", segmentsEnv.create)
	group.DELETE("/:slug", segmentsEnv.delete)

}

func newSegmentsEnv() *Segments {
	s := Segments{}
	s.db = db.NewDB()

	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	s.logger = l

	return &s
}
