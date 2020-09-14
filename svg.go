package streamdeck

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var cache = map[string]*oksvg.SvgIcon{}

func init() {
	cache = make(map[string]*oksvg.SvgIcon)
}

func getSvgIcon(icon string) (*oksvg.SvgIcon, error) {
	startTime := time.Now()
	svg, found := cache[icon]
	if !found {
		var u = url.URL{
			Scheme: "https",
			Host:   "raw.githubusercontent.com",
			Path:   fmt.Sprintf("/Templarian/MaterialDesign/master/svg/%s.svg", icon),
		}
		response, err := http.Get(u.String())
		if err != nil {
			return nil, errors.Wrap(err, "unable to download icon")
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return nil, errors.Errorf("error %d from github", response.StatusCode)
		}

		svg, err = oksvg.ReadIconStream(response.Body)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read icon: %+v")
		}

		cache[icon] = svg
	}
	log.Printf("time getting icon %s: %s", icon, time.Since(startTime))
	return svg, nil
}

func getIcon(sd *StreamDeck, icon string, fgcolor color.Color, bgcolor color.Color) (image.Image, error) {
	var svg *oksvg.SvgIcon

	svg, err := getSvgIcon(icon)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get icon '%s'", icon)
	}

	img := image.NewRGBA(image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()))
	bg := image.NewUniform(bgcolor)
	draw.Draw(img, img.Bounds(), bg, image.Point{0, 0}, draw.Src)

	mask := image.NewRGBA(image.Rect(0, 0, sd.ImageWidth(), sd.ImageHeight()))
	svg.SetTarget(0.0, 0.0, float64(sd.ImageWidth()), float64(sd.ImageHeight()))
	svg.Draw(
		rasterx.NewDasher(
			sd.ImageWidth(),
			sd.ImageHeight(),
			rasterx.NewScannerGV(
				sd.ImageWidth(),
				sd.ImageHeight(),
				mask,
				mask.Bounds(),
			),
		),
		1.0,
	)

	fg := image.NewUniform(fgcolor)
	draw.DrawMask(img, img.Bounds(), fg, image.Point{0, 0}, mask, image.Point{0, 0}, draw.Over)

	return img, nil
}
