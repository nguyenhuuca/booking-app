package logic

import (
	"booking-service/dto"
	"booking-service/storage"
	"booking-service/utils"
	"gorm.io/gorm"
	"log"
)

var whiteList = []string{"name", "price", "branch", "asc", "desc"}

type Db interface {
	findAll() []dto.ProductDto
	filterProduct(name string, branch string, price float64) []dto.ProductDto
	shortBy(name string, shortType string) []dto.ProductDto
}

type BookingServ interface {
	GetProduct(db *gorm.DB) []dto.ProductDto
	FilterProduct(db *gorm.DB) []dto.ProductDto
	Sort(db *gorm.DB) ([]dto.ProductDto, error)
}

type CyloBooking struct {
	Name     string
	Branch   string
	Price    float64
	SortType string
}

type OrmDb struct {
	Instance *gorm.DB
}

func (cylo CyloBooking) GetProduct(dbService Db) []dto.ProductDto {
	var products []dto.ProductDto
	products = dbService.findAll()
	return products
}

func (cylo CyloBooking) FilterProduct(dbService Db) []dto.ProductDto {
	var products []dto.ProductDto
	products = dbService.filterProduct(cylo.Name, cylo.Branch, cylo.Price)
	return products
}

func (cylo CyloBooking) Sort(dbService Db) ([]dto.ProductDto, error) {
	var products []dto.ProductDto
	fieldName, shortName, err := getFieldNameToOrder(cylo.Name, cylo.SortType)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	products = dbService.shortBy(fieldName, shortName)
	return products, nil

}

func getFieldNameToOrder(name string, shortType string) (string, string, error) {
	if !contains(whiteList, name) {
		return "", "", utils.NewDefaultError("error get field name to sort ")
	}
	if !contains(whiteList, shortType) {
		return "", "", utils.NewDefaultError("error get field name to sort ")
	}
	return name, shortType, nil

}

func (db OrmDb) findAll() []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(&storage.Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Scan(&products)
	return products
}

func (db OrmDb) filterProduct(name string, branch string, price float64) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(&storage.Product{}).
		Select("products. id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Where("products.name like ? and branches.name like ? and products.price >= ? ", "%"+name+"%", "%"+branch+"%", price).
		Scan(&products)
	return products
}

func (db OrmDb) shortBy(name string, shortType string) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(&storage.Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Order(name + " " + shortType).
		Scan(&products)
	return products

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
