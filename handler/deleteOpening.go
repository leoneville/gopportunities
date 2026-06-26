package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leoneville/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id type")
		return
	}

	opening := schemas.Opening{}

	if err := db.WithContext(ctx).First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	if err := db.WithContext(ctx).Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error opening with id: %v", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "delete-opening", opening)
}
