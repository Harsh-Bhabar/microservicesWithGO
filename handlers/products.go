package handlers

import (
	"log"
	"net/http"

	"github.com/Harsh-Bhabar/products-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetAllProducts(w, r)
		return
	}

	// catch all else
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productsList := data.GetAllProducts()
	err := productsList.EncodeToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusInternalServerError)
	}
	// another way by using Marshal
	// data, err := json.Marshal(productsList) // another way
	// w.Write(data)

}
