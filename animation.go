package streamdeck

import (
	"context"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"time"

	"github.com/pkg/errors"
)

// Animation .
type Animation struct {
	sd     *StreamDeck
	frames []AnimationFrame
}

// AnimationFrame .
type AnimationFrame struct {
	frame []byte
	delay time.Duration
}

// NewAnimationFromGIF .
func (sd *StreamDeck) NewAnimationFromGIF(filename string) (*Animation, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open gif")
	}
	defer r.Close()

	g, err := gif.DecodeAll(r)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode gif")
	}

	frameCount := len(g.Image)

	animation := new(Animation)
	animation.sd = sd

	img := image.NewRGBA(g.Image[0].Rect)

	for i := 0; i < frameCount; i++ {
		nextImage := g.Image[i]
		nextDelay := time.Duration(g.Delay[i]) * (time.Second / 100)
		draw.Draw(
			img,
			nextImage.Rect,
			nextImage,
			nextImage.Rect.Min,
			draw.Over,
		)
		preparedImage := prepareImage(sd, img)
		encodedImage, err := sd.encodeImage(preparedImage)
		if err != nil {
			return nil, errors.Wrap(err, "unable to encode image")
		}
		animation.frames = append(animation.frames, AnimationFrame{frame: encodedImage, delay: nextDelay})
	}
	return animation, nil
}

// RunAnimation .
func (animation *Animation) RunAnimation(ctx context.Context, x int, y int) {
	buttonIndex, err := convertXYToButtonIndex(animation.sd, x, y)
	if err != nil {
		fmt.Printf("unable to convert x y: %+v\n", err)
		return
	}
	for {
		for _, frame := range animation.frames {
			animation.sd.writeImageData(buttonIndex, frame.frame)
			select {
			case <-ctx.Done():
				return
			case <-time.After(frame.delay):
			}
		}
	}
}
