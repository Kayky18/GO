package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayky18/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error:%v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	//logger.Infof("Request Received: %+v", request)
	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating  opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error opening on database")
		return
	}
	sendSucess(ctx, "create-opening", opening)
}
