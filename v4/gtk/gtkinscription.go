// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

type InscriptionClass struct {
	ParentClass uintptr
}

// The different methods to handle text in #GtkInscription when it doesn't
// fit the available space.
type InscriptionOverflow int

const (

	// Clip the remaining text
	InscriptionOverflowClipValue InscriptionOverflow = 0
	// Omit characters at the start of the text
	InscriptionOverflowEllipsizeStartValue InscriptionOverflow = 1
	// Omit characters at the middle of the text
	InscriptionOverflowEllipsizeMiddleValue InscriptionOverflow = 2
	// Omit characters at the end of the text
	InscriptionOverflowEllipsizeEndValue InscriptionOverflow = 3
)

// `GtkInscription` is a widget to show text in a predefined area.
//
// You likely want to use `GtkLabel` instead as this widget is intended only
// for a small subset of use cases. The main scenario envisaged is inside lists
// such as `GtkColumnView`.
//
// While a `GtkLabel` sizes itself depending on the text that is displayed,
// `GtkInscription` is given a size and inscribes the given text into that
// space as well as it can.
//
// Users of this widget should take care to plan behaviour for the common case
// where the text doesn't fit exactly in the allocated space, .
type Inscription struct {
	Widget
}

func InscriptionNewFromInternalPtr(ptr uintptr) *Inscription {
	cls := &Inscription{}
	cls.Ptr = ptr
	return cls
}

var xNewInscription func(string) uintptr

// Creates a new `GtkInscription` with the given text.
func NewInscription(TextVar string) *Widget {
	NewInscriptionPtr := xNewInscription(TextVar)
	if NewInscriptionPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewInscriptionPtr)

	NewInscriptionCls := &Widget{}
	NewInscriptionCls.Ptr = NewInscriptionPtr
	return NewInscriptionCls
}

var xInscriptionGetAttributes func(uintptr) *pango.AttrList

// Gets the inscription's attribute list.
func (x *Inscription) GetAttributes() *pango.AttrList {

	return xInscriptionGetAttributes(x.GoPointer())

}

var xInscriptionGetMinChars func(uintptr) uint

// Gets the `min-chars` of the inscription.
//
// See the [property@Gtk.Inscription:min-chars] property.
func (x *Inscription) GetMinChars() uint {

	return xInscriptionGetMinChars(x.GoPointer())

}

var xInscriptionGetMinLines func(uintptr) uint

// Gets the `min-lines` of the inscription.
//
// See the [property@Gtk.Inscription:min-lines] property.
func (x *Inscription) GetMinLines() uint {

	return xInscriptionGetMinLines(x.GoPointer())

}

var xInscriptionGetNatChars func(uintptr) uint

// Gets the `nat-chars` of the inscription.
//
// See the [property@Gtk.Inscription:nat-chars] property.
func (x *Inscription) GetNatChars() uint {

	return xInscriptionGetNatChars(x.GoPointer())

}

var xInscriptionGetNatLines func(uintptr) uint

// Gets the `nat-lines` of the inscription.
//
// See the [property@Gtk.Inscription:nat-lines] property.
func (x *Inscription) GetNatLines() uint {

	return xInscriptionGetNatLines(x.GoPointer())

}

var xInscriptionGetText func(uintptr) string

// Gets the text that is displayed.
func (x *Inscription) GetText() string {

	return xInscriptionGetText(x.GoPointer())

}

var xInscriptionGetTextOverflow func(uintptr) InscriptionOverflow

// Gets the inscription's overflow method.
func (x *Inscription) GetTextOverflow() InscriptionOverflow {

	return xInscriptionGetTextOverflow(x.GoPointer())

}

var xInscriptionGetWrapMode func(uintptr) pango.WrapMode

// Returns line wrap mode used by the inscription.
//
// See [method@Gtk.Inscription.set_wrap_mode].
func (x *Inscription) GetWrapMode() pango.WrapMode {

	return xInscriptionGetWrapMode(x.GoPointer())

}

var xInscriptionGetXalign func(uintptr) float32

// Gets the `xalign` of the inscription.
//
// See the [property@Gtk.Inscription:xalign] property.
func (x *Inscription) GetXalign() float32 {

	return xInscriptionGetXalign(x.GoPointer())

}

var xInscriptionGetYalign func(uintptr) float32

// Gets the `yalign` of the inscription.
//
// See the [property@Gtk.Inscription:yalign] property.
func (x *Inscription) GetYalign() float32 {

	return xInscriptionGetYalign(x.GoPointer())

}

var xInscriptionSetAttributes func(uintptr, *pango.AttrList)

// Apply attributes to the inscription text.
//
// These attributes will not be evaluated for sizing the inscription.
func (x *Inscription) SetAttributes(AttrsVar *pango.AttrList) {

	xInscriptionSetAttributes(x.GoPointer(), AttrsVar)

}

var xInscriptionSetMarkup func(uintptr, string)

// Utility function to set the text and attributes to be displayed.
//
// See the [property@Gtk.Inscription:markup] property.
func (x *Inscription) SetMarkup(MarkupVar string) {

	xInscriptionSetMarkup(x.GoPointer(), MarkupVar)

}

var xInscriptionSetMinChars func(uintptr, uint)

// Sets the `min-chars` of the inscription.
//
// See the [property@Gtk.Inscription:min-chars] property.
func (x *Inscription) SetMinChars(MinCharsVar uint) {

	xInscriptionSetMinChars(x.GoPointer(), MinCharsVar)

}

var xInscriptionSetMinLines func(uintptr, uint)

// Sets the `min-lines` of the inscription.
//
// See the [property@Gtk.Inscription:min-lines] property.
func (x *Inscription) SetMinLines(MinLinesVar uint) {

	xInscriptionSetMinLines(x.GoPointer(), MinLinesVar)

}

var xInscriptionSetNatChars func(uintptr, uint)

// Sets the `nat-chars` of the inscription.
//
// See the [property@Gtk.Inscription:nat-chars] property.
func (x *Inscription) SetNatChars(NatCharsVar uint) {

	xInscriptionSetNatChars(x.GoPointer(), NatCharsVar)

}

var xInscriptionSetNatLines func(uintptr, uint)

// Sets the `nat-lines` of the inscription.
//
// See the [property@Gtk.Inscription:nat-lines] property.
func (x *Inscription) SetNatLines(NatLinesVar uint) {

	xInscriptionSetNatLines(x.GoPointer(), NatLinesVar)

}

var xInscriptionSetText func(uintptr, string)

// Sets the text to be displayed.
func (x *Inscription) SetText(TextVar string) {

	xInscriptionSetText(x.GoPointer(), TextVar)

}

var xInscriptionSetTextOverflow func(uintptr, InscriptionOverflow)

// Sets what to do when the text doesn't fit.
func (x *Inscription) SetTextOverflow(OverflowVar InscriptionOverflow) {

	xInscriptionSetTextOverflow(x.GoPointer(), OverflowVar)

}

var xInscriptionSetWrapMode func(uintptr, pango.WrapMode)

// Controls how line wrapping is done.
func (x *Inscription) SetWrapMode(WrapModeVar pango.WrapMode) {

	xInscriptionSetWrapMode(x.GoPointer(), WrapModeVar)

}

var xInscriptionSetXalign func(uintptr, float32)

// Sets the `xalign` of the inscription.
//
// See the [property@Gtk.Inscription:xalign] property.
func (x *Inscription) SetXalign(XalignVar float32) {

	xInscriptionSetXalign(x.GoPointer(), XalignVar)

}

var xInscriptionSetYalign func(uintptr, float32)

// Sets the `yalign` of the inscription.
//
// See the [property@Gtk.Inscription:yalign] property.
func (x *Inscription) SetYalign(YalignVar float32) {

	xInscriptionSetYalign(x.GoPointer(), YalignVar)

}

func (c *Inscription) GoPointer() uintptr {
	return c.Ptr
}

func (c *Inscription) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Inscription) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Inscription) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Inscription) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Inscription) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

}

// Updates a list of accessible properties.
//
// See the [enum@Gtk.AccessibleProperty] documentation for the
// value types of accessible properties.
//
// This function should be called by `GtkWidget` types whenever
// an accessible property change must be communicated to assistive
// technologies.
//
// Example:
// ```c
// value = gtk_adjustment_get_value (adjustment);
// gtk_accessible_update_property (GTK_ACCESSIBLE (spin_button),
//
//	GTK_ACCESSIBLE_PROPERTY_VALUE_NOW, value,
//	-1);
//
// ```
func (x *Inscription) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Inscription) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

}

// Updates a list of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// If the [enum@Gtk.AccessibleRelation] requires a list of references,
// you should pass each reference individually, followed by %NULL, e.g.
//
// ```c
// gtk_accessible_update_relation (accessible,
//
//	GTK_ACCESSIBLE_RELATION_CONTROLS,
//	  ref1, NULL,
//	GTK_ACCESSIBLE_RELATION_LABELLED_BY,
//	  ref1, ref2, ref3, NULL,
//	-1);
//
// ```
func (x *Inscription) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Inscription) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

}

// Updates a list of accessible states. See the [enum@Gtk.AccessibleState]
// documentation for the value types of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// Example:
// ```c
// value = GTK_ACCESSIBLE_TRISTATE_MIXED;
// gtk_accessible_update_state (GTK_ACCESSIBLE (check_button),
//
//	GTK_ACCESSIBLE_STATE_CHECKED, value,
//	-1);
//
// ```
func (x *Inscription) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Inscription) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Inscription) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewInscription, lib, "gtk_inscription_new")

	core.PuregoSafeRegister(&xInscriptionGetAttributes, lib, "gtk_inscription_get_attributes")
	core.PuregoSafeRegister(&xInscriptionGetMinChars, lib, "gtk_inscription_get_min_chars")
	core.PuregoSafeRegister(&xInscriptionGetMinLines, lib, "gtk_inscription_get_min_lines")
	core.PuregoSafeRegister(&xInscriptionGetNatChars, lib, "gtk_inscription_get_nat_chars")
	core.PuregoSafeRegister(&xInscriptionGetNatLines, lib, "gtk_inscription_get_nat_lines")
	core.PuregoSafeRegister(&xInscriptionGetText, lib, "gtk_inscription_get_text")
	core.PuregoSafeRegister(&xInscriptionGetTextOverflow, lib, "gtk_inscription_get_text_overflow")
	core.PuregoSafeRegister(&xInscriptionGetWrapMode, lib, "gtk_inscription_get_wrap_mode")
	core.PuregoSafeRegister(&xInscriptionGetXalign, lib, "gtk_inscription_get_xalign")
	core.PuregoSafeRegister(&xInscriptionGetYalign, lib, "gtk_inscription_get_yalign")
	core.PuregoSafeRegister(&xInscriptionSetAttributes, lib, "gtk_inscription_set_attributes")
	core.PuregoSafeRegister(&xInscriptionSetMarkup, lib, "gtk_inscription_set_markup")
	core.PuregoSafeRegister(&xInscriptionSetMinChars, lib, "gtk_inscription_set_min_chars")
	core.PuregoSafeRegister(&xInscriptionSetMinLines, lib, "gtk_inscription_set_min_lines")
	core.PuregoSafeRegister(&xInscriptionSetNatChars, lib, "gtk_inscription_set_nat_chars")
	core.PuregoSafeRegister(&xInscriptionSetNatLines, lib, "gtk_inscription_set_nat_lines")
	core.PuregoSafeRegister(&xInscriptionSetText, lib, "gtk_inscription_set_text")
	core.PuregoSafeRegister(&xInscriptionSetTextOverflow, lib, "gtk_inscription_set_text_overflow")
	core.PuregoSafeRegister(&xInscriptionSetWrapMode, lib, "gtk_inscription_set_wrap_mode")
	core.PuregoSafeRegister(&xInscriptionSetXalign, lib, "gtk_inscription_set_xalign")
	core.PuregoSafeRegister(&xInscriptionSetYalign, lib, "gtk_inscription_set_yalign")

}
