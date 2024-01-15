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
)

func ItemsCarruselHome(w fyne.Window, characterId uint, page int) *fyne.Container {
	var itemsCtl controller.ItemsController
	countPages := itemsCtl.GetTotalItemsCharacter(characterId)
	if countPages == 0 {
		pop := widget.NewPopUp(widget.NewLabel("Not Have Item! please create"), w.Canvas())
		pop.Show()
		return CreateItemHandler(w, characterId)
	}
	content := CarruselItemHandler(w, characterId, page)

	botonReturn := widget.NewButton("Return", func() {
		w.SetContent(ItemsCarruselHome(w, characterId, page-1))
	})
	buttonNext := widget.NewButton("Next", func() {
		w.SetContent(ItemsCarruselHome(w, characterId, page+1))
	})
	buttonContent := container.NewHBox()
	if page > 1 {
		buttonContent.Add(botonReturn)
	}

	if countPages > int64(page) {
		buttonContent.Add(buttonNext)
	}

	return container.NewBorder(nil, container.NewCenter(buttonContent), nil, nil, container.NewPadded(content))
}

func CarruselItemHandler(w fyne.Window, characterId uint, page int) *fyne.Container {
	var itemsCtl controller.ItemsController
	items, err := itemsCtl.OffsetItems(characterId, page)
	if err != nil {
		pop := widget.NewPopUp(widget.NewLabel("Not Have Item!"), w.Canvas())
		pop.Show()
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
			statusSave := false
			if !(items.Name == name.Text) {
				statusSave = true
			}
			if !(items.Balance == uint(balanceUint)) {
				statusSave = true
			}
			if !(items.CostItem == uint(costItemUint)) {
				statusSave = true
			}
			items.Name = name.Text
			items.Balance = uint(balanceUint)
			items.CostItem = uint(costItemUint)
			if statusSave {
				_, err := itemsCtl.SaveItems(items)
				if err != nil {
					log.Println(err)
					return
				}
				pop := widget.NewPopUp(widget.NewLabel("Update Completed!"), w.Canvas())
				pop.Show()
			} else {
				pop := widget.NewPopUp(widget.NewLabel("Not have change!"), w.Canvas())
				pop.Show()
			}

		},
		OnCancel: func() {
			w.SetContent(ItemsHanlder(w, items.CharacterID))
		},
	}
	form.CancelText = "Return"
	form.SubmitText = "Update"
	buttonHistory := widget.NewButton("History Updates", func() {
		w.SetContent(HistoryItemsHandler(w, items.ID, page))
	})
	copyButton := widget.NewButton(fmt.Sprintf("Copy %s", items.Name), func() {
		w.Clipboard().SetContent(items.Name)
		pop := widget.NewPopUp(widget.NewLabel("Copy Text Completed!"), w.Canvas())
		pop.Show()

	})
	contentResult := container.NewVBox(copyButton, widget.NewSeparator(), widget.NewCard("Update Items", "", form), widget.NewSeparator(), widget.NewSeparator(), buttonHistory)

	return container.NewBorder(
		container.NewBorder(nil, nil,
			widget.NewToolbar(widget.NewToolbarAction(theme.ContentUndoIcon(), func() { w.SetContent(ItemsHanlder(w, items.CharacterID)) })),
			widget.NewToolbar(widget.NewToolbarAction(theme.DeleteIcon(), func() {
				_, err = itemsCtl.DelItems(items)
				if err != nil {
					log.Println(err)
				}
				w.SetContent(ItemsHanlder(w, items.CharacterID))
			}))), nil, nil, nil, container.NewPadded(container.NewVBox(contentResult)))
}
