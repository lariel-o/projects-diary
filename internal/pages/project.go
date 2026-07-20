package pages

import (
	// "fmt"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func projectPageStatic() (*tview.Grid, *tview.List, *tview.TextView) {
	header := tview.NewTextView().
		SetText("Projetos").
		SetTextAlign(tview.AlignCenter).
		SetTextStyle(tcell.StyleDefault.Bold(true))

	headerWrapper := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(header, 1, 0, true).
		AddItem(nil, 0, 1, false)

	footerText := "(a) add  |  (d) delete  |  (e) edit   |  (m) move item"
	footer := tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetText(footerText)

	list := tview.NewList()
	list.AddItem("Ola mundo 1", "", '*', nil)
	list.AddItem("Ola mundo 2", "", '*', nil)
	list.AddItem("Ola mundo 3", "", '*', nil)
	list.AddItem("Ola mundo 4", "", '*', nil)
	
	// Text description 
	descText := tview.NewTextView().SetWrap(true).SetScrollable(true)

	// Description box
	description := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(
			tview.NewTextView().
				SetText("Description").
				SetTextAlign(tview.AlignCenter).
				SetTextStyle(tcell.StyleDefault.Bold(true)),
			1, 0, true,
		).
		AddItem(nil, 1, 0, false).
		AddItem(descText, 
		0, 1, true)

	layout := tview.NewGrid().
		SetRows(3, 0, 1).
		SetColumns(120, 0).
		SetBorders(true).
		AddItem(headerWrapper, 0, 0, 1, 3, 0, 0, false).
		AddItem(description, 1, 0, 1, 1, 0, 0, false).
		AddItem(list, 1, 1, 1, 2, 0, 0, true).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	return layout, list, descText
}




func projectPageDinamic() (*tview.Grid, *tview.List, *tview.TextView) {
	grid, list, desc := projectPageStatic()

	list.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
	})

	grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		current := list.GetCurrentItem()
		lastIndex := list.GetItemCount() - 1

		switch event.Rune() {
		case 'j': 
			current := current
			if current == lastIndex {
				list.SetCurrentItem(0)
			} else {
				list.SetCurrentItem(current + 1)
			}
			return nil

		case 'k': // sobe na lista
			current := current
			if current == 0 {
				list.SetCurrentItem(lastIndex)
			} else {
				list.SetCurrentItem(current - 1)
			}
			return nil

		case 'g': 
			if lastIndex + 1 > 0 {
				list.SetCurrentItem(0)
			}
			return nil

		case 'G': 
			count := lastIndex
			if count > 0 {
				list.SetCurrentItem(count)
			}
			return nil
		}



		switch event.Key() {
		case tcell.KeyUp:
			row, col := desc.GetScrollOffset()
			if row > 0 {
				desc.ScrollTo(row-1, col)
			}
			return nil

		case tcell.KeyDown:
			row, col := desc.GetScrollOffset()
			desc.ScrollTo(row+1, col)
			return nil

		case tcell.KeyPgUp:
			_, col := desc.GetScrollOffset()
			desc.ScrollTo(0, col)
			return nil

		case tcell.KeyPgDn:
			_, col := desc.GetScrollOffset()
			desc.ScrollTo(900000, col)
			return nil

		default:
			return event
		}
	})

	return grid, list, desc
}

