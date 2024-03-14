package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie-Tracker")
	myWindow.Resize(fyne.NewSize(500, 650))

	var showHomePage func()
	showArtistsPage := func() {
		artistsPageContent := widget.NewLabel("Page des artistes:")
		myWindow.SetContent(container.NewVBox(
			artistsPageContent,
			widget.NewButton("Back", func() {
				showHomePage()
			}),
		))
	}
	showHomePage = func() {
		myWindow.SetContent(container.NewVBox(
			widget.NewLabel("Hello!"),
			widget.NewButton("Artists", func() {
				showArtistsPage()
			}),
			widget.NewButton("Quit", func() {
				myApp.Quit()
			}),
		))
	}
	showHomePage()

	myWindow.ShowAndRun()
}
