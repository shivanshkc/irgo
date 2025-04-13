package main

import (
	"irgo/pkg"

	"github.com/hajimehoshi/ebiten/v2"

	_ "embed"
)

var screenSize = [2]float64{1920, 1080}

//go:embed nums/phi-100k.txt
var phi string

//go:embed nums/pi-100k.txt
var pi string

func main() {
	ebiten.SetWindowSize(int(screenSize[0]), int(screenSize[1]))
	ebiten.SetWindowTitle("Irrational Number Visualization - Golden Ratio")

	if err := ebiten.RunGame(pkg.NewIrrationalPlotter("Phi", phi, screenSize)); err != nil {
		panic("error in ebiten.RunGame: " + err.Error())
	}
}
