package streamdeck

import (
	"fmt"

	"github.com/karalabe/hid"
)

// V2 .
type V2 struct {
	StreamDeck
	device *hid.Device
}

// ImageReportHeaderLength .
func (sd *V2) ImageReportHeaderLength() int {
	return 8
}

// ImageReportPayloadLength .
func (sd *V2) ImageReportPayloadLength() int {
	return sd.ImageReportLength() - sd.ImageReportHeaderLength()
}

// ImageReportLength .
func (sd *V2) ImageReportLength() int {
	return 1024
}

// ImageWidth .
func (sd *V2) ImageWidth() int {
	return 72
}

// ImageHeight .
func (sd *V2) ImageHeight() int {
	return 72
}

// ImageFlipHorizontal .
func (sd *V2) ImageFlipHorizontal() bool {
	return true
}

// ImageFlipVertical .
func (sd *V2) ImageFlipVertical() bool {
	return true
}

// ImageRotation .
func (sd *V2) ImageRotation() ImageRotation {
	return ImageRotation0
}

// ImageFormat .
func (sd *V2) ImageFormat() ImageFormat {
	return ImageFormatJPEG
}

// NumberOfColumns returns the number of columns that the device supports
func (sd *V2) NumberOfColumns() int {
	switch sd.device.ProductID {
	case OriginalV2ProductID:
		return 5
	case XLProductID:
		return 8
	default:
		panic(fmt.Sprintf("not defined for product id %x", sd.device.ProductID))
	}
}

// NumberOfRows returns the number of rows that the device supports
func (sd *V2) NumberOfRows() int {
	switch sd.device.ProductID {
	case OriginalV2ProductID:
		return 3
	case XLProductID:
		return 4
	default:
		panic(fmt.Sprintf("not defined for product id %x", sd.device.ProductID))
	}
}

// NumberOfButtons returns the number of buttons that the device supports
func (sd *V2) NumberOfButtons() int {
	return sd.NumberOfColumns() * sd.NumberOfRows()
}
