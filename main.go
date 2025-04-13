package main

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	_ "embed"
)

const (
	screenWidth  = 1024.0
	screenHeight = 720.0
)

//go:embed nums/phi-100k.txt
var phi string

//go:embed nums/pi-100k.txt
var pi string

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Irrational Number Visualization - Golden Ratio")

	if err := ebiten.RunGame(NewIrrationalPlotter(phi)); err != nil {
		panic("error in ebiten.RunGame: " + err.Error())
	}
}

// NewIrrationalPlotter returns a new instance of the IrrationalPlotter.
//
// "irr" is the irrational number with decimal point removed.
func NewIrrationalPlotter(irr string) *IrrationalPlotter {
	return &IrrationalPlotter{
		irr:     irr,
		lineLen: 20,
		cursor:  [2]float64{screenWidth / 2, screenHeight / 2},
	}
}

// IrrationalPlotter implementes the ebiten.Game interface to plot irrational numbers.
type IrrationalPlotter struct {
	// irr is the irrational number that will be plotted.
	irr string
	// lineLen is the length of each line in the plot.
	lineLen float64

	// idx keeps track of how much of the number has been plotted.
	idx int64
	// cursor is the current cursor position. It is the starting position for the next line.
	// It is in the form of [x, y] coordinates.
	cursor [2]float64

	// lines is the list of lines to draw.
	lines []Line
}

// Line represents a single straight line.
type Line struct {
	p1, p2   [2]float64
	angleRad float64
}

func (i *IrrationalPlotter) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (i *IrrationalPlotter) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{20, 20, 20, 255})

	// Draw each line.
	for _, line := range i.lines {
		x1, y1, x2, y2 := float32(line.p1[0]), float32(line.p1[1]), float32(line.p2[0]), float32(line.p2[1])
		vector.StrokeLine(screen, x1, y1, x2, y2, 1, color.White, false)
	}
}

func (i *IrrationalPlotter) Update() error {
	// Early return if no digits are left to print.
	if i.idx+3 > int64(len(i.irr)) {
		fmt.Println("Number fully printed.")
		return nil
	}

	// Get the angle to plot.
	angleString := i.irr[i.idx : i.idx+3]
	angleDeg, err := strconv.Atoi(angleString)
	if err != nil {
		return fmt.Errorf("error in strconv.Atoi call: %w", err)
	}

	// Convert to radians to use with trig functions.
	angleRad := float64(angleDeg) * math.Pi / 180.0
	// End coordinates of the new line.
	endX := i.cursor[0] + i.lineLen*math.Cos(angleRad)
	endY := i.cursor[1] + i.lineLen*math.Sin(-angleRad)

	// Add the line.
	i.lines = append(i.lines, Line{
		p1:       [2]float64{i.cursor[0], i.cursor[1]},
		p2:       [2]float64{endX, endY},
		angleRad: angleRad,
	})

	// Update the cursor for the next line.
	i.cursor[0], i.cursor[1] = endX, endY
	// Update the index for the next 3 digits.
	i.idx += 3

	return nil
}
