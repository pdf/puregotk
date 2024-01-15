// Package graphene was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package graphene

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A size.
type Size struct {
	Width float32

	Height float32
}

func (x *Size) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xSizeAlloc func() *Size

// Allocates a new #graphene_size_t.
//
// The contents of the returned value are undefined.
func SizeAlloc() *Size {

	cret := xSizeAlloc()
	return cret
}

var xSizeEqual func(uintptr, *Size) bool

// Checks whether the two give #graphene_size_t are equal.
func (x *Size) Equal(BVar *Size) bool {

	cret := xSizeEqual(x.GoPointer(), BVar)
	return cret
}

var xSizeFree func(uintptr)

// Frees the resources allocated by graphene_size_alloc().
func (x *Size) Free() {

	xSizeFree(x.GoPointer())

}

var xSizeInit func(uintptr, float32, float32) *Size

// Initializes a #graphene_size_t using the given @width and @height.
func (x *Size) Init(WidthVar float32, HeightVar float32) *Size {

	cret := xSizeInit(x.GoPointer(), WidthVar, HeightVar)
	return cret
}

var xSizeInitFromSize func(uintptr, *Size) *Size

// Initializes a #graphene_size_t using the width and height of
// the given @src.
func (x *Size) InitFromSize(SrcVar *Size) *Size {

	cret := xSizeInitFromSize(x.GoPointer(), SrcVar)
	return cret
}

var xSizeInterpolate func(uintptr, *Size, float64, *Size)

// Linearly interpolates the two given #graphene_size_t using the given
// interpolation @factor.
func (x *Size) Interpolate(BVar *Size, FactorVar float64, ResVar *Size) {

	xSizeInterpolate(x.GoPointer(), BVar, FactorVar, ResVar)

}

var xSizeScale func(uintptr, float32, *Size)

// Scales the components of a #graphene_size_t using the given @factor.
func (x *Size) Scale(FactorVar float32, ResVar *Size) {

	xSizeScale(x.GoPointer(), FactorVar, ResVar)

}

var xSizeZero func() *Size

// A constant pointer to a zero #graphene_size_t, useful for
// equality checks and interpolations.
func SizeZero() *Size {

	cret := xSizeZero()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GRAPHENE"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xSizeZero, lib, "graphene_size_zero")

	core.PuregoSafeRegister(&xSizeAlloc, lib, "graphene_size_alloc")

	core.PuregoSafeRegister(&xSizeEqual, lib, "graphene_size_equal")
	core.PuregoSafeRegister(&xSizeFree, lib, "graphene_size_free")
	core.PuregoSafeRegister(&xSizeInit, lib, "graphene_size_init")
	core.PuregoSafeRegister(&xSizeInitFromSize, lib, "graphene_size_init_from_size")
	core.PuregoSafeRegister(&xSizeInterpolate, lib, "graphene_size_interpolate")
	core.PuregoSafeRegister(&xSizeScale, lib, "graphene_size_scale")

}
