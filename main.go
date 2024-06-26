package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// struct contains all the updatable elements of the Ui
type Ui struct {
	app   *tview.Application
	pages *tview.Pages

	helpModal  tview.Primitive
	helpWidget *HelpWidget

	rootPage *tview.Flex
}

const (
	PageHelpBox = "helpBox"
)

func main() {
	ui := InitGui()

	// run main loop
	if err := ui.Run(); err != nil {
		panic(err)
	}
}

var logList *tview.List
var showCount, hideCount int

func InitGui() (ui *Ui) {
	ui = &Ui{}

	ui.app = tview.NewApplication()
	ui.pages = tview.NewPages()

	ui.helpWidget = ui.createHelpWidget()

	// help box modal
	ui.helpModal = makeModal(ui.helpWidget.Root, 80, 30)
	ui.helpWidget.Root.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		logList.AddItem("Got an input event!", "", 0, nil)
		ui.CloseHelp()
		return event
	})

	logList = tview.NewList().ShowSecondaryText(false).
		AddItem("Some Text", "", 0, nil).
		AddItem("More text", "", 0, nil)

	logPage := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(logList, 0, 1, true)

	ui.pages.AddPage("logpage", logPage, true, true).
		AddPage(PageHelpBox, ui.helpModal, true, false)

	rootFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(ui.pages, 0, 1, true)

	// add main input handler
	rootFlex.SetInputCapture(ui.handlePageInput)

	ui.app.SetRoot(rootFlex, true).
		SetFocus(rootFlex).
		EnableMouse(true)

	return ui
}

func (ui *Ui) Run() error {
	// gui main loop (blocking)
	return ui.app.Run()
}

func (ui *Ui) ShowHelp() {
	logList.AddItem(fmt.Sprintf("ShowHelp called %d times", showCount), "", 0, nil)
	showCount++
	ui.helpWidget.RenderHelp()

	ui.pages.ShowPage(PageHelpBox)
	ui.app.SetFocus(ui.helpModal)
}

func (ui *Ui) CloseHelp() {
	logList.AddItem(fmt.Sprintf("CloseHelp called %d times", hideCount), "", 0, nil)
	hideCount++
	ui.pages.HidePage(PageHelpBox)
}

func (ui *Ui) handlePageInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case '?':
		ui.ShowHelp()

	case 'Q':
		ui.Quit()
	}

	return event
}

func (ui *Ui) ShowPage(name string) {
	ui.pages.SwitchToPage(name)
}

func (ui *Ui) Quit() {
	ui.app.Stop()
}
