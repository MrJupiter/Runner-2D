package items

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

type Runner2D struct{
	Position  collision2d.Vector
	ImgScale  float64
	Img *ebiten.Image
	RunnerBox collision2d.Box
}

func (runner2d * Runner2D) Initialize(){
	var err error
	runner2d.Img,_, err = ebitenutil.NewImageFromFile("resources/img/runner2d.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	runner2d.Position.X, runner2d.Position.Y, runner2d.ImgScale = 400, 593, 0.26
	runner2d.RunnerBox = collision2d.Box{
		Pos: collision2d.Vector{X:runner2d.Position.X + 189 * runner2d.ImgScale, Y:runner2d.Position.Y + 49 * runner2d.ImgScale},
		W:   (432 - 189) * runner2d.ImgScale,
		H:   (432 - 49) * runner2d.ImgScale,
	}

	return
}

func (runner2d *Runner2D) GetDrawOptions() (*ebiten.Image, *ebiten.DrawImageOptions){
	opRunner2D := &ebiten.DrawImageOptions{}
	opRunner2D.GeoM.Scale(runner2d.ImgScale, runner2d.ImgScale)
	opRunner2D.GeoM.Translate(runner2d.Position.X, runner2d.Position.Y)
	return runner2d.Img, opRunner2D
}

func (runner2d *Runner2D) Play(){

}

func (runner2d *Runner2D) Jump() {
	runner2d.Position.Y -= 3
	runner2d.RunnerBox.Pos.Y -= 3
}

func (runner2d *Runner2D) Descent() {
	runner2d.Position.Y += 3
	runner2d.RunnerBox.Pos.Y += 3
}

func (runner2d *Runner2D) GetCollisionBox() (*ebiten.Image, *ebiten.DrawImageOptions){
	square, _ := ebiten.NewImage(int(runner2d.RunnerBox.W), int(runner2d.RunnerBox.H), ebiten.FilterNearest)

	square.Fill(color.Black)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(runner2d.RunnerBox.Pos.X, runner2d.RunnerBox.Pos.Y)

	return square, opts
}

