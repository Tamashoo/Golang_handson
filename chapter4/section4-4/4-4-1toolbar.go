package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("This is a Sample widget.")
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			l.SetText("Select Home Icon")
		}),
		widget.NewToolbarAction(
			theme.InfoIcon(),
			func() {
				l.SetText("Select Info Icon")
			}),
	)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, tb, nil, nil,
			),
			l,
			tb,
		),
	)
	w.ShowAndRun()
}
