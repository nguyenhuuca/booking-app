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
	SendAudit(audit dto.AuditDto)
}

type Analyze struct {
	AuditRepo storage.AuditRepo
}

func (a Analyze) SendAudit(auditDto dto.AuditDto) {
	var auditBaseUrl = os.Getenv("AUDIT_SERVICE")
	if auditBaseUrl != "" {
		sendAuditViaApi(auditDto)
	} else {
		a.saveAudit(auditDto)
	}
}

func (a Analyze) saveAudit(auditDto dto.AuditDto) {
	log.Println("Save audit...")
	auditData := storage.Audit{Identifier: auditDto.Identifier, Data: auditDto.Data, Action: auditDto.Action}
	a.AuditRepo.Save(auditData)
}

func sendAuditViaApi(auditDto dto.AuditDto) {
	log.Println("Send audit...")
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")
	go func() {
		_, err := restclient.Post(os.Getenv("AUDIT_SERVICE")+"/send-audits", auditDto, headers)
		if err != nil {
			log.Println("Error", err)
		}
	}()

}
