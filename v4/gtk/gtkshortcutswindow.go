// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
)

// A `GtkShortcutsWindow` shows information about the keyboard shortcuts
// and gestures of an application.
//
// The shortcuts can be grouped, and you can have multiple sections in this
// window, corresponding to the major modes of your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a `GtkShortcutsWindow` is with
// [class@Gtk.Builder], by populating a `GtkShortcutsWindow` with one or
// more `GtkShortcutsSection` objects, which contain `GtkShortcutsGroups`
// that in turn contain objects of class `GtkShortcutsShortcut`.
//
// # A simple example:
//
// ![](gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups
// are arranged in columns, and spread across several pages if there are too
// many to find on a single page.
//
// The .ui file for this example can be found [here](https://gitlab.gnome.org/GNOME/gtk/tree/main/demos/gtk-demo/shortcuts-gedit.ui).
//
// # An example with multiple views:
//
// ![](clocks-shortcuts.png)
//
// This example shows a `GtkShortcutsWindow` that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found [here](https://gitlab.gnome.org/GNOME/gtk/tree/main/demos/gtk-demo/shortcuts-clocks.ui).
//
// # An example with multiple sections:
//
// ![](builder-shortcuts.png)
//
// This example shows a `GtkShortcutsWindow` with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found [here](https://gitlab.gnome.org/GNOME/gtk/tree/main/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	Window
}

func ShortcutsWindowNewFromInternalPtr(ptr uintptr) *ShortcutsWindow {
	cls := &ShortcutsWindow{}
	cls.Ptr = ptr
	return cls
}

func (c *ShortcutsWindow) GoPointer() uintptr {
	return c.Ptr
}

func (c *ShortcutsWindow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user uses a keybinding to close the window.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is the Escape key.
func (x *ShortcutsWindow) ConnectClose(cb func(ShortcutsWindow)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := ShortcutsWindow{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "close", purego.NewCallback(fcb))
}

// Emitted when the user uses a keybinding to start a search.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is Control-F.
func (x *ShortcutsWindow) ConnectSearch(cb func(ShortcutsWindow)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := ShortcutsWindow{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "search", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ShortcutsWindow) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ShortcutsWindow) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ShortcutsWindow) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ShortcutsWindow) ResetState(StateVar AccessibleState) {

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
func (x *ShortcutsWindow) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ShortcutsWindow) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ShortcutsWindow) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ShortcutsWindow) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ShortcutsWindow) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ShortcutsWindow) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ShortcutsWindow) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Returns the renderer that is used for this `GtkNative`.
func (x *ShortcutsWindow) GetRenderer() *gsk.Renderer {
	var cls *gsk.Renderer

	cret := XGtkNativeGetRenderer(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gsk.Renderer{}
	cls.Ptr = cret
	return cls
}

// Returns the surface of this `GtkNative`.
func (x *ShortcutsWindow) GetSurface() *gdk.Surface {
	var cls *gdk.Surface

	cret := XGtkNativeGetSurface(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Surface{}
	cls.Ptr = cret
	return cls
}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *ShortcutsWindow) GetSurfaceTransform(XVar float64, YVar float64) {

	XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *ShortcutsWindow) Realize() {

	XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *ShortcutsWindow) Unrealize() {

	XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *ShortcutsWindow) GetDisplay() *gdk.Display {
	var cls *gdk.Display

	cret := XGtkRootGetDisplay(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Display{}
	cls.Ptr = cret
	return cls
}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *ShortcutsWindow) GetFocus() *Widget {
	var cls *Widget

	cret := XGtkRootGetFocus(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *ShortcutsWindow) SetFocus(FocusVar *Widget) {

	XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}
