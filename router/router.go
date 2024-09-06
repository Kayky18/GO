package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize Router
	router := gin.Default()

	//Initialize Routes
	InitializeRoutes(router)

	//Estamos rodando a nossa api, por padrão já usa a 8080
	router.Run(":8080")
}
