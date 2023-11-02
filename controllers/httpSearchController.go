package controllers

import (
	"abstractions/views"
	"net/http"
)

type Search interface {
	GetProducts(string) string
}

func SearchController(w http.ResponseWriter, r *http.Request, searchTerm string, search Search) {
	products := search.GetProducts(searchTerm)

	views.JsonProducts(products)
}
