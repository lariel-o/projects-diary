package pages

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func warningPage(message string, app *tview.Application, onClose func()) *tview.Grid { 
	counter := 3

	footerText := tview.NewTextView().
		SetText(fmt.Sprintf("%d  press x to skip", counter)).
		SetTextAlign(tview.AlignLeft)

	msgView := tview.NewTextView().
		SetText(message).
		SetTextAlign(tview.AlignCenter).
		SetWrap(true)

	innerFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(msgView, 0, 1, false).
		AddItem(footerText, 1, 0, false)

	popup := tview.NewGrid().
		SetRows(0, 0, 0).
		SetColumns(0, 90, 0).
		AddItem(innerFlex, 1, 1, 1, 1, 0, 0, true)

	popup.SetBorder(true).
		SetTitle(" Warning ").
		SetBorderPadding(0, 0, 0, 0)

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	closePopup := func() {
		ticker.Stop()
		close(quit)
		onClose()
	}

	go func() {
		for {
			select {
			case <-ticker.C:
				counter--
				if counter >= 0 {
					app.QueueUpdateDraw(func() { 
						footerText.SetText(fmt.Sprintf("%d  press x to skip", counter))
					})
				}
				if counter < 0 {
					app.QueueUpdateDraw(func() { 
						closePopup()
					})
					return
				}
			case <-quit:
				return
			}
		}
	}()

	popup.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'x' || event.Rune() == 'X' {
			closePopup()
			return nil
		}
		return event
	})

	grid := tview.NewGrid().
		SetRows(7, 10, 0).
		SetColumns(-1, 150, -1).
		AddItem(popup, 1, 1, 1, 1, 0, 0, true)

	return grid
}
