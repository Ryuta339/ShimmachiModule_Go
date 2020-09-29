package railroad

import (
	"./direction"
)

type ChangeableSpeed interface {
	ChangeSpeed(int)
}

type ChangeableDirection interface {
	ChangeDirection(direction.DirectionState)
}
