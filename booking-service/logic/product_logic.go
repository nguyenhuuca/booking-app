package logic

import (
	"booking-service/dto"
	"booking-service/storage"
	"errors"
	"gorm.io/gorm"
	"log"
)

var whiteList = []string{"name", "price", "branch"}

func GetProduct(db *gorm.DB) []dto.ProductDto {
	var products []dto.ProductDto
	db.
		Model(&storage.Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Scan(&products)
	return products
}

func FilterProduct(db *gorm.DB, name string) []dto.ProductDto {
	var products []dto.ProductDto
	db.
		Model(&storage.Product{}).
		Select("products. id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Where("products.name like ?", "%"+name+"%").
		Scan(&products)
	return products
}
func Sort(db *gorm.DB, name string, typeSort string) ([]dto.ProductDto, error) {
	var products []dto.ProductDto
	fieldName, shortName, err := getFieldNameToOrder(name, typeSort)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.
		Model(&storage.Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins("inner join branches on products.branch_id = branches.id").
		Order(fieldName + " " + shortName).
		Scan(&products)
	return products, nil

}

func getFieldNameToOrder(name string, shortType string) (string, string, error) {
	if contains(whiteList, name) {
		return name, shortType, nil
	}
	return "", "", errors.New("error get field name ")

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
