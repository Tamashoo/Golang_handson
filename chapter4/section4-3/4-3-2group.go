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
	//ck1 := widget.NewCheck("Check 1", nil)
	//ck2 := widget.NewCheck("Check 2", nil)
	cg := widget.NewCheckGroup([]string{"Check 1", "Check 2"}, func(selected []string) {})
	w.SetContent(
		container.NewVBox(
			l,
			cg,
			widget.NewButton("OK", func() {
				re := "result: "
				len := len(cg.Selected)
				for i, v := range cg.Selected {
					re += v
					if i != len-1 {
						re += ", "
					}
				}
				l.SetText(re)
			}),
		),
	)
	w.ShowAndRun()
}
