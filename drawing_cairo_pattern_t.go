package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"

type Pattern struct {
	pattern *C.cairo_pattern_t
}

type PatternType int
// cairo_pattern_type_t
const (
	PATTERN_TYPE_SOLID PatternType = iota
	PATTERN_TYPE_SURFACE
	PATTERN_TYPE_LINEAR
	PATTERN_TYPE_RADIAL
)

type Extent int
// cairo_extend_t
const (
	EXTEND_NONE Extent = iota
	EXTEND_REPEAT
	EXTEND_REFLECT
	EXTEND_PAD
)

type Filter int
// cairo_filter_t
const (
	CAIRO_FILTER_FAST Filter = iota
	CAIRO_FILTER_GOOD
	CAIRO_FILTER_BEST
	CAIRO_FILTER_NEAREST
	CAIRO_FILTER_BILINEAR
	CAIRO_FILTER_GAUSSIAN
)


// cairo_pattern_t* cairo_pattern_create_linear (double x0, double y0, double x1, double y1);
func LinearGradient(x0, y0, x1, y1 float64) (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_linear( C.double(x0), C.double(y0), C.double(x1), C.double(y1) )}
}


// void cairo_pattern_set_extend (cairo_pattern_t *pattern, cairo_extend_t extend);
func (self *Pattern) SetExtend(extend Extent) {
	C.cairo_pattern_set_extend( self.pattern, C.cairo_extend_t(extend) )
}

// void cairo_pattern_add_color_stop_rgb (cairo_pattern_t *pattern, double offset, double red, double green, double blue);
func (self *Pattern) AddColorStopRGB(offset, red, green, blue float64) {
	C.cairo_pattern_add_color_stop_rgb( self.pattern, C.double(offset), C.double(red), C.double(green), C.double(blue) )
}



