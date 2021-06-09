package ximagick

//
// export CGO_CFLAGS_ALLOW='-Xpreprocessor' go get gopkg.in/gographics/imagick.v3/imagick
//func ConvertPdfToJpg(pdfName string, imageName string) error {
//
//	// Setup
//	imagick.Initialize()
//	defer imagick.Terminate()
//
//	mw := imagick.NewMagickWand()
//	defer mw.Destroy()
//
//	// Must be *before* ReadImageFile
//	// Make sure our image is high quality
//	if err := mw.SetResolution(300, 300); err != nil {
//		return err
//	}
//
//	// Load the image file into imagick
//	if err := mw.ReadImage(pdfName); err != nil {
//		return err
//	}
//
//	// Must be *after* ReadImageFile
//	// Flatten image and remove alpha channel, to prevent alpha turning black in jpg
//	if err := mw.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_COPY); err != nil {
//		return err
//	}
//
//	// Set any compression (100 = max quality)
//	if err := mw.SetCompressionQuality(95); err != nil {
//		return err
//	}
//
//	// Select only first page of pdf
//	mw.SetIteratorIndex(0)
//
//	// Convert into JPG
//	if err := mw.SetFormat("jpg"); err != nil {
//		return err
//	}
//
//	// Save File
//	return mw.WriteImage(imageName)
//}
//
