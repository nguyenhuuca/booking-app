package api

import (
	"booking-service/logic"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// getProduct responds with the list of all product as JSON.
func (h handler) getProduct(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, logic.GetProduct(h.DB))
}

func (h handler) getFilter(c *gin.Context) {
	name := c.Query("name")
	c.IndentedJSON(http.StatusOK, logic.FilterProduct(h.DB, name))
}

func (h handler) sort(c *gin.Context) {
	sortCond := c.Query("name")
	sortType := c.Query("type")
	rs, err := logic.Sort(h.DB, sortCond, sortType)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, rs)
	}
}
