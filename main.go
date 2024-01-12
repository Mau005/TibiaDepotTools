package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Mau005/TibiaDepotTools/db"
	"github.com/Mau005/TibiaDepotTools/handler"
)

func main() {
	//init layout
	a := app.New()
	w := a.NewWindow("TibiaDepotTools V1.0")

	err := db.ConectionSqlite()
	if err != nil {
		log.Panic(err)
	}

	w.Resize(fyne.NewSize(280, 360))
	w.SetContent(handler.Lobby(w))
	w.ShowAndRun()
	//end
}
