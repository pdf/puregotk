// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A widget with two panes, arranged either horizontally or vertically.
//
// ![An example GtkPaned](panes.png)
//
// The division between the two panes is adjustable by the user
// by dragging a handle.
//
// Child widgets are added to the panes of the widget with
// [method@Gtk.Paned.set_start_child] and [method@Gtk.Paned.set_end_child].
// The division between the two children is set by default from the size
// requests of the children, but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a
// small handle that the user can drag to adjust the division. It does not
// draw any relief around the children or around the separator. (The space
// in which the separator is called the gutter.) Often, it is useful to put
// each child inside a [class@Gtk.Frame] so that the gutter appears as a
// ridge. No separator is drawn if one of the children is missing.
//
// Each child has two options that can be set, "resize" and "shrink". If
// "resize" is true then, when the `GtkPaned` is resized, that child will
// expand or shrink along with the paned widget. If "shrink" is true, then
// that child can be made smaller than its requisition by the user.
// Setting "shrink" to false allows the application to set a minimum size.
// If "resize" is false for both children, then this is treated as if
// "resize" is true for both children.
//
// The application can set the position of the slider as if it were set
// by the user, by calling [method@Gtk.Paned.set_position].
//
// # CSS nodes
//
// ```
// paned
// ├── &lt;child&gt;
// ├── separator[.wide]
// ╰── &lt;child&gt;
// ```
//
// `GtkPaned` has a main CSS node with name paned, and a subnode for
// the separator with name separator. The subnode gets a .wide style
// class when the paned is supposed to be wide.
//
// In horizontal orientation, the nodes are arranged based on the text
// direction, so in left-to-right mode, :first-child will select the
// leftmost child, while it will select the rightmost child in
// RTL layouts.
//
// ## Creating a paned widget with minimum sizes.
//
// ```c
// GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
// GtkWidget *frame1 = gtk_frame_new (NULL);
// GtkWidget *frame2 = gtk_frame_new (NULL);
//
// gtk_widget_set_size_request (hpaned, 200, -1);
//
// gtk_paned_set_start_child (GTK_PANED (hpaned), frame1);
// gtk_paned_set_start_child_resize (GTK_PANED (hpaned), TRUE);
// gtk_paned_set_start_child_shrink (GTK_PANED (hpaned), FALSE);
// gtk_widget_set_size_request (frame1, 50, -1);
//
// gtk_paned_set_end_child (GTK_PANED (hpaned), frame2);
// gtk_paned_set_end_child_resize (GTK_PANED (hpaned), FALSE);
// gtk_paned_set_end_child_shrink (GTK_PANED (hpaned), FALSE);
// gtk_widget_set_size_request (frame2, 50, -1);
// ```
type Paned struct {
	Widget
}

func PanedNewFromInternalPtr(ptr uintptr) *Paned {
	cls := &Paned{}
	cls.Ptr = ptr
	return cls
}

var xNewPaned func(Orientation) uintptr

// Creates a new `GtkPaned` widget.
func NewPaned(OrientationVar Orientation) *Widget {
	NewPanedPtr := xNewPaned(OrientationVar)
	if NewPanedPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewPanedPtr)

	NewPanedCls := &Widget{}
	NewPanedCls.Ptr = NewPanedPtr
	return NewPanedCls
}

var xPanedGetEndChild func(uintptr) uintptr

// Retrieves the end child of the given `GtkPaned`.
func (x *Paned) GetEndChild() *Widget {

	GetEndChildPtr := xPanedGetEndChild(x.GoPointer())
	if GetEndChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetEndChildPtr)

	GetEndChildCls := &Widget{}
	GetEndChildCls.Ptr = GetEndChildPtr
	return GetEndChildCls

}

var xPanedGetPosition func(uintptr) int

// Obtains the position of the divider between the two panes.
func (x *Paned) GetPosition() int {

	return xPanedGetPosition(x.GoPointer())

}

var xPanedGetResizeEndChild func(uintptr) bool

// Returns whether the [property@Gtk.Paned:end-child] can be resized.
func (x *Paned) GetResizeEndChild() bool {

	return xPanedGetResizeEndChild(x.GoPointer())

}

var xPanedGetResizeStartChild func(uintptr) bool

// Returns whether the [property@Gtk.Paned:start-child] can be resized.
func (x *Paned) GetResizeStartChild() bool {

	return xPanedGetResizeStartChild(x.GoPointer())

}

var xPanedGetShrinkEndChild func(uintptr) bool

// Returns whether the [property@Gtk.Paned:end-child] can shrink.
func (x *Paned) GetShrinkEndChild() bool {

	return xPanedGetShrinkEndChild(x.GoPointer())

}

var xPanedGetShrinkStartChild func(uintptr) bool

// Returns whether the [property@Gtk.Paned:start-child] can shrink.
func (x *Paned) GetShrinkStartChild() bool {

	return xPanedGetShrinkStartChild(x.GoPointer())

}

var xPanedGetStartChild func(uintptr) uintptr

// Retrieves the start child of the given `GtkPaned`.
func (x *Paned) GetStartChild() *Widget {

	GetStartChildPtr := xPanedGetStartChild(x.GoPointer())
	if GetStartChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetStartChildPtr)

	GetStartChildCls := &Widget{}
	GetStartChildCls.Ptr = GetStartChildPtr
	return GetStartChildCls

}

var xPanedGetWideHandle func(uintptr) bool

// Gets whether the separator should be wide.
func (x *Paned) GetWideHandle() bool {

	return xPanedGetWideHandle(x.GoPointer())

}

var xPanedSetEndChild func(uintptr, uintptr)

// Sets the end child of @paned to @child.
//
// If @child is `NULL`, the existing child will be removed.
func (x *Paned) SetEndChild(ChildVar *Widget) {

	xPanedSetEndChild(x.GoPointer(), ChildVar.GoPointer())

}

var xPanedSetPosition func(uintptr, int)

// Sets the position of the divider between the two panes.
func (x *Paned) SetPosition(PositionVar int) {

	xPanedSetPosition(x.GoPointer(), PositionVar)

}

var xPanedSetResizeEndChild func(uintptr, bool)

// Sets whether the [property@Gtk.Paned:end-child] can be resized.
func (x *Paned) SetResizeEndChild(ResizeVar bool) {

	xPanedSetResizeEndChild(x.GoPointer(), ResizeVar)

}

var xPanedSetResizeStartChild func(uintptr, bool)

// Sets whether the [property@Gtk.Paned:start-child] can be resized.
func (x *Paned) SetResizeStartChild(ResizeVar bool) {

	xPanedSetResizeStartChild(x.GoPointer(), ResizeVar)

}

var xPanedSetShrinkEndChild func(uintptr, bool)

// Sets whether the [property@Gtk.Paned:end-child] can shrink.
func (x *Paned) SetShrinkEndChild(ResizeVar bool) {

	xPanedSetShrinkEndChild(x.GoPointer(), ResizeVar)

}

var xPanedSetShrinkStartChild func(uintptr, bool)

// Sets whether the [property@Gtk.Paned:start-child] can shrink.
func (x *Paned) SetShrinkStartChild(ResizeVar bool) {

	xPanedSetShrinkStartChild(x.GoPointer(), ResizeVar)

}

var xPanedSetStartChild func(uintptr, uintptr)

// Sets the start child of @paned to @child.
//
// If @child is `NULL`, the existing child will be removed.
func (x *Paned) SetStartChild(ChildVar *Widget) {

	xPanedSetStartChild(x.GoPointer(), ChildVar.GoPointer())

}

var xPanedSetWideHandle func(uintptr, bool)

// Sets whether the separator should be wide.
func (x *Paned) SetWideHandle(WideVar bool) {

	xPanedSetWideHandle(x.GoPointer(), WideVar)

}

func (c *Paned) GoPointer() uintptr {
	return c.Ptr
}

func (c *Paned) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to accept the current position of the handle when
// moving it using key bindings.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;Return&lt;/kbd&gt; or
// &lt;kbd&gt;Space&lt;/kbd&gt;.
func (x *Paned) ConnectAcceptPosition(cb func(Paned) bool) {
	fcb := func(clsPtr uintptr) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::accept-position", purego.NewCallback(fcb))
}

// Emitted to cancel moving the position of the handle using key
// bindings.
//
// The position of the handle will be reset to the value prior to
// moving it.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;Escape&lt;/kbd&gt;.
func (x *Paned) ConnectCancelPosition(cb func(Paned) bool) {
	fcb := func(clsPtr uintptr) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::cancel-position", purego.NewCallback(fcb))
}

// Emitted to cycle the focus between the children of the paned.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding is &lt;kbd&gt;F6&lt;/kbd&gt;.
func (x *Paned) ConnectCycleChildFocus(cb func(Paned, bool) bool) {
	fcb := func(clsPtr uintptr, ReversedVarp bool) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa, ReversedVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::cycle-child-focus", purego.NewCallback(fcb))
}

// Emitted to cycle whether the paned should grab focus to allow
// the user to change position of the handle by using key bindings.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is &lt;kbd&gt;F8&lt;/kbd&gt;.
func (x *Paned) ConnectCycleHandleFocus(cb func(Paned, bool) bool) {
	fcb := func(clsPtr uintptr, ReversedVarp bool) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa, ReversedVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::cycle-handle-focus", purego.NewCallback(fcb))
}

// Emitted to move the handle with key bindings.
//
// This is a [keybinding signal](class.SignalAction.html).
func (x *Paned) ConnectMoveHandle(cb func(Paned, ScrollType) bool) {
	fcb := func(clsPtr uintptr, ScrollTypeVarp ScrollType) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa, ScrollTypeVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::move-handle", purego.NewCallback(fcb))
}

// Emitted to accept the current position of the handle and then
// move focus to the next widget in the focus chain.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding is &lt;kbd&gt;Tab&lt;/kbd&gt;.
func (x *Paned) ConnectToggleHandleFocus(cb func(Paned) bool) {
	fcb := func(clsPtr uintptr) bool {
		fa := Paned{}
		fa.Ptr = clsPtr

		return cb(fa)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::toggle-handle-focus", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Paned) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Paned) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Paned) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Paned) ResetState(StateVar AccessibleState) {

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
func (x *Paned) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Paned) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Paned) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Paned) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Paned) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Paned) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Paned) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

// Retrieves the orientation of the @orientable.
func (x *Paned) GetOrientation() Orientation {

	return XGtkOrientableGetOrientation(x.GoPointer())

}

// Sets the orientation of the @orientable.
func (x *Paned) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewPaned, lib, "gtk_paned_new")

	core.PuregoSafeRegister(&xPanedGetEndChild, lib, "gtk_paned_get_end_child")
	core.PuregoSafeRegister(&xPanedGetPosition, lib, "gtk_paned_get_position")
	core.PuregoSafeRegister(&xPanedGetResizeEndChild, lib, "gtk_paned_get_resize_end_child")
	core.PuregoSafeRegister(&xPanedGetResizeStartChild, lib, "gtk_paned_get_resize_start_child")
	core.PuregoSafeRegister(&xPanedGetShrinkEndChild, lib, "gtk_paned_get_shrink_end_child")
	core.PuregoSafeRegister(&xPanedGetShrinkStartChild, lib, "gtk_paned_get_shrink_start_child")
	core.PuregoSafeRegister(&xPanedGetStartChild, lib, "gtk_paned_get_start_child")
	core.PuregoSafeRegister(&xPanedGetWideHandle, lib, "gtk_paned_get_wide_handle")
	core.PuregoSafeRegister(&xPanedSetEndChild, lib, "gtk_paned_set_end_child")
	core.PuregoSafeRegister(&xPanedSetPosition, lib, "gtk_paned_set_position")
	core.PuregoSafeRegister(&xPanedSetResizeEndChild, lib, "gtk_paned_set_resize_end_child")
	core.PuregoSafeRegister(&xPanedSetResizeStartChild, lib, "gtk_paned_set_resize_start_child")
	core.PuregoSafeRegister(&xPanedSetShrinkEndChild, lib, "gtk_paned_set_shrink_end_child")
	core.PuregoSafeRegister(&xPanedSetShrinkStartChild, lib, "gtk_paned_set_shrink_start_child")
	core.PuregoSafeRegister(&xPanedSetStartChild, lib, "gtk_paned_set_start_child")
	core.PuregoSafeRegister(&xPanedSetWideHandle, lib, "gtk_paned_set_wide_handle")

}
