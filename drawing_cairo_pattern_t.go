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