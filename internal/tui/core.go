package tui

import (
	"G2I/internal/components/layout"
	"G2I/internal/input"
	"G2I/internal/page"
	"G2I/internal/page/modal"
	"github.com/rivo/tview"
)

type App struct {
	app   *tview.Application
	pages *tview.Pages
}

func NewApp() *App {
	app := tview.NewApplication()
	pages := tview.NewPages()

	home := page.NewHome(app)
	quitModal := modal.NewQuit(app, pages)

	menuBar := layout.NewMenuBar()

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(home, 0, 1, true).
		AddItem(menuBar, 1, 0, false)

	pages.AddPage("home", flex, true, true)
	pages.AddPage("quit", quitModal, true, false)

	inputHandler := input.NewHandler(app, pages)
	inputHandler.CaptureKeys()

	return &App{
		app:   app,
		pages: pages,
	}
}

func (a *App) GetRoot() tview.Primitive {
	return a.pages
}

func (a *App) Run() error {
	return a.app.SetRoot(a.pages, true).Run()
}
