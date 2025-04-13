package pkg

import (
	"fmt"
	"image/color"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// NewIrrationalPlotter returns a new instance of the IrrationalPlotter.
//
// "irr" is the irrational number with decimal point removed.
func NewIrrationalPlotter(irr string, screenSize [2]float64) *IrrationalPlotter {
	return &IrrationalPlotter{
		irr:        irr,
		lineLen:    10,
		screenSize: screenSize,
		cursor:     [2]float64{screenSize[0] / 2, screenSize[1] / 2},
		zoom:       1,
	}
}

// IrrationalPlotter implementes the ebiten.Game interface to plot irrational numbers.
type IrrationalPlotter struct {
	// irr is the irrational number that will be plotted.
	irr string
	// lineLen is the length of each line in the plot.
	lineLen float64
	// screenSize is width and height of the screen.
	screenSize [2]float64

	// idx keeps track of how much of the number has been plotted.
	idx int64
	// cursor is the current cursor position. It is the starting position for the next line.
	// It is in the form of [x, y] coordinates.
	cursor [2]float64

	// Zoom level.
	zoom float64
	// For Panning.
	offset [2]float64
	// For tracking mouse drag.
	dragStart  [2]int
	isDragging bool

	// lines is the list of lines to draw.
	lines []Line
}

// Line represents a single straight line.
type Line struct {
	p1, p2   [2]float64
	angleRad float64
}

func (i *IrrationalPlotter) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return int(i.screenSize[0]), int(i.screenSize[1])
}

func (i *IrrationalPlotter) Draw(screen *ebiten.Image) {
	// Clear the screen
	screen.Fill(color.RGBA{20, 20, 20, 255})

	// Draw each line.
	for _, line := range i.lines {
		x1, y1 := i.worldToScreen(line.p1[0], line.p1[1])
		x2, y2 := i.worldToScreen(line.p2[0], line.p2[1])
		vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, color.White, false)
	}
}

func (i *IrrationalPlotter) Update() error {
	// Handle user inputs.
	i.handleZoom()
	i.handleDrag()

	// Early return if no digits are left to print.
	if i.idx+3 > int64(len(i.irr)) {
		fmt.Println("Number fully plotted.")
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

// handleZoom registers and handles the zoom input.
func (i *IrrationalPlotter) handleZoom() {
	// Zoom-in control.
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) {
		i.zoom *= 1.2
	}

	// Zoom-out control.
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) {
		// Zoom out
		i.zoom *= 0.8
	}
}

// handleDrag registers and handles the mouse drag input.
func (i *IrrationalPlotter) handleDrag() {
	// Detect the beginning of mouse drag.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i.dragStart[0] = x
		i.dragStart[1] = y
		i.isDragging = true
	}

	// If the mouse is not being dragged, nothing to do.
	if !i.isDragging {
		return
	}

	// Mark dragging as false when the mouse stops.
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.isDragging = false
		return
	}

	// Mouse is being dragged, update view.
	x, y := ebiten.CursorPosition()
	i.offset[0] += float64(x-i.dragStart[0]) / i.zoom
	i.offset[1] += float64(y-i.dragStart[1]) / i.zoom
	i.dragStart[0] = x
	i.dragStart[1] = y
}

// worldToScreen transforms the given point as per the offset and zoom levels.
func (i *IrrationalPlotter) worldToScreen(x, y float64) (float64, float64) {
	x -= i.screenSize[0] / 2
	y -= i.screenSize[1] / 2

	screenX := (x+i.offset[0])*i.zoom + i.screenSize[0]/2
	screenY := (y+i.offset[1])*i.zoom + i.screenSize[1]/2
	return screenX, screenY
}
