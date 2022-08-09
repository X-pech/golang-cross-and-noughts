package state

type State int

const (
  NOT_STARTED = iota
  GAMEPLAY = iota
  WIN_PLAYER_ONE = iota
  WIN_PLAYER_TWO = iota
  TIE = iota
)