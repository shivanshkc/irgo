package pkg

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// saveScreenAsImage saves the given screen as an image.
//
//nolint:unused // Will be used in the future so save high-quality screenshots.
func saveScreenAsImage(screen *ebiten.Image, path string) error {
	// Create a new RGBA image.
	bounds := screen.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	// Create Go's image object that can be saved to a file.
	goImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw the Ebiten image to the Go image.
	for y := range height {
		for x := range width {
			goImage.Set(x, y, screen.At(x, y))
		}
	}

	// Create a file for the image.
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error in os.Create call: %w", err)
	}
	defer f.Close()

	// Save as PNG.
	if err := png.Encode(f, goImage); err != nil {
		return fmt.Errorf("error in png.Encode call: %w", err)
	}

	return nil
}
