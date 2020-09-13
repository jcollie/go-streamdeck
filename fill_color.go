package streamdeck

import (
	"image"
	"image/color"
	"image/draw"
)

func fillColor(sd Device, x int, y int, fgcolor color.Color) error {
	if err := checkValidButtonXY(sd, x, y); err != nil {
		return err
	}

	img := image.NewRGBA(image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()))
	draw.Draw(img, img.Bounds(), image.NewUniform(fgcolor), image.Point{0, 0}, draw.Src)

	return sd.writeImage(x, y, img)
}

// FillColor fills the button with a solid color.
func (sd *V2) FillColor(x int, y int, fgcolor color.Color) error {
	return fillColor(sd, x, y, fgcolor)
}
