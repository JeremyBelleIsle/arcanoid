package ball

import (
	"ARCANOID/block"
	"ARCANOID/paddel"
	"image/color"
	"math"
	"math/rand"
	"slices"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	x, y, r float64
	vx, vy  float64
	clr     color.RGBA
}

func (b *Ball) Init() {
	b.x = paddel.ScreenWidth / 2
	b.y = paddel.ScreenHeight / 2
	b.r = 13
	b.vx = rand.Float64() + 1.7
	b.vy = rand.Float64() + 1.7
	b.clr = color.RGBA{255, 0, 0, 255}
}

func (b *Ball) Mouvement() {
	b.x += b.vx
	b.y += b.vy
}

func (b *Ball) Coll(p paddel.Paddel, blocks *[]block.Block) {
	// paddel coll

	if gameutil.CircleRectCollision(b.x, b.y, b.r, p.X, p.Y, p.W, p.H) {
		b.vy = -b.vy
	}
	// window coll

	if b.x-(b.r/2) >= paddel.ScreenWidth {
		b.vx = -b.vx
	} else if b.x-(b.r/2) <= 0 {
		b.vx = math.Abs(b.vx)
	}

	if b.y-(b.r/2) <= 0 {
		b.vy = math.Abs(b.vy)
	}
	// blocks coll

	for i, bl := range *blocks {
		if gameutil.CircleRectCollision(b.x, b.y, b.r, bl.X, bl.Y, bl.W, bl.H) {
			blockCenterX := bl.X + bl.W/2
			blockCenterY := bl.Y + bl.H/2

			distX := math.Abs(b.x - blockCenterX)
			distY := math.Abs(b.y - blockCenterY)

			overlapX := (bl.W / 2) + b.r - distX
			overlapY := (bl.H / 2) + b.r - distY

			if overlapX < overlapY {
				// X
				if b.x < blockCenterX {
					b.vx = -math.Abs(b.vx)
				} else {
					b.vx = math.Abs(b.vx)
				}
			} else {
				// Y
				if b.y < blockCenterY {
					b.vy = -math.Abs(b.vy)
				} else {
					b.vy = math.Abs(b.vy)
				}
			}

			switch bl.Speciality {
			case "+1 ball":
			case "pallet extender":
			}

			*blocks = slices.Delete(*blocks, i, i+1)

			break
		}
	}
}

func (b *Ball) Lose() bool {
	return b.y > paddel.ScreenHeight+160
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.x), float32(b.y), float32(b.r), b.clr, false)
}
