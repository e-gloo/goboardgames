package game

import (
	"fmt"

	"github.com/e-gloo/goboardgames/utils"
)

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

func NewFleet() *[]Ship {
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
	return &ships
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

func debug(cells []uint8) {
	fmt.Println("|-------------------|")
	for i := uint8(0); i < MAP_HEIGHT; i++ {
		fmt.Print("|")
		for j := uint8(0); j < MAP_WIDTH; j++ {
			if contains(cells, i*MAP_HEIGHT+j) {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("|")
		}
		fmt.Println()
	}
	fmt.Println("|-------------------|")
}

func debugMap(cells map[uint8]uint8) {
	arrCells := []uint8{}
	for cell := range cells {
		arrCells = append(arrCells, cell)
	}
	debug(arrCells)
}

func contains(cells []uint8, val uint8) bool {
	for _, a := range cells {
		if a == val {
			return true
		}
	}
	return false
}

// 1. éliminer toutes les cellules connues "mortes"
// 2. éliminer toutes les cellules où on peut pas se mettre à cause d'autres bateaux
// 3. faire un random sur le reste des cellules pour savoir où on va mettre le bato
