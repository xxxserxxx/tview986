package main

import (
	"strings"

	"github.com/rivo/tview"
)

type HelpWidget struct {
	Root *tview.Flex

	helpBook                *tview.Flex
	leftColumn, rightColumn *tview.TextView

	// external references
	ui *Ui
}

func (ui *Ui) createHelpWidget() (m *HelpWidget) {
	m = &HelpWidget{
		ui: ui,
	}

	// two help columns side by side
	m.leftColumn = tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
	m.rightColumn = tview.NewTextView().
		SetTextAlign(tview.AlignLeft).
		SetDynamicColors(true)
	m.helpBook = tview.NewFlex().
		SetDirection(tview.FlexColumn)

	// button at the bottom
	closeButton := tview.NewButton("Close")
	closeButton.SetSelectedFunc(func() {
		ui.CloseHelp()
	})

	m.Root = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(m.helpBook, 0, 1, false).
		AddItem(closeButton, 1, 0, true)

	m.Root.Box.SetBorder(true).SetTitle(" Help ")

	return
}

func (h *HelpWidget) RenderHelp() {
	leftText := "[::b]Playback[::-]\n" + tview.Escape(strings.TrimSpace(helpPlayback))
	h.leftColumn.SetText(leftText)

	rightText := "[::b]Browser[::-]\n" + tview.Escape(strings.TrimSpace(helpPageBrowser))

	h.rightColumn.SetText(rightText)

	h.helpBook.Clear()
	if rightText != "" {
		h.helpBook.AddItem(h.leftColumn, 38, 0, false).
			AddItem(h.rightColumn, 0, 1, true) // gets focus for scrolling
	} else {
		h.helpBook.AddItem(h.leftColumn, 0, 1, false)
	}
}

const helpPlayback = `
p      play/pause
P      stop
>      next song
-/=(+) volume down/volume up
,/.    seek -10/+10 seconds
r      add 50 random songs to queue
`

const helpPageBrowser = `
artist tab
  R     refresh the list
  /     Search artists
  a     Add all artist songs to queue
  n     Continue search forward
  N     Continue search backwards
song tab
  ENTER play song (clears current queue)
  a     add album or song to queue
  A     add song to playlist
  y     toggle star on song/album
  R     refresh the list
ESC   Close search
`

const helpPageQueue = `
d/DEL remove currently selected song from the queue
D     remove all songs from queue
y     toggle star on song
`

const helpPagePlaylists = `
n     new playlist
d     delete playlist
a     add playlist or song to queue
`
