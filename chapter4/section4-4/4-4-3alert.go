package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	b := widget.NewButton("Click", func() {
		dialog.ShowInformation("Alert",
			"This is a sample alert dialog.", w)
	})
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				nil, b, nil, nil,
			),
			l, b,
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
