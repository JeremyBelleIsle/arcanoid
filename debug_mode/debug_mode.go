package debugmode

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func CheckWordTyped(motTapé *string) {
	// mode Debug
	*motTapé += string(ebiten.AppendInputChars(nil))
}

func DoDebug(motTapé string, paX *float64, paW float64, bx float64, paSpeed float64) {
	if motTapé == "debug mode" {
		println("Debug mode actived!")
		if *paX+paW/2 > bx {
			*paX -= paSpeed
		} else {
			*paX += paSpeed
		}
	}
}
