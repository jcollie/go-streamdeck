package streamdeck

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

func itob(i byte) ButtonState {
	if i == 0 {
		return ButtonReleased
	}
	return ButtonPressed
}

// Read .
func (sd *StreamDeck) Read() {
	switch sd.device.ProductID {
	case OriginalV2ProductID:
		sd.readV2()
	default:
		log.Panicf("read is not implemented for product id %04x", sd.device.ProductID)
	}
}

// readV1 will read button presses from an API v1 Stream Deck and dispatch button calls
func (sd *StreamDeck) readV1() {
	for {
		data := make([]byte, 1+sd.NumberOfButtons())
		_, err := sd.device.Read(data)
		timestamp := time.Now()
		if err != nil {
			fmt.Printf("unable to read: %v", err)
		}

		sd.dispatchButtons(timestamp, data[1:])
	}
}

// readV2 will read button presses from an API v2 Stream Deck and dispatch button calls
func (sd *StreamDeck) readV2() {
	for {
		data := make([]byte, 4+sd.NumberOfButtons())

		start := time.Now()
		_, err := sd.device.Read(data)
		end := time.Now()
		log.Printf("waited %s for a button to be pressed", end.Sub(start))

		if err != nil {
			log.Printf("unable to read: %v", err.Error())
			continue
		}

		// op := binary.LittleEndian.Uint16(data[0:2])

		nButtons := int(binary.LittleEndian.Uint16(data[2:4]))
		if nButtons != sd.NumberOfButtons() {
			log.Printf("wrong number of buttons: %d, was expecting %d", nButtons, sd.NumberOfButtons())
			continue
		}

		sd.dispatchButtons(end, data[4:])
	}
}

func (sd *StreamDeck) dispatchButtons(timestamp time.Time, data []byte) {

	sd.Lock()
	log.Printf("locked by dispatchButtons")
	defer func() {
		log.Printf("unlocked by dispatchButtons")
		sd.Unlock()
	}()

	for buttonIndex, rawButtonState := range data {
		x, y, err := sd.convertButtonIndexToXY(buttonIndex)
		if err != nil {
			log.Printf("err: %+v", err)
			continue
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
}
