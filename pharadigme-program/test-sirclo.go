package main

import (
	"fmt"

	"github.com/salimkun/sirclo-test/pharadigme-program/lib"
)

//Test function
func main() {
	// pharadigme-program
	var p lib.ProgrammingPharadigme = &lib.Ship{}
	resp := p.GetShipByID(212)
	if resp == nil {
		fmt.Println("Error Id not Found")
	} else {
		fmt.Println("SHIP ID : ", resp.ShipID)
		fmt.Println("TYPE ID : ", resp.Type)
	}
}
