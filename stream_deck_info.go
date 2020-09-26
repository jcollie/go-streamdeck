package streamdeck

// Info .
// type Info struct {
// 	Description              string
// 	APIVersion               APIVersion
// 	ButtonDirection          KeyDirection
// 	ButtonColumns            int
// 	ButtonRows               int
// 	ImageWidth               int
// 	ImageHeight              int
// 	ImageFormat              ImageFormat
// 	ImageFlipHorizontal      bool
// 	ImageFlipVertical        bool
// 	ImageRotation            ImageRotation
// 	ImageReportLength        int
// 	ImageReportHeaderLength  int
// 	ImageReportPayloadLength int
// 	SpacingHorizontal        int
// 	SpacingVertical          int
// }

// ProductIDs .
// var ProductIDs = map[uint16]Info{
// 	OriginalProductID: {
// 		Description:              "Stream Deck Original",
// 		APIVersion:               1,
// 		ButtonDirection:          KeyDirectionLeftToRight,
// 		ButtonColumns:            5,
// 		ButtonRows:               3,
// 		ImageWidth:               72,
// 		ImageHeight:              72,
// 		ImageFormat:              ImageFormatBMP,
// 		ImageFlipHorizontal:      true,
// 		ImageFlipVertical:        true,
// 		ImageRotation:            ImageRotation0,
// 		ImageReportLength:        1024,
// 		ImageReportHeaderLength:  8,
// 		ImageReportPayloadLength: 1024 - 8,
// 		SpacingHorizontal:        19,
// 		SpacingVertical:          19,
// 	},
// 	MiniProductID: {
// 		Description:              "Stream Deck Mini",
// 		APIVersion:               1,
// 		ButtonDirection:          KeyDirectionLeftToRight,
// 		ButtonColumns:            3,
// 		ButtonRows:               2,
// 		ImageWidth:               80,
// 		ImageHeight:              80,
// 		ImageFormat:              ImageFormatBMP,
// 		ImageFlipHorizontal:      false,
// 		ImageFlipVertical:        true,
// 		ImageRotation:            ImageRotation90,
// 		ImageReportLength:        1024,
// 		ImageReportHeaderLength:  8,
// 		ImageReportPayloadLength: 1024 - 8,
// 		SpacingHorizontal:        19,
// 		SpacingVertical:          19,
// 	},
// 	XLProductID: {
// 		Description:              "Stream Deck XL",
// 		APIVersion:               APIVersion2,
// 		ButtonDirection:          KeyDirectionLeftToRight,
// 		ButtonColumns:            8,
// 		ButtonRows:               4,
// 		ImageWidth:               96,
// 		ImageHeight:              96,
// 		ImageFormat:              ImageFormatJPEG,
// 		ImageFlipHorizontal:      true,
// 		ImageFlipVertical:        true,
// 		ImageRotation:            ImageRotation0,
// 		ImageReportLength:        1024,
// 		ImageReportHeaderLength:  8,
// 		ImageReportPayloadLength: 1024 - 8,
// 		SpacingHorizontal:        19,
// 		SpacingVertical:          19,
// 	},
// 	OriginalV2ProductID: {
// 		Description:              "Stream Deck Original (V2)",
// 		APIVersion:               APIVersion2,
// 		ButtonDirection:          KeyDirectionLeftToRight,
// 		ButtonColumns:            5,
// 		ButtonRows:               3,
// 		ImageWidth:               72,
// 		ImageHeight:              72,
// 		ImageFormat:              ImageFormatJPEG,
// 		ImageFlipHorizontal:      true,
// 		ImageFlipVertical:        true,
// 		ImageRotation:            ImageRotation0,
// 		ImageReportLength:        1024,
// 		ImageReportHeaderLength:  8,
// 		ImageReportPayloadLength: 1024 - 8,
// 		SpacingHorizontal:        19,
// 		SpacingVertical:          19,
// 	},
// }

// NumberOfButtons .
// func (sdi *Info) NumberOfButtons() int {
// 	return sdi.ButtonColumns * sdi.ButtonRows
// }

// // PanelWidth .
// func (sdi *Info) PanelWidth() int {
// 	return sdi.ButtonColumns*sdi.ImageWidth + sdi.SpacingHorizontal*(sdi.ButtonColumns-1)
// }

// // PanelHeight .
// func (sdi *Info) PanelHeight() int {
// 	return sdi.ButtonRows*sdi.ImageHeight + sdi.SpacingVertical*(sdi.ButtonRows-1)
// }
