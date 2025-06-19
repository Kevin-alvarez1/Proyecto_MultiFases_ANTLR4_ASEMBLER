package main

import (
	"VAsm/backend"
	"VAsm/frontend"

	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	backend.Run("")
	frontend.BuildMainWindow(w)
	w.ShowAndRun()
}
