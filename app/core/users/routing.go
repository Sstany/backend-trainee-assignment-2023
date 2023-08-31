package users

import (
	"segmenty/app/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Users struct {
	db     *db.Database
	logger *zap.Logger
}

func AttachToGroup(group *gin.RouterGroup) {
	usersEnv := newUsersEnv()

	group.GET("/:id", usersEnv.fetch)
	group.POST("/", usersEnv.create)
	group.DELETE("/:id", usersEnv.delete)

	group.GET("/:id/segments", usersEnv.listSegments)
	group.POST("/:id/segments", usersEnv.updateSegments)
}

func newUsersEnv() *Users {
	u := Users{}
	u.db = db.NewDB()

	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	u.logger = l

	return &u
}
