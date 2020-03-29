package ui

import (
	"github.com/MrJupiter/Runner-2D/internal/items"
	"github.com/MrJupiter/Runner-2D/internal/ui/components"
	"github.com/MrJupiter/Runner-2D/resources/fonts"
	"github.com/Tarliton/collision2d"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Runner2DGame struct{
	Background components.Background
	Floor components.Floor
	Runner2D *items.Runner2D
	Barrels[] *items.Barrel
	Runner2DScore int
	WindowDimensions dimension
	GameOver components.GameOver
}

type dimension struct {
	Width, Height int
}

var (
	fontsRunner2DScore font.Face
	descent bool
	jumpPressed bool
	startGame bool
	jumpCounter int
)

const barrelsNumber  = 2

func (game *Runner2DGame) createBarrels(number int) {
	barrel := new(items.Barrel)
	barrel.Initialize(1024)
	game.Barrels = append(game.Barrels, barrel)

	for i:=1; i<number; i++ {
		barrelLoop := new(items.Barrel)
		barrelLoop.Initialize(game.Barrels[i-1].Position.X + (400 + rand.Float64() * (900 - 400)))
		game.Barrels = append(game.Barrels, barrelLoop)
	}
	return
}

func (game *Runner2DGame) drawBarrels(screen *ebiten.Image){
	for i:=0; i<len(game.Barrels) ; i++ {
		screen.DrawImage(game.Barrels[i].GetDrawOptions())
		game.Barrels[i].Play()
	}
}

func initializeFont(){
	tt, err := freetype.ParseFont(fonts.GetFont())
	if err != nil {
		log.Fatal(err)
	}

	fontsRunner2DScore =  truetype.NewFace(tt, &truetype.Options{
		Size:    34,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func (game *Runner2DGame) Initialize(){
	initializeFont()

	game.WindowDimensions = dimension{Width: 1024, Height: 768}
	game.Runner2DScore = 0

	game.Background.Initialize()
	game.Floor.Initialize()
	game.GameOver.Initialize()

	game.Runner2D = new(items.Runner2D)
	game.Runner2D.Initialize()

	game.createBarrels(barrelsNumber)

	return
}

func (game *Runner2DGame) incrementRunner2DScore() {
	for i:=0; i<len(game.Barrels); i++{
		birdPassed, _ := collision2d.TestPolygonPolygon(game.Runner2D.RunnerBox.ToPolygon(), game.Barrels[i].VoidBox.ToPolygon())
		if birdPassed {
			game.Barrels[i].Passed = birdPassed
		}
		if game.Barrels[i].Passed && game.Runner2D.RunnerBox.Pos.X > game.Barrels[i].BarrelBox.Pos.X + game.Barrels[i].BarrelBox.W {
			game.Runner2DScore++
			game.Barrels[i].Ignored = true // used to compare collision between the bird and the nearest pipe (not bypassed yet)
			game.Barrels[i].Passed = false // used to increment the score
		}
	}
}

func (game *Runner2DGame) getIndexOfNextBomb() int{
	for i:=0; i<len(game.Barrels) ; i++ {
		if !game.Barrels[i].Ignored {
			return i
		}
	}
	return 0
}

func (game *Runner2DGame) checkGameOverTrigger() bool {
	index := game.getIndexOfNextBomb()
	checkUpperPipeBodyBox, _ := collision2d.TestPolygonPolygon(game.Runner2D.RunnerBox.ToPolygon(), game.Barrels[index].BarrelBox.ToPolygon())
	return checkUpperPipeBodyBox
}

func (game *Runner2DGame) Update(screen *ebiten.Image) error {
	rand.Seed(time.Now().UnixNano())
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.DrawImage(game.Background.GetDrawOptions())
	defer screen.DrawImage(game.Floor.GetDrawOptions())
	defer text.Draw(screen,  strconv.Itoa(game.Runner2DScore), fontsRunner2DScore, 15,40, color.White)

	if game.checkGameOverTrigger() {
		screen.DrawImage(game.GameOver.Img, game.GameOver.GetDrawOptions(game.WindowDimensions.Width, game.WindowDimensions.Height))

		startGame = false

		if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			game.Runner2DScore = 0
			game.Barrels = nil
			game.Runner2D.Initialize()
			game.createBarrels(barrelsNumber)
		}
	}else {
		screen.DrawImage(game.Runner2D.GetDrawOptions())

		if game.Barrels[0].Undisplayed == true {
			game.Barrels = game.Barrels[1:]
			barrelLoop := new(items.Barrel)
			barrelLoop.Initialize(game.Barrels[len(game.Barrels)-1].Position.X + (400 + rand.Float64() * (900 - 400)))
			game.Barrels = append(game.Barrels, barrelLoop)
		}

		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			startGame = true
		}

		if startGame {
			game.incrementRunner2DScore()
			game.Runner2D.Play()
			game.drawBarrels(screen)

			if ebiten.IsKeyPressed(ebiten.KeyUp) {
				jumpPressed = true
			}
			if jumpPressed {
				game.Runner2D.Jump()
				jumpCounter++
			}
			if jumpCounter > 60 {
				jumpPressed = false
				if game.Runner2D.Position.Y < 593{
					game.Runner2D.Descent()
				}
				if game.Runner2D.Position.Y > 592{
					jumpCounter = 0
				}
			}
		}
	}

	return nil
}