package logic

import (
	"audit-service/dto"
	"audit-service/storage"
)

type AuditServ interface {
	SaveAudit(audit dto.AuditDto)
}

type Analyze struct {
	AuditRepo storage.AuditRepo
}

func (a Analyze) SaveAudit(auditDto dto.AuditDto) {
	auditData := storage.Audit{Identifier: auditDto.Identifier, Data: auditDto.Data, Action: auditDto.Action}
	a.AuditRepo.Save(auditData)
}
