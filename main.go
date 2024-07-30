package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Val1, Val2, Score  int
	Action, UserResult string
	IsError            bool
}

type Solver struct {
	State
	Input                 *widget.Entry
	ScoreDisplay, Display *widget.Label
	AppError              string
}

func main() {
	solverApp := app.New()
	solver := Solver{
		Input:        widget.NewEntry(),
		ScoreDisplay: widget.NewLabel("Score:"),
		Display:      widget.NewLabel(""),
		AppError:     "incorrect data, try again. please.",
	}
	window := solverApp.NewWindow("Solver")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter result")

	btnExit := widget.NewButton("Exit", func() {
		solverApp.Quit()
	})

	values := getRandomValues()
	solver.Val1 = values[0]
	solver.Val2 = values[1]
	var strBuilder strings.Builder
	strBuilder.WriteString(strconv.Itoa(solver.Val1))
	strBuilder.WriteString("+")
	strBuilder.WriteString(strconv.Itoa(solver.Val2))
	newVal := strBuilder.String()
	strBuilder.Reset()
	solver.Display.SetText(newVal)
	solver.ScoreDisplay.SetText(strconv.Itoa(solver.Score))

	window.SetContent(
		container.NewGridWithColumns(
			1,
			container.NewGridWithColumns(2,
				solver.ScoreDisplay,
				solver.Display,
			),
			solver.Input,
			solver.enterBtnHandler(),
			btnExit,
		),
	)
	window.Resize(fyne.NewSize(300, 200))
	window.ShowAndRun()
}

func getRandomValues() []int {
	resp := make([]int, 2)
	resp[0] = rand.Intn(11)
	resp[1] = rand.Intn(11)
	return resp
}

func (solver *Solver) enterBtnHandler() *widget.Button {
	return widget.NewButton("Enter", func() {
		fmt.Println(solver.Input.Text)
		solver.UserResult = solver.Input.Text
		res := solver.Val1 + solver.Val2
		userRes, err := strconv.Atoi(solver.UserResult)
		if err == nil {
			if res == userRes {
				solver.Display.SetText("Right!")
				solver.Score = solver.Score + 1
				solver.ScoreDisplay.SetText(strconv.Itoa(solver.Score))
			} else {
				solver.Display.SetText("Wrong!")
			}
		} else {
			solver.Display.SetText(err.Error())
		}
	})
}
