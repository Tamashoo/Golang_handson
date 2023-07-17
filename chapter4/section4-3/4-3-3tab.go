package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	t1 := container.NewTabItem("Tab 1", widget.NewLabel("This is tab 1"))
	t2 := container.NewTabItem("Tab 2", widget.NewLabel("This is tab 2"))
	w.SetContent(
		container.NewAppTabs(
			t1, t2,
		),
	)
	w.ShowAndRun()
}
