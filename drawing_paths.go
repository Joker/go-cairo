package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
	"math"
)


type PathData []float64
type PathDataPoint struct { x, y float64 }

// cairo_path_t* cairo_copy_path (cairo_t *cr);
func (self *Surface) CopyPath() (PathData, Status) {
	path := C.cairo_copy_path(self.context)
	defer C.cairo_path_destroy(path)

	raw_data := make(PathData, int(path.num_data*2))
	C.memcpy(unsafe.Pointer(&raw_data[0]), unsafe.Pointer(path.data), C.size_t(path.num_data*16))

	return raw_data, Status(path.status)
}
// cairo_path_t* cairo_copy_path_flat (cairo_t *cr);
func (self *Surface) CopyPathFlat() (PathData, Status) {
	path := C.cairo_copy_path_flat(self.context)
	defer C.cairo_path_destroy(path)

	raw_data := make(PathData, int(path.num_data*2))
	C.memcpy(unsafe.Pointer(&raw_data[0]), unsafe.Pointer(path.data), C.size_t(path.num_data*16))

	return raw_data, Status(path.status)
}

type PathDataType int
// cairo_path_data_type_t values
const (
	PATH_MOVE_TO PathDataType = iota
	PATH_LINE_TO
	PATH_CURVE_TO
	PATH_CLOSE_PATH
)
type NextElement func()(PathDataType, []PathDataPoint)

func (self PathData) Interator() NextElement {
	count    := 0
	num_data := len(self)

	return func() (PathDataType, []PathDataPoint) {
		if count >= num_data  {return -1, nil}

		length    := int(math.Float64bits(self[count]) >> 32)
		path_type := int( (math.Float64bits(self[count]) << 32) >> 32 )
		count += 2

		var pathPoints []PathDataPoint
		if length >= 2 {
			length--
			pathPoints = make([]PathDataPoint, length)
			for index   := range pathPoints {
				pathPoints[index].x = self[count]
				pathPoints[index].y = self[count+1]
			}
			count += (length * 2)
		} else {
			pathPoints = []PathDataPoint{}
		}

		return PathDataType(path_type), pathPoints
	}
}



// cairo_bool_t cairo_has_current_point (cairo_t *cr);
func (self *Surface) HasCurrentPoint() bool {
	return C.cairo_has_current_point(self.context) != 0
}

// void cairo_get_current_point (cairo_t *cr, double *x, double *y);
func (self *Surface) GetCurrentPoint( x, y float64 ) (float64, float64) {
	C.cairo_get_current_point( self.context, (*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)) )
	return x, y
}



func (self *Surface) NewPath()		{ C.cairo_new_path(self.context) }

func (self *Surface) NewSubPath()	{ C.cairo_new_sub_path(self.context) }

func (self *Surface) ClosePath()	{ C.cairo_close_path(self.context) }

func (self *Surface) PathExtents() (left, top, right, bottom float64) {
	C.cairo_path_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}



func (self *Surface) MoveTo(x, y float64) { C.cairo_move_to(self.context, C.double(x), C.double(y)) }

func (self *Surface) LineTo(x, y float64) { C.cairo_line_to(self.context, C.double(x), C.double(y)) }

func (self *Surface) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	C.cairo_curve_to(self.context,
		C.double(x1), C.double(y1),
		C.double(x2), C.double(y2),
		C.double(x3), C.double(y3))
}



func (self *Surface) RelMoveTo(dx, dy float64) { C.cairo_rel_move_to(self.context, C.double(dx), C.double(dy)) }

func (self *Surface) RelLineTo(dx, dy float64) { C.cairo_rel_line_to(self.context, C.double(dx), C.double(dy)) }

func (self *Surface) RelCurveTo(dx1, dy1, dx2, dy2, dx3, dy3 float64) {
	C.cairo_rel_curve_to(self.context,
		C.double(dx1), C.double(dy1),
		C.double(dx2), C.double(dy2),
		C.double(dx3), C.double(dy3))
}



func (self *Surface) Arc(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}
func (self *Surface) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	C.cairo_arc_negative(self.context,
		C.double(xc), C.double(yc),
		C.double(radius),
		C.double(angle1), C.double(angle2))
}



func (self *Surface) Rectangle(x, y, width, height float64) {
	C.cairo_rectangle(self.context, C.double(x), C.double(y), C.double(width), C.double(height))
}


func (self *Surface) TextPath(text string) {
	cs := C.CString(text)
	defer C.free(unsafe.Pointer(cs))
	C.cairo_text_path(self.context, cs)
}
// void cairo_glyph_path (cairo_t *cr, const cairo_glyph_t *glyphs, int num_glyphs);
func (self *Surface) GlyphPath(glyphs []Glyph) {
	C.cairo_glyph_path(self.context, (*C.cairo_glyph_t)(unsafe.Pointer(&glyphs[0])), C.int(len(glyphs)))
}



