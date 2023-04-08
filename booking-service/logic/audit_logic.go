package logic

import (
	"booking-service/dto"
	"booking-service/storage"
	"encoding/json"
	"log"
)

type AuditServ interface {
	sendAudit(audit dto.AuditDto)
}

type Analyze struct {
	AuditRepo storage.AuditRepo
}

func (a Analyze) sendAudit(auditDto dto.AuditDto) {
	product, err := json.Marshal(auditDto.Data)
	if err != nil {
		log.Printf("Error to parse auditDto data")
		return
	}
	auditData := storage.Audit{Identifier: auditDto.Identifier, Data: string(product), Action: auditDto.Action}
	a.AuditRepo.Save(auditData)
}
