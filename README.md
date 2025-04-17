# Irgo

**Irgo** is a fun and visually captivating Go program that generates beautiful visualizations of irrational numbers using angles and line drawings. It uses [Ebiten](https://ebiten.org/), a powerful game library in Go, to render animated and interactive line plots.

## 🌀 What Does It Do?

Given a file containing a long string of digits (typically from an irrational number like π or φ), Irgo:

1. **Reads the digits in chunks of three** (e.g., "123", "456", "789", ...).
2. **Draws a line** with an angle corresponding to the current 3-digit number.
3. **Adds the previous angle** to the current one to keep turning the drawing.
4. **Repeats** this process until it runs out of digits.
5. **Renders the result** as a continuously turning polyline that shows the “shape” of the number.

The result is a mesmerizing, abstract pattern unique to each number.

## 📸 Showcase

Here are a few visualizations generated using the digits of the golden ratio:

<p float="left">
  <img src="showcase/Golden%20Ratio%20Random%201.png" width="45%" alt="Golden Ratio Random 1">
  <img src="showcase/Golden%20Ratio%20Random%202.png" width="45%" alt="Golden Ratio Random 2">
</p>

## 📦 How to Run

### 1. Prerequisites

- Go 1.18 or later
- A file with a long string of digits in the range `[0-9]` (included in `nums/`)

### 2. Install Dependencies

If you're new to Ebiten, install it via:

```bash
go get github.com/hajimehoshi/ebiten/v2
```

### 3. Run Irgo

```bash
go run cmd/irgo/main.go <path-to-digit-file>
```

Example:

```bash
go run cmd/irgo/main.go nums/phi-100k.txt
```

This will open an interactive window where the visualization is drawn in real time.

## 🖱️ Controls

- **Zoom In:** `+` key
- **Zoom Out:** `-` key
- **Pan:** Click and drag with the mouse

## 📁 Directory Structure

```
irgo/
├── cmd/irgo/          # Main entrypoint
├── pkg/               # Core logic and rendering
├── nums/              # Sample digit files (e.g. pi, phi)
├── showcase/          # Example output images
├── go.mod / go.sum    # Module definitions
```

## 💡 Why Irrational Numbers?

Irrational numbers like π, φ, and √2, by definition, never end and never repeat.
This makes them perfect for generating infinite, non-repeating visual patterns.

## 📜 License

MIT License. See the [LICENSE](LICENSE) file for more details.