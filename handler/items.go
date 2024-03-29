package handler

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/Mau005/TibiaDepotTools/controller"
	"github.com/Mau005/TibiaDepotTools/model"
)

func ItemsHanlder(w fyne.Window, idCharacter uint) *fyne.Container {
	var charaterCtl controller.CharacterController
	character, err := charaterCtl.GetCharacter(idCharacter)
	if err != nil {
		log.Println(err)
		return nil
	}
	content := container.NewVBox(

		widget.NewButton("Create new Item", func() {
			w.SetContent(CreateItemHandler(w, idCharacter))
		}))

	for _, values := range character.Items {
		item := values
		content.Add(container.NewHBox(
			widget.NewButton(item.Name, func() {
				w.Clipboard().SetContent(item.Name)
			}),
			widget.NewLabel(fmt.Sprintf("My Value: %d", item.Balance)),
			widget.NewLabel(fmt.Sprintf("Origin Value: %d", item.CostItem)),
			widget.NewButton("Edit", func() {
				w.SetContent(ModifyItemHandler(w, item.ID))
			})))
	}
	scroll := container.NewVScroll(content)

	return container.NewBorder(widget.NewToolbar(widget.NewToolbarAction(theme.VisibilityIcon(), func() {
		w.SetContent(ItemsCarruselHome(w, idCharacter, 1))
	})), widget.NewButton("Return", func() {
		w.SetContent(Lobby(w))
	}), nil, nil, container.NewPadded(widget.NewCard("List Item", "", scroll)))
}

func CreateItemHandler(w fyne.Window, idCharacter uint) *fyne.Container {
	name := widget.NewEntry()
	value := widget.NewEntry()
	valueOrigin := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name Item", Widget: name},
			{Text: "My Value", Widget: value},
			{Text: "Value Origin", Widget: valueOrigin}},
		OnSubmit: func() {
			var itemsCtl controller.ItemsController
			costItem, _ := strconv.ParseUint(valueOrigin.Text, 10, 64)
			balance, _ := strconv.ParseUint(value.Text, 10, 64)
			_, err := itemsCtl.CreateItems(model.Items{Name: name.Text,
				CostItem:    uint(costItem),
				Balance:     uint(balance),
				CharacterID: idCharacter})
			if err != nil {
				log.Println(err)
				return
			}
			w.SetContent(ItemsHanlder(w, idCharacter))
		},
		OnCancel: func() {
			w.SetContent(ItemsHanlder(w, idCharacter))
		},
	}
	return container.NewPadded(widget.NewCard("Create Item", "", form))
}

func ModifyItemHandler(w fyne.Window, idItem uint) *fyne.Container {
	var itemsCtl controller.ItemsController
	items, err := itemsCtl.GetItems(idItem)
	if err != nil {
		log.Println(err)
		return nil
	}

	name := widget.NewEntry()
	name.Text = items.Name
	costItem := widget.NewEntry()
	costItem.Text = strconv.Itoa(int(items.CostItem))
	balance := widget.NewEntry()
	balance.Text = strconv.Itoa(int(items.Balance))

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name Item", Widget: name},
			{Text: "My Value", Widget: balance},
			{Text: "Value Origin", Widget: costItem}},
		OnSubmit: func() {
			balanceUint, _ := strconv.ParseUint(balance.Text, 10, 64)
			costItemUint, _ := strconv.ParseUint(costItem.Text, 10, 64)

			items.Name = name.Text
			items.Balance = uint(balanceUint)
			items.CostItem = uint(costItemUint)
			_, err := itemsCtl.SaveItems(items)
			if err != nil {
				log.Println(err)
				return
			}
			w.SetContent(ItemsHanlder(w, items.CharacterID))

		},
		OnCancel: func() {
			w.SetContent(ItemsHanlder(w, items.CharacterID))
		},
	}
	form.CancelText = "Return"
	form.SubmitText = "Update"

	return container.NewPadded(container.NewBorder(
		container.NewBorder(nil, nil, nil, widget.NewToolbar(
			widget.NewToolbarAction(theme.DeleteIcon(), func() {
				_, err = itemsCtl.DelItems(items)
				if err != nil {
					log.Println(err)
				}
				w.SetContent(ItemsHanlder(w, items.CharacterID))
			}),
		)), nil, nil, nil, widget.NewCard("Update Character", "", form)))
}

func HistoryItemsHandler(w fyne.Window, idItems uint, page int) *fyne.Container {
	var itemsCtl controller.ItemsController
	item, err := itemsCtl.GetItems(idItems)
	if err != nil {
		log.Println(err)
	}
	containerListItems := container.NewVBox()
	if len(item.HistoryItems) > 0 {
		for _, historyItems := range item.HistoryItems {
			history := historyItems
			containerListItems.Add(
				container.NewHBox(
					widget.NewLabel(fmt.Sprintf("Balance New: %d", history.BalanceNew)),
					widget.NewLabel(fmt.Sprintf("Balance Old: %d", history.BalanceOld))),
			)
		}
	} else {
		containerListItems.Add(widget.NewLabel("This item has not changed since it was created."))
	}

	card := widget.NewCard(fmt.Sprintf("History Items: %s", item.Name), "This module is ordered from the most current date to the oldest date.", containerListItems)

	return container.NewBorder(nil, widget.NewButton("Return", func() {
		w.SetContent(ItemsCarruselHome(w, item.CharacterID, page))
	}), nil, nil, card)
}
