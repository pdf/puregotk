// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type SignalListItemFactoryClass struct {
}

// `GtkSignalListItemFactory` is a `GtkListItemFactory` that emits signals
// to to manage listitems.
//
// Signals are emitted for every listitem in the same order:
//
//  1. [signal@Gtk.SignalListItemFactory::setup] is emitted to set up permanent
//     things on the listitem. This usually means constructing the widgets used in
//     the row and adding them to the listitem.
//
//  2. [signal@Gtk.SignalListItemFactory::bind] is emitted to bind the item passed
//     via [property@Gtk.ListItem:item] to the widgets that have been created in
//     step 1 or to add item-specific widgets. Signals are connected to listen to
//     changes - both to changes in the item to update the widgets or to changes
//     in the widgets to update the item. After this signal has been called, the
//     listitem may be shown in a list widget.
//
//  3. [signal@Gtk.SignalListItemFactory::unbind] is emitted to undo everything
//     done in step 2. Usually this means disconnecting signal handlers. Once this
//     signal has been called, the listitem will no longer be used in a list
//     widget.
//
//  4. [signal@Gtk.SignalListItemFactory::bind] and
//     [signal@Gtk.SignalListItemFactory::unbind] may be emitted multiple times
//     again to bind the listitem for use with new items. By reusing listitems,
//     potentially costly setup can be avoided. However, it means code needs to
//     make sure to properly clean up the listitem in step 3 so that no information
//     from the previous use leaks into the next use.
//
// 5. [signal@Gtk.SignalListItemFactory::teardown] is emitted to allow undoing
// the effects of [signal@Gtk.SignalListItemFactory::setup]. After this signal
// was emitted on a listitem, the listitem will be destroyed and not be used again.
//
// Note that during the signal emissions, changing properties on the
// `GtkListItem`s passed will not trigger notify signals as the listitem's
// notifications are frozen. See g_object_freeze_notify() for details.
//
// For tracking changes in other properties in the `GtkListItem`, the
// ::notify signal is recommended. The signal can be connected in the
// [signal@Gtk.SignalListItemFactory::setup] signal and removed again during
// [signal@Gtk.SignalListItemFactory::teardown].
type SignalListItemFactory struct {
	ListItemFactory
}

func SignalListItemFactoryNewFromInternalPtr(ptr uintptr) *SignalListItemFactory {
	cls := &SignalListItemFactory{}
	cls.Ptr = ptr
	return cls
}

var xNewSignalListItemFactory func() uintptr

// Creates a new `GtkSignalListItemFactory`.
//
// You need to connect signal handlers before you use it.
func NewSignalListItemFactory() *ListItemFactory {
	var cls *ListItemFactory

	cret := xNewSignalListItemFactory()

	if cret == 0 {
		return nil
	}
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

func (c *SignalListItemFactory) GoPointer() uintptr {
	return c.Ptr
}

func (c *SignalListItemFactory) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when an object has been bound, for example when a
// new [property@Gtk.ListItem:item] has been set on a
// `GtkListItem` and should be bound for use.
//
// After this signal was emitted, the object might be shown in
// a [class@Gtk.ListView] or other widget.
//
// The [signal@Gtk.SignalListItemFactory::unbind] signal is the
// opposite of this signal and can be used to undo everything done
// in this signal.
func (x *SignalListItemFactory) ConnectBind(cb func(SignalListItemFactory, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, ObjectVarp uintptr) {
		fa := SignalListItemFactory{}
		fa.Ptr = clsPtr

		cb(fa, ObjectVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "bind", purego.NewCallback(fcb))
}

// Emitted when a new listitem has been created and needs to be setup for use.
//
// It is the first signal emitted for every listitem.
//
// The [signal@Gtk.SignalListItemFactory::teardown] signal is the opposite
// of this signal and can be used to undo everything done in this signal.
func (x *SignalListItemFactory) ConnectSetup(cb func(SignalListItemFactory, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, ObjectVarp uintptr) {
		fa := SignalListItemFactory{}
		fa.Ptr = clsPtr

		cb(fa, ObjectVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "setup", purego.NewCallback(fcb))
}

// Emitted when an object is about to be destroyed.
//
// It is the last signal ever emitted for this @object.
//
// This signal is the opposite of the [signal@Gtk.SignalListItemFactory::setup]
// signal and should be used to undo everything done in that signal.
func (x *SignalListItemFactory) ConnectTeardown(cb func(SignalListItemFactory, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, ObjectVarp uintptr) {
		fa := SignalListItemFactory{}
		fa.Ptr = clsPtr

		cb(fa, ObjectVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "teardown", purego.NewCallback(fcb))
}

// Emitted when a object has been unbound from its item, for example when
// a listitem was removed from use in a list widget
// and its new [property@Gtk.ListItem:item] is about to be unset.
//
// This signal is the opposite of the [signal@Gtk.SignalListItemFactory::bind]
// signal and should be used to undo everything done in that signal.
func (x *SignalListItemFactory) ConnectUnbind(cb func(SignalListItemFactory, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, ObjectVarp uintptr) {
		fa := SignalListItemFactory{}
		fa.Ptr = clsPtr

		cb(fa, ObjectVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "unbind", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSignalListItemFactory, lib, "gtk_signal_list_item_factory_new")

}
