package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leoneville/gopportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid id type")
		return
	}

	opening := schemas.Opening{}

	err = db.WithContext(ctx).First(&opening, id).Error
	if err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %v not found", id))
		return
	}

	sendSuccess(ctx, http.StatusOK, "show-opening", opening)
}
