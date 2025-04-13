package main

import (
	"irgo/pkg"
	"os"
	"path/filepath"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenW, screenH = 1024, 720
)

func main() {
	path := mustParseInputs()

	// Extract file name from the path (even the extension is excluded).
	// This will be used as the name of the window, save files etc.
	filename := filepath.Base(path)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	// Read the number that will be plotted.
	number := mustReadFile(path)

	// Initiate an Ebiten window.
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Irgo - " + filename)

	// Math objects.
	plotter := pkg.NewPlotter(number)
	game := pkg.NewGame(plotter, screenW, screenH)

	// Start Ebiten game.
	if err := ebiten.RunGame(game); err != nil {
		panic("error in ebiten.RunGame: " + err.Error())
	}
}

// mustParseInputs reads command line user inputs.
// It panics if it's not found or on any other error.
func mustParseInputs() string {
	if len(os.Args) < 2 {
		panic("Provide path to file that contains the number.")
	}

	return os.Args[1]
}

// mustReadFile reads the entire file present at the given path.
// It panics on error.
func mustReadFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		panic("error in os.ReadFile call: " + err.Error())
	}
	return string(data)
}
