package paddel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddel struct {
	X, Y, W, H float64
	speed      float64
	clr        color.RGBA
}

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

func (p *Paddel) Init() {
	p.W = 100
	p.H = 14
	p.X = ScreenWidth/2 - p.W/2
	p.Y = ScreenHeight - 100
	p.speed = 3.5
	p.clr = color.RGBA{0, 0, 255, 255}
}

func (p *Paddel) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.X-p.speed > 0 {
			p.X -= p.speed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.X+p.W+p.speed < ScreenWidth {
			p.X += p.speed
		}
	}
}

func (p *Paddel) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.W), float32(p.H), p.clr, false)
}
