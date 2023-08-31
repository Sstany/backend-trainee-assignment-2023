package users

import (
	"net/http"

	"segmenty/app/db/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Users) fetch(ctx *gin.Context) {
	var user *models.User

	err := ctx.ShouldBindUri(&user)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}
	newUser, isNoRows, err := r.db.FetchUser(ctx, user.ID)
	if err != nil {
		r.logger.Error("", zap.Error(err))

		if isNoRows {
			ctx.String(http.StatusNotFound, "User with ID '%d' is not found", user.ID)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, newUser)

}
func (r *Users) create(ctx *gin.Context) {

}
func (r *Users) delete(ctx *gin.Context) {

}
func (r *Users) updateSegments(ctx *gin.Context) {

}
func (r *Users) listSegments(ctx *gin.Context) {

}
