package cairo

// #include <cairo/cairo-pdf.h>
// #include <cairo/cairo-ps.h>
// #include <cairo/cairo-svg.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)


type PDFVersion int

func (self PDFVersion) String() string {
	return C.GoString(C.cairo_pdf_version_to_string(C.cairo_pdf_version_t(self)))
}

const (
	PDF_VERSION_1_4 PDFVersion = iota
	PDF_VERSION_1_5
)

type PSLevel int

func (self PSLevel) String() string {
	return C.GoString(C.cairo_ps_level_to_string(C.cairo_ps_level_t(self)))
}

const (
	PS_LEVEL_2 PSLevel = iota
	PS_LEVEL_3
)

type SVGVersion int

func (self SVGVersion) String() string {
	return C.GoString(C.cairo_svg_version_to_string(C.cairo_svg_version_t(self)))
}

const (
	SVG_VERSION_1_1 SVGVersion = iota
	SVG_VERSION_1_2
)



func NewPDFSurface(filename string, widthInPoints, heightInPoints float64, version PDFVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_pdf_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_pdf_surface_restrict_to_version(s, C.cairo_pdf_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func NewPSSurface(filename string, widthInPoints, heightInPoints float64, level PSLevel) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_ps_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_ps_surface_restrict_to_level(s, C.cairo_ps_level_t(level))
	return &Surface{surface: s, context: C.cairo_create(s)}
}

func NewSVGSurface(filename string, widthInPoints, heightInPoints float64, version SVGVersion) *Surface {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	s := C.cairo_svg_surface_create(cs, C.double(widthInPoints), C.double(heightInPoints))
	C.cairo_svg_surface_restrict_to_version(s, C.cairo_svg_version_t(version))
	return &Surface{surface: s, context: C.cairo_create(s)}
}