package streamdeck

//ImageFormat is used to indicate what format images should be in
type ImageFormat int

const (
	//ImageFormatJPEG idicates that the StreamDeck wants images in JPEG format
	ImageFormatJPEG ImageFormat = iota
	//ImageFormatBMP indicates that the StreamDeck wants images in BMP format
	ImageFormatBMP
)
