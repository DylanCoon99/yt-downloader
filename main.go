
package main

import (
	"log"
	"image/color"
	"github.com/DylanCoon99/yt-downloader/downloader"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)




func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")


	content := container.NewVBox(input, widget.NewButton("Download video", func() {
		log.Println("Your url was:", input.Text)
		msg := downloader.Download(input.Text)
			retText := canvas.NewText(msg, color.White)
			myWindow.SetContent(retText)
	}))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()

	

}
