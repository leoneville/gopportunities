package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leoneville/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param id path int true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
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
