// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gsk"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type WindowClass struct {
	ParentClass uintptr

	Padding uintptr
}

// A freeform window.
//
// &lt;picture&gt;
//
//	&lt;source srcset="window-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="window.png" alt="window"&gt;
//
// &lt;/picture&gt;
//
// The `AdwWindow` widget is a subclass of [class@Gtk.Window] which has no
// titlebar area. It means [class@Gtk.HeaderBar] can be used as follows:
//
// ```xml
// &lt;object class="AdwWindow"&gt;
//
//	&lt;property name="content"&gt;
//	  &lt;object class="GtkBox"&gt;
//	    &lt;property name="orientation"&gt;vertical&lt;/property&gt;
//	    &lt;child&gt;
//	      &lt;object class="GtkHeaderBar"/&gt;
//	    &lt;/child&gt;
//	    &lt;child&gt;
//	      &lt;!-- ... --&gt;
//	    &lt;/child&gt;
//	  &lt;/object&gt;
//	&lt;/property&gt;
//
// &lt;/object&gt;
// ```
//
// Using [property@Gtk.Window:titlebar] or [property@Gtk.Window:child]
// is not supported and will result in a crash. Use [property@Window:content]
// instead.
type Window struct {
	gtk.Window
}

func WindowNewFromInternalPtr(ptr uintptr) *Window {
	cls := &Window{}
	cls.Ptr = ptr
	return cls
}

var xNewWindow func() uintptr

// Creates a new `AdwWindow`.
func NewWindow() *gtk.Widget {
	NewWindowPtr := xNewWindow()
	if NewWindowPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewWindowPtr)

	NewWindowCls := &gtk.Widget{}
	NewWindowCls.Ptr = NewWindowPtr
	return NewWindowCls
}

var xWindowGetContent func(uintptr) uintptr

// Gets the content widget of @self.
//
// This method should always be used instead of [method@Gtk.Window.get_child].
func (x *Window) GetContent() *gtk.Widget {

	GetContentPtr := xWindowGetContent(x.GoPointer())
	if GetContentPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetContentPtr)

	GetContentCls := &gtk.Widget{}
	GetContentCls.Ptr = GetContentPtr
	return GetContentCls

}

var xWindowSetContent func(uintptr, uintptr)

// Sets the content widget of @self.
//
// This method should always be used instead of [method@Gtk.Window.set_child].
func (x *Window) SetContent(ContentVar *gtk.Widget) {

	xWindowSetContent(x.GoPointer(), ContentVar.GoPointer())

}

func (c *Window) GoPointer() uintptr {
	return c.Ptr
}

func (c *Window) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Window) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Window) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Window) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Window) ResetState(StateVar gtk.AccessibleState) {

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
func (x *Window) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Window) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Window) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Window) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Window) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Window) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Window) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

// Returns the renderer that is used for this `GtkNative`.
func (x *Window) GetRenderer() *gsk.Renderer {

	GetRendererPtr := gtk.XGtkNativeGetRenderer(x.GoPointer())
	if GetRendererPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetRendererPtr)

	GetRendererCls := &gsk.Renderer{}
	GetRendererCls.Ptr = GetRendererPtr
	return GetRendererCls

}

// Returns the surface of this `GtkNative`.
func (x *Window) GetSurface() *gdk.Surface {

	GetSurfacePtr := gtk.XGtkNativeGetSurface(x.GoPointer())
	if GetSurfacePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSurfacePtr)

	GetSurfaceCls := &gdk.Surface{}
	GetSurfaceCls.Ptr = GetSurfacePtr
	return GetSurfaceCls

}

// Retrieves the surface transform of @self.
//
// This is the translation from @self's surface coordinates into
// @self's widget coordinates.
func (x *Window) GetSurfaceTransform(XVar float64, YVar float64) {

	gtk.XGtkNativeGetSurfaceTransform(x.GoPointer(), XVar, YVar)

}

// Realizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Window) Realize() {

	gtk.XGtkNativeRealize(x.GoPointer())

}

// Unrealizes a `GtkNative`.
//
// This should only be used by subclasses.
func (x *Window) Unrealize() {

	gtk.XGtkNativeUnrealize(x.GoPointer())

}

// Returns the display that this `GtkRoot` is on.
func (x *Window) GetDisplay() *gdk.Display {

	GetDisplayPtr := gtk.XGtkRootGetDisplay(x.GoPointer())
	if GetDisplayPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetDisplayPtr)

	GetDisplayCls := &gdk.Display{}
	GetDisplayCls.Ptr = GetDisplayPtr
	return GetDisplayCls

}

// Retrieves the current focused widget within the root.
//
// Note that this is the widget that would have the focus
// if the root is active; if the root is not focused then
// `gtk_widget_has_focus (widget)` will be %FALSE for the
// widget.
func (x *Window) GetFocus() *gtk.Widget {

	GetFocusPtr := gtk.XGtkRootGetFocus(x.GoPointer())
	if GetFocusPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFocusPtr)

	GetFocusCls := &gtk.Widget{}
	GetFocusCls.Ptr = GetFocusPtr
	return GetFocusCls

}

// If @focus is not the current focus widget, and is focusable, sets
// it as the focus widget for the root.
//
// If @focus is %NULL, unsets the focus widget for the root.
//
// To set the focus to a particular widget in the root, it is usually
// more convenient to use [method@Gtk.Widget.grab_focus] instead of
// this function.
func (x *Window) SetFocus(FocusVar *gtk.Widget) {

	gtk.XGtkRootSetFocus(x.GoPointer(), FocusVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewWindow, lib, "adw_window_new")

	core.PuregoSafeRegister(&xWindowGetContent, lib, "adw_window_get_content")
	core.PuregoSafeRegister(&xWindowSetContent, lib, "adw_window_set_content")

}
