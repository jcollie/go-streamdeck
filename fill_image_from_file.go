package streamdeck

import (
	"image"
	_ "image/gif"  // needed to register image format
	_ "image/jpeg" // needed to register image format
	_ "image/png"  // needed to register image format
	"os"

	_ "golang.org/x/image/bmp"  // needed to register image format
	_ "golang.org/x/image/webp" // needed to register image format

	"github.com/pkg/errors"
)

func fillImageFromFile(sd Device, x int, y int, path string) error {
	if err := checkValidButtonXY(sd, x, y); err != nil {
		return errors.Wrap(err, "unable to fill image")
	}
	reader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return errors.Wrap(err, "unable to decode image")
	}

	return sd.FillImage(x, y, img)
}

// FillImageFromFile fills the button with an image from a file.
func (sd *V2) FillImageFromFile(x int, y int, path string) error {
	return fillImageFromFile(sd, x, y, path)
}
