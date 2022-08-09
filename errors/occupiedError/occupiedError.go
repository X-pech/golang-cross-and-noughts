package occupiedError

import "fmt"

type OccupiedError struct {
	x int
	y int
}

func New(x int, y int) OccupiedError {
	return OccupiedError{
		x: x, y: y,
	}
}

func (err *OccupiedError) Error() string {
	return fmt.Sprintf("Cell %d %d is alreadu occupied", err.x, err.y)
}
