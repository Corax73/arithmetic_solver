package main

import (
	"arithmetic_solver/customTheme"
	"arithmetic_solver/randomizer"
	"image/color"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const SIMPLE_DIFFICULT_LEVEL = 2

type State struct {
	Val1, Val2, Val3, Score, Difficult int
	Action1, Action2, UserResult, Lang string
	IsError                            bool
}

type Internationalization struct {
	DataByLang map[string]map[string]string
}

type Solver struct {
	State
	Internationalization
	Input                                *widget.Entry
	ExpDisplay, ResDisplay, ScoreDisplay *canvas.Text
	BtnEnter, BtnNewExp, BtnExit         *widget.Button
	SelectDifficult                      *widget.Select
	SolverTheme                          fyne.Theme
	LangToggler                          *widget.RadioGroup
	TextSize                             float32
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
		SolverTheme:  appTheme,
	}
	solver.Lang = "ru"
	solver.DataByLang = map[string]map[string]string{
		"ru": map[string]string{
			"ScoreDisplay":     "Баллы: ",
			"BtnExit":          "Выход",
			"EnterBtn":         "Ввод",
			"ResultRight":      "Правильно!",
			"ResultWrong":      "Ошибка!",
			"InputPlaceHolder": "Введите результат",
			"NewExp":           "Новый пример",
			"Difficult":        "Выберите сложность",
			"AppError":         "Неверные данные",
		},
		"en": map[string]string{
			"ScoreDisplay":     "Score: ",
			"BtnExit":          "Exit",
			"EnterBtn":         "Enter",
			"ResultRight":      "Right!",
			"ResultWrong":      "Wrong!",
			"InputPlaceHolder": "Enter result",
			"NewExp":           "New expression",
			"Difficult":        "Select difficult",
			"AppError":         "Incorrect data",
		},
	}
	solver.LangToggler = solver.langTogglerHandler()
	solver.ScoreDisplay.Text = solver.DataByLang[solver.Lang]["ScoreDisplay"]
	solverApp.Settings().SetTheme(solver.SolverTheme)
	solver.ScoreDisplay.TextSize, solver.ExpDisplay.TextSize, solver.ResDisplay.TextSize = solver.TextSize, solver.TextSize, solver.TextSize
	solver.BtnEnter = solver.EnterBtnHandler()
	solver.SelectDifficult = solver.GetSelectDifficult()
	window := solverApp.NewWindow("Solver")

	solver.BtnNewExp = solver.newBtnHandler()
	solver.BtnExit = widget.NewButton(solver.DataByLang[solver.Lang]["BtnExit"], func() {
		solverApp.Quit()
	})

	content := container.NewGridWithColumns(
		1,
		container.NewGridWithColumns(
			4,
			solver.SelectDifficult,
			solver.ScoreDisplay,
			solver.ExpDisplay,
			solver.ResDisplay,
		),
		solver.LangToggler,
		solver.Input,
		solver.BtnEnter,
		solver.BtnNewExp,
		solver.BtnExit,
	)

	solver.newExpression()
	window.SetContent(content)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()
}

func (solver *Solver) EnterBtnHandler() *widget.Button {
	return widget.NewButton(solver.DataByLang[solver.Lang]["EnterBtn"], func() {
		solver.UserResult = solver.Input.Text
		var res int
		if solver.Action1 == " + " {
			res = solver.Val1 + solver.Val2
		} else if solver.Action1 == " - " {
			res = solver.Val1 - solver.Val2
		} else if solver.Action1 == " * " {
			res = solver.Val1 * solver.Val2
		} else if solver.Action1 == " / " {
			res = solver.Val1 / solver.Val2
		}
		if solver.Difficult > SIMPLE_DIFFICULT_LEVEL {
			if solver.Action2 == " + " {
				res = res + solver.Val3
			} else if solver.Action2 == " - " {
				res = res - solver.Val3
			} else if solver.Action2 == " * " {
				res = res * solver.Val3
			} else if solver.Action2 == " / " {
				res = res / solver.Val3
			}
		}
		userRes, err := strconv.Atoi(solver.UserResult)
		if err == nil {
			if res == userRes {
				solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["ResultRight"]
				accruedPoint := solver.Difficult
				if accruedPoint == 0 {
					accruedPoint += 1
				}
				solver.Score = solver.Score + accruedPoint
				solver.setScoreVal()
				solver.btnDisable(solver.BtnEnter)
				solver.ResDisplay.Refresh()
			} else {
				solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["ResultWrong"]
				solver.ResDisplay.Refresh()
			}
		} else {
			solver.IsError = true
			solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["AppError"]
			solver.ResDisplay.Refresh()
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
	solver.Val1, solver.Val2, solver.Val3, solver.Action1, solver.Action2 = randomizer.GetRandomValues(solver.Difficult)
	solver.Input.SetPlaceHolder(solver.DataByLang[solver.Lang]["InputPlaceHolder"])
	solver.Input.SetText("")
	var strBuilder strings.Builder
	strBuilder.WriteString(strconv.Itoa(solver.Val1))
	strBuilder.WriteString(solver.Action1)
	strBuilder.WriteString(strconv.Itoa(solver.Val2))
	if solver.Difficult > SIMPLE_DIFFICULT_LEVEL {
		strBuilder.WriteString(solver.Action2)
		strBuilder.WriteString(strconv.Itoa(solver.Val3))
	}
	solver.ExpDisplay.Text = strBuilder.String()
	strBuilder.Reset()
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
	strBuilder.Reset()
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
	strBuilder.Reset()
	solver.SelectDifficult.PlaceHolder = solver.DataByLang[solver.Lang]["Difficult"]
	solver.BtnEnter.Text = solver.DataByLang[solver.Lang]["EnterBtn"]
	solver.BtnExit.Text = solver.DataByLang[solver.Lang]["BtnExit"]
	solver.BtnNewExp.Text = solver.DataByLang[solver.Lang]["NewExp"]
	solver.Input.PlaceHolder = solver.DataByLang[solver.Lang]["InputPlaceHolder"]
	solver.BtnEnter.Refresh()
	solver.BtnExit.Refresh()
	solver.BtnNewExp.Refresh()
	solver.BtnNewExp.Refresh()
	solver.Input.Refresh()
	solver.ScoreDisplay.Refresh()
	solver.SelectDifficult.Refresh()
	if solver.IsError {
		solver.ResDisplay.Text = solver.DataByLang[solver.Lang]["AppError"]
		solver.ResDisplay.Refresh()
	}
}

func (solver *Solver) GetSelectDifficult() *widget.Select {
	resp := widget.NewSelect([]string{"1", "2", "3", "4"}, func(value string) {
		mappedDiff := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4}
		if _, ok := mappedDiff[value]; ok {
			solver.Difficult = mappedDiff[value]
		} else {
			solver.Difficult = mappedDiff["1"]
		}
	})
	resp.PlaceHolder = solver.DataByLang[solver.Lang]["Difficult"]
	return resp
}
