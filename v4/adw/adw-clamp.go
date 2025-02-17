// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type ClampClass struct {
	ParentClass uintptr
}

// A widget constraining its child to a given size.
//
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-wide-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-wide.png" alt="clamp-wide"&gt;
//
// &lt;/picture&gt;
// &lt;picture&gt;
//
//	&lt;source srcset="clamp-narrow-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="clamp-narrow.png" alt="clamp-narrow"&gt;
//
// &lt;/picture&gt;
//
// The `AdwClamp` widget constrains the size of the widget it contains to a
// given maximum size. It will constrain the width if it is horizontal, or the
// height if it is vertical. The expansion of the child from its minimum to its
// maximum size is eased out for a smooth transition.
//
// If the child requires more than the requested maximum size, it will be
// allocated the minimum size it can fit in instead.
//
// ## CSS nodes
//
// `AdwClamp` has a single CSS node with name `clamp`.
//
// Its children will receive the style classes `.large` when the child reached
// its maximum size, `.small` when the clamp allocates its full size to the
// child, `.medium` in-between, or none if it hasn't computed its size yet.
type Clamp struct {
	gtk.Widget
}

func ClampNewFromInternalPtr(ptr uintptr) *Clamp {
	cls := &Clamp{}
	cls.Ptr = ptr
	return cls
}

var xNewClamp func() uintptr

// Creates a new `AdwClamp`.
func NewClamp() *gtk.Widget {
	var cls *gtk.Widget

	cret := xNewClamp()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xClampGetChild func(uintptr) uintptr

// Gets the child widget of @self.
func (x *Clamp) GetChild() *gtk.Widget {
	var cls *gtk.Widget

	cret := xClampGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gtk.Widget{}
	cls.Ptr = cret
	return cls
}

var xClampGetMaximumSize func(uintptr) int

// Gets the maximum size allocated to the child.
func (x *Clamp) GetMaximumSize() int {

	cret := xClampGetMaximumSize(x.GoPointer())
	return cret
}

var xClampGetTighteningThreshold func(uintptr) int

// Gets the size above which the child is clamped.
func (x *Clamp) GetTighteningThreshold() int {

	cret := xClampGetTighteningThreshold(x.GoPointer())
	return cret
}

var xClampSetChild func(uintptr, uintptr)

// Sets the child widget of @self.
func (x *Clamp) SetChild(ChildVar *gtk.Widget) {

	xClampSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xClampSetMaximumSize func(uintptr, int)

// Sets the maximum size allocated to the child.
//
// It is the width if the clamp is horizontal, or the height if it is vertical.
func (x *Clamp) SetMaximumSize(MaximumSizeVar int) {

	xClampSetMaximumSize(x.GoPointer(), MaximumSizeVar)

}

var xClampSetTighteningThreshold func(uintptr, int)

// Sets the size above which the child is clamped.
//
// Starting from this size, the clamp will tighten its grip on the child, slowly
// allocating less and less of the available size up to the maximum allocated
// size. Below that threshold and below the maximum size, the child will be
// allocated all the available size.
//
// If the threshold is greater than the maximum size to allocate to the child,
// the child will be allocated all the size up to the maximum. If the threshold
// is lower than the minimum size to allocate to the child, that size will be
// used as the tightening threshold.
//
// Effectively, tightening the grip on the child before it reaches its maximum
// size makes transitions to and from the maximum size smoother when resizing.
func (x *Clamp) SetTighteningThreshold(TighteningThresholdVar int) {

	xClampSetTighteningThreshold(x.GoPointer(), TighteningThresholdVar)

}

func (c *Clamp) GoPointer() uintptr {
	return c.Ptr
}

func (c *Clamp) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Clamp) GetAccessibleRole() gtk.AccessibleRole {

	cret := gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Clamp) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Clamp) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Clamp) ResetState(StateVar gtk.AccessibleState) {

	gtk.XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *Clamp) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Clamp) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *Clamp) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Clamp) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *Clamp) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Clamp) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Clamp) GetBuildableId() string {

	cret := gtk.XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *Clamp) GetOrientation() gtk.Orientation {

	cret := gtk.XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *Clamp) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewClamp, lib, "adw_clamp_new")

	core.PuregoSafeRegister(&xClampGetChild, lib, "adw_clamp_get_child")
	core.PuregoSafeRegister(&xClampGetMaximumSize, lib, "adw_clamp_get_maximum_size")
	core.PuregoSafeRegister(&xClampGetTighteningThreshold, lib, "adw_clamp_get_tightening_threshold")
	core.PuregoSafeRegister(&xClampSetChild, lib, "adw_clamp_set_child")
	core.PuregoSafeRegister(&xClampSetMaximumSize, lib, "adw_clamp_set_maximum_size")
	core.PuregoSafeRegister(&xClampSetTighteningThreshold, lib, "adw_clamp_set_tightening_threshold")

}
