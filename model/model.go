package model

import (
	"main/interfaces"
	"main/interfaces/marker"
	"main/interfaces/state"
	"main/model/board"
	"main/model/player"
)

type Model struct {
	players      [2]player.Player
	board        *board.Board
	currentState state.State
	view         interfaces.View
}

func New(side int, nameOne string, nameTwo string, view interfaces.View) Model {
	var m *Model = new(Model)
	m.view = view
	b := board.New(side)
	m.board = &b
	m.currentState = state.NOT_STARTED

	m.view.InitBoard(m.board.GetMarkers())

	m.players[0] = player.New(marker.CROSS, nameOne)
	m.players[1] = player.New(marker.NOUGHT, nameTwo)
	return *m
}

func (model *Model) Turn(player int, x int, y int) (state.State, error) {

	var err = model.board.SetPlayer(x, y, &model.players[player])

	if err != nil {
		return state.NOT_STARTED, err
	}

	if model.board.IsOnFullLine(x, y) {
		if player == 0 {
			model.currentState = state.WIN_PLAYER_ONE
		} else {
			model.currentState = state.WIN_PLAYER_TWO
		}
	} else if model.board.EmptyRemain() == 0 {
		model.currentState = state.TIE
	} else {
		model.currentState = state.GAMEPLAY
	}

	model.view.OnBoardChanged(x, y, model.board.GetMarker(x, y))

	return model.currentState, nil
}
