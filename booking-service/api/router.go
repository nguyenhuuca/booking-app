package api

import (
	"booking-service/db"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	DB := db.Init()
	h := New(DB)
	router := gin.Default()

	// Define endpoints
	router.GET("/products", h.getProduct)
	router.GET("/products/filter", h.getFilter)
	router.GET("/products/short", h.sort)

	return router
}
