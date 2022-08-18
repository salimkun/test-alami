package main

import (
	"fmt"
	"log"
	"os"

	pp "github.com/salimkun/sirclo-test/pharadigme-program/lib"
	sc "github.com/salimkun/sirclo-test/shopping-chart/lib"
)

func pharadigmeProgram() {
	// pharadigme-program
	var p pp.ProgrammingPharadigme = &pp.Ship{}
	resp := p.GetShipByID(212)
	if resp == nil {
		fmt.Println("Error Id not Found")
	} else {
		fmt.Println("SHIP ID : ", resp.ShipID)
		fmt.Println("TYPE ID : ", resp.Type)
	}
}

func createFileTemp() {
	// Shopping chart
	// create file to temp database
	_, err := os.Create("shopping-db.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
}

func viewAllproduct() {
	// get file
	var sc sc.ProductInterface = &sc.ListProduct{}

	data := sc.ViewAllProduct()
	fmt.Println("============ DATA PRODUCT ===============")
	for _, a := range data {
		fmt.Println("Product Code : ", a.ProductCode)
		fmt.Println("Product Name : ", a.ProductName)
		fmt.Println("QTY : ", a.QTY)
	}
}

func addProduct(product *sc.Product) {
	var sc sc.ProductInterface = &sc.ListProduct{}

	sc.AddProduct(product)
}

func deleteProduct(code string) {
	var sc sc.ProductInterface = &sc.ListProduct{}

	sc.DeleteProduct(code)
}

//Test function
func main() {
	pharadigmeProgram()

	//shopping chart
	dataProduct := []*sc.Product{
		{ProductCode: "UC2000", ProductName: "UC 2000", QTY: 21},
		{ProductCode: "KPAPI21", ProductName: "Kapal API Hitam 21", QTY: 22},
		{ProductName: "MYRSK", ProductCode: "My Roti Selai Kacang", QTY: 1},
		{ProductName: "MYRSK", ProductCode: "My Roti Selai Kacang", QTY: 1},
		{ProductName: "MYRSK", ProductCode: "My Roti Selai Kacang", QTY: 1},
	}

	createFileTemp()

	// add product
	addProduct(dataProduct[0])
	addProduct(dataProduct[1])
	addProduct(dataProduct[2])
	addProduct(dataProduct[3])
	addProduct(dataProduct[4])

	// // delete product
	deleteProduct("UC2000")

	// view All product
	viewAllproduct()
}
