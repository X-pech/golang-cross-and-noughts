package player

import "main/interfaces/marker"

type Player struct {
	Marker marker.Marker
	Name   string
}

func New(marker marker.Marker, name string) Player {
	p := new(Player)
	p.Marker = marker
	p.Name = name
	return *p
}
