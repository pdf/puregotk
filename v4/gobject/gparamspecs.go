// Package gobject was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gobject

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

var xNewParamSpecBoolean func(string, string, string, bool, ParamFlags) uintptr

// Creates a new #GParamSpecBoolean instance specifying a %G_TYPE_BOOLEAN
// property. In many cases, it may be more appropriate to use an enum with
// g_param_spec_enum(), both to improve code clarity by using explicitly named
// values, and to allow for more values to be added in future without breaking
// API.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecBoolean(NameVar string, NickVar string, BlurbVar string, DefaultValueVar bool, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecBooleanPtr := xNewParamSpecBoolean(NameVar, NickVar, BlurbVar, DefaultValueVar, FlagsVar)
	if NewParamSpecBooleanPtr == 0 {
		return nil
	}

	NewParamSpecBooleanCls := &ParamSpec{}
	NewParamSpecBooleanCls.Ptr = NewParamSpecBooleanPtr
	return NewParamSpecBooleanCls

}

var xNewParamSpecBoxed func(string, string, string, []interface{}, ParamFlags) uintptr

// Creates a new #GParamSpecBoxed instance specifying a %G_TYPE_BOXED
// derived property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecBoxed(NameVar string, NickVar string, BlurbVar string, BoxedTypeVar []interface{}, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecBoxedPtr := xNewParamSpecBoxed(NameVar, NickVar, BlurbVar, BoxedTypeVar, FlagsVar)
	if NewParamSpecBoxedPtr == 0 {
		return nil
	}

	NewParamSpecBoxedCls := &ParamSpec{}
	NewParamSpecBoxedCls.Ptr = NewParamSpecBoxedPtr
	return NewParamSpecBoxedCls

}

var xNewParamSpecChar func(string, string, string, int8, int8, int8, ParamFlags) uintptr

// Creates a new #GParamSpecChar instance specifying a %G_TYPE_CHAR property.
func NewParamSpecChar(NameVar string, NickVar string, BlurbVar string, MinimumVar int8, MaximumVar int8, DefaultValueVar int8, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecCharPtr := xNewParamSpecChar(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecCharPtr == 0 {
		return nil
	}

	NewParamSpecCharCls := &ParamSpec{}
	NewParamSpecCharCls.Ptr = NewParamSpecCharPtr
	return NewParamSpecCharCls

}

var xNewParamSpecDouble func(string, string, string, float64, float64, float64, ParamFlags) uintptr

// Creates a new #GParamSpecDouble instance specifying a %G_TYPE_DOUBLE
// property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecDouble(NameVar string, NickVar string, BlurbVar string, MinimumVar float64, MaximumVar float64, DefaultValueVar float64, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecDoublePtr := xNewParamSpecDouble(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecDoublePtr == 0 {
		return nil
	}

	NewParamSpecDoubleCls := &ParamSpec{}
	NewParamSpecDoubleCls.Ptr = NewParamSpecDoublePtr
	return NewParamSpecDoubleCls

}

var xNewParamSpecEnum func(string, string, string, []interface{}, int, ParamFlags) uintptr

// Creates a new #GParamSpecEnum instance specifying a %G_TYPE_ENUM
// property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecEnum(NameVar string, NickVar string, BlurbVar string, EnumTypeVar []interface{}, DefaultValueVar int, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecEnumPtr := xNewParamSpecEnum(NameVar, NickVar, BlurbVar, EnumTypeVar, DefaultValueVar, FlagsVar)
	if NewParamSpecEnumPtr == 0 {
		return nil
	}

	NewParamSpecEnumCls := &ParamSpec{}
	NewParamSpecEnumCls.Ptr = NewParamSpecEnumPtr
	return NewParamSpecEnumCls

}

var xNewParamSpecFlags func(string, string, string, []interface{}, uint, ParamFlags) uintptr

// Creates a new #GParamSpecFlags instance specifying a %G_TYPE_FLAGS
// property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecFlags(NameVar string, NickVar string, BlurbVar string, FlagsTypeVar []interface{}, DefaultValueVar uint, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecFlagsPtr := xNewParamSpecFlags(NameVar, NickVar, BlurbVar, FlagsTypeVar, DefaultValueVar, FlagsVar)
	if NewParamSpecFlagsPtr == 0 {
		return nil
	}

	NewParamSpecFlagsCls := &ParamSpec{}
	NewParamSpecFlagsCls.Ptr = NewParamSpecFlagsPtr
	return NewParamSpecFlagsCls

}

var xNewParamSpecFloat func(string, string, string, float32, float32, float32, ParamFlags) uintptr

// Creates a new #GParamSpecFloat instance specifying a %G_TYPE_FLOAT property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecFloat(NameVar string, NickVar string, BlurbVar string, MinimumVar float32, MaximumVar float32, DefaultValueVar float32, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecFloatPtr := xNewParamSpecFloat(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecFloatPtr == 0 {
		return nil
	}

	NewParamSpecFloatCls := &ParamSpec{}
	NewParamSpecFloatCls.Ptr = NewParamSpecFloatPtr
	return NewParamSpecFloatCls

}

var xParamSpecGtype func(string, string, string, []interface{}, ParamFlags) uintptr

// Creates a new #GParamSpecGType instance specifying a
// %G_TYPE_GTYPE property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecGtype(NameVar string, NickVar string, BlurbVar string, IsATypeVar []interface{}, FlagsVar ParamFlags) *ParamSpec {

	ParamSpecGtypePtr := xParamSpecGtype(NameVar, NickVar, BlurbVar, IsATypeVar, FlagsVar)
	if ParamSpecGtypePtr == 0 {
		return nil
	}

	ParamSpecGtypeCls := &ParamSpec{}
	ParamSpecGtypeCls.Ptr = ParamSpecGtypePtr
	return ParamSpecGtypeCls

}

var xNewParamSpecInt func(string, string, string, int, int, int, ParamFlags) uintptr

// Creates a new #GParamSpecInt instance specifying a %G_TYPE_INT property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecInt(NameVar string, NickVar string, BlurbVar string, MinimumVar int, MaximumVar int, DefaultValueVar int, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecIntPtr := xNewParamSpecInt(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecIntPtr == 0 {
		return nil
	}

	NewParamSpecIntCls := &ParamSpec{}
	NewParamSpecIntCls.Ptr = NewParamSpecIntPtr
	return NewParamSpecIntCls

}

var xNewParamSpecInt64 func(string, string, string, int64, int64, int64, ParamFlags) uintptr

// Creates a new #GParamSpecInt64 instance specifying a %G_TYPE_INT64 property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecInt64(NameVar string, NickVar string, BlurbVar string, MinimumVar int64, MaximumVar int64, DefaultValueVar int64, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecInt64Ptr := xNewParamSpecInt64(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecInt64Ptr == 0 {
		return nil
	}

	NewParamSpecInt64Cls := &ParamSpec{}
	NewParamSpecInt64Cls.Ptr = NewParamSpecInt64Ptr
	return NewParamSpecInt64Cls

}

var xNewParamSpecLong func(string, string, string, int32, int32, int32, ParamFlags) uintptr

// Creates a new #GParamSpecLong instance specifying a %G_TYPE_LONG property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecLong(NameVar string, NickVar string, BlurbVar string, MinimumVar int32, MaximumVar int32, DefaultValueVar int32, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecLongPtr := xNewParamSpecLong(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if NewParamSpecLongPtr == 0 {
		return nil
	}

	NewParamSpecLongCls := &ParamSpec{}
	NewParamSpecLongCls.Ptr = NewParamSpecLongPtr
	return NewParamSpecLongCls

}

var xNewParamSpecObject func(string, string, string, []interface{}, ParamFlags) uintptr

// Creates a new #GParamSpecBoxed instance specifying a %G_TYPE_OBJECT
// derived property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecObject(NameVar string, NickVar string, BlurbVar string, ObjectTypeVar []interface{}, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecObjectPtr := xNewParamSpecObject(NameVar, NickVar, BlurbVar, ObjectTypeVar, FlagsVar)
	if NewParamSpecObjectPtr == 0 {
		return nil
	}

	NewParamSpecObjectCls := &ParamSpec{}
	NewParamSpecObjectCls.Ptr = NewParamSpecObjectPtr
	return NewParamSpecObjectCls

}

var xNewParamSpecOverride func(string, uintptr) uintptr

// Creates a new property of type #GParamSpecOverride. This is used
// to direct operations to another paramspec, and will not be directly
// useful unless you are implementing a new base type similar to GObject.
func NewParamSpecOverride(NameVar string, OverriddenVar *ParamSpec) *ParamSpec {

	NewParamSpecOverridePtr := xNewParamSpecOverride(NameVar, OverriddenVar.GoPointer())
	if NewParamSpecOverridePtr == 0 {
		return nil
	}

	NewParamSpecOverrideCls := &ParamSpec{}
	NewParamSpecOverrideCls.Ptr = NewParamSpecOverridePtr
	return NewParamSpecOverrideCls

}

var xNewParamSpecParam func(string, string, string, []interface{}, ParamFlags) uintptr

// Creates a new #GParamSpecParam instance specifying a %G_TYPE_PARAM
// property.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecParam(NameVar string, NickVar string, BlurbVar string, ParamTypeVar []interface{}, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecParamPtr := xNewParamSpecParam(NameVar, NickVar, BlurbVar, ParamTypeVar, FlagsVar)
	if NewParamSpecParamPtr == 0 {
		return nil
	}

	NewParamSpecParamCls := &ParamSpec{}
	NewParamSpecParamCls.Ptr = NewParamSpecParamPtr
	return NewParamSpecParamCls

}

var xNewParamSpecPointer func(string, string, string, ParamFlags) uintptr

// Creates a new #GParamSpecPointer instance specifying a pointer property.
// Where possible, it is better to use g_param_spec_object() or
// g_param_spec_boxed() to expose memory management information.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecPointer(NameVar string, NickVar string, BlurbVar string, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecPointerPtr := xNewParamSpecPointer(NameVar, NickVar, BlurbVar, FlagsVar)
	if NewParamSpecPointerPtr == 0 {
		return nil
	}

	NewParamSpecPointerCls := &ParamSpec{}
	NewParamSpecPointerCls.Ptr = NewParamSpecPointerPtr
	return NewParamSpecPointerCls

}

var xNewParamSpecString func(string, string, string, string, ParamFlags) uintptr

// Creates a new #GParamSpecString instance.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecString(NameVar string, NickVar string, BlurbVar string, DefaultValueVar string, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecStringPtr := xNewParamSpecString(NameVar, NickVar, BlurbVar, DefaultValueVar, FlagsVar)
	if NewParamSpecStringPtr == 0 {
		return nil
	}

	NewParamSpecStringCls := &ParamSpec{}
	NewParamSpecStringCls.Ptr = NewParamSpecStringPtr
	return NewParamSpecStringCls

}

var xParamSpecUchar func(string, string, string, byte, byte, byte, ParamFlags) uintptr

// Creates a new #GParamSpecUChar instance specifying a %G_TYPE_UCHAR property.
func ParamSpecUchar(NameVar string, NickVar string, BlurbVar string, MinimumVar byte, MaximumVar byte, DefaultValueVar byte, FlagsVar ParamFlags) *ParamSpec {

	ParamSpecUcharPtr := xParamSpecUchar(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if ParamSpecUcharPtr == 0 {
		return nil
	}

	ParamSpecUcharCls := &ParamSpec{}
	ParamSpecUcharCls.Ptr = ParamSpecUcharPtr
	return ParamSpecUcharCls

}

var xParamSpecUint func(string, string, string, uint, uint, uint, ParamFlags) uintptr

// Creates a new #GParamSpecUInt instance specifying a %G_TYPE_UINT property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint(NameVar string, NickVar string, BlurbVar string, MinimumVar uint, MaximumVar uint, DefaultValueVar uint, FlagsVar ParamFlags) *ParamSpec {

	ParamSpecUintPtr := xParamSpecUint(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if ParamSpecUintPtr == 0 {
		return nil
	}

	ParamSpecUintCls := &ParamSpec{}
	ParamSpecUintCls.Ptr = ParamSpecUintPtr
	return ParamSpecUintCls

}

var xParamSpecUint64 func(string, string, string, uint64, uint64, uint64, ParamFlags) uintptr

// Creates a new #GParamSpecUInt64 instance specifying a %G_TYPE_UINT64
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint64(NameVar string, NickVar string, BlurbVar string, MinimumVar uint64, MaximumVar uint64, DefaultValueVar uint64, FlagsVar ParamFlags) *ParamSpec {

	ParamSpecUint64Ptr := xParamSpecUint64(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if ParamSpecUint64Ptr == 0 {
		return nil
	}

	ParamSpecUint64Cls := &ParamSpec{}
	ParamSpecUint64Cls.Ptr = ParamSpecUint64Ptr
	return ParamSpecUint64Cls

}

var xParamSpecUlong func(string, string, string, uint32, uint32, uint32, ParamFlags) uintptr

// Creates a new #GParamSpecULong instance specifying a %G_TYPE_ULONG
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUlong(NameVar string, NickVar string, BlurbVar string, MinimumVar uint32, MaximumVar uint32, DefaultValueVar uint32, FlagsVar ParamFlags) *ParamSpec {

	ParamSpecUlongPtr := xParamSpecUlong(NameVar, NickVar, BlurbVar, MinimumVar, MaximumVar, DefaultValueVar, FlagsVar)
	if ParamSpecUlongPtr == 0 {
		return nil
	}

	ParamSpecUlongCls := &ParamSpec{}
	ParamSpecUlongCls.Ptr = ParamSpecUlongPtr
	return ParamSpecUlongCls

}

var xNewParamSpecUnichar func(string, string, string, uint32, ParamFlags) uintptr

// Creates a new #GParamSpecUnichar instance specifying a %G_TYPE_UINT
// property. #GValue structures for this property can be accessed with
// g_value_set_uint() and g_value_get_uint().
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecUnichar(NameVar string, NickVar string, BlurbVar string, DefaultValueVar uint32, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecUnicharPtr := xNewParamSpecUnichar(NameVar, NickVar, BlurbVar, DefaultValueVar, FlagsVar)
	if NewParamSpecUnicharPtr == 0 {
		return nil
	}

	NewParamSpecUnicharCls := &ParamSpec{}
	NewParamSpecUnicharCls.Ptr = NewParamSpecUnicharPtr
	return NewParamSpecUnicharCls

}

var xNewParamSpecValueArray func(string, string, string, uintptr, ParamFlags) uintptr

// Creates a new #GParamSpecValueArray instance specifying a
// %G_TYPE_VALUE_ARRAY property. %G_TYPE_VALUE_ARRAY is a
// %G_TYPE_BOXED type, as such, #GValue structures for this property
// can be accessed with g_value_set_boxed() and g_value_get_boxed().
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecValueArray(NameVar string, NickVar string, BlurbVar string, ElementSpecVar *ParamSpec, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecValueArrayPtr := xNewParamSpecValueArray(NameVar, NickVar, BlurbVar, ElementSpecVar.GoPointer(), FlagsVar)
	if NewParamSpecValueArrayPtr == 0 {
		return nil
	}

	NewParamSpecValueArrayCls := &ParamSpec{}
	NewParamSpecValueArrayCls.Ptr = NewParamSpecValueArrayPtr
	return NewParamSpecValueArrayCls

}

var xNewParamSpecVariant func(string, string, string, *glib.VariantType, *glib.Variant, ParamFlags) uintptr

// Creates a new #GParamSpecVariant instance specifying a #GVariant
// property.
//
// If @default_value is floating, it is consumed.
//
// See g_param_spec_internal() for details on property names.
func NewParamSpecVariant(NameVar string, NickVar string, BlurbVar string, TypeVar *glib.VariantType, DefaultValueVar *glib.Variant, FlagsVar ParamFlags) *ParamSpec {

	NewParamSpecVariantPtr := xNewParamSpecVariant(NameVar, NickVar, BlurbVar, TypeVar, DefaultValueVar, FlagsVar)
	if NewParamSpecVariantPtr == 0 {
		return nil
	}

	NewParamSpecVariantCls := &ParamSpec{}
	NewParamSpecVariantCls.Ptr = NewParamSpecVariantPtr
	return NewParamSpecVariantCls

}

// A #GParamSpec derived structure that contains the meta data for boolean properties.
type ParamSpecBoolean struct {
	ParamSpec
}

func ParamSpecBooleanNewFromInternalPtr(ptr uintptr) *ParamSpecBoolean {
	cls := &ParamSpecBoolean{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecBoolean) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecBoolean) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for boxed properties.
type ParamSpecBoxed struct {
	ParamSpec
}

func ParamSpecBoxedNewFromInternalPtr(ptr uintptr) *ParamSpecBoxed {
	cls := &ParamSpecBoxed{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecBoxed) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecBoxed) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for character properties.
type ParamSpecChar struct {
	ParamSpec
}

func ParamSpecCharNewFromInternalPtr(ptr uintptr) *ParamSpecChar {
	cls := &ParamSpecChar{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecChar) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecChar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for double properties.
type ParamSpecDouble struct {
	ParamSpec
}

func ParamSpecDoubleNewFromInternalPtr(ptr uintptr) *ParamSpecDouble {
	cls := &ParamSpecDouble{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecDouble) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecDouble) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for enum
// properties.
type ParamSpecEnum struct {
	ParamSpec
}

func ParamSpecEnumNewFromInternalPtr(ptr uintptr) *ParamSpecEnum {
	cls := &ParamSpecEnum{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecEnum) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecEnum) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for flags
// properties.
type ParamSpecFlags struct {
	ParamSpec
}

func ParamSpecFlagsNewFromInternalPtr(ptr uintptr) *ParamSpecFlags {
	cls := &ParamSpecFlags{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecFlags) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecFlags) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for float properties.
type ParamSpecFloat struct {
	ParamSpec
}

func ParamSpecFloatNewFromInternalPtr(ptr uintptr) *ParamSpecFloat {
	cls := &ParamSpecFloat{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecFloat) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecFloat) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for #GType properties.
type ParamSpecGType struct {
	ParamSpec
}

func ParamSpecGTypeNewFromInternalPtr(ptr uintptr) *ParamSpecGType {
	cls := &ParamSpecGType{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecGType) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecGType) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for integer properties.
type ParamSpecInt struct {
	ParamSpec
}

func ParamSpecIntNewFromInternalPtr(ptr uintptr) *ParamSpecInt {
	cls := &ParamSpecInt{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecInt) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecInt) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for 64bit integer properties.
type ParamSpecInt64 struct {
	ParamSpec
}

func ParamSpecInt64NewFromInternalPtr(ptr uintptr) *ParamSpecInt64 {
	cls := &ParamSpecInt64{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecInt64) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecInt64) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for long integer properties.
type ParamSpecLong struct {
	ParamSpec
}

func ParamSpecLongNewFromInternalPtr(ptr uintptr) *ParamSpecLong {
	cls := &ParamSpecLong{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecLong) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecLong) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for object properties.
type ParamSpecObject struct {
	ParamSpec
}

func ParamSpecObjectNewFromInternalPtr(ptr uintptr) *ParamSpecObject {
	cls := &ParamSpecObject{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecObject) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecObject) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that redirects operations to
// other types of #GParamSpec.
//
// All operations other than getting or setting the value are redirected,
// including accessing the nick and blurb, validating a value, and so
// forth.
//
// See g_param_spec_get_redirect_target() for retrieving the overridden
// property. #GParamSpecOverride is used in implementing
// g_object_class_override_property(), and will not be directly useful
// unless you are implementing a new base type similar to GObject.
type ParamSpecOverride struct {
	ParamSpec
}

func ParamSpecOverrideNewFromInternalPtr(ptr uintptr) *ParamSpecOverride {
	cls := &ParamSpecOverride{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecOverride) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecOverride) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for %G_TYPE_PARAM
// properties.
type ParamSpecParam struct {
	ParamSpec
}

func ParamSpecParamNewFromInternalPtr(ptr uintptr) *ParamSpecParam {
	cls := &ParamSpecParam{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecParam) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecParam) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for pointer properties.
type ParamSpecPointer struct {
	ParamSpec
}

func ParamSpecPointerNewFromInternalPtr(ptr uintptr) *ParamSpecPointer {
	cls := &ParamSpecPointer{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecPointer) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecPointer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for string
// properties.
type ParamSpecString struct {
	ParamSpec
}

func ParamSpecStringNewFromInternalPtr(ptr uintptr) *ParamSpecString {
	cls := &ParamSpecString{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecString) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecString) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for unsigned character properties.
type ParamSpecUChar struct {
	ParamSpec
}

func ParamSpecUCharNewFromInternalPtr(ptr uintptr) *ParamSpecUChar {
	cls := &ParamSpecUChar{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecUChar) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecUChar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for unsigned integer properties.
type ParamSpecUInt struct {
	ParamSpec
}

func ParamSpecUIntNewFromInternalPtr(ptr uintptr) *ParamSpecUInt {
	cls := &ParamSpecUInt{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecUInt) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecUInt) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for unsigned 64bit integer properties.
type ParamSpecUInt64 struct {
	ParamSpec
}

func ParamSpecUInt64NewFromInternalPtr(ptr uintptr) *ParamSpecUInt64 {
	cls := &ParamSpecUInt64{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecUInt64) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecUInt64) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for unsigned long integer properties.
type ParamSpecULong struct {
	ParamSpec
}

func ParamSpecULongNewFromInternalPtr(ptr uintptr) *ParamSpecULong {
	cls := &ParamSpecULong{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecULong) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecULong) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for unichar (unsigned integer) properties.
type ParamSpecUnichar struct {
	ParamSpec
}

func ParamSpecUnicharNewFromInternalPtr(ptr uintptr) *ParamSpecUnichar {
	cls := &ParamSpecUnichar{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecUnichar) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecUnichar) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for #GValueArray properties.
type ParamSpecValueArray struct {
	ParamSpec
}

func ParamSpecValueArrayNewFromInternalPtr(ptr uintptr) *ParamSpecValueArray {
	cls := &ParamSpecValueArray{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecValueArray) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecValueArray) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// A #GParamSpec derived structure that contains the meta data for #GVariant properties.
//
// When comparing values with g_param_values_cmp(), scalar values with the same
// type will be compared with g_variant_compare(). Other non-%NULL variants will
// be checked for equality with g_variant_equal(), and their sort order is
// otherwise undefined. %NULL is ordered before non-%NULL variants. Two %NULL
// values compare equal.
type ParamSpecVariant struct {
	ParamSpec
}

func ParamSpecVariantNewFromInternalPtr(ptr uintptr) *ParamSpecVariant {
	cls := &ParamSpecVariant{}
	cls.Ptr = ptr
	return cls
}

func (c *ParamSpecVariant) GoPointer() uintptr {
	return c.Ptr
}

func (c *ParamSpecVariant) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GOBJECT"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xNewParamSpecBoolean, lib, "g_param_spec_boolean")
	core.PuregoSafeRegister(&xNewParamSpecBoxed, lib, "g_param_spec_boxed")
	core.PuregoSafeRegister(&xNewParamSpecChar, lib, "g_param_spec_char")
	core.PuregoSafeRegister(&xNewParamSpecDouble, lib, "g_param_spec_double")
	core.PuregoSafeRegister(&xNewParamSpecEnum, lib, "g_param_spec_enum")
	core.PuregoSafeRegister(&xNewParamSpecFlags, lib, "g_param_spec_flags")
	core.PuregoSafeRegister(&xNewParamSpecFloat, lib, "g_param_spec_float")
	core.PuregoSafeRegister(&xParamSpecGtype, lib, "g_param_spec_gtype")
	core.PuregoSafeRegister(&xNewParamSpecInt, lib, "g_param_spec_int")
	core.PuregoSafeRegister(&xNewParamSpecInt64, lib, "g_param_spec_int64")
	core.PuregoSafeRegister(&xNewParamSpecLong, lib, "g_param_spec_long")
	core.PuregoSafeRegister(&xNewParamSpecObject, lib, "g_param_spec_object")
	core.PuregoSafeRegister(&xNewParamSpecOverride, lib, "g_param_spec_override")
	core.PuregoSafeRegister(&xNewParamSpecParam, lib, "g_param_spec_param")
	core.PuregoSafeRegister(&xNewParamSpecPointer, lib, "g_param_spec_pointer")
	core.PuregoSafeRegister(&xNewParamSpecString, lib, "g_param_spec_string")
	core.PuregoSafeRegister(&xParamSpecUchar, lib, "g_param_spec_uchar")
	core.PuregoSafeRegister(&xParamSpecUint, lib, "g_param_spec_uint")
	core.PuregoSafeRegister(&xParamSpecUint64, lib, "g_param_spec_uint64")
	core.PuregoSafeRegister(&xParamSpecUlong, lib, "g_param_spec_ulong")
	core.PuregoSafeRegister(&xNewParamSpecUnichar, lib, "g_param_spec_unichar")
	core.PuregoSafeRegister(&xNewParamSpecValueArray, lib, "g_param_spec_value_array")
	core.PuregoSafeRegister(&xNewParamSpecVariant, lib, "g_param_spec_variant")

}
