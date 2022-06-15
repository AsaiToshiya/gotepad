package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var input = widget.NewMultiLineEntry()

func createFileMenu() *fyne.Menu {
	newMenuItem := fyne.NewMenuItem("New", func() { input.SetText("") })
	fileMenu := fyne.NewMenu("File",
		newMenuItem,
	)
	return fileMenu
}

func createMenus() *fyne.MainMenu {
	mainMenu := fyne.NewMainMenu(
		createFileMenu(),
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
