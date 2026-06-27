package power_up

import (
	"math/rand"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
)

type PowerUp struct {
	x, y  float64
	speed float64
	img   *ebiten.Image
}

var (
	plus1Img *ebiten.Image
)

func InitImgs() {
	plus1Img = gameutil.LoadImage("power_up/+1 icon.png")
}

func Spawn(powerUps *[]PowerUp, x float64, y float64) {
	var img *ebiten.Image
	switch rand.Intn(2) {
	case 1:
		img = plus1Img
	case 2:
	}

	*powerUps = append(*powerUps, PowerUp{x: x, y: y, speed: float64(rand.Intn(3) + 3), img: img})
}
