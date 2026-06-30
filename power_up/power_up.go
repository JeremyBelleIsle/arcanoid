package power_up

import (
	"ARCANOID/paddel"
	"image/color"
	"math/rand"
	"slices"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type PowerUp struct {
	x, y  float64
	speed float64
	power string
	img   *ebiten.Image
}

var (
	plus1Img *ebiten.Image
)

func InitImgs() {
	plus1Img = gameutil.LoadImage("power_up/+1 icon.png")
}

func Spawn(powerUps []PowerUp, x float64, y float64, power string) []PowerUp {
	var img *ebiten.Image
	switch power {
	case "+1 ball":
		img = plus1Img
	}

	powerUps = append(powerUps, PowerUp{x: x, y: y, speed: float64(rand.Intn(3) + 1), img: img, power: power})

	return powerUps
}

func (p *PowerUp) Mouv() {
	p.y += p.speed
}

// Coll retourne maintenant une chaîne de caractères contenant le pouvoir si touché, sinon une chaîne vide ""
func Coll(paddel paddel.Paddel, powerUps *[]PowerUp) string {
	for i, p := range *powerUps {
		hitBoxX, hitBoxY, hitBoxR := p.x+20, p.y+20, 20

		if gameutil.CircleRectCollision(hitBoxX, hitBoxY, float64(hitBoxR), paddel.X, paddel.Y, paddel.W, paddel.H) {
			*powerUps = slices.Delete(*powerUps, i, i+1)
			return p.power
		}
	}
	return ""
}

func (p *PowerUp) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.02, 0.02)
	op.GeoM.Translate(p.x+7, p.y)
	screen.DrawImage(p.img, op)

	vector.StrokeCircle(screen, float32(p.x+29), float32(p.y+20), 19, 1, color.RGBA{255, 0, 0, 255}, false)
}
