// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The `GtkColorChooserWidget` widget lets the user select a color.
//
// By default, the chooser presents a predefined palette of colors,
// plus a small number of settable custom colors. It is also possible
// to select a different color with the single-color editor.
//
// To enter the single-color editing mode, use the context menu of any
// color of the palette, or use the '+' button to add a new custom color.
//
// The chooser automatically remembers the last selection, as well
// as custom colors.
//
// To create a `GtkColorChooserWidget`, use [ctor@Gtk.ColorChooserWidget.new].
//
// To change the initially selected color, use
// [method@Gtk.ColorChooser.set_rgba]. To get the selected color use
// [method@Gtk.ColorChooser.get_rgba].
//
// The `GtkColorChooserWidget` is used in the [class@Gtk.ColorChooserDialog]
// to provide a dialog for selecting colors.
//
// # CSS names
//
// `GtkColorChooserWidget` has a single CSS node with name colorchooser.
type ColorChooserWidget struct {
	Widget
}

func ColorChooserWidgetNewFromInternalPtr(ptr uintptr) *ColorChooserWidget {
	cls := &ColorChooserWidget{}
	cls.Ptr = ptr
	return cls
}

var xNewColorChooserWidget func() uintptr

// Creates a new `GtkColorChooserWidget`.
func NewColorChooserWidget() *Widget {
	NewColorChooserWidgetPtr := xNewColorChooserWidget()
	if NewColorChooserWidgetPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewColorChooserWidgetPtr)

	NewColorChooserWidgetCls := &Widget{}
	NewColorChooserWidgetCls.Ptr = NewColorChooserWidgetPtr
	return NewColorChooserWidgetCls
}

func (c *ColorChooserWidget) GoPointer() uintptr {
	return c.Ptr
}

func (c *ColorChooserWidget) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ColorChooserWidget) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *ColorChooserWidget) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ColorChooserWidget) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ColorChooserWidget) ResetState(StateVar AccessibleState) {

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
func (x *ColorChooserWidget) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorChooserWidget) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ColorChooserWidget) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorChooserWidget) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ColorChooserWidget) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ColorChooserWidget) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ColorChooserWidget) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Adds a palette to the color chooser.
//
// If @orientation is horizontal, the colors are grouped in rows,
// with @colors_per_line colors in each row. If @horizontal is %FALSE,
// the colors are grouped in columns instead.
//
// The default color palette of [class@Gtk.ColorChooserWidget] has
// 45 colors, organized in columns of 5 colors (this includes some
// grays).
//
// The layout of the color chooser widget works best when the
// palettes have 9-10 columns.
//
// Calling this function for the first time has the side effect
// of removing the default color palette from the color chooser.
//
// If @colors is %NULL, removes all previously added palettes.
func (x *ColorChooserWidget) AddPalette(OrientationVar Orientation, ColorsPerLineVar int, NColorsVar int, ColorsVar uintptr) {

	XGtkColorChooserAddPalette(x.GoPointer(), OrientationVar, ColorsPerLineVar, NColorsVar, ColorsVar)

}

// Gets the currently-selected color.
func (x *ColorChooserWidget) GetRgba(ColorVar *gdk.RGBA) {

	XGtkColorChooserGetRgba(x.GoPointer(), ColorVar)

}

// Returns whether the color chooser shows the alpha channel.
func (x *ColorChooserWidget) GetUseAlpha() bool {

	return XGtkColorChooserGetUseAlpha(x.GoPointer())

}

// Sets the color.
func (x *ColorChooserWidget) SetRgba(ColorVar *gdk.RGBA) {

	XGtkColorChooserSetRgba(x.GoPointer(), ColorVar)

}

// Sets whether or not the color chooser should use the alpha channel.
func (x *ColorChooserWidget) SetUseAlpha(UseAlphaVar bool) {

	XGtkColorChooserSetUseAlpha(x.GoPointer(), UseAlphaVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewColorChooserWidget, lib, "gtk_color_chooser_widget_new")

}
