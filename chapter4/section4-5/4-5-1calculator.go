package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type cdata struct {
	mem int
	cal string
	flg bool
}

func createNumButtons(f func(v int)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(3),
		widget.NewButton("7", func() { f(7) }),
		widget.NewButton("8", func() { f(8) }),
		widget.NewButton("9", func() { f(9) }),
		widget.NewButton("4", func() { f(4) }),
		widget.NewButton("5", func() { f(5) }),
		widget.NewButton("6", func() { f(6) }),
		widget.NewButton("1", func() { f(1) }),
		widget.NewButton("2", func() { f(2) }),
		widget.NewButton("3", func() { f(3) }),
		widget.NewButton("0", func() { f(0) }),
	)
	return c
}

func createCalcButtons(f func(c string)) *fyne.Container {
	c := fyne.NewContainerWithLayout(
		layout.NewGridLayout(1),
		widget.NewButton("CL", func() { f("CL") }),
		widget.NewButton("/", func() { f("/") }),
		widget.NewButton("*", func() { f("*") }),
		widget.NewButton("+", func() { f("+") }),
		widget.NewButton("-", func() { f("-") }),
	)
	return c
}

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	w.SetFixedSize(true)
	l := widget.NewLabel("0")
	l.Alignment = fyne.TextAlignTrailing

	data := cdata{
		mem: 0,
		cal: "",
		flg: false,
	}

	calc := func(n int) {
		switch data.cal {
		case "":
			data.mem = n
		case "+":
			data.mem += n
		case "-":
			data.mem -= n
		case "*":
			data.mem *= n
		case "/":
			data.mem /= n
		}
		l.SetText(strconv.Itoa(data.mem))
		data.flg = true
	}

	pushNum := func(v int) {
		s := l.Text
		if data.flg {
			s = "0"
			data.flg = false
		}
		s += strconv.Itoa(v)
		n, err := strconv.Atoi(s)
		if err == nil {
			l.SetText(strconv.Itoa(n))
		}
	}

	pushCalc := func(c string) {
		if c == "CL" {
			l.SetText("0")
			data.mem = 0
			data.cal = ""
			data.flg = false
			return
		}
		n, err := strconv.Atoi(l.Text)
		if err != nil {
			return
		}
		calc(n)
		data.cal = c
	}

	pushEnter := func() {
		n, err := strconv.Atoi(l.Text)
		if err != nil {
			return
		}
		calc(n)
		data.cal = ""
	}

	k := createNumButtons(pushNum)
	c := createCalcButtons(pushCalc)
	e := widget.NewButton("Enter", pushEnter)

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				l, e, nil, c,
			),
			l, e, k, c,
		),
	)
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
