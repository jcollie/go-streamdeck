package streamdeck

import (
	"github.com/pkg/errors"
)

// ResetToLogo resets the keys to the detault logo.
func (sd *StreamDeck) ResetToLogo() error {
	payload := make([]byte, 32)
	payload[0] = 0x03
	payload[1] = 0x02
	_, err := sd.device.SendFeatureReport(payload)
	if err != nil {
		return errors.Wrap(err, "unable to reset to logo")
	}
	return nil
}
