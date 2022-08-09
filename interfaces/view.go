package interfaces

import "main/interfaces/marker"

type View interface {
	InitBoard(b [][]marker.Marker)
	OnBoardChanged(x int, y int, m marker.Marker)
}
