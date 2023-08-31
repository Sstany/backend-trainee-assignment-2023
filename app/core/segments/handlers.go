package segments

import (
	"net/http"
	"segmenty/app/db/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *Segments) fetch(ctx *gin.Context) {
	var segment *models.Segment

	err := ctx.ShouldBindUri(&segment)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	newSegment, isNoRows, err := r.db.FetchSegment(ctx, segment.Name)
	if err != nil {
		r.logger.Error("", zap.Error(err))

		if isNoRows {
			ctx.String(http.StatusNotFound, "Segment with name '%s' is not found", segment.Name)

			return
		}

		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}
	ctx.JSON(http.StatusOK, newSegment)

}

func (r *Segments) create(ctx *gin.Context) {
	var segment *models.Segment

	err := ctx.ShouldBindUri(&segment)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusBadRequest, err)

		return
	}

	segmentID, isNotUnique, err := r.db.InsertSegment(ctx, segment)
	if err != nil {
		if isNotUnique {
			ctx.String(http.StatusConflict, "Segment with name '%s' is already exists", segment.Name)

			return
		}

		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	segment.ID = segmentID

	ctx.JSON(http.StatusCreated, segment)

}

func (r *Segments) delete(ctx *gin.Context) {
	var segment *models.Segment

	err := ctx.ShouldBindUri(&segment)
	if err != nil {
		r.logger.Error("", zap.Error(err))

		return
	}

	newSegment, err := r.db.DeleteSegment(ctx, segment.Name)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, newSegment)
}

func (r *Segments) list(ctx *gin.Context) {
	allSegments, err := r.db.ListSegments(ctx)
	if err != nil {
		r.logger.Error("", zap.Error(err))
		ctx.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	ctx.JSON(http.StatusOK, allSegments)
}
