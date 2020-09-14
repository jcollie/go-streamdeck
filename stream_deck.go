package streamdeck

import (
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

// // Device .
// type Device interface {
// 	GetSerialNumber() (string, error)
// 	ResetToLogo() error
// 	ResetKeyStream() error
// 	ImageReportPayloadLength() int
// 	NumberOfButtons() int
// 	NumberOfColumns() int
// 	NumberOfRows() int
// 	ImageWidth() int
// 	ImageHeight() int
// 	ImageFlipHorizontal() bool
// 	ImageFlipVertical() bool
// 	ImageRotation() ImageRotation
// 	ImageFormat() ImageFormat
// 	FillImage(int, int, image.Image) error
// 	FillText(int, int, TextButton) error
// 	FillIcon(int, int, IconButton) error
// 	FillIconText(int, int, IconTextButton) error
// 	SetCallback(int, int, ButtonCallback) error
// 	Read()
// 	writeImage(int, int, image.Image) error
// 	encodeImage(image.Image) ([]byte, error)
// 	writeImageData(buttonIndex int, data []byte) error
// }
