package cell

import (
	"main/interfaces/marker"
	"main/model/player"
)

type Cell struct {
	p *player.Player
}

func New() Cell {
	result := new(Cell)
	return *result
}

func (c Cell) Empty() bool {
	return c.p == nil
}

func (c *Cell) SetPlayer(p *player.Player) {
	c.p = p
}

func (c Cell) GetMarker() marker.Marker {
	if c.Empty() {
		return marker.EMPTY
	}
	return c.p.Marker
}
