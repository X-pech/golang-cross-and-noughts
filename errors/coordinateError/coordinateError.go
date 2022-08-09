package coordinateError

import "fmt"

type CoordinateError struct {
	x int
	y int
}

func New(x int, y int) CoordinateError {
	var err = CoordinateError{
		x: x, y: y,
	}
	return err
}

func (err *CoordinateError) Error() string {
	return fmt.Sprintf("Incorrect coordinates: %d %d", err.x, err.y)
}
