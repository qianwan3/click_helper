package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func MainMenu(myApp fyne.App, myWin fyne.Window) *fyne.MainMenu {
	return fyne.NewMainMenu(AboutMenu(myWin))
}

func AboutMenu(win fyne.Window) *fyne.Menu {
	return fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			info := dialog.NewInformation("About", "Version 1.0", win)
			info.Resize(fyne.NewSize(200, 200))
			info.Show()
		}),
	)
}
