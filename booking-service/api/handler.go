package api

import (
	"booking-service/logic"
	"booking-service/storage"
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

	productRepo := storage.ProductGorm{Instance: h.DB}
	booking := logic.CyloBooking{ProductRepo: productRepo}
	c.IndentedJSON(http.StatusOK, booking.GetProduct())
}

func (h handler) getFilter(c *gin.Context) {

	name := c.Query("name")
	branch := c.Query("branch")
	var price = 0.0
	if c.Query("price") != "" {
		price, _ = strconv.ParseFloat(c.Query("price"), 2)
	}

	productRepo := storage.ProductGorm{Instance: h.DB}
	auditRepo := storage.AuditOrm{Instance: h.DB}
	analyze := logic.Analyze{AuditRepo: auditRepo}
	booking := logic.CyloBooking{Name: name, Branch: branch, Price: price, ProductRepo: productRepo, AuditServ: analyze}
	c.IndentedJSON(http.StatusOK, booking.FilterProduct())
}

func (h handler) sort(c *gin.Context) {
	sortCond := c.Query("name")
	sortType := c.Query("type")

	productRepo := storage.ProductGorm{Instance: h.DB}
	auditRepo := storage.AuditOrm{Instance: h.DB}
	analyze := logic.Analyze{AuditRepo: auditRepo}

	booking := logic.CyloBooking{Name: sortCond, SortType: sortType, ProductRepo: productRepo, AuditServ: analyze}

	rs, err := booking.Sort(productRepo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusOK, rs)
	}
}
