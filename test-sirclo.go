package main

import (
	"fmt"

	lib "./lib"
)

//Test function
func Test() {
	// pharadigme-programm

	var e1 lib.ProgrammingPharadigme
	resp, err := e1.GetShipByID(212)
	if err != nil {
		fmt.Println("Error ", err)
	}

	fmt.Println("SHIP ID : ", resp.ShipID)
	fmt.Println("TYPE ID : ", resp.Type)
}
