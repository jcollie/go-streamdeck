package streamdeck

import (
	"fmt"

	"github.com/pkg/errors"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func checkRGB(value int) error {
	if value < 0 || value > 255 {
		return fmt.Errorf("invalid color range")
	}
	return nil
}

// checkValidButtonIndex checks that the button index is valid
func (sd *StreamDeck) checkValidButtonIndex(buttonIndex int) error {
	if buttonIndex < 0 || buttonIndex > sd.NumberOfButtons() {
		return errors.Errorf("invalid key index")
	}
	return nil
}

func (sd *StreamDeck) checkValidButtonXY(x int, y int) error {
	if x < 0 || x >= sd.NumberOfColumns() {
		return errors.Errorf("invalid x coordinate")
	}
	if x < 0 || y >= sd.NumberOfRows() {
		return errors.Errorf("invalid y coordinate")
	}
	return nil
}

func (sd *StreamDeck) convertButtonIndexToXY(buttonIndex int) (int, int, error) {
	err := sd.checkValidButtonIndex(buttonIndex)
	if err != nil {
		return 0, 0, errors.Wrap(err, "can't convert to XY")
	}
	x := buttonIndex % sd.NumberOfColumns()
	y := sd.NumberOfRows() - 1 - (buttonIndex / sd.NumberOfColumns())
	return x, y, nil
}

func (sd *StreamDeck) convertXYToButtonIndex(x int, y int) (int, error) {
	if err := sd.checkValidButtonXY(x, y); err != nil {
		return 0, errors.Wrap(err, "can't convert X, Y to button index")
	}
	buttonIndex := (sd.NumberOfRows()-y-1)*sd.NumberOfColumns() + x
	if err := sd.checkValidButtonIndex(buttonIndex); err != nil {
		return 0, errors.Wrap(err, "unable to convert to button index")
	}
	return buttonIndex, nil
}
