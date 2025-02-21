// Package bunnyhue provides preconfigured color palettes for dark and light themed sites.
package bunnyhue

import "image/color"

// Palette represents a collection of standard colors for a theme.
type Palette struct {
	Name string

	Background color.Color // Background color
	Contrast   color.Color // Color for contrasting elements
	Primary    color.Color // Primary color
	Secondary  color.Color // Secondary color

	Blue   color.Color
	Red    color.Color
	Green  color.Color
	Pink   color.Color
	Purple color.Color
	Cyan   color.Color
	Orange color.Color
	Teal   color.Color
	Brown  color.Color
	Lime   color.Color
}

// Dark is the palette for a dark themed site.
var Dark = Palette{
	Name: "Dark",

	Background: color.RGBA{R: 18, G: 18, B: 18, A: 255},
	Contrast:   color.RGBA{R: 30, G: 30, B: 30, A: 255},
	Primary:    color.RGBA{R: 224, G: 224, B: 224, A: 255},
	Secondary:  color.RGBA{R: 218, G: 218, B: 218, A: 255},

	Blue:   color.RGBA{R: 30, G: 144, B: 255, A: 255},
	Red:    color.RGBA{R: 255, G: 69, B: 0, A: 255},
	Green:  color.RGBA{R: 50, G: 205, B: 50, A: 255},
	Pink:   color.RGBA{R: 218, G: 112, B: 214, A: 255},
	Purple: color.RGBA{R: 147, G: 112, B: 219, A: 255},
	Cyan:   color.RGBA{R: 0, G: 255, B: 255, A: 255},
	Orange: color.RGBA{R: 255, G: 140, B: 0, A: 255},
	Teal:   color.RGBA{R: 0, G: 128, B: 128, A: 255},
	Brown:  color.RGBA{R: 139, G: 69, B: 19, A: 255},
	Lime:   color.RGBA{R: 191, G: 255, B: 0, A: 255},
}

// Light is the palette for a light themed site.
var Light = Palette{
	Name: "Light",

	Background: color.RGBA{R: 245, G: 245, B: 245, A: 255},
	Contrast:   color.RGBA{R: 235, G: 235, B: 235, A: 255},
	Primary:    color.RGBA{R: 20, G: 20, B: 20, A: 255},
	Secondary:  color.RGBA{R: 26, G: 26, B: 26, A: 255},

	Blue:   color.RGBA{R: 65, G: 105, B: 225, A: 255},
	Red:    color.RGBA{R: 220, G: 20, B: 60, A: 255},
	Green:  color.RGBA{R: 34, G: 139, B: 34, A: 255},
	Pink:   color.RGBA{R: 255, G: 105, B: 180, A: 255},
	Purple: color.RGBA{R: 138, G: 43, B: 226, A: 255},
	Cyan:   color.RGBA{R: 0, G: 206, B: 209, A: 255},
	Orange: color.RGBA{R: 255, G: 165, B: 0, A: 255},
	Teal:   color.RGBA{R: 32, G: 178, B: 170, A: 255},
	Brown:  color.RGBA{R: 160, G: 82, B: 45, A: 255},
	Lime:   color.RGBA{R: 50, G: 205, B: 50, A: 255},
}
