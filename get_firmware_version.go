package streamdeck

import (
	"github.com/pkg/errors"
)

// GetFirmwareVersion .
func (sd *V2) GetFirmwareVersion() (string, error) {
	data := make([]byte, 32)
	data[0] = 0x05
	_, err := sd.device.GetFeatureReport(data)
	if err != nil {
		return "", errors.Errorf("can't get firmware version: %s", err.Error())
	}
	length := int(data[1])
	return string(data[6 : 2+length]), nil
}
