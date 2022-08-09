package interfaces

import (
	"main/interfaces/state"
)

type Model interface {
	Turn(player int, x int, y int) (state.State, error)
}
