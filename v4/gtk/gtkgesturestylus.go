// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type GestureStylusClass struct {
}

// `GtkGestureStylus` is a `GtkGesture` specific to stylus input.
//
// The provided signals just relay the basic information of the
// stylus events.
type GestureStylus struct {
	GestureSingle
}

func GestureStylusNewFromInternalPtr(ptr uintptr) *GestureStylus {
	cls := &GestureStylus{}
	cls.Ptr = ptr
	return cls
}

var xNewGestureStylus func() uintptr

// Creates a new `GtkGestureStylus`.
func NewGestureStylus() *Gesture {
	var cls *Gesture

	cret := xNewGestureStylus()

	if cret == 0 {
		return nil
	}
	cls = &Gesture{}
	cls.Ptr = cret
	return cls
}

var xGestureStylusGetAxes func(uintptr, uintptr, uintptr) bool

// Returns the current values for the requested @axes.
//
// This function must be called from the handler of one of the
// [signal@Gtk.GestureStylus::down], [signal@Gtk.GestureStylus::motion],
// [signal@Gtk.GestureStylus::up] or [signal@Gtk.GestureStylus::proximity]
// signals.
func (x *GestureStylus) GetAxes(AxesVar uintptr, ValuesVar uintptr) bool {

	cret := xGestureStylusGetAxes(x.GoPointer(), AxesVar, ValuesVar)
	return cret
}

var xGestureStylusGetAxis func(uintptr, gdk.AxisUse, float64) bool

// Returns the current value for the requested @axis.
//
// This function must be called from the handler of one of the
// [signal@Gtk.GestureStylus::down], [signal@Gtk.GestureStylus::motion],
// [signal@Gtk.GestureStylus::up] or [signal@Gtk.GestureStylus::proximity]
// signals.
func (x *GestureStylus) GetAxis(AxisVar gdk.AxisUse, ValueVar float64) bool {

	cret := xGestureStylusGetAxis(x.GoPointer(), AxisVar, ValueVar)
	return cret
}

var xGestureStylusGetBacklog func(uintptr, uintptr, uint) bool

// Returns the accumulated backlog of tracking information.
//
// By default, GTK will limit rate of input events. On stylus input
// where accuracy of strokes is paramount, this function returns the
// accumulated coordinate/timing state before the emission of the
// current [Gtk.GestureStylus::motion] signal.
//
// This function may only be called within a [signal@Gtk.GestureStylus::motion]
// signal handler, the state given in this signal and obtainable through
// [method@Gtk.GestureStylus.get_axis] express the latest (most up-to-date)
// state in motion history.
//
// The @backlog is provided in chronological order.
func (x *GestureStylus) GetBacklog(BacklogVar uintptr, NElemsVar uint) bool {

	cret := xGestureStylusGetBacklog(x.GoPointer(), BacklogVar, NElemsVar)
	return cret
}

var xGestureStylusGetDeviceTool func(uintptr) uintptr

// Returns the `GdkDeviceTool` currently driving input through this gesture.
//
// This function must be called from the handler of one of the
// [signal@Gtk.GestureStylus::down], [signal@Gtk.GestureStylus::motion],
// [signal@Gtk.GestureStylus::up] or [signal@Gtk.GestureStylus::proximity]
// signals.
func (x *GestureStylus) GetDeviceTool() *gdk.DeviceTool {
	var cls *gdk.DeviceTool

	cret := xGestureStylusGetDeviceTool(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gdk.DeviceTool{}
	cls.Ptr = cret
	return cls
}

func (c *GestureStylus) GoPointer() uintptr {
	return c.Ptr
}

func (c *GestureStylus) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the stylus touches the device.
func (x *GestureStylus) ConnectDown(cb func(GestureStylus, float64, float64)) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := GestureStylus{}
		fa.Ptr = clsPtr

		cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "down", purego.NewCallback(fcb))
}

// Emitted when the stylus moves while touching the device.
func (x *GestureStylus) ConnectMotion(cb func(GestureStylus, float64, float64)) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := GestureStylus{}
		fa.Ptr = clsPtr

		cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "motion", purego.NewCallback(fcb))
}

// Emitted when the stylus is in proximity of the device.
func (x *GestureStylus) ConnectProximity(cb func(GestureStylus, float64, float64)) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := GestureStylus{}
		fa.Ptr = clsPtr

		cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "proximity", purego.NewCallback(fcb))
}

// Emitted when the stylus no longer touches the device.
func (x *GestureStylus) ConnectUp(cb func(GestureStylus, float64, float64)) uint32 {
	fcb := func(clsPtr uintptr, XVarp float64, YVarp float64) {
		fa := GestureStylus{}
		fa.Ptr = clsPtr

		cb(fa, XVarp, YVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "up", purego.NewCallback(fcb))
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGestureStylus, lib, "gtk_gesture_stylus_new")

	core.PuregoSafeRegister(&xGestureStylusGetAxes, lib, "gtk_gesture_stylus_get_axes")
	core.PuregoSafeRegister(&xGestureStylusGetAxis, lib, "gtk_gesture_stylus_get_axis")
	core.PuregoSafeRegister(&xGestureStylusGetBacklog, lib, "gtk_gesture_stylus_get_backlog")
	core.PuregoSafeRegister(&xGestureStylusGetDeviceTool, lib, "gtk_gesture_stylus_get_device_tool")

}
