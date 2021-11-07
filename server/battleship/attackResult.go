package battleship

type AttackResultEnum uint8

const (
	MISS AttackResultEnum = iota
	HIT
	SUNK
)

type AttackResult struct {
	Position uint8
	PlayerNb uint8
	Result   AttackResultEnum
}
