package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1200
	screenHeight = 675
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ai Đưa Em Về - TIA | MelodyCore (Go)")
	ebiten.SetTPS(60)

	game := NewApp()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
