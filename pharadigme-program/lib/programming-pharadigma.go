package lib

type ProgrammingPharadigme interface {
	GetShipByID(id int64) *Ship
}

//Ship struct
type Ship struct {
	ShipID int64  `json:"ship_id"`
	Type   string `json:"type"`
}

func (p *Ship) GetShipByID(id int64) *Ship {
	dataShip := []*Ship{
		{ShipID: 212, Type: "Perahu Motor"},
		{ShipID: 321, Type: "Perahu Layar"},
		{ShipID: 923, Type: "Kapal Pesiar"},
	}

	for _, i := range dataShip {
		if id == i.ShipID {
			return i
		}
	}

	return nil
}
