package page

import (
	"G2I/internal/components/layout"
	"github.com/rivo/tview"
)

type Home struct {
	*tview.Flex
	app *tview.Application
}

func NewHome(app *tview.Application) *Home {
	sidebar := layout.NewSidebar("Menu")

	formText := tview.NewTextView().SetText("Create a new public and private key").SetSize(1, 40)
	nameInput := tview.NewInputField().SetLabel("Name: ")
	emailInput := tview.NewInputField().SetLabel("Email: ")
	commentInput := tview.NewInputField().SetLabel("Comment: ")

	form := tview.NewForm().
		AddFormItem(formText).
		AddFormItem(nameInput).
		AddFormItem(emailInput).
		AddFormItem(commentInput).
		AddButton("Generate", func() {
			// Generate logic here
		}).SetButtonsAlign(tview.AlignRight)

	content := tview.NewFlex().SetDirection(tview.FlexRow).SetDirection(tview.FlexColumn)
	content.SetBorder(true)
	content.SetTitle("Pretty Good Privacy (PGP)")

	flex := tview.NewFlex().
		AddItem(sidebar, 20, 1, true).
		AddItem(content, 0, 1, false)

	sidebar.SetItems([]string{"Create new alias"}, func(index int, text string) {
		switch index {
		case 0:
			content.Clear()
			spacer := tview.NewBox()
			content.AddItem(spacer, 24, 1, false)
			content.AddItem(form, 0, 1, true)
			content.AddItem(spacer, 28, 1, false)
			app.SetFocus(form)
		}
	})

	return &Home{
		Flex: flex,
		app:  app,
	}
}
