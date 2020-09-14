package streamdeck

import (
	"fmt"

	"github.com/karalabe/hid"
	"github.com/pkg/errors"
)

func enumerateStreamDecks() []hid.DeviceInfo {
	allDeviceInfos := hid.Enumerate(VendorID, 0)
	streamDeckDeviceInfos := []hid.DeviceInfo{}

	for _, deviceInfo := range allDeviceInfos {
		if streamDeckInfo, ok := ProductIDs[deviceInfo.ProductID]; ok {
			streamDeckDeviceInfos = append(streamDeckDeviceInfos, deviceInfo)
			fmt.Printf("Detected %s with serial number %s\n", streamDeckInfo.Description, deviceInfo.Serial)
		}
	}
	return streamDeckDeviceInfos
}

func initializeStreamDeck(deviceInfo hid.DeviceInfo) (*StreamDeck, error) {
	device, err := deviceInfo.Open()
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to open Stream Deck %s", deviceInfo.Serial)
	}

	switch device.ProductID {
	case OriginalProductID, MiniProductID, OriginalV2ProductID, XLProductID:
		sd := new(StreamDeck)
		sd.device = device
		sd.previousState = make([]ButtonState, sd.NumberOfButtons())
		sd.buttonCallbacks = make([]ButtonCallback, sd.NumberOfButtons())
		return sd, nil
	default:
		return nil, errors.Errorf("not implemented for product id %x", device.ProductID)
	}
}

// New .
func New() (*StreamDeck, error) {
	deviceInfos := enumerateStreamDecks()

	if len(deviceInfos) == 0 {
		return nil, errors.Errorf("No Stream Decks found")
	}

	return initializeStreamDeck(deviceInfos[0])
}

// NewWithSerial .
func NewWithSerial(serial string) (*StreamDeck, error) {
	deviceInfos := enumerateStreamDecks()

	if len(deviceInfos) == 0 {
		return nil, errors.Errorf("no StreamDecks found")
	}

	for _, deviceInfo := range deviceInfos {
		if deviceInfo.Serial == serial {
			return initializeStreamDeck(deviceInfo)
		}
	}

	return nil, errors.Errorf("Unable to locate Stream Deck with serial %s", serial)
}
