package streamdeck

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
)

// TextButton .
type TextButton struct {
	Lines   []TextLine
	BgColor color.Color
}

// TextLine holds the content of one text line.
type TextLine struct {
	Text      string
	PosX      int
	PosY      int
	Font      *truetype.Font
	FontSize  float64
	FontColor color.Color
}

func fillLines(img *image.RGBA, lines []TextLine) error {
	for i, line := range lines {
		fontColor := image.NewUniform(line.FontColor)
		c := freetype.NewContext()
		c.SetDPI(72)
		c.SetFont(line.Font)
		c.SetFontSize(line.FontSize)
		c.SetClip(img.Bounds())
		c.SetDst(img)
		c.SetSrc(fontColor)
		pt := freetype.Pt(line.PosX, img.Bounds().Max.Y-line.PosY)
		// pt := freetype.Pt(line.PosX, line.PosY+int(c.PointToFixed(line.FontSize)>>6))
		if _, err := c.DrawString(line.Text, pt); err != nil {
			return errors.Wrapf(err, "unable to write line %d", i)
		}
	}
	return nil
}

// FillText .
func (sd *StreamDeck) FillText(x int, y int, text TextButton) error {
	img := image.NewRGBA(image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()))
	bg := image.NewUniform(text.BgColor)
	draw.Draw(img, img.Bounds(), bg, image.Point{0, 0}, draw.Src)

	err := fillLines(img, text.Lines)
	if err != nil {
		return errors.Wrap(err, "unable to write lines")
	}

	return sd.FillImage(x, y, img)
}
