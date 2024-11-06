package controller

import (
	"click-helper/global"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func StartHandler() {
	if global.IntervalMs == 0 {
		global.ShowErr("Interval number must be greater than zero")
		return
	}
	if global.ClickTotal == 0 {
		global.ShowErr("Total Times must be greater than zero")
		return
	}

	posStr, err := global.BindPosData.Get()
	if err != nil {
		global.ShowErr(err.Error())
		return
	}
	if posStr == "" {
		global.ShowErr("The coordinates to be clicked must be set first")
		return
	}

	global.ClickTimes = 0
	global.ProcessBarData.Set(0.00)
	global.IsRunning = true

	go func() {
		hook.Register(hook.KeyDown, []string{"q", "ctrl"}, func(e hook.Event) {
			global.IsRunning = false
		})
		s := hook.Start()
		<-hook.Process(s)
	}()

	posSlice := strings.Split(posStr, ",")
	if len(posSlice) != 2 {
		global.ShowErr("The Position data is error!")
		return
	}
	x, _ := strconv.Atoi(posSlice[0])
	y, _ := strconv.Atoi(posSlice[1])

	for i := 0; i < global.ClickTotal; i++ {
		if !global.IsRunning {
			break
		}

		offsetX := 0
		offsetY := 0
		if global.RandomOffset > 0 {
			offsetX = rand.Intn(global.RandomOffset*2) - global.RandomOffset
			offsetY = rand.Intn(global.RandomOffset*2) - global.RandomOffset
		}

		robotgo.Move(x+offsetX, y+offsetY)
		robotgo.Click("left", false)
		global.ClickTimes++
		rate := global.ClickTimes * 100 / global.ClickTotal
		global.ProcessBarData.Set(float64(rate) / float64(100))

		randomDelay := 0
		if global.RandomIntervalMs > 0 {
			randomDelay = rand.Intn(global.RandomIntervalMs)
		}
		time.Sleep(time.Duration(global.IntervalMs+randomDelay) * time.Millisecond)
	}

	hook.End()
	global.IsRunning = false
	global.MainWindow.Show()
}
