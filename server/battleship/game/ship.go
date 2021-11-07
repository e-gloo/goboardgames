package game

type Ship struct {
	Cells          []uint8
	RemainingAlive uint8
}

type Direction bool

const (
	HORIZONTAL Direction = false
	VERTICAL             = true
)

func NewShip(startCell uint8, length uint8, direction Direction) Ship {
	cells := make([]uint8, length)
	ship := Ship{
		Cells:          cells,
		RemainingAlive: length,
	}
	for i := range ship.Cells {
		if direction == HORIZONTAL {
			ship.Cells[i] = startCell + uint8(i)
		} else {
			ship.Cells[i] = startCell + uint8(i)*MAP_WIDTH
		}
	}
	return ship
}

func checkVertical(cells map[uint8]uint8, start uint8, shipLen uint8) bool {
	// premier filtre sur les bordures
	y := start / MAP_WIDTH
	if shipLen+y >= MAP_HEIGHT {
		return false
	}

	for i := uint8(1); i < shipLen; i++ {
		if _, ok := cells[start+i*MAP_WIDTH]; !ok {
			return false
		}
	}
	return true
}

func checkHorizontal(cells map[uint8]uint8, start uint8, shipLen uint8) bool {
	// premier filtre sur les bordures
	x := start % MAP_WIDTH
	if shipLen+x >= MAP_WIDTH {
		return false
	}

	for i := uint8(1); i < shipLen; i++ {
		if _, ok := cells[start+i]; !ok {
			return false
		}
	}
	return true
}
