package main

import (
	"click-helper/global"
	"click-helper/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Click Helper V1.2.0")
	global.MainWindow = myWindow
	myWindow.SetMainMenu(view.MainMenu(myApp, myWindow))
	myWindow.SetContent(view.MainWindows(myWindow))

	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.ShowAndRun()
}
