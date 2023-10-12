package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 2.3,
		SKU:   "abc-abc-asg",
	}
	err := p.ValidateProduct()

	if err != nil {
		t.Fatal(err)
	}
}
