package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leoneville/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.WithContext(ctx).Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, http.StatusOK, "list-openings", openings)
}
