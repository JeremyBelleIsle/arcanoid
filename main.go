package main

import (
	"ARCANOID/ball"
	"ARCANOID/block"
	"ARCANOID/paddel"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	paddel paddel.Paddel
	level  int
	ball   ball.Ball
	blocks []block.Block
}

func (g *Game) Update() error {
	g.paddel.Input()

	g.ball.Mouvement()

	g.ball.Coll(g.paddel, &g.blocks)

	if block.Win(g.blocks) {
		g.level++
		time.Sleep(1500 * time.Millisecond)
		block.LoadLevel(g.level, &g.blocks)
	}

	if g.ball.Lose() {
		panic("YOU LOSE")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.paddel.Draw(screen)

	g.ball.Draw(screen)

	for _, b := range g.blocks {
		b.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return paddel.ScreenWidth, paddel.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ARCANOID!!!")

	g := &Game{
		level: 1,
	}

	g.paddel.Init()

	println("Paddel init|check")

	g.ball.Init()

	println("Ball init|check")

	block.LoadLevel(g.level, &g.blocks)

	println("Level load|check")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
