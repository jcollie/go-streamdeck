package streamdeck

import "github.com/pkg/errors"

//SetBrightness sets the background brightness
func (sd *StreamDeck) SetBrightness(percent int) error {
	payload := make([]byte, 32)
	payload[0] = 0x03
	payload[1] = 0x08
	payload[2] = byte(min(max(percent, 0), 100))
	_, err := sd.device.SendFeatureReport(payload)
	if err != nil {
		return errors.Errorf("unable to set brightness: %s", err.Error())
	}
	return nil
}
