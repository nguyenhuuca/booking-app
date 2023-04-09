package logic

import (
	"booking-service/dto"
	"log"
	"testing"
)

type ProductRepoMock struct {
}

type AnalyzeMock struct {
}

func TestService_GetDataSuccess(t *testing.T) {
	expectValue := []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}

	productRepoMock := ProductRepoMock{}
	booking := CyloBooking{ProductRepo: productRepoMock}
	var rs = booking.GetProduct()

	if len(rs) != len(expectValue) {
		t.Errorf("got %d, wanted %d", len(rs), len(expectValue))
	}
	if rs[0].Branch != expectValue[0].Branch {
		t.Errorf("got %q, wanted %q", rs[0].Branch, expectValue[0].Branch)
	}
}

func TestService_SortSuccess(t *testing.T) {
	expectValue := []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}

	productRepoMock := ProductRepoMock{}
	analyzeMock := AnalyzeMock{}
	booking := CyloBooking{Name: "branch", SortType: "desc", AuditServ: analyzeMock}
	var rs, _ = booking.Sort(productRepoMock)

	if len(rs) != len(expectValue) {
		t.Errorf("got %d, wanted %d", len(rs), len(expectValue))
	}
}

func TestService_SortReturnErr(t *testing.T) {
	dbMock := ProductRepoMock{}
	booking := CyloBooking{Name: "branch11", SortType: "desc"}
	var _, err = booking.Sort(dbMock)
	if err == nil {
		t.Errorf("got success but want get error")
	}
}

func TestService_filter(t *testing.T) {
	productRepoMock := ProductRepoMock{}
	analyzeMock := AnalyzeMock{}

	booking := CyloBooking{Name: "branch", Branch: "test", ProductRepo: productRepoMock, AuditServ: analyzeMock}
	var rs = booking.FilterProduct()
	if len(rs) != 1 {
		t.Errorf("got %d, wanted %d", len(rs), 1)
	}
}

func (db ProductRepoMock) FindAll() []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}
	return products
}

func (db ProductRepoMock) FilterProduct(name string, branch string, price float64) []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: name, Branch: branch, Price: price},
	}
	return products
}

func (db ProductRepoMock) ShortBy(name string, shortType string) []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}
	return products
}

func (aMock AnalyzeMock) SendAudit(auditDto dto.AuditDto) {
	// do nothing
	log.Printf("Mocking SaveAudit")
}
