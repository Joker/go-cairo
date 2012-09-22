package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe";
)


/*
typedef struct {
    double ascent;
    double descent;
    double height;
    double max_x_advance;
    double max_y_advance;
} cairo_font_extents_t;

type FontExtents struct {
    extents *C.cairo_font_extents_t;
}
*/


type FontExtents struct {
    Ascent        float64
    Descent       float64
    Height        float64
    Max_x_advance float64
    Max_y_advance float64
}
// void cairo_font_extents(cairo_t *cr, cairo_font_extents_t *extents);
func (self *Surface) FontExtents(extents *FontExtents){
	C.cairo_font_extents(self.context, (*C.cairo_font_extents_t)(unsafe.Pointer(extents)));
}


type TextExtents struct {
    X_bearing float64
    Y_bearing float64
    Width     float64
    Height    float64
    X_advance float64
    Y_advance float64
}
// void cairo_text_extents (cairo_t *cr, const char *utf8, cairo_text_extents_t *extents);
func (self *Surface) TextExtents(utf8 string, extents *TextExtents){
    C.cairo_text_extents(self.context, C.CString(utf8),(*C.cairo_text_extents_t)(unsafe.Pointer(extents)));
}


func (self *Surface) GlyphExtents(glyphs []Glyph) *TextExtents {
    panic("not implemented") // todo
    //C.cairo_glyph_extents
    return nil
}


type TextClusterFlag int
// cairo_text_cluster_flag_t
const (
    // TextClusterFlagBackward TextClusterFlag = 1 << iota
    TEXT_CLUSTER_FLAG_BACKWARD TextClusterFlag = 0x00000001
)

type FontSlant int
// cairo_font_slant_t 
const (
    FONT_SLANT_NORMAL FontSlant = iota
    FONT_SLANT_ITALIC
    FONT_SLANT_OBLIQUE
)

type FontWeight int
// cairo_font_weight_t 
const (
    FONT_WEIGHT_NORMAL FontWeight = iota
    FONT_WEIGHT_BOLD
)

// void cairo_select_font_face (cairo_t *cr, const char *family, cairo_font_slant_t slant, cairo_font_weight_t weight);
func (self *Surface) SelectFontFace(name string, font_slant FontSlant, font_weight FontWeight) {
    p := C.CString(name);
    C.cairo_select_font_face(self.context, p, C.cairo_font_slant_t(font_slant), C.cairo_font_weight_t(font_weight));
    C.free(unsafe.Pointer(p));
}

// void cairo_set_font_size (cairo_t *cr, double size);
func (self *Surface) SetFontSize(size float64) {
    C.cairo_set_font_size(self.context, C.double(size))
}

// void cairo_show_text (cairo_t *cr, const char *utf8);
func (self *Surface) ShowText(text string) {
    p := C.CString(text);
    C.cairo_show_text(self.context, p);
    C.free(unsafe.Pointer(p));
}

func (self *Surface) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flags TextClusterFlag) {
    panic("not implemented") // todo
}
func (self *Surface) ShowGlyphs(glyphs []Glyph) {
    panic("not implemented") // todo
}
