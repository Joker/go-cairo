package cairo

// #cgo LDFLAGS: -lcairo -lxcb
// #include <cairo/cairo-xcb.h>
// #include <stdlib.h>
import "C"
import "github.com/BurntSushi/xgb/xproto"



func NewSurfaceFromXCB(xcb_drawable xproto.Drawable, xcb_VI xproto.VisualInfo, width, height int) *Surface {

	var xcb_visualtype  C.xcb_visualtype_t
		xcb_visualtype.visual_id		  = C.xcb_visualid_t(xcb_VI.VisualId)
		xcb_visualtype._class 			  = C.uint8_t(xcb_VI.Class)
		xcb_visualtype.bits_per_rgb_value = C.uint8_t(xcb_VI.BitsPerRgbValue)
		xcb_visualtype.colormap_entries   = C.uint16_t(xcb_VI.ColormapEntries)
		xcb_visualtype.red_mask 		  = C.uint32_t(xcb_VI.RedMask)
		xcb_visualtype.green_mask 		  = C.uint32_t(xcb_VI.GreenMask)
		xcb_visualtype.blue_mask 		  = C.uint32_t(xcb_VI.BlueMask)

	var connect_xcb (*C.xcb_connection_t) = C.xcb_connect(nil, nil);

	surface := new(Surface);
	surface.surface = C.cairo_xcb_surface_create( connect_xcb, C.xcb_drawable_t(xcb_drawable), &xcb_visualtype, C.int(width), C.int(height));
	surface.context = C.cairo_create(surface.surface);
	
	return surface;
}
