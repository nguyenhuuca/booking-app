package logic

import (
	"booking-service/dto"
	"booking-service/restclient"
	"booking-service/storage"
	"log"
	"net/http"
	"os"
)

type AuditServ interface {
	SendAudit(audit dto.AuditDto) bool
}

type Analyze struct {
	AuditRepo storage.AuditRepo
}

func (a Analyze) SendAudit(auditDto dto.AuditDto) bool {
	var auditBaseUrl = os.Getenv("AUDIT_SERVICE")
	if auditBaseUrl != "" {
		err := sendAuditViaApi(auditDto)
		if err != nil {
			return false
		}
		return true
	} else {
		log.Println("Save audit...")
		auditData := storage.Audit{Identifier: auditDto.Identifier, Data: auditDto.Data, Action: auditDto.Action}
		a.AuditRepo.Save(auditData)
		return true
	}
}

func sendAuditViaApi(auditDto dto.AuditDto) error {
	log.Println("Send audit...")
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")
	_, err := restclient.Post(os.Getenv("AUDIT_SERVICE")+"/send-audits", auditDto, headers)
	if err != nil {
		log.Println("Error", err)
	}
	return err

}
