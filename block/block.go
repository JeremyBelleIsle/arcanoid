package block

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Block struct {
	X, Y, W, H float64
	Speciality string
	clr        color.RGBA
}

func LoadLevel(level int, blocks *[]Block) {
	switch level {
	case 1:
		for i := 0; i < 10; i++ {
			for j := 0; j < 3; j++ {
				*blocks = append(*blocks, Block{
					X:          float64(i*60 + 20),
					Y:          float64(j*25 + 10),
					W:          50,
					H:          20,
					Speciality: "normal",
					clr:        color.RGBA{211, 211, 211, 255},
				})
			}
		}
	case 2:
		for i := 0; i < 10; i++ {
			*blocks = append(*blocks, Block{
				X:          float64(i*60 + 20),
				Y:          10,
				W:          50,
				H:          20,
				Speciality: "normal",
				clr:        color.RGBA{211, 211, 211, 255},
			})
			*blocks = append(*blocks, Block{
				X:          float64(i*60 + 20),
				Y:          35,
				W:          50,
				H:          20,
				Speciality: "normal",
				clr:        color.RGBA{211, 211, 211, 255},
			})
		}
	case 3:
		for i := 0; i < 10; i++ {
			for j := 0; j < 5; j++ {
				if (i+j)%2 == 0 {
					*blocks = append(*blocks, Block{
						X:          float64(i*60 + 20),
						Y:          float64(j*25 + 10),
						W:          50,
						H:          20,
						Speciality: "normal",
						clr:        color.RGBA{211, 211, 211, 255},
					})
				}
			}
		}
	case 4:
		for j := 0; j < 5; j++ {
			for i := j; i < 10-j; i++ {
				*blocks = append(*blocks, Block{
					X:          float64(i*60 + 20),
					Y:          float64(j*25 + 10),
					W:          50,
					H:          20,
					Speciality: "normal",
					clr:        color.RGBA{211, 211, 211, 255},
				})
			}
		}
	case 5:
		for j := 0; j < 5; j++ {
			for i := 0; i < 10; i++ {
				if i == j || i == 9-j {
					*blocks = append(*blocks, Block{
						X:          float64(i*60 + 20),
						Y:          float64(j*25 + 10),
						W:          50,
						H:          20,
						Speciality: "normal",
						clr:        color.RGBA{211, 211, 211, 255},
					})
				}
			}
		}
	case 6:
		for i := 0; i < 10; i++ {
			*blocks = append(*blocks, Block{X: float64(i*60 + 20), Y: 10, W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
			*blocks = append(*blocks, Block{X: float64(i*60 + 20), Y: 110, W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
		}
		for j := 1; j < 4; j++ {
			*blocks = append(*blocks, Block{X: 20, Y: float64(j*25 + 10), W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
			*blocks = append(*blocks, Block{X: 560, Y: float64(j*25 + 10), W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
		}
	case 7:
		for i := 0; i < 10; i++ {
			if i > 2 && i < 7 {
				*blocks = append(*blocks, Block{X: float64(i*60 + 20), Y: 60, W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
			}
		}
		for j := 0; j < 5; j++ {
			if j > 0 && j < 4 {
				*blocks = append(*blocks, Block{X: 290, Y: float64(j*25 + 10), W: 50, H: 20, Speciality: "normal", clr: color.RGBA{211, 211, 211, 255}})
			}
		}
	case 8:
		for i := 0; i < 5; i++ {
			*blocks = append(*blocks, Block{
				X:          float64(i*60 + 20),
				Y:          float64(i*25 + 10),
				W:          50,
				H:          20,
				Speciality: "normal",
				clr:        color.RGBA{211, 211, 211, 255},
			})
			*blocks = append(*blocks, Block{
				X:          float64((9-i)*60 + 20),
				Y:          float64(i*25 + 10),
				W:          50,
				H:          20,
				Speciality: "normal",
				clr:        color.RGBA{211, 211, 211, 255},
			})
		}
	case 9:
		for i := 0; i < 10; i++ {
			for j := 0; j < 5; j++ {
				if !((i > 2 && i < 7) && (j > 0 && j < 4)) {
					*blocks = append(*blocks, Block{
						X:          float64(i*60 + 20),
						Y:          float64(j*25 + 10),
						W:          50,
						H:          20,
						Speciality: "normal",
						clr:        color.RGBA{211, 211, 211, 255},
					})
				}
			}
		}
	case 10:
		for i := 0; i < 10; i++ {
			for j := 0; j < 5; j++ {
				if (i+j)%3 == 0 || (i+j)%4 == 0 {
					*blocks = append(*blocks, Block{
						X:          float64(i*60 + 20),
						Y:          float64(j*25 + 10),
						W:          50,
						H:          20,
						Speciality: "normal",
						clr:        color.RGBA{211, 211, 211, 255},
					})
				}
			}
		}
	}
}

func Win(blocks []Block) bool {
	return len(blocks) == 0
}

func (b *Block) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, float32(b.X), float32(b.Y), float32(b.W), float32(b.H), b.clr, false)
}
