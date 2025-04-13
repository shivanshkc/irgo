package pkg

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// Helper function to save Ebiten screen as image.
func saveImage(img *ebiten.Image, path string) error {
	// Create a new RGBA image.
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	// Create Go's image object that can be saved to a file.
	goImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Draw the Ebiten image to the Go image.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			goImage.Set(x, y, img.At(x, y))
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
