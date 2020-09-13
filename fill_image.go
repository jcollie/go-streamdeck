package streamdeck

import (
	"image"

	"github.com/pkg/errors"
)

func fillImage(sd Device, x int, y int, img image.Image) error {
	if err := checkValidButtonXY(sd, x, y); err != nil {
		return err
	}

	img = prepareImage(sd, img)

	if err := sd.writeImage(x, y, img); err != nil {
		return errors.Wrap(err, "unable to write image")
	}

	return nil
}

// FillImage fills the given button with an image. For best performance, provide
// the image in the native size. Otherwise it will be automatically
// resized.
func (sd *V2) FillImage(x int, y int, img image.Image) error {
	return fillImage(sd, x, y, img)
}
