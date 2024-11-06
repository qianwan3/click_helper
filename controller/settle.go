package controller

import (
	"click-helper/global"
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"log"
)

func PosSelectHandler() {
	global.MainWindow.Hide()
	hook.Register(hook.MouseDown, []string{}, func(event hook.Event) {
		x, y := robotgo.Location()
		global.BindPosData.Set(fmt.Sprintf("%d,%d", x, y))
		log.Println(global.BindPosData.Get())
		global.MainWindow.Show()
		hook.End()
	})
	s := hook.Start()
	<-hook.Process(s)
}
