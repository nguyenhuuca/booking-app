package api

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// Define endpoints
	router.GET("/products", getProduct)
	router.GET("/products/filter", getFilter)

	return router
}
