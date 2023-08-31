package users

import (
	"errors"
	"net/http"

	"segmenty/app/db/models"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
	var user models.User

	if err := json.NewDecoder(ctx.Request.Body).Decode(&user); err != nil {
		r.logger.Error("", zap.Error(err))

		var unmarshalError *json.UnmarshalTypeError

		if errors.As(err, &unmarshalError) {
			ctx.AbortWithError(http.StatusBadRequest, err)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	userID, err := r.db.InsertUser(ctx, &user)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	if user.ID != userID {
		user.ID = userID
	}

	ctx.JSON(http.StatusCreated, user)

}

func (r *Users) delete(ctx *gin.Context) {

}

func (r *Users) updateSegments(ctx *gin.Context) {

	var update models.Update

	if err := json.NewDecoder(ctx.Request.Body).Decode(&update); err != nil {
		r.logger.Error("", zap.Error(err))

		var unmarshalError *json.UnmarshalTypeError
		if errors.As(err, &unmarshalError) {
			ctx.AbortWithError(http.StatusBadRequest, err)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

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
			ctx.String(http.StatusNotFound, "User with id '%d' is not found", user.ID)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusAccepted, r.db.UpdateUserSegments(ctx, newUser, &update))

}

func (r *Users) listSegments(ctx *gin.Context) {
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
			ctx.String(http.StatusNotFound, "User with id '%d' is not found", user.ID)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	allSegments, err := r.db.ListAllUserSegments(ctx, newUser)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, allSegments)
}
