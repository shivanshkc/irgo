package main

import (
	"irgo/pkg"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var screenSize = [2]float64{1920, 1080}

//go:embed nums/phi-1M.txt
var num string

func main() {
	var name = "Phi-1M"

	ebiten.SetWindowSize(int(screenSize[0]), int(screenSize[1]))
	ebiten.SetWindowTitle("Irrational Number Visualization - " + name)

	if err := ebiten.RunGame(pkg.NewIrrationalPlotter(name, num, screenSize)); err != nil {
		panic("error in ebiten.RunGame: " + err.Error())
	}
}
