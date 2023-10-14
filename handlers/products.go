// package classification for product-api

// Documentation for product-api
// schemes: http
// Basepath: /products
// Version: 1.0.0
// Consumes: application/json
// Produces: application/json
//

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Harsh-Bhabar/products-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("COming in GET ")

	productsList := data.GetAllProducts()
	err := productsList.EncodeToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Coming in POST")

	// middleware
	// product := r.Context().Value(KeyProduct{}).(data.Product)
	// data.AddProductToStaticDB(&product)

	//normally
	product := &data.Product{}
	if err := product.Decoder(r.Body); err != nil {
		p.l.Println("Error decoding product:", err)
		http.Error(w, "Unable to decode", http.StatusBadRequest)
		return
	}

	p.l.Println("product", product)

	data.AddProductToStaticDB(product)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	p.l.Println("COming in pUT ")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Unable to parse id", http.StatusBadRequest)
		return
	}

	p.l.Println("Updating product", id)

	product := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &product)

	if err == data.ProdcutNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		// Validate the product data
		product := &data.Product{}
		err := product.Decoder(r.Body)
		if err != nil {
			http.Error(rw, "Unable to decode", http.StatusBadRequest)
			return
		}

		//validate the product
		err = product.ValidateProduct()
		if err != nil {
			p.l.Println("Error validating product")
			http.Error(rw,
				fmt.Sprintf("Unable to validate, error - %s\n", err),
				http.StatusBadRequest,
			)
			return
		}

		// Set the product data in the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, *product)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
