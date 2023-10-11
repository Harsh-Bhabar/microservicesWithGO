package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	Price     float32 `json:"price"`
	SKU       string  `json:"sku"`
	CreatedOn string  `json:"-"`
	DeletedOn string  `json:"-"`
	UpdatedOn string  `json:"-"`
}

type Products []*Product

func (p *Products) EncodeToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetAllProducts() Products {
	return ProductList
}

var ProductList = []*Product{
	&Product{ID: 1, Name: "Product 1", Desc: "Product-1 Description", Price: 111, SKU: "Product-1 SKU", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
	&Product{ID: 2, Name: "Product 2", Desc: "Product-2 Description", Price: 222, SKU: "Product-2 SKU", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
}
