package streamdeck

import (
	"github.com/pkg/errors"
)

// GetSerialNumber .
func (sd *StreamDeck) GetSerialNumber() (string, error) {
	switch sd.device.ProductID {
	case OriginalProductID:
		payload := make([]byte, 17)
		payload[0] = 0x03
		_, err := sd.device.GetFeatureReport(payload)
		if err != nil {
			return "", errors.Errorf("can't get serial number")
		}
		return string(payload[5:]), nil

	case OriginalV2ProductID:
		payload := make([]byte, 32)
		payload[0] = 0x06
		_, err := sd.device.GetFeatureReport(payload)
		if err != nil {
			return "", errors.Errorf("can't get serial number")
		}

		length := int(payload[1])
		return string(payload[2 : 2+length]), nil

	default:
		panic("not implemented")
	}
}
