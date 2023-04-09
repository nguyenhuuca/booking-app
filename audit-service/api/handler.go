package api

import (
	"audit-service/dto"
	"audit-service/logic"
	"audit-service/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// getProduct responds with the list of all product as JSON.
func (h handler) sendAudit(c *gin.Context) {
	var auditDto dto.AuditDto
	// BinJson
	if err := c.BindJSON(&auditDto); err != nil {
		log.Println("sendAudit err:", err)
		return
	}
	auditRepo := storage.AuditOrm{Instance: h.DB}
	analyze := logic.Analyze{AuditRepo: auditRepo}
	analyze.SaveAudit(auditDto)
	c.IndentedJSON(http.StatusOK, "Success")
}
