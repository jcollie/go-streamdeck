package streamdeck

import (
	"github.com/pkg/errors"
)

// GetSerialNumber .
func (sd *V2) GetSerialNumber() (string, error) {
	data := make([]byte, 32)
	data[0] = 0x06
	_, err := sd.device.GetFeatureReport(data)
	if err != nil {
		return "", errors.Errorf("can't get serial number")
	}

	length := int(data[1])
	return string(data[2 : 2+length]), nil
}
