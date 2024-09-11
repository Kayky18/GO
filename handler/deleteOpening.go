package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayky18/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "QueryParameter").Error())
	}

	opening := schemas.Opening{}

	//Find Opening
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("Error to find opening with id %s", id))
		return
	}

	//Delete Opening
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("Error to delete opening with id %s", id))
		return
	}
	sendSucess(ctx, "delete-opening", id)
}
