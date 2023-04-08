package api

import (
	"booking-service/logic"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// getProduct responds with the list of all product as JSON.
func (h handler) getProduct(c *gin.Context) {
	booking := logic.CyloBooking{}

	dbService := logic.OrmDb{Instance: h.DB}
	c.IndentedJSON(http.StatusOK, booking.GetProduct(dbService))
}

func (h handler) getFilter(c *gin.Context) {

	name := c.Query("name")
	branch := c.Query("branch")
	var price = 0.0
	if c.Query("price") != "" {
		price, _ = strconv.ParseFloat(c.Query("price"), 2)
	}
	booking := logic.CyloBooking{Name: name, Branch: branch, Price: price}
	dbService := logic.OrmDb{Instance: h.DB}
	c.IndentedJSON(http.StatusOK, booking.FilterProduct(dbService))
}

func (h handler) sort(c *gin.Context) {
	sortCond := c.Query("name")
	sortType := c.Query("type")

	booking := logic.CyloBooking{Name: sortCond, SortType: sortType}
	dbService := logic.OrmDb{Instance: h.DB}
	rs, err := booking.Sort(dbService)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusOK, rs)
	}
}
