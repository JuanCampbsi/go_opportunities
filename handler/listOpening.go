package handler

import (
	"net/http"

	"github.com/JuanCampbsi/go-opportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "erro listing openings")
		return
	}
	lenght := len(openings)
	if lenght == 0 {
		sendSuccess(ctx, "listings openings is empty", openings)
	} else {
		sendSuccess(ctx, "listings openings", openings)
	}
}
