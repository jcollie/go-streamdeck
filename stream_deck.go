package streamdeck

import (
	"fmt"
	"sync"
	"time"

	"github.com/karalabe/hid"
)

// ButtonCallback .
type ButtonCallback interface {
	ButtonPressed(sd *StreamDeck, x int, y int, timestamp time.Time)
	ButtonReleased(sd *StreamDeck, x int, y int, timestamp time.Time)
}

// StreamDeck .
type StreamDeck struct {
	sync.Mutex
	device          *hid.Device
	previousState   []ButtonState
	buttonCallbacks []ButtonCallback
}

// Description .
func (sd *StreamDeck) Description() string {
	switch sd.device.ProductID {
	case OriginalProductID:
		return "Stream Deck Original"
	case MiniProductID:
		return "Stream Deck Mini"
	case XLProductID:
		return "Stream Deck XL"
	case OriginalV2ProductID:
		return "Stream Deck Original (V2)"
	default:
		panic(fmt.Sprintf("not implemented for product id %x", sd.device.ProductID))
	}
}

// ImageReportHeaderLength .
func (sd *StreamDeck) ImageReportHeaderLength() int {
	switch sd.device.ProductID {
	case OriginalProductID:
		return 16
	case OriginalV2ProductID, XLProductID:
		return 8
	default:
		panic("not implemented")
	}
}

// ImageReportPayloadLength .
func (sd *StreamDeck) ImageReportPayloadLength() int {
	return sd.ImageReportLength() - sd.ImageReportHeaderLength()
}

// ImageReportLength .
func (sd *StreamDeck) ImageReportLength() int {
	switch sd.device.ProductID {
	case OriginalProductID:
		return 8191
	case OriginalV2ProductID, XLProductID:
		return 1024
	default:
		panic("not implemented")
	}
}

// ImageWidth .
func (sd *StreamDeck) ImageWidth() int {
	return 72
}

// ImageHeight .
func (sd *StreamDeck) ImageHeight() int {
	return 72
}

// ImageFlipHorizontal .
func (sd *StreamDeck) ImageFlipHorizontal() bool {
	return true
}

// ImageFlipVertical .
func (sd *StreamDeck) ImageFlipVertical() bool {
	return true
}

// ImageRotation .
func (sd *StreamDeck) ImageRotation() ImageRotation {
	return ImageRotation0
}

// ImageFormat .
func (sd *StreamDeck) ImageFormat() ImageFormat {
	switch sd.device.ProductID {
	case OriginalProductID:
		return ImageFormatBMP
	case OriginalV2ProductID:
		return ImageFormatJPEG
	default:
		panic("not implemented")
	}
}

// NumberOfColumns returns the number of columns that the device supports
func (sd *StreamDeck) NumberOfColumns() int {
	switch sd.device.ProductID {
	case OriginalProductID:
		return 5
	case MiniProductID:
		return 3
	case OriginalV2ProductID:
		return 5
	case XLProductID:
		return 8
	default:
		panic(fmt.Sprintf("not defined for product id %x", sd.device.ProductID))
	}
}

// NumberOfRows returns the number of rows that the device supports
func (sd *StreamDeck) NumberOfRows() int {
	switch sd.device.ProductID {
	case OriginalProductID:
		return 3
	case MiniProductID:
		return 2
	case OriginalV2ProductID:
		return 3
	case XLProductID:
		return 4
	default:
		panic(fmt.Sprintf("not defined for product id %x", sd.device.ProductID))
	}
}

// NumberOfButtons returns the number of buttons that the device supports
func (sd *StreamDeck) NumberOfButtons() int {
	return sd.NumberOfColumns() * sd.NumberOfRows()
}
