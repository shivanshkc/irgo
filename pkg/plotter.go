package pkg

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
)

// vec2 represents a 2D vector or any pair of values.
type vec2 struct {
	x, y float64
}

// line represents a line between 2 points.
type line struct {
	start, end vec2
}

// Plotter is responsible to tokenize the given number into lines.
type Plotter struct {
	num      string
	idx      int
	angleSum int64
	cursor   vec2

	lines []line
}

// NewPlotter returns a new Plotter object.
func NewPlotter(num string) *Plotter {
	return &Plotter{num: num}
}

// AddLine executes one step of tokenization and so it adds one new line to the list of lines.
func (p *Plotter) AddLine() error {
	// Do nothing if no digits are left.
	if p.idx+3 > len(p.num) {
		return nil
	}

	// Get the angle to plot.
	angleString := p.num[p.idx : p.idx+3]
	angleDeg, err := strconv.Atoi(angleString)
	if err != nil {
		return fmt.Errorf("error in strconv.Atoi call: %w", err)
	}

	// For a "fibonacci" plot.
	p.angleSum += int64(angleDeg)
	var lineLength = 10.0

	// Convert to radians to use with trig functions.
	angleRad := float64(p.angleSum%360) * math.Pi / 180.0
	// End coordinates of the new line.
	endX := p.cursor.x + lineLength*math.Cos(angleRad)
	endY := p.cursor.y + lineLength*math.Sin(-angleRad)

	// Add the line.
	p.lines = append(p.lines, line{
		start: vec2{x: p.cursor.x, y: p.cursor.y},
		end:   vec2{x: endX, y: endY},
	})

	// Update the cursor for the next line.
	p.cursor.x, p.cursor.y = endX, endY
	// Update the index for the next 3 digits.
	p.idx += 3

	return nil
}

// Lines returns the list of lines.
func (p *Plotter) Lines() []line { return p.lines }

// LineColor returns appropriate color for the given line.
func (p *Plotter) LineColor(l line) color.Color { return color.White }

// Progress returns how many digits and what percent of the number has been tokenized.
func (p *Plotter) Progress() (int, float64) {
	return p.idx + 1, 100 * float64(p.idx) / float64(len(p.num))
}
