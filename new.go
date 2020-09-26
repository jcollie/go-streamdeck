package streamdeck

import (
	"fmt"
	"log"
	"time"

	"github.com/karalabe/hid"
	"github.com/pkg/errors"
)

func enumerateStreamDecks() []hid.DeviceInfo {
	if !hid.Supported() {
		log.Panicf("USB HID is not supported in this build!")
	}

	allDeviceInfos := hid.Enumerate(0, 0)
	fmt.Println(len(allDeviceInfos))
	streamDeckDeviceInfos := []hid.DeviceInfo{}

	for _, deviceInfo := range allDeviceInfos {
		fmt.Printf("%04x %04x\n", deviceInfo.VendorID, deviceInfo.ProductID)
		if isStreamDeck(deviceInfo) {
			streamDeckDeviceInfos = append(streamDeckDeviceInfos, deviceInfo)
			// fmt.Printf("Detected %s with serial number %s\n", streamDeckInfo.Description, deviceInfo.Serial)
		}
	}
	return streamDeckDeviceInfos
}

func isStreamDeck(deviceInfo hid.DeviceInfo) bool {
	switch deviceInfo.VendorID {
	case ElgatoVendorID:
		switch deviceInfo.ProductID {
		case OriginalProductID:
			return true

		case MiniProductID:
			return true

		case XLProductID:
			return true

		case OriginalV2ProductID:
			return true
		}
	}
	return false
}

func initializeStreamDeck(deviceInfo hid.DeviceInfo) (*StreamDeck, error) {

	var device *hid.Device

	type result struct {
		device *hid.Device
		err    error
	}

	done := make(chan result)

	go func(done chan<- result) {
		start := time.Now()
		log.Printf("attempting to open device %s", deviceInfo.Path)
		device, err := deviceInfo.Open()
		end := time.Now()
		log.Printf("took %s to open device", end.Sub(start))
		if err != nil {
			done <- result{nil, errors.Wrapf(err, "unable to open Stream Deck %s", deviceInfo.Serial)}
			close(done)
			return
		}
		done <- result{device, nil}
		close(done)
	}(done)

	select {
	case <-time.After(30 * time.Second):
		return nil, errors.Errorf("unable to open device after 30 seconds")

	case r := <-done:
		if r.err != nil {
			return nil, r.err
		}
		device = r.device
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
		return nil, errors.Errorf("No Stream Decks found!")
	}

	return initializeStreamDeck(deviceInfos[0])
}

// NewWithSerial .
func NewWithSerial(serial string) (*StreamDeck, error) {
	deviceInfos := enumerateStreamDecks()

	if len(deviceInfos) == 0 {
		return nil, errors.Errorf("No Stream Decks found!")
	}

	for _, deviceInfo := range deviceInfos {
		if deviceInfo.Serial == serial {
			return initializeStreamDeck(deviceInfo)
		}
	}

	return nil, errors.Errorf("Unable to locate Stream Deck with serial %s", serial)
}
