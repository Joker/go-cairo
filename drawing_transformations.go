package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"

import (
	"unsafe"
)

// void cairo_translate (cairo_t *cr, double tx, double ty);
func (self *Surface) Translate(tx, ty float64) {
	C.cairo_translate(self.context, C.double(tx), C.double(ty))
}
// void cairo_scale (cairo_t *cr, double sx, double sy);
func (self *Surface) Scale(sx, sy float64) {
	C.cairo_scale(self.context, C.double(sx), C.double(sy))
}
// void cairo_rotate (cairo_t *cr, double angle);
func (self *Surface) Rotate(angle float64) {
	C.cairo_rotate(self.context, C.double(angle))
}


// void cairo_transform (cairo_t *cr, const cairo_matrix_t *matrix);
func (self *Surface) Transform(matrix Matrix) {
	C.cairo_transform(self.context, matrix.cairo_matrix_t())
}
// void cairo_set_matrix (cairo_t *cr, const cairo_matrix_t *matrix);
func (self *Surface) SetMatrix(matrix Matrix) {
	C.cairo_set_matrix(self.context, matrix.cairo_matrix_t())
}
// void cairo_get_matrix (cairo_t *cr, cairo_matrix_t *matrix);
func (self *Surface) GetMatrix(matrix Matrix) {
	C.cairo_get_matrix(self.context, matrix.cairo_matrix_t())
}
// void cairo_identity_matrix (cairo_t *cr);
func (self *Surface) IdentityMatrix() {
	C.cairo_identity_matrix(self.context)
}


// void cairo_user_to_device (cairo_t *cr, double *x, double *y);
func (self *Surface) UserToDevice( x, y float64 ) (float64, float64) {
	C.cairo_user_to_device( self.context, (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)) )
	return x, y
}
// void cairo_user_to_device_distance (cairo_t *cr, double *dx, double *dy);
func (self *Surface) UserToDeviceDistance( dx, dy float64 ) (float64, float64) {
	C.cairo_user_to_device_distance( self.context, (*C.double)(unsafe.Pointer(&dx)), (*C.double)(unsafe.Pointer(&dy)) )
	return dx, dy
}
// void cairo_device_to_user (cairo_t *cr, double *x, double *y);
func (self *Surface) DeviceToUser( x, y float64 ) (float64, float64) {
	C.cairo_device_to_user( self.context, (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)) )
	return x, y
}
// void cairo_device_to_user_distance (cairo_t *cr, double *dx, double *dy);
func (self *Surface) DeviceToUserDistance( dx, dy float64 ) (float64, float64) {
	C.cairo_device_to_user_distance( self.context, (*C.double)(unsafe.Pointer(&dx)), (*C.double)(unsafe.Pointer(&dy)) )
	return dx, dy
}
