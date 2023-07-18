package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type MyEntry struct {
	widget.Entry
	entered func(e *MyEntry)
}

func NewMyEntry(f func(e *MyEntry)) *MyEntry {
	e := &MyEntry{}
	e.ExtendBaseWidget(e)
	e.entered = f
	return e
}

func (e *MyEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyReturn, fyne.KeyEnter:
		e.entered(e)
	default:
		e.Entry.KeyDown(key)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	e := NewMyEntry(func(e *MyEntry) {
		s := e.Text
		e.SetText("")
		l.SetText("you type '" + s + "'")
	})

	w.SetContent(
		container.NewVBox(
			l, e,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
