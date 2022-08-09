package boardView

import (
	"fmt"
	"main/consoleView/cellView"
	"main/interfaces/marker"
) 

type BoardView struct {
  board [][]marker.Marker;
}

func New() BoardView {
  return *(new(BoardView))
}

func (bv *BoardView) InitBoard(board [][]marker.Marker) {
  bv.board = board;
}

func (bv *BoardView) DrawBoard() {
  for i := range bv.board {
    for j := range bv.board[i] {
      fmt.Print(cellView.String(bv.board[i][j]));
    }
    fmt.Println();
  }
}

func (bv *BoardView) OnBoardChanged(x int, y int, m marker.Marker) {
  bv.board[x][y] = m;
  bv.DrawBoard()
}