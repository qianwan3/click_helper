package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

var ProcessBarData = binding.NewFloat()
var BindPosData = binding.NewString()
var ClickTimes int
var ClickTotal int
var IntervalMs int
var RandomIntervalMs = 0 // 随机增加的间隔时间
var MainWindow fyne.Window
var RandomOffset = 0  // 随机偏移量
var IsRunning = false // 控制点击循环的标志
