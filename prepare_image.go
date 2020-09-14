package streamdeck

import (
	"image"

	"github.com/disintegration/imaging"
)

func prepareImage(sd *StreamDeck, img image.Image) image.Image {
	// first resize image height, preserving aspect rato
	rect := img.Bounds()
	if rect.Dy() != sd.ImageHeight() {
		img = imaging.Resize(img, 0, sd.ImageHeight(), imaging.Lanczos)
	}

	// then resize image width, preserving aspect ratio
	rect = img.Bounds()
	if rect.Dx() > sd.ImageWidth() {
		img = imaging.Resize(img, sd.ImageWidth(), 0, imaging.Lanczos)
	}

	// then center the image
	rect = img.Bounds()
	if rect.Dx() != sd.ImageWidth() || rect.Dy() != sd.ImageHeight() {
		img = imaging.OverlayCenter(
			image.NewRGBA(
				image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()),
			),
			img, 1.0)
	}

	// flip the image horizontally if needed
	if sd.ImageFlipHorizontal() {
		img = imaging.FlipH(img)
	}

	// flip the image vertically if needed
	if sd.ImageFlipVertical() {
		img = imaging.FlipV(img)
	}

	return img
}
