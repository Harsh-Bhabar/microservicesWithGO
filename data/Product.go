package data

import (
	"encoding/json"
	"fmt"
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

func (p *Product) Decoder(r io.Reader) error {
	err := json.NewDecoder(r)
	return err.Decode(p)
}

type Products []*Product

func (p *Products) EncodeToJSON(w io.Writer) error {
	err := json.NewEncoder(w)
	return err.Encode(p)
}

func GetAllProducts() Products {
	return ProductList
}

func AddProductToStaticDB(p *Product) {
	p.ID = getNextId()
	ProductList = append(ProductList, p)
}

func UpdateProduct(id int, p *Product) error {
	pos, err := findProduct(id)
	if err != nil {
		return err
	}
	// fp.ID = id
	ProductList[pos] = p
	return nil
}

var ProdcutNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return i, nil
		}
	}
	return -1, ProdcutNotFound
}

func getNextId() int {
	lastProd := ProductList[len(ProductList)-1]
	return lastProd.ID + 1
}

var ProductList = []*Product{
	&Product{ID: 1, Name: "Product 1", Desc: "Product-1 Description", Price: 111, SKU: "Product-1 SKU", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
	&Product{ID: 2, Name: "Product 2", Desc: "Product-2 Description", Price: 222, SKU: "Product-2 SKU", CreatedOn: time.Now().UTC().String(), UpdatedOn: time.Now().UTC().String()},
}
