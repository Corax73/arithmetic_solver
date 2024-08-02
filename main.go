package main

import (
	"arithmetic_solver/customTheme"
	"image/color"
	"math/rand"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Val1, Val2, Score        int
	Action, UserResult, Lang string
	IsError                  bool
}

type Internationalization struct {
	DataByLang map[string]map[string]string
}

type Solver struct {
	State
	Internationalization
	Input                                *widget.Entry
	TextSize                             float32
	ExpDisplay, ResDisplay, ScoreDisplay *canvas.Text
	BtnEnter                             *widget.Button
	AppError                             string
	SolverTheme                          fyne.Theme
	LangToggler                          *widget.RadioGroup
}

func main() {
	solverApp := app.New()
	color := color.White
	appTheme := customTheme.NewCustomTheme()
	solver := Solver{
		Input:        widget.NewEntry(),
		TextSize:     32,
		ScoreDisplay: canvas.NewText("", color),
		ExpDisplay:   canvas.NewText("", color),
		ResDisplay:   canvas.NewText("", color),
		AppError:     "incorrect data, try again. please.",
		SolverTheme:  appTheme,
	}
	solver.Lang = "ru"
	solver.DataByLang = map[string]map[string]string{
		"ru": map[string]string{
			"ScoreDisplay":     "Баллы: ",
			"btnExit":          "Выход",
			"enterBtn":         "Ввод",
			"ResultRight":      "Правильно!",
			"ResultWrong":      "Ошибка!",
			"InputPlaceHolder": "Введите результат",
			"NewExp":           "Новый пример",
		},
		"en": map[string]string{
			"ScoreDisplay":     "Score: ",
			"btnExit":          "Exit",
			"enterBtn":         "Enter",
			"ResultRight":      "Right!",
			"ResultWrong":      "Wrong!",
			"InputPlaceHolder": "Enter result",
			"NewExp":           "New expression",
		},
	}
	solver.LangToggler = solver.langTogglerHandler()
	solver.ScoreDisplay.Text = solver.DataByLang[solver.Lang]["ScoreDisplay"]
	solverApp.Settings().SetTheme(solver.SolverTheme)
	solver.ScoreDisplay.TextSize, solver.ExpDisplay.TextSize, solver.ResDisplay.TextSize = solver.TextSize, solver.TextSize, solver.TextSize
	solver.BtnEnter = solver.enterBtnHandler()
	window := solverApp.NewWindow("Solver")

	btnExit := widget.NewButton(solver.DataByLang[solver.Lang]["btnExit"], func() {
		solverApp.Quit()
	})

	content := container.NewGridWithColumns(
		1,
		container.NewGridWithColumns(
			3,
			solver.ScoreDisplay,
			solver.ExpDisplay,
			solver.ResDisplay,
		),
		solver.LangToggler,
		solver.Input,
		solver.BtnEnter,
		solver.newBtnHandler(),
		btnExit,
	)

	solver.newExpression()
	window.SetContent(content)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(500, 400))
	window.ShowAndRun()
}

func (solver *Solver) getRandomValues() {
	actions := []string{" + ", " - "}
	solver.Val1 = rand.Intn(11)
	val2 := rand.Intn(11)
	if solver.Val1 > val2 {
		solver.Val2 = val2
	} else if solver.Val1 < val2 {
		valTemp := solver.Val1
		solver.Val1 = val2
		solver.Val2 = valTemp
	} else {
		solver.Val2 = val2
	}
	solver.Action = actions[rand.Intn(2)]
}

func (solver *Solver) enterBtnHandler() *widget.Button {
	return widget.NewButton(solver.DataByLang[solver.Lang]["enterBtn"], func() {
		solver.UserResult = solver.Input.Text
		var res int
		if solver.Action == " + " {
			res = solver.Val1 + solver.Val2
		} else {
			res = solver.Val1 - solver.Val2
		}
		userRes, err := strconv.Atoi(solver.UserResult)
		if err == nil {
			if res == userRes {
				solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["ResultRight"]
				solver.Score = solver.Score + 1
				solver.setScoreVal()
				solver.btnDisable(solver.BtnEnter)
				solver.ResDisplay.Refresh()
			} else {
				solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["ResultWrong"]
				solver.ResDisplay.Refresh()
			}
		} else {
			solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["ResultWrong"]
			solver.newExpression()
		}
	})
}

func (solver *Solver) btnDisable(btn *widget.Button) {
	btn.Disable()
}

func (solver *Solver) btnEnable(btn *widget.Button) {
	btn.Enable()
}

func (solver *Solver) newExpression() {
	solver.getRandomValues()
	solver.Input.SetPlaceHolder(solver.DataByLang[solver.Lang]["InputPlaceHolder"])
	solver.Input.SetText("")
	var strBuilder strings.Builder
	strBuilder.WriteString(strconv.Itoa(solver.Val1))
	strBuilder.WriteString(solver.Action)
	strBuilder.WriteString(strconv.Itoa(solver.Val2))
	solver.ExpDisplay.Text = strBuilder.String()
	solver.ExpDisplay.Refresh()
	solver.ResDisplay.Text = ""
	solver.ResDisplay.Refresh()
	solver.btnEnable(solver.BtnEnter)
	solver.setScoreVal()
}

func (solver *Solver) setScoreVal() {
	var strBuilder strings.Builder
	strBuilder.WriteString(solver.DataByLang[solver.Lang]["ScoreDisplay"])
	strBuilder.WriteString(strconv.Itoa(solver.Score))
	solver.ScoreDisplay.Text = strBuilder.String()
	solver.ScoreDisplay.Refresh()
}

func (solver *Solver) newBtnHandler() *widget.Button {
	return widget.NewButton(solver.DataByLang[solver.Lang]["NewExp"], func() {
		solver.newExpression()
	})
}

func (solver *Solver) langTogglerHandler() *widget.RadioGroup {
	return widget.NewRadioGroup([]string{"ru", "en"}, func(value string) {
		solver.Lang = value
		solver.refreshAllCanvas()
	})
}

func (solver *Solver) refreshAllCanvas() {
	var strBuilder strings.Builder
	strBuilder.WriteString(solver.DataByLang[solver.Lang]["ScoreDisplay"])
	strBuilder.WriteString(strconv.Itoa(solver.Score))
	solver.ScoreDisplay.Text = strBuilder.String()
	solver.ExpDisplay.Refresh()
	solver.ResDisplay.Refresh()
	solver.ScoreDisplay.Refresh()
}
