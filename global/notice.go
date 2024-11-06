package global

import "fyne.io/fyne/v2/dialog"

func ShowErr(msg string) {
	dialog.ShowInformation("Error", msg, MainWindow)
}
