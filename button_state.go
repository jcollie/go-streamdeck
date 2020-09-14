package streamdeck

import "github.com/pkg/errors"

// ButtonState .
type ButtonState int

const (
	// ButtonInvalid .
	ButtonInvalid ButtonState = -1
	// ButtonReleased .
	ButtonReleased ButtonState = 0
	// ButtonPressed .
	ButtonPressed ButtonState = 1
)

// GetButtonState .
func (sd *StreamDeck) GetButtonState(x int, y int) (ButtonState, error) {
	buttonIndex, err := convertXYToButtonIndex(sd, x, y)
	if err != nil {
		return ButtonInvalid, errors.Wrap(err, "unable to get button state")
	}
	sd.Lock()
	defer sd.Unlock()
	buttonState := sd.previousState[buttonIndex]
	return buttonState, nil
}
