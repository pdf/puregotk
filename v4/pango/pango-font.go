// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type FontClass struct {
	ParentClass uintptr
}

// A `PangoFontDescription` describes a font in an implementation-independent
// manner.
//
// `PangoFontDescription` structures are used both to list what fonts are
// available on the system and also for specifying the characteristics of
// a font to load.
type FontDescription struct {
}

type FontFaceClass struct {
	ParentClass uintptr
}

type FontFamilyClass struct {
	ParentClass uintptr
}

// A `PangoFontMetrics` structure holds the overall metric information
// for a font.
//
// The information in a `PangoFontMetrics` structure may be restricted
// to a script. The fields of this structure are private to implementations
// of a font backend. See the documentation of the corresponding getters
// for documentation of their meaning.
//
// For an overview of the most important metrics, see:
//
// &lt;picture&gt;
//
//	&lt;source srcset="fontmetrics-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img alt="Font metrics" src="fontmetrics-light.png"&gt;
//
// &lt;/picture&gt;
type FontMetrics struct {
	RefCount uint

	Ascent int

	Descent int

	Height int

	ApproximateCharWidth int

	ApproximateDigitWidth int

	UnderlinePosition int

	UnderlineThickness int

	StrikethroughPosition int

	StrikethroughThickness int
}

// The bits in a `PangoFontMask` correspond to the set fields in a
// `PangoFontDescription`.
type FontMask int

const (

	// the font family is specified.
	FontMaskFamilyValue FontMask = 1
	// the font style is specified.
	FontMaskStyleValue FontMask = 2
	// the font variant is specified.
	FontMaskVariantValue FontMask = 4
	// the font weight is specified.
	FontMaskWeightValue FontMask = 8
	// the font stretch is specified.
	FontMaskStretchValue FontMask = 16
	// the font size is specified.
	FontMaskSizeValue FontMask = 32
	// the font gravity is specified (Since: 1.16.)
	FontMaskGravityValue FontMask = 64
	// OpenType font variations are specified (Since: 1.42)
	FontMaskVariationsValue FontMask = 128
)

// An enumeration specifying the width of the font relative to other designs
// within a family.
type Stretch int

const (

	// ultra condensed width
	StretchUltraCondensedValue Stretch = 0
	// extra condensed width
	StretchExtraCondensedValue Stretch = 1
	// condensed width
	StretchCondensedValue Stretch = 2
	// semi condensed width
	StretchSemiCondensedValue Stretch = 3
	// the normal width
	StretchNormalValue Stretch = 4
	// semi expanded width
	StretchSemiExpandedValue Stretch = 5
	// expanded width
	StretchExpandedValue Stretch = 6
	// extra expanded width
	StretchExtraExpandedValue Stretch = 7
	// ultra expanded width
	StretchUltraExpandedValue Stretch = 8
)

// An enumeration specifying the various slant styles possible for a font.
type Style int

const (

	// the font is upright.
	StyleNormalValue Style = 0
	// the font is slanted, but in a roman style.
	StyleObliqueValue Style = 1
	// the font is slanted in an italic style.
	StyleItalicValue Style = 2
)

// An enumeration specifying capitalization variant of the font.
//
// Values other than `PANGO_VARIANT_NORMAL` and `PANGO_VARIANT_SMALL_CAPS` are
// available since 1.50.
type Variant int

const (

	// A normal font.
	VariantNormalValue Variant = 0
	// A font with the lower case characters
	//   replaced by smaller variants of the capital characters.
	VariantSmallCapsValue Variant = 1
	// A font with all characters
	//   replaced by smaller variants of the capital characters.
	VariantAllSmallCapsValue Variant = 2
	// A font with the lower case characters
	//   replaced by smaller variants of the capital characters.
	//   Petite Caps can be even smaller than Small Caps.
	VariantPetiteCapsValue Variant = 3
	// A font with all characters
	//   replaced by smaller variants of the capital characters.
	//   Petite Caps can be even smaller than Small Caps.
	VariantAllPetiteCapsValue Variant = 4
	// A font with the upper case characters
	//   replaced by smaller variants of the capital letters.
	VariantUnicaseValue Variant = 5
	// A font with capital letters that
	//   are more suitable for all-uppercase titles.
	VariantTitleCapsValue Variant = 6
)

// An enumeration specifying the weight (boldness) of a font.
//
// Weight is specified as a numeric value ranging from 100 to 1000.
// This enumeration simply provides some common, predefined values.
type Weight int

const (

	// the thin weight (= 100) Since: 1.24
	WeightThinValue Weight = 100
	// the ultralight weight (= 200)
	WeightUltralightValue Weight = 200
	// the light weight (= 300)
	WeightLightValue Weight = 300
	// the semilight weight (= 350) Since: 1.36.7
	WeightSemilightValue Weight = 350
	// the book weight (= 380) Since: 1.24)
	WeightBookValue Weight = 380
	// the default weight (= 400)
	WeightNormalValue Weight = 400
	// the normal weight (= 500) Since: 1.24
	WeightMediumValue Weight = 500
	// the semibold weight (= 600)
	WeightSemiboldValue Weight = 600
	// the bold weight (= 700)
	WeightBoldValue Weight = 700
	// the ultrabold weight (= 800)
	WeightUltraboldValue Weight = 800
	// the heavy weight (= 900)
	WeightHeavyValue Weight = 900
	// the ultraheavy weight (= 1000) Since: 1.24
	WeightUltraheavyValue Weight = 1000
)

var xFontDescriptionFromString func(string) *FontDescription

// Creates a new font description from a string representation.
//
// The string must have the form
//
//	"\[FAMILY-LIST] \[STYLE-OPTIONS] \[SIZE] \[VARIATIONS]",
//
// where FAMILY-LIST is a comma-separated list of families optionally
// terminated by a comma, STYLE_OPTIONS is a whitespace-separated list
// of words where each word describes one of style, variant, weight,
// stretch, or gravity, and SIZE is a decimal number (size in points)
// or optionally followed by the unit modifier "px" for absolute size.
// VARIATIONS is a comma-separated list of font variation
// specifications of the form "\@axis=value" (the = sign is optional).
//
// The following words are understood as styles:
// "Normal", "Roman", "Oblique", "Italic".
//
// The following words are understood as variants:
// "Small-Caps", "All-Small-Caps", "Petite-Caps", "All-Petite-Caps",
// "Unicase", "Title-Caps".
//
// The following words are understood as weights:
// "Thin", "Ultra-Light", "Extra-Light", "Light", "Semi-Light",
// "Demi-Light", "Book", "Regular", "Medium", "Semi-Bold", "Demi-Bold",
// "Bold", "Ultra-Bold", "Extra-Bold", "Heavy", "Black", "Ultra-Black",
// "Extra-Black".
//
// The following words are understood as stretch values:
// "Ultra-Condensed", "Extra-Condensed", "Condensed", "Semi-Condensed",
// "Semi-Expanded", "Expanded", "Extra-Expanded", "Ultra-Expanded".
//
// The following words are understood as gravity values:
// "Not-Rotated", "South", "Upside-Down", "North", "Rotated-Left",
// "East", "Rotated-Right", "West".
//
// Any one of the options may be absent. If FAMILY-LIST is absent, then
// the family_name field of the resulting font description will be
// initialized to %NULL. If STYLE-OPTIONS is missing, then all style
// options will be set to the default values. If SIZE is missing, the
// size in the resulting font description will be set to 0.
//
// A typical example:
//
//	"Cantarell Italic Light 15 \@wght=200"
func FontDescriptionFromString(StrVar string) *FontDescription {

	return xFontDescriptionFromString(StrVar)

}

// A `PangoFont` is used to represent a font in a
// rendering-system-independent manner.
type Font struct {
	gobject.Object
}

func FontNewFromInternalPtr(ptr uintptr) *Font {
	cls := &Font{}
	cls.Ptr = ptr
	return cls
}

var xFontDescribe func(uintptr) *FontDescription

// Returns a description of the font, with font size set in points.
//
// Use [method@Pango.Font.describe_with_absolute_size] if you want
// the font size in device units.
func (x *Font) Describe() *FontDescription {

	return xFontDescribe(x.GoPointer())

}

var xFontDescribeWithAbsoluteSize func(uintptr) *FontDescription

// Returns a description of the font, with absolute font size set
// in device units.
//
// Use [method@Pango.Font.describe] if you want the font size in points.
func (x *Font) DescribeWithAbsoluteSize() *FontDescription {

	return xFontDescribeWithAbsoluteSize(x.GoPointer())

}

var xFontGetCoverage func(uintptr, *Language) uintptr

// Computes the coverage map for a given font and language tag.
func (x *Font) GetCoverage(LanguageVar *Language) *Coverage {

	GetCoveragePtr := xFontGetCoverage(x.GoPointer(), LanguageVar)
	if GetCoveragePtr == 0 {
		return nil
	}

	GetCoverageCls := &Coverage{}
	GetCoverageCls.Ptr = GetCoveragePtr
	return GetCoverageCls

}

var xFontGetFace func(uintptr) uintptr

// Gets the `PangoFontFace` to which @font belongs.
func (x *Font) GetFace() *FontFace {

	GetFacePtr := xFontGetFace(x.GoPointer())
	if GetFacePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFacePtr)

	GetFaceCls := &FontFace{}
	GetFaceCls.Ptr = GetFacePtr
	return GetFaceCls

}

var xFontGetFeatures func(uintptr, uintptr, uint, uint)

// Obtain the OpenType features that are provided by the font.
//
// These are passed to the rendering system, together with features
// that have been explicitly set via attributes.
//
// Note that this does not include OpenType features which the
// rendering system enables by default.
func (x *Font) GetFeatures(FeaturesVar uintptr, LenVar uint, NumFeaturesVar uint) {

	xFontGetFeatures(x.GoPointer(), FeaturesVar, LenVar, NumFeaturesVar)

}

var xFontGetFontMap func(uintptr) uintptr

// Gets the font map for which the font was created.
//
// Note that the font maintains a *weak* reference to
// the font map, so if all references to font map are
// dropped, the font map will be finalized even if there
// are fonts created with the font map that are still alive.
// In that case this function will return %NULL.
//
// It is the responsibility of the user to ensure that the
// font map is kept alive. In most uses this is not an issue
// as a `PangoContext` holds a reference to the font map.
func (x *Font) GetFontMap() *FontMap {

	GetFontMapPtr := xFontGetFontMap(x.GoPointer())
	if GetFontMapPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFontMapPtr)

	GetFontMapCls := &FontMap{}
	GetFontMapCls.Ptr = GetFontMapPtr
	return GetFontMapCls

}

var xFontGetGlyphExtents func(uintptr, Glyph, *Rectangle, *Rectangle)

// Gets the logical and ink extents of a glyph within a font.
//
// The coordinate system for each rectangle has its origin at the
// base line and horizontal origin of the character with increasing
// coordinates extending to the right and down. The macros PANGO_ASCENT(),
// PANGO_DESCENT(), PANGO_LBEARING(), and PANGO_RBEARING() can be used to convert
// from the extents rectangle to more traditional font metrics. The units
// of the rectangles are in 1/PANGO_SCALE of a device unit.
//
// If @font is %NULL, this function gracefully sets some sane values in the
// output variables and returns.
func (x *Font) GetGlyphExtents(GlyphVar Glyph, InkRectVar *Rectangle, LogicalRectVar *Rectangle) {

	xFontGetGlyphExtents(x.GoPointer(), GlyphVar, InkRectVar, LogicalRectVar)

}

var xFontGetHbFont func(uintptr) uintptr

// Get a `hb_font_t` object backing this font.
//
// Note that the objects returned by this function are cached
// and immutable. If you need to make changes to the `hb_font_t`,
// use [hb_font_create_sub_font()](https://harfbuzz.github.io/harfbuzz-hb-font.html#hb-font-create-sub-font).
func (x *Font) GetHbFont() uintptr {

	return xFontGetHbFont(x.GoPointer())

}

var xFontGetLanguages func(uintptr) uintptr

// Returns the languages that are supported by @font.
//
// If the font backend does not provide this information,
// %NULL is returned. For the fontconfig backend, this
// corresponds to the FC_LANG member of the FcPattern.
//
// The returned array is only valid as long as the font
// and its fontmap are valid.
func (x *Font) GetLanguages() uintptr {

	return xFontGetLanguages(x.GoPointer())

}

var xFontGetMetrics func(uintptr, *Language) *FontMetrics

// Gets overall metric information for a font.
//
// Since the metrics may be substantially different for different scripts,
// a language tag can be provided to indicate that the metrics should be
// retrieved that correspond to the script(s) used by that language.
//
// If @font is %NULL, this function gracefully sets some sane values in the
// output variables and returns.
func (x *Font) GetMetrics(LanguageVar *Language) *FontMetrics {

	return xFontGetMetrics(x.GoPointer(), LanguageVar)

}

var xFontHasChar func(uintptr, uint32) bool

// Returns whether the font provides a glyph for this character.
//
// Returns %TRUE if @font can render @wc
func (x *Font) HasChar(WcVar uint32) bool {

	return xFontHasChar(x.GoPointer(), WcVar)

}

var xFontSerialize func(uintptr) *glib.Bytes

// Serializes the @font in a way that can be uniquely identified.
//
// There are no guarantees about the format of the output across different
// versions of Pango.
//
// The intended use of this function is testing, benchmarking and debugging.
// The format is not meant as a permanent storage format.
//
// To recreate a font from its serialized form, use [func@Pango.Font.deserialize].
func (x *Font) Serialize() *glib.Bytes {

	return xFontSerialize(x.GoPointer())

}

func (c *Font) GoPointer() uintptr {
	return c.Ptr
}

func (c *Font) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `PangoFontFace` is used to represent a group of fonts with
// the same family, slant, weight, and width, but varying sizes.
type FontFace struct {
	gobject.Object
}

func FontFaceNewFromInternalPtr(ptr uintptr) *FontFace {
	cls := &FontFace{}
	cls.Ptr = ptr
	return cls
}

var xFontFaceDescribe func(uintptr) *FontDescription

// Returns a font description that matches the face.
//
// The resulting font description will have the family, style,
// variant, weight and stretch of the face, but its size field
// will be unset.
func (x *FontFace) Describe() *FontDescription {

	return xFontFaceDescribe(x.GoPointer())

}

var xFontFaceGetFaceName func(uintptr) string

// Gets a name representing the style of this face.
//
// The name identifies the face among the different faces
// in the `PangoFontFamily` for the face. It is suitable
// for displaying to users.
func (x *FontFace) GetFaceName() string {

	return xFontFaceGetFaceName(x.GoPointer())

}

var xFontFaceGetFamily func(uintptr) uintptr

// Gets the `PangoFontFamily` that @face belongs to.
func (x *FontFace) GetFamily() *FontFamily {

	GetFamilyPtr := xFontFaceGetFamily(x.GoPointer())
	if GetFamilyPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFamilyPtr)

	GetFamilyCls := &FontFamily{}
	GetFamilyCls.Ptr = GetFamilyPtr
	return GetFamilyCls

}

var xFontFaceIsSynthesized func(uintptr) bool

// Returns whether a `PangoFontFace` is synthesized.
//
// This will be the case if the underlying font rendering engine
// creates this face from another face, by shearing, emboldening,
// lightening or modifying it in some other way.
func (x *FontFace) IsSynthesized() bool {

	return xFontFaceIsSynthesized(x.GoPointer())

}

var xFontFaceListSizes func(uintptr, uintptr, int)

// List the available sizes for a font.
//
// This is only applicable to bitmap fonts. For scalable fonts, stores
// %NULL at the location pointed to by @sizes and 0 at the location pointed
// to by @n_sizes. The sizes returned are in Pango units and are sorted
// in ascending order.
func (x *FontFace) ListSizes(SizesVar uintptr, NSizesVar int) {

	xFontFaceListSizes(x.GoPointer(), SizesVar, NSizesVar)

}

func (c *FontFace) GoPointer() uintptr {
	return c.Ptr
}

func (c *FontFace) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A `PangoFontFamily` is used to represent a family of related
// font faces.
//
// The font faces in a family share a common design, but differ in
// slant, weight, width or other aspects.
type FontFamily struct {
	gobject.Object
}

func FontFamilyNewFromInternalPtr(ptr uintptr) *FontFamily {
	cls := &FontFamily{}
	cls.Ptr = ptr
	return cls
}

var xFontFamilyGetFace func(uintptr, string) uintptr

// Gets the `PangoFontFace` of @family with the given name.
func (x *FontFamily) GetFace(NameVar string) *FontFace {

	GetFacePtr := xFontFamilyGetFace(x.GoPointer(), NameVar)
	if GetFacePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFacePtr)

	GetFaceCls := &FontFace{}
	GetFaceCls.Ptr = GetFacePtr
	return GetFaceCls

}

var xFontFamilyGetName func(uintptr) string

// Gets the name of the family.
//
// The name is unique among all fonts for the font backend and can
// be used in a `PangoFontDescription` to specify that a face from
// this family is desired.
func (x *FontFamily) GetName() string {

	return xFontFamilyGetName(x.GoPointer())

}

var xFontFamilyIsMonospace func(uintptr) bool

// A monospace font is a font designed for text display where the the
// characters form a regular grid.
//
// For Western languages this would
// mean that the advance width of all characters are the same, but
// this categorization also includes Asian fonts which include
// double-width characters: characters that occupy two grid cells.
// g_unichar_iswide() returns a result that indicates whether a
// character is typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// [method@Pango.FontMetrics.get_approximate_digit_width], since the
// results of [method@Pango.FontMetrics.get_approximate_char_width] may
// be affected by double-width characters.
func (x *FontFamily) IsMonospace() bool {

	return xFontFamilyIsMonospace(x.GoPointer())

}

var xFontFamilyIsVariable func(uintptr) bool

// A variable font is a font which has axes that can be modified to
// produce different faces.
//
// Such axes are also known as _variations_; see
// [method@Pango.FontDescription.set_variations] for more information.
func (x *FontFamily) IsVariable() bool {

	return xFontFamilyIsVariable(x.GoPointer())

}

var xFontFamilyListFaces func(uintptr, uintptr, int)

// Lists the different font faces that make up @family.
//
// The faces in a family share a common design, but differ in slant, weight,
// width and other aspects.
func (x *FontFamily) ListFaces(FacesVar uintptr, NFacesVar int) {

	xFontFamilyListFaces(x.GoPointer(), FacesVar, NFacesVar)

}

func (c *FontFamily) GoPointer() uintptr {
	return c.Ptr
}

func (c *FontFamily) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xFontDescriptionFromString, lib, "pango_font_description_from_string")

	core.PuregoSafeRegister(&xFontDescribe, lib, "pango_font_describe")
	core.PuregoSafeRegister(&xFontDescribeWithAbsoluteSize, lib, "pango_font_describe_with_absolute_size")
	core.PuregoSafeRegister(&xFontGetCoverage, lib, "pango_font_get_coverage")
	core.PuregoSafeRegister(&xFontGetFace, lib, "pango_font_get_face")
	core.PuregoSafeRegister(&xFontGetFeatures, lib, "pango_font_get_features")
	core.PuregoSafeRegister(&xFontGetFontMap, lib, "pango_font_get_font_map")
	core.PuregoSafeRegister(&xFontGetGlyphExtents, lib, "pango_font_get_glyph_extents")
	core.PuregoSafeRegister(&xFontGetHbFont, lib, "pango_font_get_hb_font")
	core.PuregoSafeRegister(&xFontGetLanguages, lib, "pango_font_get_languages")
	core.PuregoSafeRegister(&xFontGetMetrics, lib, "pango_font_get_metrics")
	core.PuregoSafeRegister(&xFontHasChar, lib, "pango_font_has_char")
	core.PuregoSafeRegister(&xFontSerialize, lib, "pango_font_serialize")

	core.PuregoSafeRegister(&xFontFaceDescribe, lib, "pango_font_face_describe")
	core.PuregoSafeRegister(&xFontFaceGetFaceName, lib, "pango_font_face_get_face_name")
	core.PuregoSafeRegister(&xFontFaceGetFamily, lib, "pango_font_face_get_family")
	core.PuregoSafeRegister(&xFontFaceIsSynthesized, lib, "pango_font_face_is_synthesized")
	core.PuregoSafeRegister(&xFontFaceListSizes, lib, "pango_font_face_list_sizes")

	core.PuregoSafeRegister(&xFontFamilyGetFace, lib, "pango_font_family_get_face")
	core.PuregoSafeRegister(&xFontFamilyGetName, lib, "pango_font_family_get_name")
	core.PuregoSafeRegister(&xFontFamilyIsMonospace, lib, "pango_font_family_is_monospace")
	core.PuregoSafeRegister(&xFontFamilyIsVariable, lib, "pango_font_family_is_variable")
	core.PuregoSafeRegister(&xFontFamilyListFaces, lib, "pango_font_family_list_faces")

}
