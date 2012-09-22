package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe";
)


func NewSurfaceFromPNG(filename string) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_image_surface_create_from_png(cs)
	return &Surface{surface: s, context: C.cairo_create(s)}
}


func (self *Surface) WriteToPNG(filename string) {
	p := C.CString(filename);
	C.cairo_surface_write_to_png(self.surface, p);
	C.free(unsafe.Pointer(p));
}
