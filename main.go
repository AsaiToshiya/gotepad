package main

import (
	"io"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var w fyne.Window
var input = widget.NewMultiLineEntry()

func createFileMenu() *fyne.Menu {
	newMenuItem := fyne.NewMenuItem("New", func() { input.SetText("") })
	openMenuItem := fyne.NewMenuItem("Open...", func() {
		fd := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				if r != nil {
					b, _ := io.ReadAll(r)
					input.SetText(string(b))
				}
			}, w)
		fd.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		fd.Show()
	})
	fileMenu := fyne.NewMenu("File",
		newMenuItem,
		openMenuItem,
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
	w = a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(1366, 768))

	w.SetMainMenu(createMenus())

	w.SetContent(input)
	w.ShowAndRun()
}
