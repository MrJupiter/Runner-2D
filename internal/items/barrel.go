package items

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
)

type Barrel struct {
	Position collision2d.Vector
	ImgScale collision2d.Vector
	Img *ebiten.Image
	BarrelBox collision2d.Box
	VoidBox collision2d.Box
	Passed bool
	Ignored bool
	Undisplayed bool
}

func (barrel * Barrel) Initialize(barrelStartPosition float64){
	var err error
	barrel.Img,_, err = ebitenutil.NewImageFromFile("resources/img/barrel.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	barrel.ImgScale = collision2d.Vector{
		X: 0.22,
		Y: 0.17,
	}
	barrel.Position.X, barrel.Position.Y = barrelStartPosition, 580
	barrel.BarrelBox = collision2d.Box{
		Pos: collision2d.Vector{X:barrel.Position.X + 281 * barrel.ImgScale.X, Y:barrel.Position.Y + 165 * barrel.ImgScale.Y},
		W:   (617 - 281) * barrel.ImgScale.X,
		H:   (755 - 165) * barrel.ImgScale.Y,
	}

	barrel.VoidBox = collision2d.Box{
		Pos: collision2d.Vector{X:barrel.Position.X + 281 * barrel.ImgScale.X, Y:0},
		W:   (617 - 281) * barrel.ImgScale.X,
		H:   barrel.BarrelBox.Pos.Y,
	}

	return
}

func (barrel *Barrel) GetDrawOptions() (*ebiten.Image, *ebiten.DrawImageOptions){
	opBarrel := &ebiten.DrawImageOptions{}
	opBarrel.GeoM.Scale(barrel.ImgScale.X, barrel.ImgScale.Y)
	opBarrel.GeoM.Translate(barrel.Position.X, barrel.Position.Y)
	return barrel.Img, opBarrel
}

func (barrel *Barrel) Play() {
	barrel.Position.X-=3.5
	barrel.BarrelBox.Pos.X-=3.5
	barrel.VoidBox.Pos.X-=3.5

	if barrel.BarrelBox.Pos.X + barrel.BarrelBox.W < 0 {
		barrel.Undisplayed = true
	}
}

func (barrel *Barrel) GetCollisionBox() (*ebiten.Image, *ebiten.DrawImageOptions){
	rectangle, _ := ebiten.NewImage(int(barrel.BarrelBox.W), int(barrel.BarrelBox.H), ebiten.FilterNearest)

	rectangle.Fill(color.White)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(barrel.BarrelBox.Pos.X, barrel.BarrelBox.Pos.Y)

	return rectangle, opts
}

func (barrel *Barrel) GetVoidCollisionBox() (*ebiten.Image, *ebiten.DrawImageOptions){
	rectangle, _ := ebiten.NewImage(int(barrel.VoidBox.W), int(barrel.VoidBox.H), ebiten.FilterNearest)

	rectangle.Fill(color.Black)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(barrel.VoidBox.Pos.X, barrel.VoidBox.Pos.Y)

	return rectangle, opts
}
