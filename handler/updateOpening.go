package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayky18/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Update Opening
// @Description Update a new opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string  true "Opening identification"
// @Param request body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400  {object} ErrorResponse
// @Failure 404  {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening id not found")
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSucess(ctx, "sucess-put", opening)
}
