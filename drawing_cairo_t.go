package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe";
)


func (self *Surface) Save()				{ C.cairo_save(self.context) 	}
func (self *Surface) Restore()			{ C.cairo_restore(self.context) }
func (self *Surface) ContextDestroy()	{ C.cairo_destroy(self.context) }
//  GetStatus() >>> Status()   and   Status() >>> ContextStatus()
func (self *Surface) ContextStatus() Status {
	return Status(C.cairo_status(self.context))
}


func (self *Surface) CopyPage()	{ C.cairo_copy_page(self.context) }
func (self *Surface) ShowPage()	{ C.cairo_show_page(self.context) }



func (self *Surface) Paint()					   { C.cairo_paint(self.context) }
func (self *Surface) PaintWithAlpha(alpha float64) { C.cairo_paint_with_alpha(self.context, C.double(alpha)) }



func (self *Surface) SetSource(pattern *Pattern) {
	C.cairo_set_source(self.context, pattern.pattern)
}
func (self *Surface) SetSourceRGB(red, green, blue float64) {
	C.cairo_set_source_rgb(self.context, C.double(red), C.double(green), C.double(blue))
}
func (self *Surface) SetSourceRGBA(red, green, blue, alpha float64) {
	C.cairo_set_source_rgba(self.context, C.double(red), C.double(green), C.double(blue), C.double(alpha))
}
func (self *Surface) SetSourceSurface(surface *Surface, x, y float64) {
	C.cairo_set_source_surface(self.context, surface.surface, C.double(x), C.double(y))
}



func (self *Surface) Mask(pattern *Pattern) { 
	C.cairo_mask(self.context, pattern.pattern) 
}
func (self *Surface) MaskSurface(surface *Surface, surface_x, surface_y float64) {
	C.cairo_mask_surface(self.context, surface.surface, C.double(surface_x), C.double(surface_y))
}



func (self *Surface) Stroke()			{ C.cairo_stroke(self.context) }
func (self *Surface) Fill()				{ C.cairo_fill(self.context) }

func (self *Surface) StrokePreserve()	{ C.cairo_stroke_preserve(self.context) }
func (self *Surface) FillPreserve()		{ C.cairo_fill_preserve(self.context) }

// Insideness testing
func (self *Surface) InStroke(x, y float64) bool {
	return C.cairo_in_stroke(self.context, C.double(x), C.double(y)) != 0
}
func (self *Surface) InFill(x, y float64) bool {
	return C.cairo_in_fill(self.context, C.double(x), C.double(y)) != 0
}

// Rectangular extents
func (self *Surface) StrokeExtents() (left, top, right, bottom float64) {
	C.cairo_stroke_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}
func (self *Surface) FillExtents() (left, top, right, bottom float64) {
	C.cairo_fill_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}



func (self *Surface) Clip()			{ C.cairo_clip(self.context) }

func (self *Surface) ResetClip()	{ C.cairo_reset_clip(self.context) }

func (self *Surface) ClipPreserve()	{ C.cairo_clip_preserve(self.context) }

func (self *Surface) ClipExtents() (left, top, right, bottom float64) {
	C.cairo_clip_extents(self.context,
		(*C.double)(&left), (*C.double)(&top),
		(*C.double)(&right), (*C.double)(&bottom))
	return left, top, right, bottom
}
func (self *Surface) ClipRectangleList() ([]Rectangle, Status) {
	list := C.cairo_copy_clip_rectangle_list(self.context)
	defer C.cairo_rectangle_list_destroy(list)
	rects := make([]Rectangle, int(list.num_rectangles))
	C.memcpy(unsafe.Pointer(&rects[0]), unsafe.Pointer(list.rectangles), C.size_t(list.num_rectangles*8))
	return rects, Status(list.status)
}



func (self *Surface) PushGroup() {
	C.cairo_push_group(self.context)
}
func (self *Surface) PushGroupWithContent(content Content) {
	C.cairo_push_group_with_content(self.context, C.cairo_content_t(content))
}
func (self *Surface) PopGroup() (pattern *Pattern) {
	return &Pattern{C.cairo_pop_group(self.context)}
}
func (self *Surface) PopGroupToSource() {
	C.cairo_pop_group_to_source(self.context)
}



func (self *Surface) SetMiterLimit(limit float64) 	 { 
	C.cairo_set_miter_limit(self.context, C.double(limit)) 
}
func (self *Surface) SetTolerance(tolerance float64) { 
	C.cairo_set_tolerance(self.context, C.double(tolerance)) 
}



type FillRule int
// cairo_fill_rule_t 
const (
	FILL_RULE_WINDING FillRule = iota
	FILL_RULE_EVEN_ODD
)
func (self *Surface) SetFillRule(fill_rule FillRule) {
	C.cairo_set_fill_rule(self.context, C.cairo_fill_rule_t(fill_rule))
}


type Antialias int
// cairo_antialias_t 
const (
	ANTIALIAS_DEFAULT Antialias = iota
	ANTIALIAS_NONE
	ANTIALIAS_GRAY
	ANTIALIAS_SUBPIXEL
)
func (self *Surface) SetAntialias(antialias Antialias) {
	C.cairo_set_antialias(self.context, C.cairo_antialias_t(antialias))
}


type Operator int
// cairo_operator_t 
const (
	OPERATOR_CLEAR Operator = iota

	OPERATOR_SOURCE
	OPERATOR_OVER
	OPERATOR_IN
	OPERATOR_OUT
	OPERATOR_ATOP

	OPERATOR_DEST
	OPERATOR_DEST_OVER
	OPERATOR_DEST_IN
	OPERATOR_DEST_OUT
	OPERATOR_DEST_ATOP

	OPERATOR_XOR
	OPERATOR_ADD
	OPERATOR_SATURATE

	OPERATOR_MULTIPLY
	OPERATOR_SCREEN
	OPERATOR_OVERLAY
	OPERATOR_DARKEN
	OPERATOR_LIGHTEN
	OPERATOR_COLOR_DODGE
	OPERATOR_COLOR_BURN
	OPERATOR_HARD_LIGHT
	OPERATOR_SOFT_LIGHT
	OPERATOR_DIFFERENCE
	OPERATOR_EXCLUSION
	OPERATOR_HSL_HUE
	OPERATOR_HSL_SATURATION
	OPERATOR_HSL_COLOR
	OPERATOR_HSL_LUMINOSITY
)
func (self *Surface) SetOperator(operator Operator) {
	C.cairo_set_operator(self.context, C.cairo_operator_t(operator))
}


type LineCap int
// cairo_line_cap_t 
const (
	LINE_CAP_BUTT LineCap = iota
	LINE_CAP_ROUND
	LINE_CAP_SQUARE
)
type LineJoin int
// cairo_line_cap_join_t 
const (
	LINE_JOIN_MITER LineJoin = iota
	LINE_JOIN_ROUND
	LINE_JOIN_BEVEL
)
func (self *Surface) SetLineWidth(width float64) {
	C.cairo_set_line_width(self.context, C.double(width))
}
func (self *Surface) SetLineCap(line_cap LineCap) {
	C.cairo_set_line_cap(self.context, C.cairo_line_cap_t(line_cap))
}
func (self *Surface) SetLineJoin(line_join LineJoin) {
	C.cairo_set_line_join(self.context, C.cairo_line_join_t(line_join))
}


// void cairo_set_dash (cairo_t *cr, const double *dashes, int num_dashes, double offset);
func (self *Surface) SetDash(dashes *float64, num_dashes int, offset float64){
    C.cairo_set_dash(self.context, (*C.double)(unsafe.Pointer(dashes)), C.int(num_dashes), C.double(offset));
}

// func (self *Surface) SetDash(dashes []float64, num_dashes int, offset float64) {
// 	dashesp := (*C.double)(&dashes[0])
// 	C.cairo_set_dash(self.context, dashesp, C.int(num_dashes), C.double(offset))
// }