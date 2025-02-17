// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type OrientableIface struct {
	BaseIface uintptr
}

// The `GtkOrientable` interface is implemented by all widgets that can be
// oriented horizontally or vertically.
//
// `GtkOrientable` is more flexible in that it allows the orientation to be
// changed at runtime, allowing the widgets to “flip”.
type Orientable interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
	GetOrientation() Orientation
	SetOrientation(OrientationVar Orientation)
}
type OrientableBase struct {
	Ptr uintptr
}

func (x *OrientableBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *OrientableBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

// Retrieves the orientation of the @orientable.
func (x *OrientableBase) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *OrientableBase) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

var XGtkOrientableGetOrientation func(uintptr) Orientation
var XGtkOrientableSetOrientation func(uintptr, Orientation)

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&XGtkOrientableGetOrientation, lib, "gtk_orientable_get_orientation")
	core.PuregoSafeRegister(&XGtkOrientableSetOrientation, lib, "gtk_orientable_set_orientation")

}
