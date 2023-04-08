package logic

import (
	"booking-service/dto"
	"testing"
)

type DBServiceMock struct {
}

func TestService_GetDataSuccess(t *testing.T) {
	expectValue := []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}

	dbMock := DBServiceMock{}
	booking := CyloBooking{}
	var rs = booking.GetProduct(dbMock)

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

	dbMock := DBServiceMock{}
	booking := CyloBooking{Name: "branch", SortType: "desc"}
	var rs, _ = booking.Sort(dbMock)

	if len(rs) != len(expectValue) {
		t.Errorf("got %d, wanted %d", len(rs), len(expectValue))
	}
}

func TestService_SortReturnErr(t *testing.T) {
	dbMock := DBServiceMock{}
	booking := CyloBooking{Name: "branch11", SortType: "desc"}
	var _, err = booking.Sort(dbMock)
	if err == nil {
		t.Errorf("got success but want get error")
	}
}

func TestService_filter(t *testing.T) {
	dbMock := DBServiceMock{}
	booking := CyloBooking{Name: "branch", Branch: "test"}
	var rs = booking.FilterProduct(dbMock)
	if len(rs) != 1 {
		t.Errorf("got %d, wanted %d", len(rs), 1)
	}
}

func (db DBServiceMock) FindAll() []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}
	return products
}

func (db DBServiceMock) FilterProduct(name string, branch string, price float64) []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: name, Branch: branch, Price: price},
	}
	return products
}

func (db DBServiceMock) ShortBy(name string, shortType string) []dto.ProductDto {
	var products = []dto.ProductDto{
		{ID: "1", Name: "Apple", Branch: "Test1", Price: 15.0},
		{ID: "2", Name: "Samsung", Branch: "Test1", Price: 16.0},
	}
	return products
}
