package logic

import (
	"booking-service/dto"
	"booking-service/restclient"
	"booking-service/storage"
	"errors"
	"log"
	"net/http"
	"os"
	"testing"
)

type AuditRepoMock struct {
}
type ClientMock struct {
}

type ClientMockFalseCall struct {
}

func TestAnalyze_SendAuditSaveMode(t *testing.T) {
	mockDto := dto.AuditDto{Identifier: "t", Data: "test", Action: "Filter"}
	auditRepoMock := AuditRepoMock{}
	analyze := Analyze{AuditRepo: auditRepoMock}
	rs := analyze.SendAudit(mockDto)
	if rs == false {
		t.Fail()
	}
}

func TestAnalyze_SendAuditSendApiMode(t *testing.T) {
	err := os.Setenv("AUDIT_SERVICE", "http://localhost:8081")
	if err != nil {
		t.Fail()
	}
	mockDto := dto.AuditDto{Identifier: "t", Data: "test", Action: "Filter"}
	auditRepoMock := AuditRepoMock{}
	analyze := Analyze{AuditRepo: auditRepoMock}
	restclient.Client = ClientMock{}
	rs := analyze.SendAudit(mockDto)
	if rs == false {
		t.Fail()
	}
}

func TestAnalyze_SendAuditSendApiModeFalse(t *testing.T) {
	err := os.Setenv("AUDIT_SERVICE", "http://localhost:8081")
	if err != nil {
		t.Fail()
	}
	mockDto := dto.AuditDto{Identifier: "t", Data: "test", Action: "Filter"}
	auditRepoMock := AuditRepoMock{}
	analyze := Analyze{AuditRepo: auditRepoMock}
	restclient.Client = &ClientMockFalseCall{}
	rs := analyze.SendAudit(mockDto)
	if rs == true {
		t.Fail()
	}
}

func (auditRepoMock AuditRepoMock) Save(audit storage.Audit) {
	log.Printf("Call nothing from mock")
}

func (clientMock ClientMock) Do(req *http.Request) (*http.Response, error) {
	log.Println("do no thing")
	return nil, nil
}

func (clientMock ClientMockFalseCall) Do(req *http.Request) (*http.Response, error) {
	log.Println("do no thing")
	return nil, errors.New("call api exception")
}
