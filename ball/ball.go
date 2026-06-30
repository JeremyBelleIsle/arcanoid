package ball

import (
	"ARCANOID/block"
	"ARCANOID/paddel"
	"ARCANOID/power_up"
	"image/color"
	"math"
	"math/rand"
	"slices"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	X, y, r float64
	vx, vy  float64
	clr     color.RGBA
}

func Init(balls *[]Ball) {
	vy := float64(rand.Intn(2) + 3)
	vy = -vy

	*balls = append(*balls, Ball{X: paddel.ScreenWidth / 2, y: paddel.ScreenHeight / 2, r: 13, vx: float64(rand.Intn(2)) + 3, vy: vy, clr: color.RGBA{255, 0, 0, 255}})

	if (*balls)[0].vx < .4 && (*balls)[0].vx > -.4 {
		(*balls)[0].vx *= rand.Float64() + 1
	}
}

func (b *Ball) Mouvement() {
	b.X += b.vx
	b.y += b.vy
}

func ResetBallsForTheNextLevel(balls *[]Ball) {
	*balls = []Ball{}

	Init(balls)
}

func (b *Ball) Coll(p paddel.Paddel, blocks *[]block.Block, power_ups *[]power_up.PowerUp) {
	// paddel coll

	if gameutil.CircleRectCollision(b.X, b.y, b.r, p.X, p.Y, p.W, p.H) {
		b.vy = -b.vy

		b.vx /= rand.Float64() + .5
	}
	// window coll

	if b.X+b.vx >= paddel.ScreenWidth {
		b.vx = -b.vx
	} else if b.X-(b.r/2)-b.vx <= 0 {
		b.vx = math.Abs(b.vx)
	}

	if b.y-(b.r/2) <= 0 {
		b.vy = math.Abs(b.vy)
	}
	// blocks coll

	for i, bl := range *blocks {
		if gameutil.CircleRectCollision(b.X, b.y, b.r, bl.X, bl.Y, bl.W, bl.H) {
			blockCenterX := bl.X + bl.W/2
			blockCenterY := bl.Y + bl.H/2

			distX := math.Abs(b.X - blockCenterX)
			distY := math.Abs(b.y - blockCenterY)

			overlapX := (bl.W / 2) + b.r - distX
			overlapY := (bl.H / 2) + b.r - distY

			if overlapX < overlapY {
				// X
				if b.X < blockCenterX {
					b.vx = -math.Abs(b.vx)
				} else {
					b.vx = math.Abs(b.vx)
				}
			} else {
				// Y
				if b.y < blockCenterY {
					b.vy = -math.Abs(b.vy)

					b.vx *= rand.Float64() + .5
				} else {
					b.vy = math.Abs(b.vy)

					b.vx *= rand.Float64() + .5
				}
			}

			switch bl.Speciality {
			case "+1 ball":
				*power_ups = power_up.Spawn(*power_ups, bl.X, bl.Y, "+1 ball")
			case "pallet extender":
			}

			*blocks = slices.Delete(*blocks, i, i+1)

			break
		}
	}
}

func Lose(balls []Ball) bool {
	lose := true

	for _, b := range balls {
		if b.y < paddel.ScreenHeight {
			lose = false

			break
		}
	}

	return lose
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(b.X), float32(b.y), float32(b.r), b.clr, false)
}
