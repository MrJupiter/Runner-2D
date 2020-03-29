package main

import (
	"github.com/MrJupiter/Runner-2D/internal/ui"
	"github.com/hajimehoshi/ebiten"
	_ "image/png"
	"log"
)

func main() {
	game := new(ui.Runner2DGame)
	game.Initialize()
	if err := ebiten.Run(game.Update, game.WindowDimensions.Width, game.WindowDimensions.Height, 1, "Runner 2D"); err != nil {
		log.Fatal(err)
	}
}