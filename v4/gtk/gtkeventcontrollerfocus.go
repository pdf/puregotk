// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type EventControllerFocusClass struct {
}

// `GtkEventControllerFocus` is an event controller to keep track of
// keyboard focus.
//
// The event controller offers [signal@Gtk.EventControllerFocus::enter]
// and [signal@Gtk.EventControllerFocus::leave] signals, as well as
// [property@Gtk.EventControllerFocus:is-focus] and
// [property@Gtk.EventControllerFocus:contains-focus] properties
// which are updated to reflect focus changes inside the widget hierarchy
// that is rooted at the controllers widget.
type EventControllerFocus struct {
	EventController
}

func EventControllerFocusNewFromInternalPtr(ptr uintptr) *EventControllerFocus {
	cls := &EventControllerFocus{}
	cls.Ptr = ptr
	return cls
}

var xNewEventControllerFocus func() uintptr

// Creates a new event controller that will handle focus events.
func NewEventControllerFocus() *EventController {
	var cls *EventController

	cret := xNewEventControllerFocus()

	if cret == 0 {
		return nil
	}
	cls = &EventController{}
	cls.Ptr = cret
	return cls
}

var xEventControllerFocusContainsFocus func(uintptr) bool

// Returns %TRUE if focus is within @self or one of its children.
func (x *EventControllerFocus) ContainsFocus() bool {

	cret := xEventControllerFocusContainsFocus(x.GoPointer())
	return cret
}

var xEventControllerFocusIsFocus func(uintptr) bool

// Returns %TRUE if focus is within @self, but not one of its children.
func (x *EventControllerFocus) IsFocus() bool {

	cret := xEventControllerFocusIsFocus(x.GoPointer())
	return cret
}

func (c *EventControllerFocus) GoPointer() uintptr {
	return c.Ptr
}

func (c *EventControllerFocus) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever the focus enters into the widget or one
// of its descendents.
//
// Note that this means you may not get an ::enter signal
// even though the widget becomes the focus location, in
// certain cases (such as when the focus moves from a descendent
// of the widget to the widget itself). If you are interested
// in these cases, you can monitor the
// [property@Gtk.EventControllerFocus:is-focus]
// property for changes.
func (x *EventControllerFocus) ConnectEnter(cb func(EventControllerFocus)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := EventControllerFocus{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "enter", purego.NewCallback(fcb))
}

// Emitted whenever the focus leaves the widget hierarchy
// that is rooted at the widget that the controller is attached to.
//
// Note that this means you may not get a ::leave signal
// even though the focus moves away from the widget, in
// certain cases (such as when the focus moves from the widget
// to a descendent). If you are interested in these cases, you
// can monitor the [property@Gtk.EventControllerFocus:is-focus]
// property for changes.
func (x *EventControllerFocus) ConnectLeave(cb func(EventControllerFocus)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := EventControllerFocus{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "leave", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEventControllerFocus, lib, "gtk_event_controller_focus_new")

	core.PuregoSafeRegister(&xEventControllerFocusContainsFocus, lib, "gtk_event_controller_focus_contains_focus")
	core.PuregoSafeRegister(&xEventControllerFocusIsFocus, lib, "gtk_event_controller_focus_is_focus")

}
