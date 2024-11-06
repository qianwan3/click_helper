package view

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func MainWindows(win fyne.Window) *fyne.Container {
	box := container.NewVBox()
	box.Add(SettleForm(win))
	box.Add(widget.NewSeparator())
	box.Add(ProgressBar())
	box.Add(StartButton())
	box.Add(MainBottom())
	return box
}

func MainBottom() fyne.CanvasObject {
	//bottom := container.NewCenter()
	contactBox := container.NewVBox(layout.NewSpacer())

	titleLabel := widget.NewLabel("Contact:   (Get free software)")
	contactBox.Add(titleLabel)
	emailLabel := widget.NewLabel("Email: clickhelpersoft@gmail.com    ")
	contactBox.Add(emailLabel)
	telegramGroup := widget.NewLabel("TelegramGroup: @click_helper    ")
	uri, _ := url.ParseRequestURI("https://t.me/+vdhjH6i6Eg9lM2Vl")
	tgJoin := widget.NewHyperlink("JOIN", uri)
	tgLine := container.NewHBox(telegramGroup, tgJoin)
	contactBox.Add(tgLine)
	//bottom.Add(contactBox)
	return contactBox
}
