package cellView

import (
	"main/interfaces/marker"
)

func String(m marker.Marker) string {
	switch m {
	case marker.CROSS:
		return "X"
	case marker.NOUGHT:
		return "O"
	default:
		return "."
	}
}
