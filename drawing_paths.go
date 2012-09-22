package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe";
)


type PathDataType int
// cairo_path_data_type_t values
const (
	PATH_MOVE_TO PathDataType = iota
	PATH_LINE_TO
	PATH_CURVE_TO
	PATH_CLOSE_PATH
)



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
	C.cairo_rectangle(self.context,
		C.double(x), C.double(y),
		C.double(width), C.double(height))
}


func (self *Surface) TextPath(text string) {
	cs := C.CString(text)
	C.cairo_text_path(self.context, cs)
	C.free(unsafe.Pointer(cs))
}
func (self *Surface) GlyphPath(glyphs []Glyph) {
	panic("not implemented") // todo
}



func (self *Surface) SetFontMatrix(matrix Matrix) {
	C.cairo_set_font_matrix(self.context, matrix.cairo_matrix_t())
}

func (self *Surface) SetFontOptions(fontOptions *FontOptions) {
	panic("not implemented") // todo
}
func (self *Surface) GetFontOptions() *FontOptions {
	panic("not implemented") // todo
	return nil
}

func (self *Surface) SetFontFace(fontFace *FontFace) {
	panic("not implemented") // todo
}
func (self *Surface) GetFontFace() *FontFace {
	panic("not implemented") // todo
	return nil
}

func (self *Surface) SetScaledFont(scaledFont *ScaledFont) {
	panic("not implemented") // todo
}
func (self *Surface) GetScaledFont() *ScaledFont {
	panic("not implemented") // todo
	return nil
}