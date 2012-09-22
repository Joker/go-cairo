package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
import "C"


// Golang struct to hold both a cairo surface and a cairo context
type Surface struct {
	surface	*C.cairo_surface_t;
	context	*C.cairo_t;
}

type Content int
// cairo_content_t
const (
	CONTENT_COLOR       Content = 0x1000
	CONTENT_ALPHA       Content = 0x2000
	CONTENT_COLOR_ALPHA Content = 0x3000
)

type SurfaceType int
// cairo_surface_type_t
const (
	SURFACE_TYPE_IMAGE SurfaceType = iota
	SURFACE_TYPE_PDF
	SURFACE_TYPE_PS
	SURFACE_TYPE_XLIB
	SURFACE_TYPE_XCB
	SURFACE_TYPE_GLITZ
	SURFACE_TYPE_QUARTZ
	SURFACE_TYPE_WIN32
	SURFACE_TYPE_BEOS
	SURFACE_TYPE_DIRECTFB
	SURFACE_TYPE_SVG
	SURFACE_TYPE_OS2
	SURFACE_TYPE_WIN32_PRINTING
	SURFACE_TYPE_QUARTZ_IMAGE
	SURFACE_TYPE_SCRIPT
	SURFACE_TYPE_QT
	SURFACE_TYPE_RECORDING
	SURFACE_TYPE_VG
	SURFACE_TYPE_GL
	SURFACE_TYPE_DRM
	SURFACE_TYPE_TEE
	SURFACE_TYPE_XML
	SURFACE_TYPE_SKIA 
	SURFACE_TYPE_SUBSURFACE
)



func (self *Surface) CreateForRectangle(x, y, width, height float64) *Surface {
	return &Surface{
		context: self.context,
		surface: C.cairo_surface_create_for_rectangle(self.surface,
			C.double(x), C.double(y), C.double(width), C.double(height)),
	}
}



//  GetStatus() >>> Status()   and   Status() >>> ContextStatus()
func (self *Surface) Status() Status {
	return Status(C.cairo_surface_status(self.surface))
}
func (self *Surface) Destroy() {
	C.cairo_surface_destroy(self.surface)
}
func (self *Surface) Finish() {
	C.cairo_surface_finish(self.surface)
}
func (self *Surface) Flush() {
	C.cairo_surface_flush(self.surface)
}



func (self *Surface) GetType() SurfaceType {
	return SurfaceType(C.cairo_surface_get_type(self.surface))
}
func (self *Surface) GetContent() Content {
	return Content(C.cairo_surface_get_content(self.surface))
}
func (self *Surface) GetDevice() *Device {
	//C.cairo_surface_get_device
	panic("not implemented") // todo
	return nil
}
func (self *Surface) GetReferenceCount() int {
	return int(C.cairo_surface_get_reference_count(self.surface))
}


func (self *Surface) MarkDirty() {
	C.cairo_surface_mark_dirty(self.surface)
}
func (self *Surface) MarkDirtyRectangle(x, y, width, height int) {
	C.cairo_surface_mark_dirty_rectangle(self.surface,
		C.int(x), C.int(y), C.int(width), C.int(height))
}

func (self *Surface) SetDeviceOffset(x, y float64) {
	C.cairo_surface_set_device_offset(self.surface, C.double(x), C.double(y))
}
func (self *Surface) GetDeviceOffset() (x, y float64) {
	C.cairo_surface_get_device_offset(self.surface, (*C.double)(&x), (*C.double)(&y))
	return x, y
}

func (self *Surface) SetFallbackResolution(xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_set_fallback_resolution(self.surface,
		C.double(xPixelPerInch), C.double(yPixelPerInch))
}
func (self *Surface) GetFallbackResolution() (xPixelPerInch, yPixelPerInch float64) {
	C.cairo_surface_get_fallback_resolution(self.surface,
		(*C.double)(&xPixelPerInch), (*C.double)(&yPixelPerInch))
	return xPixelPerInch, yPixelPerInch
}



func (self *Surface) HasShowTextGlyphs() bool {
	return C.cairo_surface_has_show_text_glyphs(self.surface) != 0
}