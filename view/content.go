package view

import (
	"click-helper/controller"
	"click-helper/global"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func SettleForm(win fyne.Window) fyne.CanvasObject {
	settleBox := container.NewVBox()

	// Create left column for Position
	posLabel := widget.NewLabel("Position:")
	posInput := widget.NewEntryWithData(global.BindPosData)
	posSelectButton := widget.NewButtonWithIcon("select", theme.ContentAddIcon(), controller.PosSelectHandler)
	positionCol := container.NewVBox(
		posLabel,
		posInput,
		posSelectButton,
	)

	// Create right column for Random Offset
	helpIcon := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		dialog.NewInformation("Random Offset Help",
			"This value determines the random pixel offset range.\n"+
				"For example, if set to 50, the click position will randomly\n"+
				"shift between -50 to +50 pixels on both X and Y axes.\n"+
				"This helps to make clicks look more natural.",
			win).Show()
	})
	helpIcon.Importance = widget.LowImportance

	randomOffsetLabelBox := container.NewHBox(
		widget.NewLabel("Offset:"),
		helpIcon,
	)

	randomOffsetEntry := widget.NewEntry()
	randomOffsetEntry.SetPlaceHolder("0-100 pixels")
	randomOffsetEntry.OnChanged = func(value string) {
		if value == "" {
			global.RandomOffset = 0
			return
		}
		offset, err := strconv.Atoi(value)
		if err != nil || offset < 0 {
			global.ShowErr("Please enter a valid positive number")
			return
		}
		global.RandomOffset = offset
	}

	offsetCol := container.NewVBox(
		randomOffsetLabelBox,
		randomOffsetEntry,
	)

	// Create grid layout with two equal columns
	topRow := container.NewGridWithColumns(2, positionCol, offsetCol)
	settleBox.Add(topRow)

	// Interval settings
	intervalLabel := widget.NewLabel("Interval(MS):")
	intervalInput := widget.NewEntry()
	intervalInput.SetText("200")
	global.IntervalMs = 200
	intervalInput.OnChanged = func(s string) {
		intData, _ := strconv.Atoi(s)
		global.IntervalMs = intData
	}
	intervalCol := container.NewVBox(
		intervalLabel,
		intervalInput,
	)

	// Random Interval settings with help icon
	randomIntervalHelpIcon := widget.NewButtonWithIcon("", theme.InfoIcon(), func() {
		dialog.NewInformation("Random Interval Help",
			"This value adds additional random delay to the base interval.\n"+
				"For example, if base interval is 200ms and random interval is 100ms,\n"+
				"the actual delay between clicks will be between 200ms to 300ms.\n"+
				"This helps to make clicking pattern more natural.",
			win).Show()
	})
	randomIntervalHelpIcon.Importance = widget.LowImportance

	randomIntervalLabelBox := container.NewHBox(
		widget.NewLabel("Random:"),
		randomIntervalHelpIcon,
	)

	randomIntervalInput := widget.NewEntry()
	randomIntervalInput.SetPlaceHolder("Additional delay")
	randomIntervalInput.OnChanged = func(s string) {
		if s == "" {
			global.RandomIntervalMs = 0
			return
		}
		intData, err := strconv.Atoi(s)
		if err != nil || intData < 0 {
			global.ShowErr("Please enter a valid positive number")
			return
		}
		global.RandomIntervalMs = intData
	}
	randomIntervalCol := container.NewVBox(
		randomIntervalLabelBox, // 使用包含帮助图标的标签容器
		randomIntervalInput,
	)

	// Create grid layout for interval settings
	intervalRow := container.NewGridWithColumns(2, intervalCol, randomIntervalCol)
	settleBox.Add(intervalRow)

	// Total times settings
	timesLabel := widget.NewLabel("Total Times:")
	timesInput := widget.NewEntry()
	timesInput.OnChanged = func(s string) {
		intData, _ := strconv.Atoi(s)
		global.ClickTotal = intData
	}

	settleBox.Add(timesLabel)
	settleBox.Add(timesInput)

	return settleBox
}

func StartButton() fyne.CanvasObject {
	buttonBox := container.NewHBox()

	// Start 按钮
	startBtn := widget.NewButton("Start", controller.StartHandler)

	// Cancel 按钮，更新快捷键提示
	cancelBtn := widget.NewButton("Cancel (Ctrl+Q)", func() {
		global.IsRunning = false
	})

	buttonBox.Add(startBtn)
	buttonBox.Add(cancelBtn)

	return container.NewCenter(buttonBox)
}

func ProgressBar() fyne.CanvasObject {
	box := container.NewVBox()
	progress := widget.NewProgressBarWithData(global.ProcessBarData)
	box.Add(container.NewVBox())
	box.Add(progress)
	return progress
}
