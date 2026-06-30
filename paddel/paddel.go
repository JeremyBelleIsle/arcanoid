package paddel

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddel struct {
	X, Y, W, H float64
	Speed      float64
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
	p.Speed = 3.5
	p.clr = color.RGBA{0, 0, 255, 255}
}

func (p *Paddel) Input() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.X-p.Speed > 0 {
			p.X -= p.Speed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.X+p.W+p.Speed < ScreenWidth {
			p.X += p.Speed
		}
	}
}

func (p *Paddel) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.W), float32(p.H), p.clr, false)
}
