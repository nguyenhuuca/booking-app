package logic

import (
	"booking-service/dto"
	"booking-service/restclient"
	"booking-service/storage"
	"encoding/json"
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
	product, err := json.Marshal(auditDto.Data)
	if err != nil {
		log.Printf("Error to parse auditDto data")
		return
	}
	auditData := storage.Audit{Identifier: auditDto.Identifier, Data: string(product), Action: auditDto.Action}
	a.AuditRepo.Save(auditData)
}

func sendAuditViaApi(auditDto dto.AuditDto) {
	log.Println("Send audit...")
	var headers = http.Header{}
	headers.Add("Content-Type", "application/json")

	_, err := restclient.Post(os.Getenv("AUDIT_SERVICE")+"/send-audits", auditDto, headers)
	if err != nil {
		log.Println("Error", err)
	}
	//go func() {
	//	_, err := http.Post(os.Getenv("AUDIT_SERVICE")+"/send-audits", "application/json", responseBody)
	//	if err != nil {
	//		log.Println("Error", err)
	//	}
	//}()

}
