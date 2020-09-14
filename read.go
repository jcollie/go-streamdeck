package streamdeck

import (
	"encoding/binary"
	"fmt"
	"time"
)

func itob(i byte) ButtonState {
	if i == 0 {
		return ButtonReleased
	}
	return ButtonPressed
}

func (sd *StreamDeck) Read() {
	for {
		data := make([]byte, 4+sd.NumberOfButtons())
		_, err := sd.device.Read(data)
		timestamp := time.Now()
		if err != nil {
			fmt.Printf("unable to read: %s", err.Error())
			continue
		}

		// op := binary.LittleEndian.Uint16(data[0:2])

		nButtons := int(binary.LittleEndian.Uint16(data[2:4]))
		if nButtons != sd.NumberOfButtons() {
			fmt.Printf("wrong number of buttons!")
			continue
		}

		func() {
			sd.Lock()
			defer sd.Unlock()

			for buttonIndex, rawButtonState := range data[4:] {
				x, y, err := convertButtonIndexToXY(sd, buttonIndex)
				if err != nil {
					fmt.Printf("err: %+v\n", err)
				}
				buttonState := itob(rawButtonState)
				if sd.buttonCallbacks[buttonIndex] != nil {
					if sd.previousState[buttonIndex] != buttonState {
						if buttonState == ButtonPressed {
							go sd.buttonCallbacks[buttonIndex].ButtonPressed(sd, x, y, timestamp)
						} else {
							go sd.buttonCallbacks[buttonIndex].ButtonReleased(sd, x, y, timestamp)
						}
					}
				}
				sd.previousState[buttonIndex] = buttonState
			}
		}()
	}
}
