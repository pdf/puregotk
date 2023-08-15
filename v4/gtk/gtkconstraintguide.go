// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ConstraintGuideClass struct {
	ParentClass uintptr
}

// A `GtkConstraintGuide` is an invisible layout element in a
// `GtkConstraintLayout`.
//
// The `GtkConstraintLayout` treats guides like widgets. They
// can be used as the source or target of a `GtkConstraint`.
//
// Guides have a minimum, maximum and natural size. Depending
// on the constraints that are applied, they can act like a
// guideline that widgets can be aligned to, or like *flexible
// space*.
//
// Unlike a `GtkWidget`, a `GtkConstraintGuide` will not be drawn.
type ConstraintGuide struct {
	gobject.Object
}

func ConstraintGuideNewFromInternalPtr(ptr uintptr) *ConstraintGuide {
	cls := &ConstraintGuide{}
	cls.Ptr = ptr
	return cls
}

var xNewConstraintGuide func() uintptr

// Creates a new `GtkConstraintGuide` object.
func NewConstraintGuide() *ConstraintGuide {
	NewConstraintGuidePtr := xNewConstraintGuide()
	if NewConstraintGuidePtr == 0 {
		return nil
	}

	NewConstraintGuideCls := &ConstraintGuide{}
	NewConstraintGuideCls.Ptr = NewConstraintGuidePtr
	return NewConstraintGuideCls
}

var xConstraintGuideGetMaxSize func(uintptr, int, int)

// Gets the maximum size of @guide.
func (x *ConstraintGuide) GetMaxSize(WidthVar int, HeightVar int) {

	xConstraintGuideGetMaxSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideGetMinSize func(uintptr, int, int)

// Gets the minimum size of @guide.
func (x *ConstraintGuide) GetMinSize(WidthVar int, HeightVar int) {

	xConstraintGuideGetMinSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideGetName func(uintptr) string

// Retrieves the name set using gtk_constraint_guide_set_name().
func (x *ConstraintGuide) GetName() string {

	return xConstraintGuideGetName(x.GoPointer())

}

var xConstraintGuideGetNatSize func(uintptr, int, int)

// Gets the natural size of @guide.
func (x *ConstraintGuide) GetNatSize(WidthVar int, HeightVar int) {

	xConstraintGuideGetNatSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideGetStrength func(uintptr) ConstraintStrength

// Retrieves the strength set using gtk_constraint_guide_set_strength().
func (x *ConstraintGuide) GetStrength() ConstraintStrength {

	return xConstraintGuideGetStrength(x.GoPointer())

}

var xConstraintGuideSetMaxSize func(uintptr, int, int)

// Sets the maximum size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`,
// the constraints will be updated to reflect the new size.
func (x *ConstraintGuide) SetMaxSize(WidthVar int, HeightVar int) {

	xConstraintGuideSetMaxSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideSetMinSize func(uintptr, int, int)

// Sets the minimum size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`,
// the constraints will be updated to reflect the new size.
func (x *ConstraintGuide) SetMinSize(WidthVar int, HeightVar int) {

	xConstraintGuideSetMinSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideSetName func(uintptr, string)

// Sets a name for the given `GtkConstraintGuide`.
//
// The name is useful for debugging purposes.
func (x *ConstraintGuide) SetName(NameVar string) {

	xConstraintGuideSetName(x.GoPointer(), NameVar)

}

var xConstraintGuideSetNatSize func(uintptr, int, int)

// Sets the natural size of @guide.
//
// If @guide is attached to a `GtkConstraintLayout`,
// the constraints will be updated to reflect the new size.
func (x *ConstraintGuide) SetNatSize(WidthVar int, HeightVar int) {

	xConstraintGuideSetNatSize(x.GoPointer(), WidthVar, HeightVar)

}

var xConstraintGuideSetStrength func(uintptr, ConstraintStrength)

// Sets the strength of the constraint on the natural size of the
// given `GtkConstraintGuide`.
func (x *ConstraintGuide) SetStrength(StrengthVar ConstraintStrength) {

	xConstraintGuideSetStrength(x.GoPointer(), StrengthVar)

}

func (c *ConstraintGuide) GoPointer() uintptr {
	return c.Ptr
}

func (c *ConstraintGuide) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewConstraintGuide, lib, "gtk_constraint_guide_new")

	core.PuregoSafeRegister(&xConstraintGuideGetMaxSize, lib, "gtk_constraint_guide_get_max_size")
	core.PuregoSafeRegister(&xConstraintGuideGetMinSize, lib, "gtk_constraint_guide_get_min_size")
	core.PuregoSafeRegister(&xConstraintGuideGetName, lib, "gtk_constraint_guide_get_name")
	core.PuregoSafeRegister(&xConstraintGuideGetNatSize, lib, "gtk_constraint_guide_get_nat_size")
	core.PuregoSafeRegister(&xConstraintGuideGetStrength, lib, "gtk_constraint_guide_get_strength")
	core.PuregoSafeRegister(&xConstraintGuideSetMaxSize, lib, "gtk_constraint_guide_set_max_size")
	core.PuregoSafeRegister(&xConstraintGuideSetMinSize, lib, "gtk_constraint_guide_set_min_size")
	core.PuregoSafeRegister(&xConstraintGuideSetName, lib, "gtk_constraint_guide_set_name")
	core.PuregoSafeRegister(&xConstraintGuideSetNatSize, lib, "gtk_constraint_guide_set_nat_size")
	core.PuregoSafeRegister(&xConstraintGuideSetStrength, lib, "gtk_constraint_guide_set_strength")

}
