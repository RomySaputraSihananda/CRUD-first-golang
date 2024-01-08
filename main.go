package main

import (
	"fmt"
	"net/http"

	docs "app/docs"
	services "app/libs/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	PORT int = 4444
	TAG string = "/api/v1"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	router.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger-ui/index.html")
	})

	router.POST(fmt.Sprintf("%s/student", TAG), services.Create)
	router.GET(fmt.Sprintf("%s/student", TAG), services.Read)

	router.Run(fmt.Sprintf(":%d", PORT))
}