package storage

/**
Using to get data from database.
*/
import (
	"gorm.io/gorm"
)

type AuditRepo interface {
	Save(audit Audit)
}
type AuditOrm struct {
	Instance *gorm.DB
}
type ProductGorm struct {
	Instance *gorm.DB
}

func (db AuditOrm) Save(audit Audit) {
	db.Instance.Create(&audit)
}
