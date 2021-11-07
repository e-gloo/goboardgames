package game

import (
	"github.com/e-gloo/goboardgames/utils"
)

type PlayerFleet struct {
	Fleet []Ship
	Hits  []uint8
}

func NewFleet() []Ship {
	cells := map[uint8]uint8{}
	ships := make([]Ship, len(SHIPS_LENGTH))
	for i := uint8(0); i < MAP_WIDTH*MAP_HEIGHT; i++ {
		cells[i] = i
	}

	for i := 0; i < len(SHIPS_LENGTH); i++ {
		shipLen := SHIPS_LENGTH[i]
		dir := Direction(utils.RandBool())

		possibleStarts := []uint8{}
		for _, cell := range cells {
			if (dir == HORIZONTAL && checkHorizontal(cells, cell, shipLen)) ||
				(dir == VERTICAL && checkVertical(cells, cell, shipLen)) {
				possibleStarts = append(possibleStarts, uint8(cell))
			}
		}
		randVal := possibleStarts[utils.RandInt(0, len(possibleStarts)-1)]
		ship := NewShip(randVal, shipLen, dir)
		for _, val := range ship.Cells {
			delete(cells, val)
			if val%MAP_WIDTH > 0 {
				delete(cells, val-1)
			}
			if val%MAP_WIDTH < MAP_WIDTH-1 {
				delete(cells, val+1)
			}
			if val/MAP_WIDTH > 0 {
				delete(cells, val-MAP_WIDTH)
			}
			if val/MAP_WIDTH < MAP_HEIGHT-1 {
				delete(cells, val+MAP_WIDTH)
			}
		}
		ships[i] = ship
	}
	return ships
}

func (pf *PlayerFleet) IsAlreadyHit(position uint8) bool {
	for _, v := range pf.Hits {
		if position == v {
			return true
		}
	}
	return false
}

func (pf *PlayerFleet) GetShipOnCell(position uint8) *Ship {
	for skey := range pf.Fleet {
		for ckey := range pf.Fleet[skey].Cells {
			if pf.Fleet[skey].Cells[ckey] == position {
				return &pf.Fleet[skey]
			}
		}
	}
	return nil
}

func (pf *PlayerFleet) Attack(position uint8) *Ship {
	ship := pf.GetShipOnCell(position)
	pf.Hits = append(pf.Hits, position)
	if ship == nil {
		return nil
	}
	ship.RemainingAlive -= 1
	return ship
}
