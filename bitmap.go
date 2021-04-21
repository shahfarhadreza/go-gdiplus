package gdiplus

import (
	"log"
	"syscall"
)

type Bitmap struct {
	Image
}

func NewBitmap(width, height int32, format PixelFormat) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	status := GdipCreateBitmapFromScan0(width, height, 0, format, nil, &nativeBitmap)
	if status != Ok {
		log.Panicln(status.String())
	}
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapEx(width, height, stride int32, format PixelFormat, scan0 *byte) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromScan0(width, height, stride, format, scan0, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapFromHBITMAP(hbitmap HBITMAP) *Bitmap {
	bitmap := &Bitmap{}
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromHBITMAP(hbitmap, 0, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func NewBitmapFromFile(fileName string) *Bitmap {
	bitmap := &Bitmap{}
	fileNameUTF16, _ := syscall.UTF16PtrFromString(fileName)
	var nativeBitmap *GpBitmap
	GdipCreateBitmapFromFile(fileNameUTF16, &nativeBitmap)
	bitmap.nativeImage = (*GpImage)(nativeBitmap)
	return bitmap
}

func (bitmap *Bitmap) Dispose() {
	GdipDisposeImage(bitmap.nativeImage)
}
