package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ProductInterface interface {
	AddProduct(*Product)
	DeleteProduct(string)
	ViewAllProduct() []*Product
}

//Ship struct
type Product struct {
	ProductCode string `json:"product_code"`
	ProductName string `json:"product_name"`
	QTY         int32  `json:"qty"`
}

type ListProduct struct {
	Products []*Product `json:"product"`
}

func (p *ListProduct) AddProduct(pds *Product) {
	dt := p.ViewAllProduct()
	for c, i := range dt {
		if pds.ProductCode == i.ProductCode {
			dt[c].QTY = i.QTY + pds.QTY
			// write file
			p.Products = dt
			file, _ := json.MarshalIndent(p, "", " ")
			_ = ioutil.WriteFile("shopping-db.txt", file, 0644)
			return
		}
	}
	dt = append(dt, pds)
	p.Products = dt
	// write file
	file, _ := json.MarshalIndent(p, "", " ")
	_ = ioutil.WriteFile("shopping-db.txt", file, 0644)
}

func (p *ListProduct) DeleteProduct(code string) {
	dt := p.ViewAllProduct()
	dataProductNew := ListProduct{}

	for _, i := range dt {
		if code != i.ProductCode {
			dataProductNew.Products = append(dataProductNew.Products, i)
		}
	}

	// write file
	file, _ := json.MarshalIndent(dataProductNew, "", " ")
	_ = ioutil.WriteFile("shopping-db.txt", file, 0644)
}

func (p *ListProduct) ViewAllProduct() []*Product {
	// get all data
	data, err := ioutil.ReadFile("shopping-db.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	dataProduct := ListProduct{}
	_ = json.Unmarshal([]byte(data), &dataProduct)
	return dataProduct.Products
}
