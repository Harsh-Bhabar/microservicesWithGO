package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPost {
		p.AddProduct(w, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("Coming in PUT")
		// gorilla mux should be used here, will refractor
		regex := regexp.MustCompile(`/([0-9]+)`)
		g := regex.FindAllStringSubmatch(r.URL.Path, -1)

		p.l.Println("regex ", regex)
		p.l.Println("g", g)

		if len(g) != 1 {
			p.l.Println("Invalid URL, more than one ID.")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("Invalid URL, more than one capture group.")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URL, while converting.")
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		p.l.Println("Id - ", id)
		p.UpdateProduct(id, w, r)
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

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.Decoder(r.Body)

	if err != nil {
		http.Error(w, "Unable to decode", http.StatusBadRequest)
	}

	p.l.Println("product", product)

	data.AddProductToStaticDB(product)
}

func (p *Products) UpdateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Updating product", id)

	product := &data.Product{}

	err := product.Decoder(r.Body)
	if err != nil {
		http.Error(w, "Unable to encode", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)

	if err == data.ProdcutNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
