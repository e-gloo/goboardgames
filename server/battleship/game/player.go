package game

type PlayerFleet struct {
	Ready bool
	Fleet *[]Ship
	Hits  []uint8
}
