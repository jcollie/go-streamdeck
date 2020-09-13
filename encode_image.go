package streamdeck

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/pkg/errors"
	"golang.org/x/image/bmp"
)

func encodeImage(sd Device, img image.Image) ([]byte, error) {
	rect := img.Bounds()

	if rect.Dy() != sd.ImageHeight() {
		return []byte{}, errors.Errorf("image height does not match the size of the button")
	}

	if rect.Dx() != sd.ImageWidth() {
		return []byte{}, errors.Errorf("image width does not match the size of the button")
	}

	var buf bytes.Buffer

	switch sd.ImageFormat() {
	case ImageFormatJPEG:
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 95}); err != nil {
			return []byte{}, errors.Wrap(err, "unable to encode jpeg")
		}
	case ImageFormatBMP:
		if err := bmp.Encode(&buf, img); err != nil {
			return []byte{}, errors.Wrap(err, "unable to encode bmp")
		}
	default:
		return []byte{}, errors.Errorf("unknown image encoding format")
	}

	return buf.Bytes(), nil
}

func (sd *V2) encodeImage(img image.Image) ([]byte, error) {
	return encodeImage(sd, img)
}
