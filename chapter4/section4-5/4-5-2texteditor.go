package main

import (
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calc")
	edit := widget.NewEntry()
	edit.MultiLine = true
	sc := container.NewScroll(edit)
	inf := widget.NewLabel("infomation bar.")

	nf := func() {
		dialog.ShowConfirm("Alert", "Create new document?", func(f bool) {
			if f {
				edit.SetText("")
				inf.SetText("create new document.")
			}
		}, w)
	}

	of := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Open file name.", "OK",
			"Cancel", f, func(b bool) {
				if b {
					fn := f.Text + ".txt"
					ba, er := ioutil.ReadFile(fn)
					if er != nil {
						dialog.ShowError(er, w)
					} else {
						edit.SetText(string(ba))
						inf.SetText("open file '" + fn + "'")
					}
				}
			}, w)
	}

	sf := func() {
		f := widget.NewEntry()
		dialog.ShowCustomConfirm("Save file name.", "OK",
			"Cancel", f, func(b bool) {
				if b {
					fn := f.Text + ".txt"
					er := ioutil.WriteFile(fn, []byte(edit.Text), os.ModePerm)
					if er != nil {
						dialog.ShowError(er, w)
						return
					}
					inf.SetText("save file '" + fn + "'")
				}
			}, w)
	}

	qf := func() {
		dialog.ShowConfirm("Alert", "Quit this application?", func(f bool) {
			if f {
				a.Quit()
			}
		}, w)
	}

	tf := true

	cf := func() {
		if tf {
			a.Settings().SetTheme(theme.DarkTheme())
			inf.SetText("change theme to dark.")
		} else {
			a.Settings().SetTheme(theme.LightTheme())
			inf.SetText("change theme to light.")
		}
		tf = !tf
	}

	createMenubar := func() *fyne.MainMenu {
		return fyne.NewMainMenu(
			fyne.NewMenu("File",
				fyne.NewMenuItem("New", func() { nf() }),
				fyne.NewMenuItem("Open", func() { of() }),
				fyne.NewMenuItem("Save", func() { sf() }),
				fyne.NewMenuItem("Change Theme", func() { cf() }),
				fyne.NewMenuItem("Quit", func() { qf() }),
			),
			fyne.NewMenu("Edit",
				fyne.NewMenuItem("Cut", func() {
					edit.TypedShortcut(&fyne.ShortcutCut{
						Clipboard: w.Clipboard(),
					})
					inf.SetText("Cut text.")
				}),
				fyne.NewMenuItem("Copy", func() {
					edit.TypedShortcut(&fyne.ShortcutCopy{
						Clipboard: w.Clipboard(),
					})
					inf.SetText("Copy text.")
				}),
				fyne.NewMenuItem("Paste", func() {
					edit.TypedShortcut(&fyne.ShortcutPaste{
						Clipboard: w.Clipboard(),
					})
					inf.SetText("Paste text.")
				}),
			),
		)
	}

	createToolbar := func() *widget.Toolbar {
		return widget.NewToolbar(
			widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { nf() }),
			widget.NewToolbarAction(theme.FolderOpenIcon(), func() { of() }),
			widget.NewToolbarAction(theme.DocumentSaveIcon(), func() { sf() }),
		)
	}

	mb := createMenubar()
	tb := createToolbar()

	w.SetMainMenu(mb)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(tb, inf, nil, nil),
			tb, inf, sc,
		),
	)
	w.Resize(fyne.NewSize(500, 500))
	w.ShowAndRun()
}
