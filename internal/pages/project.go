package pages

import (
	// "fmt"

	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
)

func ProjectPage(app *tview.Application) *tview.Grid {
	// ============ DESINIG LOGIC ==============
	// header and its container
	header := tview.NewTextView().
		SetText("Projetos").
		SetTextAlign(tview.AlignCenter).
		SetTextStyle(tcell.StyleDefault.Bold(true))

	headerWrapper := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).      
		AddItem(header, 1, 0, true).    
		AddItem(nil, 0, 1, false)       



	// footer text and its container
	footerText := "(a) add  |  (d) delete  |  (e) edit   |  (m) move item"
	footer := tview.NewTextView().
        SetTextAlign(tview.AlignLeft).   
        SetText(footerText)



	// list 
	list := tview.NewList().
		AddItem("Item 1", "", '*', nil).
		AddItem("item 2", "", '*', nil).
		AddItem("item 3", "", '*', nil).
		AddItem("item 4", "", '*', nil).
		AddItem("item 5", "", '*', nil).
		AddItem("item 6", "", '*', nil)



	// description box
	descText := tview.NewTextView().
		SetText("Something as example").
		SetWrap(true).
		SetScrollable(true)



	// Description
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
		AddItem(descText, 0, 1, true)


	// ============ KEYS LOGIC ==============
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		// move down at list
		case 'j':
			current := list.GetCurrentItem()
			if current == list.GetItemCount() - 1 {
				list.SetCurrentItem(0)
			} else {
				list.SetCurrentItem(current + 1)
			}
			return nil

		// move up at list
		case 'k':
			current := list.GetCurrentItem()
			if current == 0 {
				list.SetCurrentItem(list.GetItemCount() - 1)
			} else {
				list.SetCurrentItem(current - 1)
			}
			return nil

		// move to the thirst item
		case 'g':
			list.SetCurrentItem(0)

		// move to the last item
		case 'G':
			lastIndex := list.GetItemCount() - 1
			list.SetCurrentItem(lastIndex)
		}

		switch event.Key() {
		// move description up
		case tcell.KeyUp:
			row, col := descText.GetScrollOffset()
			if row > 0 {
				descText.ScrollTo(row-1, col)
			}
			return nil

		// move description donw
		case tcell.KeyDown:
			row, col := descText.GetScrollOffset()
			descText.ScrollTo(row+1, col)
			return nil

		// move the descripion to the begin
		case tcell.KeyPgUp:
			_, col := descText.GetScrollOffset()
			descText.ScrollTo(0, col)
			return nil

		// move the description to the end
		case tcell.KeyPgDn:
			_, col := descText.GetScrollOffset()
			descText.ScrollTo(999999, col)
			return nil


		default:
			return event
		}
	})



	// ============ GRID LOGIC ==============
	layout := tview.NewGrid().
		SetRows(3, 0, 1).
		SetColumns(120, 0).
		SetBorders(true).
		AddItem(headerWrapper, 0, 0, 1, 3, 0, 0, false).
		AddItem(description, 1, 0, 1, 1, 0, 0, false).
		AddItem(list, 1, 1, 1, 2, 0, 0, true).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	return layout
}

