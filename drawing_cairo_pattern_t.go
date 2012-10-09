package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"

type Pattern struct {
	pattern *C.cairo_pattern_t
}


// cairo_pattern_t* cairo_pattern_create_for_surface (cairo_surface_t *surface);
func (self *Surface) CreatePattern() (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_for_surface( self.surface )}
}

// cairo_pattern_t* cairo_pattern_create_rgb (double red, double green, double blue);
func SolidPatternRGB(red, green, blue float64) (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_rgb(C.double(red), C.double(green), C.double(blue) )}
}
// cairo_pattern_t* cairo_pattern_create_rgba (double red, double green, double blue, double alpha);
func SolidPatternRGBA(red, green, blue, alpha float64) (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_rgba( C.double(red), C.double(green), C.double(blue), C.double(alpha) )}
}

// cairo_pattern_t* cairo_pattern_create_linear (double x0, double y0, double x1, double y1);
func LinearGradient(x0, y0, x1, y1 float64) (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_linear( C.double(x0), C.double(y0),
							C.double(x1), C.double(y1) )}
}
// cairo_pattern_t* cairo_pattern_create_radial (double cx0, double cy0, double radius0, double cx1, double cy1, double radius1);
func RadialGradient(cx0, cy0, radius0, cx1, cy1, radius1 float64) (pattern *Pattern) {
	return &Pattern{ C.cairo_pattern_create_radial( C.double(cx0), C.double(cy0), C.double(radius0),
							C.double(cx1), C.double(cy1), C.double(radius1) )}
}

// void cairo_pattern_add_color_stop_rgb (cairo_pattern_t *pattern, double offset, double red, double green, double blue);
func (self *Pattern) AddColorStopRGB(offset, red, green, blue float64) {
	C.cairo_pattern_add_color_stop_rgb( self.pattern, C.double(offset), C.double(red), C.double(green), C.double(blue) )
}
// void cairo_pattern_add_color_stop_rgba (cairo_pattern_t *pattern, double offset, double red, double green, double blue, double alpha);
func (self *Pattern) AddColorStopRGBA(offset, red, green, blue, alpha float64) {
	C.cairo_pattern_add_color_stop_rgba( self.pattern, C.double(offset), C.double(red), C.double(green), C.double(blue), C.double(alpha) )
}


type Extend int
// cairo_extend_t
const (
	EXTEND_NONE Extend = iota
	EXTEND_REPEAT
	EXTEND_REFLECT
	EXTEND_PAD
)
// void cairo_pattern_set_extend (cairo_pattern_t *pattern, cairo_extend_t extend);
func (self *Pattern) SetExtend(extend Extend) {
	C.cairo_pattern_set_extend( self.pattern, C.cairo_extend_t(extend) )
}
// cairo_extend_t cairo_pattern_get_extend (cairo_pattern_t *pattern);
func (self *Pattern) GetExtend() (Extend) {
	return Extend(C.cairo_pattern_get_extend( self.pattern ))
}

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
// void cairo_pattern_set_filter (cairo_pattern_t *pattern, cairo_filter_t filter);
func (self *Pattern) SetFilter(filter Filter) {
	C.cairo_pattern_set_filter( self.pattern, C.cairo_filter_t(filter) )
}
// cairo_filter_t cairo_pattern_get_filter (cairo_pattern_t *pattern);
func (self *Pattern) GetFilter() (Filter) {
	return Filter(C.cairo_pattern_get_filter( self.pattern ))
}


type PatternType int
// cairo_pattern_type_t
const (
	PATTERN_TYPE_SOLID PatternType = iota
	PATTERN_TYPE_SURFACE
	PATTERN_TYPE_LINEAR
	PATTERN_TYPE_RADIAL
	PATTERN_TYPE_MESH
	PATTERN_TYPE_RASTER_SOURCE
)
// cairo_pattern_type_t cairo_pattern_get_type (cairo_pattern_t *pattern);
func (self *Pattern) GetType() (PatternType) {
	return PatternType(C.cairo_pattern_get_type( self.pattern ))
}


// void cairo_pattern_destroy (cairo_pattern_t *pattern);
func (self *Pattern) Destroy() {
	C.cairo_pattern_destroy( self.pattern )
}
