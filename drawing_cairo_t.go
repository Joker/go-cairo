package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"

import (
	"unsafe";
)

// void cairo_set_dash (cairo_t *cr, const double *dashes, int num_dashes, double offset);
func (self *Surface) SetDash(dashes *float64, num_dashes int, offset float64){
    C.cairo_set_dash(self.context, (*C.double)(unsafe.Pointer(dashes)), C.int(num_dashes), C.double(offset));
}

// func (self *Surface) SetDash(dashes []float64, num_dashes int, offset float64) {
// 	dashesp := (*C.double)(&dashes[0])
// 	C.cairo_set_dash(self.context, dashesp, C.int(num_dashes), C.double(offset))
// }