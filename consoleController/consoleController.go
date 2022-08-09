package consoleController

import (
	"errors"
	"fmt"
	"main/consoleView/boardView"
	"main/errors/coordinateError"
	"main/errors/occupiedError"
	"main/interfaces/state"
	"main/model"
)

type ConsoleController struct {
}

func Run() {
	fmt.Print("Board size: ")
	var side int
	fmt.Scanln(&side)

	var name [2]string
	fmt.Print("Player for Crosses shall be named ")
	fmt.Scanln(&name[0])
	fmt.Print("Player for Noughts shall be named ")
	fmt.Scanln(&name[1])

	bv := boardView.New()
	var model = model.New(side, name[0], name[1], &bv)
	bv.DrawBoard()

	for i := 0; ; i ^= 1 {

		var curstate state.State
		var err error
		for {
			fmt.Printf("Enter your turn, %s ", name[i])
			var x, y int
			fmt.Scanln(&x, &y)
			x--
			y--

			curstate, err = model.Turn(i, x, y)

			if err == nil {
				break
			}
			var incord *coordinateError.CoordinateError
			var occord *occupiedError.OccupiedError
			switch true {
			case errors.As(err, &incord):
				fmt.Println(incord.Error())
			case errors.As(err, &occord):
				fmt.Println(occord.Error())
			default:
				fmt.Println("Something went wrong. Please try again.")
			}
		}

		if curstate != state.GAMEPLAY {
			switch curstate {
			case state.WIN_PLAYER_ONE:
				fmt.Printf("Player %s wins!\n", name[0])
			case state.WIN_PLAYER_TWO:
				fmt.Printf("Player %s wins!\n", name[1])
			case state.TIE:
				fmt.Println("Friendship wins!")
			default:
				fmt.Println("If you see this message something went wrong")
			}
			break
		}
	}

}
