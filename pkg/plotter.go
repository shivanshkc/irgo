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

// lineListNode represents a node in a linked-list of lines.
type lineListNode struct {
	start, end vec2
	angleRad   float64
	next       *lineListNode
}

// Plotter is responsible to tokenize the given number into lines.
type Plotter struct {
	num string

	lastIdx int
	lastAng int
	lastPos vec2

	lineListHead *lineListNode
	lineListTail *lineListNode
}

// NewPlotter returns a new Plotter object.
func NewPlotter(num string) *Plotter {
	return &Plotter{num: num}
}

// tokenize3 takes the next 3 digits of the number, assumes the number formed by those 3 digits to be a degree angle,
// adds that angle to "lastAng", converts "lastAng" to radians, marks those three digits as used up and returns the
// calculated radian angle.
//
// It returns -1 if less than 3 digits are left in the number.
func (p *Plotter) tokenize3() (float64, error) {
	// -1 should be detected by the caller as the end of tokenization.
	if p.lastIdx+3 > len(p.num) {
		return -1, nil
	}

	// Pick the next 3 digits.
	angleString := p.num[p.lastIdx : p.lastIdx+3]
	angleDeg, err := strconv.Atoi(angleString)
	if err != nil {
		return 0, fmt.Errorf("error in strconv.Atoi call: %w", err)
	}

	// Add and condense the angle.
	p.lastAng += angleDeg
	p.lastAng = p.lastAng % 360

	// Mark the digits as used up.
	p.lastIdx += 3

	// Convert to radians to use with trig functions.
	return float64(p.lastAng) * math.Pi / 180.0, nil
}

// tokenize1 takes the next digit of the number, maps that digit [0-9] to a degree angle [0-360], adds that angle to
// "lastAng", converts "lastAng" to radians, marks that digit as used up and returns the calculated radian angle.
//
// It returns -1 if no digits are left to tokenize.
func (p *Plotter) tokenize1() (float64, error) {
	// -1 should be detected by the caller as the end of tokenization.
	if p.lastIdx+1 > len(p.num) {
		return -1, nil
	}

	// Pick the next digit.
	angleString := string(p.num[p.lastIdx])
	angleDeg, err := strconv.Atoi(angleString)
	if err != nil {
		return 0, fmt.Errorf("error in strconv.Atoi call: %w", err)
	}

	// Map [0-9] to [0-360]
	angleDeg *= 40

	// Add and condense the angle.
	p.lastAng += angleDeg
	p.lastAng = p.lastAng % 360

	// Mark the digits as used up.
	p.lastIdx += 1

	// Convert to radians to use with trig functions.
	return float64(p.lastAng) * math.Pi / 180.0, nil
}

// addLine uses the give angle and the lastPos to determine the next line in the sequence.
// It appends that new line to the list of lines and updates lastPos for the next line to come.
func (p *Plotter) addLine(angleRad float64) {
	// Line length does not affect this plot.
	// Bigger line lengths result in (effectively) zoomed in plots.
	var lineLength = 10.0

	// End coordinates of the new line.
	endX := p.lastPos.x + lineLength*math.Cos(angleRad)
	endY := p.lastPos.y + lineLength*math.Sin(-angleRad)

	nextLineNode := &lineListNode{
		start:    vec2{x: p.lastPos.x, y: p.lastPos.y},
		end:      vec2{x: endX, y: endY},
		angleRad: angleRad,
	}

	// Add to the linked list.
	if p.lineListHead == nil {
		p.lineListHead = nextLineNode
	} else if p.lineListTail == nil {
		p.lineListTail = nextLineNode
		p.lineListHead.next = p.lineListTail
	} else {
		p.lineListTail.next = nextLineNode
		p.lineListTail = nextLineNode
	}

	// Update the lastPos for the next line.
	p.lastPos.x, p.lastPos.y = endX, endY
}

// AddLine3 obtains the line angle using tokenize3 and adds that line to the line list.
func (p *Plotter) AddLine3() error {
	// Get the next angle to plot.
	angleRad, err := p.tokenize3()
	if err != nil {
		return fmt.Errorf("error in tokenize3 call: %w", err)
	}
	// -1 suggests that tokenization is complete. So, do nothing.
	if angleRad == -1 {
		return nil
	}

	// Add the line to the line list.
	p.addLine(angleRad)
	return nil
}

// AddLine1 obtains the line angle using tokenize1 and adds that line to the line list.
func (p *Plotter) AddLine1() error {
	// Get the next angle to plot.
	angleRad, err := p.tokenize1()
	if err != nil {
		return fmt.Errorf("error in tokenize1 call: %w", err)
	}
	// -1 suggests that tokenization is complete. So, do nothing.
	if angleRad == -1 {
		return nil
	}

	// Add the line to the line list.
	p.addLine(angleRad)
	return nil
}

// lineColor returns appropriate color for the given line.
func (p *Plotter) lineColor(l *lineListNode) color.Color {
	angleColor := uint8(l.angleRad * 255 * 0.5 / math.Pi)
	return color.RGBA{R: 255, G: angleColor, B: 255, A: 255}
}

// Progress returns how many digits and what percent of the number has been tokenized.
func (p *Plotter) Progress() (int, float64) {
	return p.lastIdx + 1, 100 * float64(p.lastIdx) / float64(len(p.num))
}
