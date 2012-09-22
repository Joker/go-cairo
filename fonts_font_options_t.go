package cairo

type SubpixelOrder int
// cairo_subpixel_order_t
const (
	SUBPIXEL_ORDER_DEFAULT SubpixelOrder = iota
	SUBPIXEL_ORDER_RGB
	SUBPIXEL_ORDER_BGR
	SUBPIXEL_ORDER_VRGB
	SUBPIXEL_ORDER_VBGR
)

type HintStyle int
// cairo_hint_style_t
const (
	HINT_STYLE_DEFAULT HintStyle = iota
	HINT_STYLE_NONE
	HINT_STYLE_SLIGHT
	HINT_STYLE_MEDIUM
	HINT_STYLE_FULL
)

type HintMetrics int
// cairo_hint_metrics_t
const (
	HINT_METRICS_DEFAULT HintMetrics = iota
	HINT_METRICS_OFF
	HINT_METRICS_ON
)