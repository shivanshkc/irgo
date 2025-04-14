package pkg

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Game implements ebiten.Game.
type Game struct {
	// External inputs.
	plotter    *Plotter
	screenSize vec2

	// User controls.
	panOffset  vec2
	zoom       float64
	dragStart  vec2
	isDragging bool
}

// NewGame returns a new instance of Game.
func NewGame(plotter *Plotter, screenW, screenH int) *Game {
	return &Game{
		plotter:    plotter,
		screenSize: vec2{x: float64(screenW), y: float64(screenH)},
		panOffset:  vec2{x: float64(screenW) / 2, y: float64(screenH) / 2},
		zoom:       1,
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return int(g.screenSize.x), int(g.screenSize.y)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen.
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255})

	for _, line := range g.plotter.Lines() {
		x1, y1 := g.worldToScreen(line.start.x, line.start.y)
		x2, y2 := g.worldToScreen(line.end.x, line.end.y)

		col := g.plotter.LineColor(line)
		vector.StrokeLine(screen, float32(x1), float32(y1), float32(x2), float32(y2), 1, col, false)
	}

	// Show progress.
	doneCount, donePercent := g.plotter.Progress()
	progressInfo := fmt.Sprintf("Progress: %d (%.2f%%)", doneCount, donePercent)
	ebitenutil.DebugPrintAt(screen, progressInfo, 0, 0)
}

func (g *Game) Update() error {
	// Handle user inputs.
	g.handleZoom()
	g.handleDrag()

	if err := g.plotter.AddLine3(); err != nil {
		return fmt.Errorf("error in AddLine3 call: %w", err)
	}
	return nil
}

// handleZoom registers and handles the zoom input.
func (g *Game) handleZoom() {
	// Zoom-in control.
	if inpututil.IsKeyJustPressed(ebiten.KeyEqual) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) {
		g.zoom *= 1.2
	}

	// Zoom-out control.
	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) || inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) {
		// Zoom out
		g.zoom *= 0.8
	}
}

// handleDrag registers and handles the mouse drag input.
func (g *Game) handleDrag() {
	// Detect the beginning of mouse drag.
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.dragStart.x = float64(x)
		g.dragStart.y = float64(y)
		g.isDragging = true
	}

	// If the mouse is not being dragged, nothing to do.
	if !g.isDragging {
		return
	}

	// Mark dragging as false when the mouse stops.
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.isDragging = false
		return
	}

	// Mouse is being dragged, update view.
	xi, yi := ebiten.CursorPosition()
	x, y := float64(xi), float64(yi)

	g.panOffset.x += float64(x-g.dragStart.x) / g.zoom
	g.panOffset.y += float64(y-g.dragStart.y) / g.zoom
	g.dragStart.x = x
	g.dragStart.y = y
}

// worldToScreen transforms the given point as per the offset and zoom levels.
func (g *Game) worldToScreen(x, y float64) (float64, float64) {
	x -= g.screenSize.x / 2
	y -= g.screenSize.y / 2

	screenX := (x+g.panOffset.x)*g.zoom + g.screenSize.x/2
	screenY := (y+g.panOffset.y)*g.zoom + g.screenSize.y/2

	return screenX, screenY
}
