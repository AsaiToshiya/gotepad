package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var input = widget.NewMultiLineEntry()

func createMenus() *fyne.MainMenu {
	newMenuItem := fyne.NewMenuItem("New", func() { input.SetText("") })
	fileMenu := fyne.NewMenu("File",
		newMenuItem,
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
	)
	return mainMenu
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(1366, 768))

	w.SetMainMenu(createMenus())

	w.SetContent(input)
	w.ShowAndRun()
}
