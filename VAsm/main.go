package main

import (
	"VAsm/backend"
	"VAsm/frontend"

	"fyne.io/fyne/v2/app"
)

func main() {
	var entrada string = "print(\"Hola mundo\")"
	a := app.New()
	w := a.NewWindow("Hello")
	backend.Run(entrada)
	frontend.BuildMainWindow(w)
	w.ShowAndRun()
}
