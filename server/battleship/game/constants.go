package game

var MAP_WIDTH = uint8(10)
var MAP_HEIGHT = uint8(10)
var SHIPS_LENGTH = [5]uint8{5, 4, 3, 3, 2}

type Constants struct {
	MapWidth  uint8
	MapHeight uint8
}

var GameConstants = &Constants{MapWidth: MAP_WIDTH, MapHeight: MAP_HEIGHT}
