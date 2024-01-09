package handler

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Mau005/TibiaDepotTools/controller"
	"github.com/Mau005/TibiaDepotTools/model"
)

func Lobby(w fyne.Window) *fyne.Container {
	var character controller.CharacterController

	characters, err := character.GetAllCharacter()
	if err != nil {
		log.Println(err)
		return nil
	}
	container := container.NewVBox(widget.NewButton("Crear Character", func() {
		w.SetContent(Test(w))
	}))
	for _, value := range characters {
		container.Add(widget.NewButton(value.Name, func() {
			w.SetContent(ItemsHanlder(w, value.ID))
		}))
	}
	return container
}

func Test(w fyne.Window) *fyne.Container {

	entry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Character", Widget: entry}},
		OnSubmit: func() {
			var character model.Character
			var err error
			if entry.Text == "" {
				log.Println("El character no puede ser nulo")
				return
			}
			character.Name = entry.Text
			var characterCtl controller.CharacterController
			character, err = characterCtl.CreateCharacter(character)
			if err != nil {
				log.Println(err)
				return
			}
			w.SetContent(Lobby(w))
		},
		OnCancel: func() {
			w.SetContent(Lobby(w))
		},
	}

	// we can also append items

	return container.NewVBox(form)
}
