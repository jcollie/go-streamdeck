package streamdeck

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"github.com/srwiley/rasterx"
)

// IconButton .
type IconButton struct {
	Icon            string
	IconColor       color.Color
	BackgroundColor color.Color
}

// FillIcon fills the button with an icon drawn from the Material Design Icons found at https://materialdesignicons.com/.
func (sd *StreamDeck) FillIcon(x int, y int, icon IconButton) error {

	img, err := getIcon(sd, icon.Icon, icon.IconColor, icon.BackgroundColor)
	if err != nil {
		return errors.Wrap(err, "unable to get icon image")
	}

	return sd.FillImage(x, y, img)
}

// IconTextButton .
type IconTextButton struct {
	Icon            string
	IconColor       color.Color
	Text            string
	FontSize        float64
	FontColor       color.Color
	Font            *truetype.Font
	BackgroundColor color.Color
}

func fillIconText(sd *StreamDeck, x int, y int, icon IconTextButton) error {
	svg, err := getSvgIcon(icon.Icon)
	if err != nil {
		return errors.Wrapf(err, "unable to get icon '%s'", icon)
	}

	img := image.NewRGBA(image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()))
	backgroundColor := image.NewUniform(icon.BackgroundColor)
	iconColor := image.NewUniform(icon.IconColor)
	fontColor := image.NewUniform(icon.FontColor)

	draw.Draw(img, img.Bounds(), backgroundColor, img.Bounds().Min, draw.Src)

	r := calculateTextBounds(icon.Text, icon.FontSize, icon.Font)
	shrinkage := r.Bounds().Dy()

	mask := image.NewRGBA(img.Bounds())

	svg.SetTarget(
		float64(mask.Bounds().Min.X+(shrinkage/2)),
		float64(mask.Bounds().Min.Y),
		float64(mask.Bounds().Max.X-(shrinkage/2)),
		float64(mask.Bounds().Max.Y-shrinkage))

	svg.Draw(
		rasterx.NewDasher(
			mask.Bounds().Dx(),
			mask.Bounds().Dy(),
			rasterx.NewScannerGV(
				mask.Bounds().Dx(),
				mask.Bounds().Dy(),
				mask,
				mask.Bounds(),
			),
		),
		1.0,
	)

	draw.DrawMask(img, img.Bounds(), iconColor, img.Bounds().Min, mask, img.Bounds().Min, draw.Over)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(icon.Font)
	c.SetFontSize(icon.FontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fontColor)

	pt := freetype.Pt(
		(img.Bounds().Dx()-r.Bounds().Dx())/2+r.Min.X,
		img.Bounds().Max.Y-r.Bounds().Max.Y,
	)

	if _, err := c.DrawString(icon.Text, pt); err != nil {
		return errors.Wrapf(err, "unable to write line")
	}

	return sd.FillImage(x, y, img)
}

// FillIconText fills the button with an icon drawn from the Material Design Icons found at https://materialdesignicons.com/.
func (sd *StreamDeck) FillIconText(x int, y int, icon IconTextButton) error {
	return fillIconText(sd, x, y, icon)
}

func calculateTextBounds(text string, fontSize float64, font *truetype.Font) image.Rectangle {
	test := image.NewRGBA(image.Rect(0, 0, 1024, 1024))

	bg := image.NewUniform(color.White)
	fg := image.NewUniform(color.Black)

	draw.Draw(test, test.Bounds(), bg, image.Point{0, 0}, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(test.Bounds())
	c.SetDst(test)
	c.SetSrc(fg)

	c.DrawString(text, freetype.Pt(32, 32))

	minX := test.Bounds().Max.X
	minY := test.Bounds().Max.Y
	maxX := test.Bounds().Min.X
	maxY := test.Bounds().Min.Y

	for x := test.Bounds().Min.X; x < test.Bounds().Max.X; x++ {
		for y := test.Bounds().Min.Y; y < test.Bounds().Max.Y; y++ {
			if r, g, b, _ := test.At(x, y).RGBA(); !(r == 0xffff && g == 0xffff && b == 0xffff) {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	minX = minX - 32
	minY = minY - 32
	maxX = maxX - 32
	maxY = maxY - 32

	return image.Rect(minX, minY, maxX, maxY)
}
