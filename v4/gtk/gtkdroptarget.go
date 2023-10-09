// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type DropTargetClass struct {
}

// `GtkDropTarget` is an event controller to receive Drag-and-Drop operations.
//
// The most basic way to use a `GtkDropTarget` to receive drops on a
// widget is to create it via [ctor@Gtk.DropTarget.new], passing in the
// `GType` of the data you want to receive and connect to the
// [signal@Gtk.DropTarget::drop] signal to receive the data:
//
// ```c
// static gboolean
// on_drop (GtkDropTarget *target,
//
//	const GValue  *value,
//	double         x,
//	double         y,
//	gpointer       data)
//
//	{
//	  MyWidget *self = data;
//
//	  // Call the appropriate setter depending on the type of data
//	  // that we received
//	  if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//	    my_widget_set_file (self, g_value_get_object (value));
//	  else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//	    my_widget_set_pixbuf (self, g_value_get_object (value));
//	  else
//	    return FALSE;
//
//	  return TRUE;
//	}
//
// static void
// my_widget_init (MyWidget *self)
//
//	{
//	  GtkDropTarget *target =
//	    gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//	  // This widget accepts two types of drop types: GFile objects
//	  // and GdkPixbuf objects
//	  gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//	    G_TYPE_FILE,
//	    GDK_TYPE_PIXBUF,
//	  }, 2);
//
//	  g_signal_connect (target, "drop", G_CALLBACK (on_drop), self);
//	  gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//	}
//
// ```
//
// `GtkDropTarget` supports more options, such as:
//
//   - rejecting potential drops via the [signal@Gtk.DropTarget::accept] signal
//     and the [method@Gtk.DropTarget.reject] function to let other drop
//     targets handle the drop
//   - tracking an ongoing drag operation before the drop via the
//     [signal@Gtk.DropTarget::enter], [signal@Gtk.DropTarget::motion] and
//     [signal@Gtk.DropTarget::leave] signals
//   - configuring how to receive data by setting the
//     [property@Gtk.DropTarget:preload] property and listening for its
//     availability via the [property@Gtk.DropTarget:value] property
//
// However, `GtkDropTarget` is ultimately modeled in a synchronous way
// and only supports data transferred via `GType`. If you want full control
// over an ongoing drop, the [class@Gtk.DropTargetAsync] object gives you
// this ability.
//
// While a pointer is dragged over the drop target's widget and the drop
// has not been rejected, that widget will receive the
// %GTK_STATE_FLAG_DROP_ACTIVE state, which can be used to style the widget.
//
// If you are not interested in receiving the drop, but just want to update
// UI state during a Drag-and-Drop operation (e.g. switching tabs), you can
// use [class@Gtk.DropControllerMotion].
type DropTarget struct {
	EventController
}

func DropTargetNewFromInternalPtr(ptr uintptr) *DropTarget {
	cls := &DropTarget{}
	cls.Ptr = ptr
	return cls
}

var xNewDropTarget func([]interface{}, gdk.DragAction) uintptr

// Creates a new `GtkDropTarget` object.
//
// If the drop target should support more than 1 type, pass
// %G_TYPE_INVALID for @type and then call
// [method@Gtk.DropTarget.set_gtypes].
func NewDropTarget(TypeVar []interface{}, ActionsVar gdk.DragAction) *DropTarget {
	var cls *DropTarget

	cret := xNewDropTarget(TypeVar, ActionsVar)

	if cret == 0 {
		return nil
	}
	cls = &DropTarget{}
	cls.Ptr = cret
	return cls
}

var xDropTargetGetActions func(uintptr) gdk.DragAction

// Gets the actions that this drop target supports.
func (x *DropTarget) GetActions() gdk.DragAction {

	cret := xDropTargetGetActions(x.GoPointer())
	return cret
}

var xDropTargetGetCurrentDrop func(uintptr) uintptr

// Gets the currently handled drop operation.
//
// If no drop operation is going on, %NULL is returned.
func (x *DropTarget) GetCurrentDrop() *gdk.Drop {
	var cls *gdk.Drop

	cret := xDropTargetGetCurrentDrop(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Drop{}
	cls.Ptr = cret
	return cls
}

var xDropTargetGetDrop func(uintptr) uintptr

// Gets the currently handled drop operation.
//
// If no drop operation is going on, %NULL is returned.
func (x *DropTarget) GetDrop() *gdk.Drop {
	var cls *gdk.Drop

	cret := xDropTargetGetDrop(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.Drop{}
	cls.Ptr = cret
	return cls
}

var xDropTargetGetFormats func(uintptr) *gdk.ContentFormats

// Gets the data formats that this drop target accepts.
//
// If the result is %NULL, all formats are expected to be supported.
func (x *DropTarget) GetFormats() *gdk.ContentFormats {

	cret := xDropTargetGetFormats(x.GoPointer())
	return cret
}

var xDropTargetGetGtypes func(uintptr, uint) uintptr

// Gets the list of supported `GType`s that can be dropped on the target.
//
// If no types have been set, `NULL` will be returned.
func (x *DropTarget) GetGtypes(NTypesVar uint) uintptr {

	cret := xDropTargetGetGtypes(x.GoPointer(), NTypesVar)
	return cret
}

var xDropTargetGetPreload func(uintptr) bool

// Gets whether data should be preloaded on hover.
func (x *DropTarget) GetPreload() bool {

	cret := xDropTargetGetPreload(x.GoPointer())
	return cret
}

var xDropTargetGetValue func(uintptr) *gobject.Value

// Gets the current drop data, as a `GValue`.
func (x *DropTarget) GetValue() *gobject.Value {

	cret := xDropTargetGetValue(x.GoPointer())
	return cret
}

var xDropTargetReject func(uintptr)

// Rejects the ongoing drop operation.
//
// If no drop operation is ongoing, i.e when [property@Gtk.DropTarget:current-drop]
// is %NULL, this function does nothing.
//
// This function should be used when delaying the decision
// on whether to accept a drag or not until after reading
// the data.
func (x *DropTarget) Reject() {

	xDropTargetReject(x.GoPointer())

}

var xDropTargetSetActions func(uintptr, gdk.DragAction)

// Sets the actions that this drop target supports.
func (x *DropTarget) SetActions(ActionsVar gdk.DragAction) {

	xDropTargetSetActions(x.GoPointer(), ActionsVar)

}

var xDropTargetSetGtypes func(uintptr, uintptr, uint)

// Sets the supported `GTypes` for this drop target.
func (x *DropTarget) SetGtypes(TypesVar uintptr, NTypesVar uint) {

	xDropTargetSetGtypes(x.GoPointer(), TypesVar, NTypesVar)

}

var xDropTargetSetPreload func(uintptr, bool)

// Sets whether data should be preloaded on hover.
func (x *DropTarget) SetPreload(PreloadVar bool) {

	xDropTargetSetPreload(x.GoPointer(), PreloadVar)

}

func (c *DropTarget) GoPointer() uintptr {
	return c.Ptr
}

func (c *DropTarget) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted on the drop site when a drop operation is about to begin.
//
// If the drop is not accepted, %FALSE will be returned and the drop target
// will ignore the drop. If %TRUE is returned, the drop is accepted for now
// but may be rejected later via a call to [method@Gtk.DropTarget.reject]
// or ultimately by returning %FALSE from a [signal@Gtk.DropTarget::drop]
// handler.
//
// The default handler for this signal decides whether to accept the drop
// based on the formats provided by the @drop.
//
// If the decision whether the drop will be accepted or rejected depends
// on the data, this function should return %TRUE, the
// [property@Gtk.DropTarget:preload] property should be set and the value
// should be inspected via the ::notify:value signal, calling
// [method@Gtk.DropTarget.reject] if required.
func (x *DropTarget) ConnectAccept(cb func(DropTarget, uintptr) bool) uint32 {
	fcb := func(clsPtr uintptr, DropVarp uintptr) bool {
		fa := DropTarget{}
		fa.Ptr = clsPtr

		return cb(fa, DropVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "accept", purego.NewCallback(fcb))
}

// Emitted on the drop site when the user drops the data onto the widget.
//
// The signal handler must determine whether the pointer position is in
// a drop zone or not. If it is not in a drop zone, it returns %FALSE
// and no further processing is necessary.
//
// Otherwise, the handler returns %TRUE. In this case, this handler will
// accept the drop. The handler is responsible for using the given @value
// and performing the drop operation.
func (x *DropTarget) ConnectDrop(cb func(DropTarget, uintptr, float64, float64) bool) uint32 {
	fcb := func(clsPtr uintptr, ValueVarp uintptr, XVarp float64, YVarp float64) bool {
		fa := DropTarget{}
		fa.Ptr = clsPtr

		return cb(fa, ValueVarp, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "drop", purego.NewCallback(fcb))
}

// Emitted on the drop site when the pointer enters the widget.
//
// It can be used to set up custom highlighting.
func (x *DropTarget) ConnectEnter(cb func(DropTarget, float64, float64) gdk.DragAction) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) gdk.DragAction {
		fa := DropTarget{}
		fa.Ptr = clsPtr

		return cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "enter", purego.NewCallback(fcb))
}

// Emitted on the drop site when the pointer leaves the widget.
//
// Its main purpose it to undo things done in
// [signal@Gtk.DropTarget::enter].
func (x *DropTarget) ConnectLeave(cb func(DropTarget)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := DropTarget{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "leave", purego.NewCallback(fcb))
}

// Emitted while the pointer is moving over the drop target.
func (x *DropTarget) ConnectMotion(cb func(DropTarget, float64, float64) gdk.DragAction) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) gdk.DragAction {
		fa := DropTarget{}
		fa.Ptr = clsPtr

		return cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "motion", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDropTarget, lib, "gtk_drop_target_new")

	core.PuregoSafeRegister(&xDropTargetGetActions, lib, "gtk_drop_target_get_actions")
	core.PuregoSafeRegister(&xDropTargetGetCurrentDrop, lib, "gtk_drop_target_get_current_drop")
	core.PuregoSafeRegister(&xDropTargetGetDrop, lib, "gtk_drop_target_get_drop")
	core.PuregoSafeRegister(&xDropTargetGetFormats, lib, "gtk_drop_target_get_formats")
	core.PuregoSafeRegister(&xDropTargetGetGtypes, lib, "gtk_drop_target_get_gtypes")
	core.PuregoSafeRegister(&xDropTargetGetPreload, lib, "gtk_drop_target_get_preload")
	core.PuregoSafeRegister(&xDropTargetGetValue, lib, "gtk_drop_target_get_value")
	core.PuregoSafeRegister(&xDropTargetReject, lib, "gtk_drop_target_reject")
	core.PuregoSafeRegister(&xDropTargetSetActions, lib, "gtk_drop_target_set_actions")
	core.PuregoSafeRegister(&xDropTargetSetGtypes, lib, "gtk_drop_target_set_gtypes")
	core.PuregoSafeRegister(&xDropTargetSetPreload, lib, "gtk_drop_target_set_preload")

}
