package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"



func (self *Surface) Translate(tx, ty float64) {
	C.cairo_translate(self.context, C.double(tx), C.double(ty))
}
func (self *Surface) Scale(sx, sy float64) {
	C.cairo_scale(self.context, C.double(sx), C.double(sy))
}
func (self *Surface) Rotate(angle float64) {
	C.cairo_rotate(self.context, C.double(angle))
}


func (self *Surface) Transform(matrix Matrix) {
	C.cairo_transform(self.context, matrix.cairo_matrix_t())
}
func (self *Surface) SetMatrix(matrix Matrix) {
	C.cairo_set_matrix(self.context, matrix.cairo_matrix_t())
}
func (self *Surface) IdentityMatrix() {
	C.cairo_identity_matrix(self.context)
}


func (self *Surface) UserToDevice(x, y float64) (float64, float64) {
	C.cairo_user_to_device(self.context, (*C.double)(&x), (*C.double)(&y))
	return x, y
}
func (self *Surface) UserToDeviceDistance(dx, dy float64) (float64, float64) {
	C.cairo_user_to_device_distance(self.context, (*C.double)(&dx), (*C.double)(&dy))
	return dx, dy
}
// void cairo_user_to_device_distance (cairo_t *cr, double *dx, double *dy);
func (self *Surface) DeviceToUserDistance( dx *float64, dy *float64 ){
	C.cairo_user_to_device_distance(self.context, (*C.double)(dx), (*C.double)(dy));
}