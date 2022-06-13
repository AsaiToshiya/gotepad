package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(1366, 768))

	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Exit", func() { a.Quit() }),
	)
	mainMenu := fyne.NewMainMenu(
		fileMenu,
	)
	w.SetMainMenu(mainMenu)

	input := widget.NewMultiLineEntry()

	w.SetContent(input)
	w.ShowAndRun()
}
