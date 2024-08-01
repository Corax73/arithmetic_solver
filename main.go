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

	btnExit := widget.NewButton("Exit", func() {
		solverApp.Quit()
	})

	solver.newExpression()

	window.SetContent(
		container.NewGridWithColumns(
			1,
			container.NewGridWithColumns(2,
				solver.ScoreDisplay,
				solver.Display,
			),
			solver.Input,
			solver.enterBtnHandler(),
			solver.newBtnHandler(),
			btnExit,
		),
	)
	window.Resize(fyne.NewSize(500, 400))
	window.ShowAndRun()
}

func (solver *Solver) getRandomValues() {
	solver.Val1 = rand.Intn(11)
	solver.Val2 = rand.Intn(11)
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
				solver.setScoreVal()
			} else {
				solver.Display.SetText("Wrong!")
			}
		} else {
			solver.Display.SetText(err.Error())
		}
	})
}

func (solver *Solver) newExpression() {
	solver.getRandomValues()
	solver.Input.SetPlaceHolder("Enter result")
	solver.Input.SetText("")
	var strBuilder strings.Builder
	strBuilder.WriteString(strconv.Itoa(solver.Val1))
	strBuilder.WriteString("+")
	strBuilder.WriteString(strconv.Itoa(solver.Val2))
	solver.Display.SetText(strBuilder.String())
	solver.setScoreVal()
}

func (solver *Solver) setScoreVal() {
	var strBuilder strings.Builder
	strBuilder.WriteString("Score: ")
	strBuilder.WriteString(strconv.Itoa(solver.Score))
	solver.ScoreDisplay.SetText(strBuilder.String())
}

func (solver *Solver) newBtnHandler() *widget.Button {
	return widget.NewButton("New expression", func() {
		solver.newExpression()
	})
}
