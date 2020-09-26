package streamdeck

import (
	"image"
	"log"
	"time"

	"github.com/pkg/errors"
)

func (sd *StreamDeck) writeImage(x int, y int, img image.Image) error {

	buttonIndex, err := sd.convertXYToButtonIndex(x, y)
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

func (sd *StreamDeck) writeImageData(buttonIndex int, data []byte) error {
	switch sd.device.ProductID {
	case OriginalV2ProductID:
		return sd.writeImageDataV2(buttonIndex, data)
	default:
		panic("not implemented")
	}
}

func (sd *StreamDeck) writeImageDataV1(buttonIndex int, data []byte) error {
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
		header := make([]byte, 16)
		header[0] = 0x02
		header[1] = 0x01
		header[2] = byte(pageNumber + 1)
		header[4] = byte(last)
		header[5] = byte(buttonIndex + 1)

		payload := append(header, page...)
		padding := make([]byte, sd.ImageReportLength()-len(payload))
		payload = append(payload, padding...)
		_, err := sd.device.Write(payload)
		if err != nil {
			return errors.Wrapf(err, "unable to send page: %d", pageNumber)
		}
		pageNumber++
	}
	log.Printf("write image: %s", time.Since(startTime))
	return nil
}

func (sd *StreamDeck) writeImageDataV2(buttonIndex int, data []byte) error {
	var pageNumber int = 0

	log.Printf("locking by writeImageDataV2")
	sd.Lock()
	log.Printf("locked by writeImageDataV2")
	defer func() {
		log.Printf("unlocked by writeImageDataV2")
		sd.Unlock()
	}()

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
			log.Printf("unable to send page %d: %+v", pageNumber, err)
			return errors.Wrapf(err, "unable to send page: %d", pageNumber)
		}
		pageNumber++
	}
	log.Printf("wrote image in %s", time.Since(startTime))
	return nil
}
