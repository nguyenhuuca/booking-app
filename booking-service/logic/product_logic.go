// Package logic /*
package logic

/**
Using to handle business from booking
*/
import (
	"booking-service/dto"
	"booking-service/storage"
	"booking-service/utils"
	"gorm.io/gorm"
	"log"
)

var whiteList = []string{"name", "price", "branch", "asc", "desc"}

type BookingServ interface {
	GetProduct(db *gorm.DB) []dto.ProductDto
	FilterProduct(db *gorm.DB) []dto.ProductDto
	Sort(db *gorm.DB) ([]dto.ProductDto, error)
}

type CyloBooking struct {
	Name        string
	Branch      string
	Price       float64
	SortType    string
	ProductRepo storage.ProductRepo
	AuditServ   AuditServ
}

func (cylo CyloBooking) GetProduct() []dto.ProductDto {
	var products []dto.ProductDto
	products = cylo.ProductRepo.FindAll()
	return products
}

func (cylo CyloBooking) FilterProduct() []dto.ProductDto {
	var products []dto.ProductDto
	products = cylo.ProductRepo.FilterProduct(cylo.Name, cylo.Branch, cylo.Price)
	cylo.AuditServ.sendAudit(dto.AuditDto{Identifier: "test",
		Action: "Filter",
		Data:   dto.ProductDto{Price: cylo.Price, Name: cylo.Name, Branch: cylo.Branch}})
	return products
}

func (cylo CyloBooking) Sort(productRepo storage.ProductRepo) ([]dto.ProductDto, error) {
	var products []dto.ProductDto
	fieldName, shortName, err := getFieldNameToOrder(cylo.Name, cylo.SortType)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	products = productRepo.ShortBy(fieldName, shortName)
	return products, nil

}

// avoid sql injection when using some built-in function from GORM
// ref: https://gorm.io/docs/security.html
func getFieldNameToOrder(name string, shortType string) (string, string, error) {
	if !contains(whiteList, name) {
		return "", "", utils.NewDefaultError("error get field name to sort ")
	}
	if !contains(whiteList, shortType) {
		return "", "", utils.NewDefaultError("error get field name to sort ")
	}
	return name, shortType, nil

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
