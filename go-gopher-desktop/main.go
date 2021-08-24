package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const KuteGoAPIURL = "https://kutego-api-xxxxx-ew.a.run.app"

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Gopher")

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Aurélie Vache"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	// Define a welcome text centered
	text := canvas.NewText("Display a random Gopher!", color.White)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, _ = fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0

	// Define a "random" button
	randomBtn := widget.NewButton("Random", func() {
		resource, _ := fyne.LoadResourceFromURLString(KuteGoAPIURL + "/gopher/random/")
		gopherImg.Resource = resource

		//Redrawn the image with the new path
		gopherImg.Refresh()
	})
	randomBtn.Importance = widget.HighImportance

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		randomBtn,
	)

	// Display our content
	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}
