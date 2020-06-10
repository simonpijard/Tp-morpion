package main

import (
	"github.com/tadvi/winc"
)

func main() {

	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)
	mainWindow.SetText("Morpion")
	btn := winc.NewPushButton(mainWindow)
	btn.SetText("Fermer")
	btn.SetPos(150, 200)
	btn.SetSize(100, 40)
	btn.OnClick().Bind(wndOnClose)
	tab00 := winc.NewPushButton(mainWindow)
	tab00.SetText("")
	tab00.SetPos(10, 10)
	tab00.SetSize(50, 50)
	tab01 := winc.NewPushButton(mainWindow)
	tab01.SetText("")
	tab01.SetPos(65, 10)
	tab01.SetSize(50, 50)
	tab02 := winc.NewPushButton(mainWindow)
	tab02.SetText("")
	tab02.SetPos(120, 10)
	tab02.SetSize(50, 50)
	tab10 := winc.NewPushButton(mainWindow)
	tab10.SetText("")
	tab10.SetPos(10, 65)
	tab10.SetSize(50, 50)
	tab11 := winc.NewPushButton(mainWindow)
	tab11.SetText("")
	tab11.SetPos(65, 65)
	tab11.SetSize(50, 50)
	tab12 := winc.NewPushButton(mainWindow)
	tab12.SetText("")
	tab12.SetPos(120, 65)
	tab12.SetSize(50, 50)
	tab20 := winc.NewPushButton(mainWindow)
	tab20.SetText("")
	tab20.SetPos(10, 120)
	tab20.SetSize(50, 50)
	tab21 := winc.NewPushButton(mainWindow)
	tab21.SetText("")
	tab21.SetPos(65, 120)
	tab21.SetSize(50, 50)
	tab22 := winc.NewPushButton(mainWindow)
	tab22.SetText("")
	tab22.SetPos(120, 120)
	tab22.SetSize(50, 50)
	mainWindow.Center()
	mainWindow.Show()
	mainWindow.OnClose().Bind(wndOnClose)
	winc.RunMainLoop()
	tour := 1
	gagner := 0
	for gagner == 0 {
		if tab00.Text == "X" && tab01.Text == "X" && tab02.Text == "X" {
			gagner = 1
		}
		if tab10.Text == "X" && tab11.Text == "X" && tab12.Text == "X" {
			gagner = 1
		}
		if tab20.Text == "X" && tab21.Text == "X" && tab22.Text == "X" {
			gagner = 1
		}
		if tab00.Text == "X" && tab10.Text == "X" && tab20.Text == "X" {
			gagner = 1
		}
		if tab01.Text == "X" && tab11.Text == "X" && tab21.Text == "X" {
			gagner = 1
		}
		if tab02.Text == "X" && tab12.Text == "X" && tab22.Text == "X" {
			gagner = 1
		}
		if tab00.Text == "X" && tab11.Text == "X" && tab22.Text == "X" {
			gagner = 1
		}
		if tab02.Text == "X" && tab11.Text == "X" && tab20.Text == "X" {
			gagner = 1
		}
		if tab00.Text == "O" && tab01.Text == "O" && tab02.Text == "O" {
			gagner = 1
		}
		if tab10.Text == "O" && tab11.Text == "O" && tab12.Text == "O" {
			gagner = 1
		}
		if tab20.Text == "O" && tab21.Text == "O" && tab22.Text == "O" {
			gagner = 1
		}
		if tab00.Text == "O" && tab10.Text == "O" && tab20.Text == "O" {
			gagner = 1
		}
		if tab01.Text == "O" && tab11.Text == "O" && tab21.Text == "O" {
			gagner = 1
		}
		if tab02.Text == "O" && tab12.Text == "O" && tab22.Text == "O" {
			gagner = 1
		}
		if tab00.Text == "O" && tab11.Text == "O" && tab22.Text == "O" {
			gagner = 1
		}
		if tab02.Text == "O" && tab11.Text == "O" && tab20.Text == "O" {
			gagner = 1
		}
		if tour == 1 {
			mainWindow.SetText("Morpion: Joueur 1 jouez!")
			tab00.OnClick().Bind(func(e *winc.Event) {
				tab00.SetText("X")
				tour = 2
			})

			tab01.OnClick().Bind(func(e *winc.Event) {
				tab01.SetText("X")
				tour = 2
			})

			tab02.OnClick().Bind(func(e *winc.Event) {
				tab02.SetText("X")
				tour = 2
			})

			tab10.OnClick().Bind(func(e *winc.Event) {
				tab10.SetText("X")
				tour = 2
			})
			tab11.OnClick().Bind(func(e *winc.Event) {
				tab11.SetText("X")
				tour = 2
			})

			tab12.OnClick().Bind(func(e *winc.Event) {
				tab12.SetText("X")
				tour = 2
			})

			tab20.OnClick().Bind(func(e *winc.Event) {
				tab20.SetText("X")
				tour = 2
			})

			tab21.OnClick().Bind(func(e *winc.Event) {
				tab21.SetText("X")
				tour = 2
			})

			tab22.OnClick().Bind(func(e *winc.Event) {
				tab22.SetText("X")
				tour = 2
			})
		}
		if tour == 2 {
			mainWindow.SetText("Morpion: Joueur 2 jouez!")
			tab00.OnClick().Bind(func(e *winc.Event) {
				tab00.SetText("O")
				tour = 1
			})

			tab01.OnClick().Bind(func(e *winc.Event) {
				tab01.SetText("O")
				tour = 1
			})

			tab02.OnClick().Bind(func(e *winc.Event) {
				tab02.SetText("O")
				tour = 1
			})

			tab10.OnClick().Bind(func(e *winc.Event) {
				tab10.SetText("O")
				tour = 1
			})

			tab11.OnClick().Bind(func(e *winc.Event) {
				tab11.SetText("O")
				tour = 1
			})

			tab12.OnClick().Bind(func(e *winc.Event) {
				tab12.SetText("O")
				tour = 1
			})

			tab20.OnClick().Bind(func(e *winc.Event) {
				tab20.SetText("O")
				tour = 1
			})

			tab21.OnClick().Bind(func(e *winc.Event) {
				tab21.SetText("O")
				tour = 1
			})

			tab22.OnClick().Bind(func(e *winc.Event) {
				tab22.SetText("O")
				tour = 1
			})

		}
	}

}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
