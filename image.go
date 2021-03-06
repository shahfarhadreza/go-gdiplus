package gdiplus

import (
	"syscall"
)

type Image struct {
	nativeImage *GpImage
}

func NewImageFromFile(fileName string) *Image {
	image := &Image{}
	fileNameUTF16, _ := syscall.UTF16PtrFromString(fileName)
	GdipLoadImageFromFile(fileNameUTF16, &image.nativeImage)
	return image
}

func (image *Image) GetWidth() (width uint32) {
	GdipGetImageWidth(image.nativeImage, &width)
	return
}

func (image *Image) GetHeight() (height uint32) {
	GdipGetImageHeight(image.nativeImage, &height)
	return
}

func (image *Image) Dispose() {
	GdipDisposeImage(image.nativeImage)
}
