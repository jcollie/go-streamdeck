package streamdeck

// FillPanel fills the whole panel witn an image. The image is scaled to fit
// and then center-cropped (if necessary). The native picture size is 360px x 216px.
// func (sd *StreamDeck) FillPanel(img image.Image) error {

// 	// resize if the picture width is larger or smaller than panel
// 	rect := img.Bounds()
// 	if rect.Dx() != sd.info.PanelWidth() || rect.Dy() != sd.info.PanelHeight() {
// 		newWidthRatio := float32(rect.Dx()) / float32((sd.info.PanelWidth()))
// 		img = resize(img, sd.info.PanelWidth(), int(float32(rect.Dy())/newWidthRatio))
// 	}

// 	// if the Canvas is larger than PanelWidth x PanelHeight then we crop
// 	// the Center match PanelWidth x PanelHeight
// 	rect = img.Bounds()
// 	if rect.Dx() > sd.info.PanelWidth() || rect.Dy() > sd.info.PanelHeight() {
// 		img = cropCenter(img, sd.info.PanelWidth(), sd.info.PanelHeight())
// 	}

// 	counter := 0

// 	for row := 0; row < sd.info.KeyRows; row++ {
// 		for col := 0; col < sd.info.KeyColumns; col++ {
// 			rect := image.Rectangle{
// 				Min: image.Point{
// 					sd.info.PanelWidth() - sd.info.KeyPixelWidth - col*ButtonSize - col*Spacer,
// 					row*ButtonSize + row*Spacer,
// 				},
// 				Max: image.Point{
// 					PanelWidth - 1 - col*ButtonSize - col*Spacer,
// 					ButtonSize - 1 + row*ButtonSize + row*Spacer,
// 				},
// 			}
// 			sd.FillImage(counter, img.(*image.RGBA).SubImage(rect))
// 			counter++
// 		}
// 	}

// 	return nil
// }
