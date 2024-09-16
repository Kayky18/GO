package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/kayky18/gopportunities/docs"
	"github.com/kayky18/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	handler.InitializeHanler()
	basePath := "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		//Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}

	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
