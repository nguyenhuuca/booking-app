package service

import (
	"booking-service/model"
)

var products = []model.Product{
	{ID: "1", Title: "Blue Train", Branch: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Branch: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Branch: "Sarah Vaughan", Price: 39.99},
}

func GetProduct() []model.Product {
	return products
}

func FilterProduct(name string) []model.Product {
	filterRs := filter(products, func(p model.Product) bool {
		return p.Title == name
	})
	return filterRs
}

func filter(arr []model.Product, fn func(model.Product) bool) []model.Product {
	var filteredArr []model.Product
	for _, obj := range arr {
		if fn(obj) {
			filteredArr = append(filteredArr, obj)
		}
	}
	return filteredArr
}
