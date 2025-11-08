package main

import (
	"G2I/internal/tui"
)

func main() {
	app := tui.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
