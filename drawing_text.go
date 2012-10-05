package cairo

// #cgo LDFLAGS: -lcairo
// #include <cairo/cairo.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)



type FontExtents struct {       // typedef struct {
    Ascent        float64       //     double ascent;
    Descent       float64       //     double descent;
    Height        float64       //     double height;
    Max_x_advance float64       //     double max_x_advance;
    Max_y_advance float64       //     double max_y_advance;
}                               // } cairo_font_extents_t;

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




type Glyph struct {             // typedef struct {
    Index   uint32              //     unsigned long  index;
    X, Y    float64             //     double  x; double  y;
}                               // } cairo_glyph_t;

// void cairo_show_glyphs (cairo_t *cr, const cairo_glyph_t *glyphs, int num_glyphs);
func (self *Surface) ShowGlyphs(glyphs []Glyph) {
    C.cairo_show_glyphs( self.context, (*C.cairo_glyph_t)(unsafe.Pointer(&glyphs[0])), C.int(len(glyphs)) )
}


type TextCluster struct {       // typedef struct {
    NumBytes   int              //     int        num_bytes;
    NumGlyphs  int              //     int        num_glyphs;
}                               // } cairo_text_cluster_t;

type TextClusterFlag int
// cairo_text_cluster_flag_t
const (
    // TextClusterFlagBackward TextClusterFlag = 1 << iota
    TEXT_CLUSTER_FLAG_BACKWARD TextClusterFlag = 0x00000001
)

// void cairo_show_text_glyphs(cairo_t *cr, const char *utf8, int utf8_len,
//                             const cairo_glyph_t *glyphs, int num_glyphs,
//                             const cairo_text_cluster_t *clusters, int num_clusters,
//                             cairo_text_cluster_flags_t cluster_flags );
func (self *Surface) ShowTextGlyphs(text string, glyphs []Glyph, clusters []TextCluster, flag TextClusterFlag) {
    utf8 := C.CString(text)
    defer C.free(unsafe.Pointer(utf8))

    C.cairo_show_text_glyphs( self.context, utf8, C.int(len(text)),
                              (*C.cairo_glyph_t)(unsafe.Pointer(&glyphs[0])), C.int(len(glyphs)),
                              (*C.cairo_text_cluster_t)(unsafe.Pointer(&clusters[0])), C.int(len(clusters)),
                              C.cairo_text_cluster_flags_t(flag) );
}
