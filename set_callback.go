package streamdeck

// SetCallback .
func (sd *StreamDeck) SetCallback(x int, y int, callback ButtonCallback) error {
	buttonIndex, err := sd.convertXYToButtonIndex(x, y)
	if err != nil {
		return err
	}
	sd.Lock()
	defer sd.Unlock()
	sd.buttonCallbacks[buttonIndex] = callback
	return nil
}
