package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreatedOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorF("validation error: %v", err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.ErrorF("error creating opening: %+v", err.Error())
		return
	}
}
