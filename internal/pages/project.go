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


	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	footerText := "(a) add  |  (d) delete  |  (e) edit   |  (m) move item"
	footer := tview.NewTextView().
			SetTextAlign(tview.AlignLeft).
			SetText(footerText)

	list := tview.NewList().
		AddItem("Item 1", "", '*', nil).
		AddItem("item 2", "", '*', nil).
		AddItem("item 3", "", '*', nil).
		AddItem("item 4", "", '*', nil).
		AddItem("item 5", "", '*', nil).
		AddItem("item 6", "", '*', nil)





	// ============ KEYS LOGIC ==============
	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        switch event.Rune() {
        case 'd':
			list.RemoveItem(list.GetCurrentItem())
            return nil
        default:
            return event 
        }
    })




	// ============ GRID LOGIC ==============
	layout := tview.NewGrid().
		SetRows(3, 30, 3).
		SetColumns(120, 0).
		SetBorders(true).
		AddItem(headerWrapper, 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("Side 1"), 1, 0, 1, 1, 0, 0, false).
		AddItem(list, 1, 1, 1, 2, 0, 0, true).
		AddItem(footer, 2, 0, 1, 3, 0, 0, false)

	return layout
}

