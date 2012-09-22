package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"image"
	"image/draw"
	"unsafe"

	"./extimage"
)

type Format int
// cairo_format_t
const (
	FORMAT_INVALID   Format = -1
	FORMAT_ARGB32    Format = 0
	FORMAT_RGB24     Format = 1
	FORMAT_A8        Format = 2
	FORMAT_A1        Format = 3
	FORMAT_RGB16_565 Format = 4
	FORMAT_RGB30     Format = 5
)
func (self Format) StrideForWidth(width int) int {
	return int(C.cairo_format_stride_for_width(C.cairo_format_t(self), C.int(width)))
}


func NewSurface(format Format, width, height int) *Surface {
	s := C.cairo_image_surface_create(C.cairo_format_t(format), C.int(width), C.int(height))
	return &Surface{surface: s, context: C.cairo_create(s)}
}
func NewSurfaceFromImage(img image.Image) *Surface {
	format := FORMAT_ARGB32
	switch img.(type) {
	case *image.Alpha, *image.Alpha16:
		format = FORMAT_A8
	case *extimage.RGB, *image.Gray, *image.Gray16, *image.YCbCr:
		format = FORMAT_RGB24
	}
	surface := NewSurface(format, img.Bounds().Dx(), img.Bounds().Dy())
	surface.SetImage(img)
	return surface
}



func (self *Surface) GetFormat() Format {
	return Format(C.cairo_image_surface_get_format(self.surface))
}
func (self *Surface) GetWidth() int {
	return int(C.cairo_image_surface_get_width(self.surface))
}
func (self *Surface) GetHeight() int {
	return int(C.cairo_image_surface_get_height(self.surface))
}
func (self *Surface) GetStride() int {
	return int(C.cairo_image_surface_get_stride(self.surface))
}







///////////////////////////////////////////////////////////////////////////////

// GetData returns a copy of the surfaces raw pixel data.
// This method also calls Flush.
func (self *Surface) GetData() []byte {
	self.Flush()
	dataPtr := C.cairo_image_surface_get_data(self.surface)
	if dataPtr == nil {
		panic("cairo.Surface.GetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(self.surface)
	height := C.cairo_image_surface_get_height(self.surface)
	return C.GoBytes(unsafe.Pointer(dataPtr), stride*height)
}

// SetData sets the surfaces raw pixel data.
// This method also calls Flush and MarkDirty.
func (self *Surface) SetData(data []byte) {
	self.Flush()
	dataPtr := unsafe.Pointer(C.cairo_image_surface_get_data(self.surface))
	if dataPtr == nil {
		panic("cairo.Surface.SetData(): can't access surface pixel data")
	}
	stride := C.cairo_image_surface_get_stride(self.surface)
	height := C.cairo_image_surface_get_height(self.surface)
	if len(data) != int(stride*height) {
		panic("cairo.Surface.SetData(): invalid data size")
	}
	C.memcpy(dataPtr, unsafe.Pointer(&data[0]), C.size_t(stride*height))
	self.MarkDirty()
}

// image.Image methods

func (self *Surface) GetImage() image.Image {
	width  := self.GetWidth()
	height := self.GetHeight()
	stride := self.GetStride()
	data   := self.GetData()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		return &extimage.ARGB{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_RGB24:
		return &extimage.RGB{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_A8:
		return &image.Alpha{
			Pix:    data,
			Stride: stride,
			Rect:   image.Rect(0, 0, width, height),
		}

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}

// SetImage set the data from an image.Image with identical size.
func (self *Surface) SetImage(img image.Image) {
	width  := self.GetWidth()
	height := self.GetHeight()
	stride := self.GetStride()

	switch self.GetFormat() {
	case FORMAT_ARGB32:
		if i, ok := img.(*extimage.ARGB); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*extimage.ARGB)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_RGB24:
		if i, ok := img.(*extimage.RGB); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*extimage.RGB)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_A8:
		if i, ok := img.(*image.Alpha); ok {
			if i.Rect.Dx() == width && i.Rect.Dy() == height && i.Stride == stride {
				self.SetData(i.Pix)
				return
			}
		}
		surfImg := self.GetImage().(*image.Alpha)
		draw.Draw(surfImg, surfImg.Bounds(), img, img.Bounds().Min, draw.Src)
		self.SetData(surfImg.Pix)

	case FORMAT_A1:
		panic("Unsuppored surface format cairo.FORMAT_A1")

	case FORMAT_RGB16_565:
		panic("Unsuppored surface format cairo.FORMAT_RGB16_565")

	case FORMAT_RGB30:
		panic("Unsuppored surface format cairo.FORMAT_RGB30")

	case FORMAT_INVALID:
		panic("Invalid surface format")
	}
	panic("Unknown surface format")
}