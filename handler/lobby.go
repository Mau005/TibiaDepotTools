package handler

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
	containerCharacter := container.NewVBox(container.NewCenter(widget.NewLabel("List Character")))
	for _, value := range characters {
		character := value
		containerCharacter.Add(widget.NewButton(character.Name, func() {
			w.SetContent(ItemsHanlder(w, character.ID))
		}))
	}

	return container.NewBorder(widget.NewToolbar(widget.NewToolbarAction(theme.ContentAddIcon(), func() {
		w.SetContent(Test(w))
	})),
		nil, nil, nil, container.NewPadded(container.NewVScroll(containerCharacter)))

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

	return container.NewVBox(container.NewCenter(widget.NewLabel("Create Character")),
		widget.NewLabel("The names created have to be unique, as advice\n I can say that it is better to place\n the name of your character, as a detail to\n manage in the application"),
		container.NewPadded(widget.NewCard("Create Character", "", form)))
}
