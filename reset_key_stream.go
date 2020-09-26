package streamdeck

import "github.com/pkg/errors"

// ResetKeyStream .
func (sd *StreamDeck) ResetKeyStream() error {
	switch sd.device.ProductID {
	case OriginalProductID, OriginalV2ProductID:
		payload := make([]byte, sd.ImageReportPayloadLength())
		payload[0] = 0x02
		_, err := sd.device.Write(payload)
		if err != nil {
			return errors.Errorf("unable to reset key stream: %s", err.Error())
		}
		return nil
	default:
		panic("not implemented")
	}
}
