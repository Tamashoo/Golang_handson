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
	radio := widget.NewRadioGroup(
		[]string{"One", "Two", "Three"},
		func(selected string) {
			if selected == "" {
				l.SetText("Nothing selected!")
			} else {
				l.SetText("Selected: " + selected)
			}
		},
	)

	radio.SetSelected("One")

	w.SetContent(
		container.NewVBox(
			l, radio,
		),
	)
	w.ShowAndRun()
}
