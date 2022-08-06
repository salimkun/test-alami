package lib

import (
	"errors"
)

type ProgrammingPharadigme interface {
	GetShipByID(id int64) (*Ship, error)
}

//Ship struct
type Ship struct {
	ShipID int64  `json:"ship_id"`
	Type   string `json:"type"`
}

//Get of ShipID
func (p *Ship) GetShipID() int64 {
	return p.ShipID
}

//Get of Type
func (p *Ship) GetType() string {
	return p.Type
}

func GetShipByID(id int64) (*Ship, error) {
	dataShip := []*Ship{
		{ShipID: 212, Type: "Perahu Motor"},
		{ShipID: 321, Type: "Perahu Layar"},
		{ShipID: 923, Type: "Kapal Pesiar"},
	}

	for _, i := range dataShip {
		if id == i.ShipID {
			return i, nil
		}
	}

	return nil, errors.New("ID SHIP NOT FOUND")
}
