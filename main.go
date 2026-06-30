package main

import (
	"ARCANOID/ball"
	"ARCANOID/block"
	debugmode "ARCANOID/debug_mode"
	"ARCANOID/paddel"
	"ARCANOID/power_up"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	paddel    paddel.Paddel
	balls     []ball.Ball
	power_ups []power_up.PowerUp
	blocks    []block.Block
	modeDebug bool
	level     int
}

var motTapé = ""

func (g *Game) Update() error {
	debugmode.CheckWordTyped(&motTapé)
	debugmode.DoDebug(motTapé, &g.paddel.X, g.paddel.W, g.balls[0].X, g.paddel.Speed)

	g.paddel.Input()

	for i := range g.balls {
		ball := &g.balls[i]

		ball.Mouvement()

		ball.Coll(g.paddel, &g.blocks, &g.power_ups)
	}

	for i := range g.power_ups {
		p := &g.power_ups[i]

		p.Mouv()
	}

	switch power_up.Coll(g.paddel, &g.power_ups) {
	case "+1 ball":
		ball.Init(&g.balls)
	}

	if block.LevelIsEnd(g.blocks) {
		g.level++
		block.LoadLevel(g.level, &g.blocks)
		ball.ResetBallsForTheNextLevel(&g.balls)
	}

	if ball.Lose(g.balls) {
		panic("YOU LOSE")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	g.paddel.Draw(screen)

	for _, ball := range g.balls {
		ball.Draw(screen)
	}

	for _, b := range g.blocks {
		b.Draw(screen)
	}

	for _, p := range g.power_ups {
		p.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return paddel.ScreenWidth, paddel.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("ARCANOID!!!")

	g := &Game{
		level: 2,
	}

	g.paddel.Init()

	println("Paddel init|check")

	ball.Init(&g.balls)

	println("Ball init|check")

	block.LoadLevel(g.level, &g.blocks)

	power_up.InitImgs()

	println("Level load|check")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
