package api

import (
	"booking-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getProduct responds with the list of all product as JSON.
func getProduct(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, service.GetProduct())
}

func getFilter(c *gin.Context) {
	name := c.Query("name")
	c.IndentedJSON(http.StatusOK, service.FilterProduct(name))
}
