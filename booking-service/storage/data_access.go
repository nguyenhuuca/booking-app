package storage

/**
Using to get data from database.
*/
import (
	"booking-service/dto"
	"gorm.io/gorm"
)

const CommonJoin = "inner join branches on products.branch_id = branches.id"

type ProductRepo interface {
	FindAll() []dto.ProductDto
	FilterProduct(name string, branch string, price float64) []dto.ProductDto
	ShortBy(name string, shortType string) []dto.ProductDto
}

type AuditRepo interface {
	Save(audit Audit)
}
type AuditOrm struct {
	Instance *gorm.DB
}
type ProductGorm struct {
	Instance *gorm.DB
}

func (db ProductGorm) FilterProduct(name string, branch string, price float64) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products. id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Where("products.name like ? and branches.name like ? and products.price >= ? ", "%"+name+"%", "%"+branch+"%", price).
		Scan(&products)
	return products
}

func (db ProductGorm) ShortBy(name string, shortType string) []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Order(name + " " + shortType).
		Scan(&products)
	return products

}

func (db ProductGorm) FindAll() []dto.ProductDto {
	var products []dto.ProductDto
	db.Instance.
		Model(Product{}).
		Select("products.id, products.name, products.price, branches.name as branch").
		Joins(CommonJoin).
		Scan(&products)
	return products
}

func (db AuditOrm) Save(audit Audit) {
	db.Instance.Create(&audit)
}
