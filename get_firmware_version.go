package streamdeck

import (
	"github.com/pkg/errors"
)

// GetFirmwareVersion .
func (sd *StreamDeck) GetFirmwareVersion() (string, error) {
	switch sd.device.ProductID {
	case OriginalProductID:
		payload := make([]byte, 17)
		payload[0] = 0x04
		_, err := sd.device.GetFeatureReport(payload)
		if err != nil {
			return "", errors.Wrap(err, "can't get firmware version")
		}
		return string(payload[5:]), nil

	case OriginalV2ProductID:
		payload := make([]byte, 32)
		payload[0] = 0x05
		_, err := sd.device.GetFeatureReport(payload)
		if err != nil {
			return "", errors.Wrap(err, "can't get firmware version")
		}
		length := int(payload[1])
		return string(payload[6 : 2+length]), nil

	default:
		panic("not implemented")
	}
}
