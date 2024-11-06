package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func init() {
	app.SetMetadata(fyne.AppMetadata{
		ID: "com.clickhelper.app",
		Name: "click-helper.exe",
		Version: "0.0.1",
		Build: 1,
		Icon: &fyne.StaticResource{
	StaticName: "logo.png",
	StaticContent: []byte{
		
		Release: true,
		Custom: map[string]string{
			
		},
		
	})
}
