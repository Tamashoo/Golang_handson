package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	sl := widget.NewSelect(
		[]string{"Eins", "Zwei", "Drei"},
		func(selected string) {
			l.SetText("Selected: " + selected)
		},
	)
	w.SetContent(
		container.NewVBox(
			l, sl,
		),
	)
	w.ShowAndRun()
}
