package streamdeck

import (
	"log"

	"github.com/pkg/errors"
)

// ResetToLogo resets the keys to the detault logo.
func (sd *StreamDeck) ResetToLogo() error {
	sd.Lock()
	log.Printf("locked by ResetToLogo")
	defer func() {
		log.Printf("unlocked by ResetToLogo")
		sd.Unlock()
	}()

	switch sd.device.ProductID {
	case OriginalProductID:
		payload := make([]byte, 17)
		payload[0] = 0x0b
		payload[1] = 0x63
		_, err := sd.device.SendFeatureReport(payload)
		if err != nil {
			return errors.Wrap(err, "unable to reset to logo")
		}
		return nil
	case OriginalV2ProductID:
		payload := make([]byte, 32)
		payload[0] = 0x03
		payload[1] = 0x02
		_, err := sd.device.SendFeatureReport(payload)
		if err != nil {
			return errors.Wrap(err, "unable to reset to logo")
		}
		return nil
	default:
		panic("not implemented")
	}
}
