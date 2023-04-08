package storage

/**
Using to get data from database.
*/
import (
	"booking-service/dto"
	"gorm.io/gorm"
)

const CommonJoin = "inner join branches on products.branch_id = branches.id"

type Db interface {
	FindAll() []dto.ProductDto
	FilterProduct(name string, branch string, price float64) []dto.ProductDto
	ShortBy(name string, shortType string) []dto.ProductDto
}
type OrmDb struct {
	Instance *gorm.DB
}

func (db OrmDb) FilterProduct(name string, branch string, price float64) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products. id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Where("products.name like ? and branches.name like ? and products.price >= ? ", "%"+name+"%", "%"+branch+"%", price).
		Scan(&products)
	return products
}

func (db OrmDb) ShortBy(name string, shortType string) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Order(name + " " + shortType).
		Scan(&products)
	return products

}

func (db OrmDb) FindAll() []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Scan(&products)
	return products
}
