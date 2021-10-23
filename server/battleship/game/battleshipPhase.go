package game

type BattleshipPhase uint8;

const (
	WAITING BattleshipPhase = iota
	PREPARATION
	PLAYING
	OVER
)