package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayky18/gopportunities/schemas"
)

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	openings := &schemas.Opening{}

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "error id is empty")
		return
	}

	if err := db.First(&openings, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "error opening not found")
		return
	}

	sendSucess(ctx, "show-object-sucess", openings)
}
