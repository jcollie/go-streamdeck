package streamdeck

import (
	"image"
	"log"
	"time"

	"github.com/pkg/errors"
)

func (sd *V2) writeImage(x int, y int, img image.Image) error {

	buttonIndex, err := convertXYToButtonIndex(sd, x, y)
	if err != nil {
		return errors.Wrap(err, "unable to write image")
	}

	data, err := sd.encodeImage(img)

	if err != nil {
		return errors.Wrap(err, "unable to encode image")
	}

	if err := sd.writeImageData(buttonIndex, data); err != nil {
		return errors.Wrap(err, "unable to write data")
	}

	return nil
}

func (sd *V2) writeImageData(buttonIndex int, data []byte) error {
	var pageNumber int = 0

	sd.Lock()
	defer sd.Unlock()
	startTime := time.Now()

	for len(data) > 0 {
		page := data[:min(len(data), sd.ImageReportPayloadLength())]
		data = data[len(page):]
		last := 0
		if len(data) == 0 {
			last = 1
		}
		header := []byte{
			0x02, 0x07, byte(buttonIndex), byte(last),
			byte(len(page) & 0xff), byte(len(page) >> 8),
			byte(pageNumber & 0xff), byte(pageNumber >> 8),
		}
		payload := append(header, page...)
		padding := make([]byte, sd.ImageReportLength()-len(payload))
		payload = append(payload, padding...)
		_, err := sd.device.Write(payload)
		if err != nil {
			return errors.Wrapf(err, "unable to send page: %d", pageNumber)
		}
		pageNumber++
	}
	endTime := time.Now()
	log.Printf("write image: %s", endTime.Sub(startTime))
	return nil
}
