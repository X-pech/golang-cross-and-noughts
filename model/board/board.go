package board

import (
	"main/errors/coordinateError"
	"main/errors/occupiedError"
	"main/interfaces/marker"
	"main/model/cell"
	"main/model/player"
)

type Board struct {
	n    int
	Data [][]cell.Cell
}

func New(n int) Board {
	result := new(Board)
	result.n = n
	result.Data = make([][]cell.Cell, n)
	for i := range result.Data {
		result.Data[i] = make([]cell.Cell, n)
	}

	return *result
}

func (b *Board) IsCorrectCoordinate(x int, y int) bool {
	return x >= 0 && x < b.n && y >= 0 && y < b.n
}

func (b Board) IsEmptyCell(x int, y int) bool {
	return b.Data[x][y].Empty()
}

func (b *Board) ErrorIfCantMakeTurn(x int, y int) error {

	if !b.IsCorrectCoordinate(x, y) {
		var err = coordinateError.New(x, y)
		return &err
	}

	if !b.IsEmptyCell(x, y) {
		var err = occupiedError.New(x, y)
		return &err
	}

	return nil
}

func (b *Board) SetPlayer(x int, y int, p *player.Player) error {
	var err = b.ErrorIfCantMakeTurn(x, y)

	if err == nil {
		b.Data[x][y].SetPlayer(p)
		return nil
	}

	return err
}

func (b *Board) GetMarker(x int, y int) marker.Marker {
	return b.Data[x][y].GetMarker()
}

func (b Board) GetMarkers() [][]marker.Marker {
	var result = make([][]marker.Marker, b.n)
	for i := range result {
		result[i] = make([]marker.Marker, b.n)
		for j := range result[i] {
			result[i][j] = b.GetMarker(i, j)
		}
	}

	return result
}

func (b *Board) EmptyRemain() int {
	var c int = 0
	for i := range b.Data {
		for j := range b.Data[i] {
			if b.IsEmptyCell(i, j) {
				c++
			}
		}
	}

	return c
}

func (b *Board) checkLine(xb int, yb int, xs int, ys int, ch chan<- bool) {
	var x = xb + xs
	var y = yb + ys

	if b.GetMarker(xb, yb) == marker.EMPTY {
		ch <- false
		return
	}

	for {
		if !b.IsCorrectCoordinate(x, y) {
			break
		}

		if b.GetMarker(x, y) != b.GetMarker(xb, yb) {
			ch <- false
			return

		}

		x += xs
		y += ys
	}

	ch <- true
}

func (b *Board) IsOnFullLine(x int, y int) bool {
	results := make(chan bool, 4)

	go b.checkLine(x, 0, 0, 1, results)
	go b.checkLine(0, y, 1, 0, results)
	go b.checkLine(0, 0, 1, 1, results)
	go b.checkLine(0, b.n-1, 1, -1, results)

	return (<-results || <-results || <-results || <-results)
}
