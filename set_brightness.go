package streamdeck

import "github.com/pkg/errors"

//SetBrightness sets the background brightness
func (sd *StreamDeck) SetBrightness(percent int) error {
	switch sd.device.ProductID {
	case OriginalProductID:
		payload := make([]byte, 17)
		payload[0] = 0x05
		payload[1] = 0x55
		payload[2] = 0xaa
		payload[3] = 0xd1
		payload[4] = 0x01
		payload[5] = byte(min(max(percent, 0), 100))
		_, err := sd.device.SendFeatureReport(payload)
		if err != nil {
			return errors.Errorf("unable to set brightness: %s", err.Error())
		}
		return nil

	case OriginalV2ProductID:
		payload := make([]byte, 32)
		payload[0] = 0x03
		payload[1] = 0x08
		payload[2] = byte(min(max(percent, 0), 100))
		_, err := sd.device.SendFeatureReport(payload)
		if err != nil {
			return errors.Errorf("unable to set brightness: %s", err.Error())
		}
		return nil

	default:
		panic("not implemented")
	}
}
