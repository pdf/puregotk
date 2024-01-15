// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// `GtkPaperSize` handles paper sizes.
//
// It uses the standard called
// [PWG 5101.1-2002 PWG: Standard for Media Standardized Names](http://www.pwg.org/standards.html)
// to name the paper sizes (and to get the data for the page sizes).
// In addition to standard paper sizes, `GtkPaperSize` allows to
// construct custom paper sizes with arbitrary dimensions.
//
// The `GtkPaperSize` object stores not only the dimensions (width
// and height) of a paper size and its name, it also provides
// default print margins.
type PaperSize struct {
}

func (x *PaperSize) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewPaperSize func(string) *PaperSize

// Creates a new `GtkPaperSize` object by parsing a
// [PWG 5101.1-2002](ftp://ftp.pwg.org/pub/pwg/candidates/cs-pwgmsn10-20020226-5101.1.pdf)
// paper name.
//
// If @name is %NULL, the default paper size is returned,
// see [func@Gtk.PaperSize.get_default].
func NewPaperSize(NameVar string) *PaperSize {

	cret := xNewPaperSize(NameVar)
	return cret
}

var xNewPaperSizeCustom func(string, string, float64, float64, Unit) *PaperSize

// Creates a new `GtkPaperSize` object with the
// given parameters.
func NewPaperSizeCustom(NameVar string, DisplayNameVar string, WidthVar float64, HeightVar float64, UnitVar Unit) *PaperSize {

	cret := xNewPaperSizeCustom(NameVar, DisplayNameVar, WidthVar, HeightVar, UnitVar)
	return cret
}

var xNewPaperSizeFromGvariant func(*glib.Variant) *PaperSize

// Deserialize a paper size from a `GVariant`.
//
// The `GVariant must be in the format produced by
// [method@Gtk.PaperSize.to_gvariant].
func NewPaperSizeFromGvariant(VariantVar *glib.Variant) *PaperSize {

	cret := xNewPaperSizeFromGvariant(VariantVar)
	return cret
}

var xNewPaperSizeFromIpp func(string, float64, float64) *PaperSize

// Creates a new `GtkPaperSize` object by using
// IPP information.
//
// If @ipp_name is not a recognized paper name,
// @width and @height are used to
// construct a custom `GtkPaperSize` object.
func NewPaperSizeFromIpp(IppNameVar string, WidthVar float64, HeightVar float64) *PaperSize {

	cret := xNewPaperSizeFromIpp(IppNameVar, WidthVar, HeightVar)
	return cret
}

var xNewPaperSizeFromKeyFile func(*glib.KeyFile, string, **glib.Error) *PaperSize

// Reads a paper size from the group @group_name in the key file
// @key_file.
func NewPaperSizeFromKeyFile(KeyFileVar *glib.KeyFile, GroupNameVar string) (*PaperSize, error) {
	var cerr *glib.Error

	cret := xNewPaperSizeFromKeyFile(KeyFileVar, GroupNameVar, &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xNewPaperSizeFromPpd func(string, string, float64, float64) *PaperSize

// Creates a new `GtkPaperSize` object by using
// PPD information.
//
// If @ppd_name is not a recognized PPD paper name,
// @ppd_display_name, @width and @height are used to
// construct a custom `GtkPaperSize` object.
func NewPaperSizeFromPpd(PpdNameVar string, PpdDisplayNameVar string, WidthVar float64, HeightVar float64) *PaperSize {

	cret := xNewPaperSizeFromPpd(PpdNameVar, PpdDisplayNameVar, WidthVar, HeightVar)
	return cret
}

var xPaperSizeCopy func(uintptr) *PaperSize

// Copies an existing `GtkPaperSize`.
func (x *PaperSize) Copy() *PaperSize {

	cret := xPaperSizeCopy(x.GoPointer())
	return cret
}

var xPaperSizeFree func(uintptr)

// Free the given `GtkPaperSize` object.
func (x *PaperSize) Free() {

	xPaperSizeFree(x.GoPointer())

}

var xPaperSizeGetDefaultBottomMargin func(uintptr, Unit) float64

// Gets the default bottom margin for the `GtkPaperSize`.
func (x *PaperSize) GetDefaultBottomMargin(UnitVar Unit) float64 {

	cret := xPaperSizeGetDefaultBottomMargin(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeGetDefaultLeftMargin func(uintptr, Unit) float64

// Gets the default left margin for the `GtkPaperSize`.
func (x *PaperSize) GetDefaultLeftMargin(UnitVar Unit) float64 {

	cret := xPaperSizeGetDefaultLeftMargin(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeGetDefaultRightMargin func(uintptr, Unit) float64

// Gets the default right margin for the `GtkPaperSize`.
func (x *PaperSize) GetDefaultRightMargin(UnitVar Unit) float64 {

	cret := xPaperSizeGetDefaultRightMargin(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeGetDefaultTopMargin func(uintptr, Unit) float64

// Gets the default top margin for the `GtkPaperSize`.
func (x *PaperSize) GetDefaultTopMargin(UnitVar Unit) float64 {

	cret := xPaperSizeGetDefaultTopMargin(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeGetDisplayName func(uintptr) string

// Gets the human-readable name of the `GtkPaperSize`.
func (x *PaperSize) GetDisplayName() string {

	cret := xPaperSizeGetDisplayName(x.GoPointer())
	return cret
}

var xPaperSizeGetHeight func(uintptr, Unit) float64

// Gets the paper height of the `GtkPaperSize`, in
// units of @unit.
func (x *PaperSize) GetHeight(UnitVar Unit) float64 {

	cret := xPaperSizeGetHeight(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeGetName func(uintptr) string

// Gets the name of the `GtkPaperSize`.
func (x *PaperSize) GetName() string {

	cret := xPaperSizeGetName(x.GoPointer())
	return cret
}

var xPaperSizeGetPpdName func(uintptr) string

// Gets the PPD name of the `GtkPaperSize`, which
// may be %NULL.
func (x *PaperSize) GetPpdName() string {

	cret := xPaperSizeGetPpdName(x.GoPointer())
	return cret
}

var xPaperSizeGetWidth func(uintptr, Unit) float64

// Gets the paper width of the `GtkPaperSize`, in
// units of @unit.
func (x *PaperSize) GetWidth(UnitVar Unit) float64 {

	cret := xPaperSizeGetWidth(x.GoPointer(), UnitVar)
	return cret
}

var xPaperSizeIsCustom func(uintptr) bool

// Returns %TRUE if @size is not a standard paper size.
func (x *PaperSize) IsCustom() bool {

	cret := xPaperSizeIsCustom(x.GoPointer())
	return cret
}

var xPaperSizeIsEqual func(uintptr, *PaperSize) bool

// Compares two `GtkPaperSize` objects.
func (x *PaperSize) IsEqual(Size2Var *PaperSize) bool {

	cret := xPaperSizeIsEqual(x.GoPointer(), Size2Var)
	return cret
}

var xPaperSizeIsIpp func(uintptr) bool

// Returns %TRUE if @size is an IPP standard paper size.
func (x *PaperSize) IsIpp() bool {

	cret := xPaperSizeIsIpp(x.GoPointer())
	return cret
}

var xPaperSizeSetSize func(uintptr, float64, float64, Unit)

// Changes the dimensions of a @size to @width x @height.
func (x *PaperSize) SetSize(WidthVar float64, HeightVar float64, UnitVar Unit) {

	xPaperSizeSetSize(x.GoPointer(), WidthVar, HeightVar, UnitVar)

}

var xPaperSizeToGvariant func(uintptr) *glib.Variant

// Serialize a paper size to an `a{sv}` variant.
func (x *PaperSize) ToGvariant() *glib.Variant {

	cret := xPaperSizeToGvariant(x.GoPointer())
	return cret
}

var xPaperSizeToKeyFile func(uintptr, *glib.KeyFile, string)

// This function adds the paper size from @size to @key_file.
func (x *PaperSize) ToKeyFile(KeyFileVar *glib.KeyFile, GroupNameVar string) {

	xPaperSizeToKeyFile(x.GoPointer(), KeyFileVar, GroupNameVar)

}

const (
	// Name for the A3 paper size.
	PAPER_NAME_A3 string = "iso_a3"
	// Name for the A4 paper size.
	PAPER_NAME_A4 string = "iso_a4"
	// Name for the A5 paper size.
	PAPER_NAME_A5 string = "iso_a5"
	// Name for the B5 paper size.
	PAPER_NAME_B5 string = "iso_b5"
	// Name for the Executive paper size.
	PAPER_NAME_EXECUTIVE string = "na_executive"
	// Name for the Legal paper size.
	PAPER_NAME_LEGAL string = "na_legal"
	// Name for the Letter paper size.
	PAPER_NAME_LETTER string = "na_letter"
)

var xPaperSizeGetDefault func() string

// Returns the name of the default paper size, which
// depends on the current locale.
func PaperSizeGetDefault() string {

	cret := xPaperSizeGetDefault()
	return cret
}

var xPaperSizeGetPaperSizes func(bool) *glib.List

// Creates a list of known paper sizes.
func PaperSizeGetPaperSizes(IncludeCustomVar bool) *glib.List {

	cret := xPaperSizeGetPaperSizes(IncludeCustomVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xPaperSizeGetDefault, lib, "gtk_paper_size_get_default")
	core.PuregoSafeRegister(&xPaperSizeGetPaperSizes, lib, "gtk_paper_size_get_paper_sizes")

	core.PuregoSafeRegister(&xNewPaperSize, lib, "gtk_paper_size_new")
	core.PuregoSafeRegister(&xNewPaperSizeCustom, lib, "gtk_paper_size_new_custom")
	core.PuregoSafeRegister(&xNewPaperSizeFromGvariant, lib, "gtk_paper_size_new_from_gvariant")
	core.PuregoSafeRegister(&xNewPaperSizeFromIpp, lib, "gtk_paper_size_new_from_ipp")
	core.PuregoSafeRegister(&xNewPaperSizeFromKeyFile, lib, "gtk_paper_size_new_from_key_file")
	core.PuregoSafeRegister(&xNewPaperSizeFromPpd, lib, "gtk_paper_size_new_from_ppd")

	core.PuregoSafeRegister(&xPaperSizeCopy, lib, "gtk_paper_size_copy")
	core.PuregoSafeRegister(&xPaperSizeFree, lib, "gtk_paper_size_free")
	core.PuregoSafeRegister(&xPaperSizeGetDefaultBottomMargin, lib, "gtk_paper_size_get_default_bottom_margin")
	core.PuregoSafeRegister(&xPaperSizeGetDefaultLeftMargin, lib, "gtk_paper_size_get_default_left_margin")
	core.PuregoSafeRegister(&xPaperSizeGetDefaultRightMargin, lib, "gtk_paper_size_get_default_right_margin")
	core.PuregoSafeRegister(&xPaperSizeGetDefaultTopMargin, lib, "gtk_paper_size_get_default_top_margin")
	core.PuregoSafeRegister(&xPaperSizeGetDisplayName, lib, "gtk_paper_size_get_display_name")
	core.PuregoSafeRegister(&xPaperSizeGetHeight, lib, "gtk_paper_size_get_height")
	core.PuregoSafeRegister(&xPaperSizeGetName, lib, "gtk_paper_size_get_name")
	core.PuregoSafeRegister(&xPaperSizeGetPpdName, lib, "gtk_paper_size_get_ppd_name")
	core.PuregoSafeRegister(&xPaperSizeGetWidth, lib, "gtk_paper_size_get_width")
	core.PuregoSafeRegister(&xPaperSizeIsCustom, lib, "gtk_paper_size_is_custom")
	core.PuregoSafeRegister(&xPaperSizeIsEqual, lib, "gtk_paper_size_is_equal")
	core.PuregoSafeRegister(&xPaperSizeIsIpp, lib, "gtk_paper_size_is_ipp")
	core.PuregoSafeRegister(&xPaperSizeSetSize, lib, "gtk_paper_size_set_size")
	core.PuregoSafeRegister(&xPaperSizeToGvariant, lib, "gtk_paper_size_to_gvariant")
	core.PuregoSafeRegister(&xPaperSizeToKeyFile, lib, "gtk_paper_size_to_key_file")

}
