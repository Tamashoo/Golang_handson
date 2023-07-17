package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Bottom", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)
	ct := widget.NewLabel("Center.")

	layout := container.NewBorder(bt, bb, bl, br, ct)
	w.SetContent(
		layout,
	)
	w.ShowAndRun()
}
