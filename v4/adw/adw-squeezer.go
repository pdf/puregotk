// Package adw was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package adw

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/gtk"
)

type SqueezerClass struct {
	ParentClass uintptr
}

type SqueezerPageClass struct {
	ParentClass uintptr
}

// Describes the possible transitions in a [class@Squeezer] widget.
type SqueezerTransitionType int

const (

	// No transition
	SqueezerTransitionTypeNoneValue SqueezerTransitionType = 0
	// A cross-fade
	SqueezerTransitionTypeCrossfadeValue SqueezerTransitionType = 1
)

// A best fit container.
//
// &lt;picture&gt;
//
//	&lt;source srcset="squeezer-wide-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="squeezer-wide.png" alt="squeezer-wide"&gt;
//
// &lt;/picture&gt;
// &lt;picture&gt;
//
//	&lt;source srcset="squeezer-narrow-dark.png" media="(prefers-color-scheme: dark)"&gt;
//	&lt;img src="squeezer-narrow.png" alt="squeezer-narrow"&gt;
//
// &lt;/picture&gt;
//
// The `AdwSqueezer` widget is a container which only shows the first of its
// children that fits in the available size. It is convenient to offer different
// widgets to represent the same data with different levels of detail, making
// the widget seem to squeeze itself to fit in the available space.
//
// Transitions between children can be animated as fades. This can be controlled
// with [property@Squeezer:transition-type].
//
// ## CSS nodes
//
// `AdwSqueezer` has a single CSS node with name `squeezer`.
type Squeezer struct {
	gtk.Widget
}

func SqueezerNewFromInternalPtr(ptr uintptr) *Squeezer {
	cls := &Squeezer{}
	cls.Ptr = ptr
	return cls
}

var xNewSqueezer func() uintptr

// Creates a new `AdwSqueezer`.
func NewSqueezer() *gtk.Widget {
	NewSqueezerPtr := xNewSqueezer()
	if NewSqueezerPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewSqueezerPtr)

	NewSqueezerCls := &gtk.Widget{}
	NewSqueezerCls.Ptr = NewSqueezerPtr
	return NewSqueezerCls
}

var xSqueezerAdd func(uintptr, uintptr) uintptr

// Adds a child to @self.
func (x *Squeezer) Add(ChildVar *gtk.Widget) *SqueezerPage {

	AddPtr := xSqueezerAdd(x.GoPointer(), ChildVar.GoPointer())
	if AddPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(AddPtr)

	AddCls := &SqueezerPage{}
	AddCls.Ptr = AddPtr
	return AddCls

}

var xSqueezerGetAllowNone func(uintptr) bool

// Gets whether to allow squeezing beyond the last child's minimum size.
func (x *Squeezer) GetAllowNone() bool {

	return xSqueezerGetAllowNone(x.GoPointer())

}

var xSqueezerGetHomogeneous func(uintptr) bool

// Gets whether all children have the same size for the opposite orientation.
func (x *Squeezer) GetHomogeneous() bool {

	return xSqueezerGetHomogeneous(x.GoPointer())

}

var xSqueezerGetInterpolateSize func(uintptr) bool

// Gets whether @self interpolates its size when changing the visible child.
func (x *Squeezer) GetInterpolateSize() bool {

	return xSqueezerGetInterpolateSize(x.GoPointer())

}

var xSqueezerGetPage func(uintptr, uintptr) uintptr

// Returns the [class@SqueezerPage] object for @child.
func (x *Squeezer) GetPage(ChildVar *gtk.Widget) *SqueezerPage {

	GetPagePtr := xSqueezerGetPage(x.GoPointer(), ChildVar.GoPointer())
	if GetPagePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetPagePtr)

	GetPageCls := &SqueezerPage{}
	GetPageCls.Ptr = GetPagePtr
	return GetPageCls

}

var xSqueezerGetPages func(uintptr) uintptr

// Returns a [iface@Gio.ListModel] that contains the pages of @self.
//
// This can be used to keep an up-to-date view. The model also implements
// [iface@Gtk.SelectionModel] and can be used to track the visible page.
func (x *Squeezer) GetPages() *gtk.SelectionModelBase {

	GetPagesPtr := xSqueezerGetPages(x.GoPointer())
	if GetPagesPtr == 0 {
		return nil
	}

	GetPagesCls := &gtk.SelectionModelBase{}
	GetPagesCls.Ptr = GetPagesPtr
	return GetPagesCls

}

var xSqueezerGetSwitchThresholdPolicy func(uintptr) FoldThresholdPolicy

// Gets the switch threshold policy for @self.
func (x *Squeezer) GetSwitchThresholdPolicy() FoldThresholdPolicy {

	return xSqueezerGetSwitchThresholdPolicy(x.GoPointer())

}

var xSqueezerGetTransitionDuration func(uintptr) uint

// Gets the transition animation duration for @self.
func (x *Squeezer) GetTransitionDuration() uint {

	return xSqueezerGetTransitionDuration(x.GoPointer())

}

var xSqueezerGetTransitionRunning func(uintptr) bool

// Gets whether a transition is currently running for @self.
//
// If a transition is impossible, the property value will be set to `TRUE` and
// then immediately to `FALSE`, so it's possible to rely on its notifications
// to know that a transition has happened.
func (x *Squeezer) GetTransitionRunning() bool {

	return xSqueezerGetTransitionRunning(x.GoPointer())

}

var xSqueezerGetTransitionType func(uintptr) SqueezerTransitionType

// Gets the type of animation used for transitions between children in @self.
func (x *Squeezer) GetTransitionType() SqueezerTransitionType {

	return xSqueezerGetTransitionType(x.GoPointer())

}

var xSqueezerGetVisibleChild func(uintptr) uintptr

// Gets the currently visible child of @self.
func (x *Squeezer) GetVisibleChild() *gtk.Widget {

	GetVisibleChildPtr := xSqueezerGetVisibleChild(x.GoPointer())
	if GetVisibleChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetVisibleChildPtr)

	GetVisibleChildCls := &gtk.Widget{}
	GetVisibleChildCls.Ptr = GetVisibleChildPtr
	return GetVisibleChildCls

}

var xSqueezerGetXalign func(uintptr) float32

// Gets the horizontal alignment, from 0 (start) to 1 (end).
func (x *Squeezer) GetXalign() float32 {

	return xSqueezerGetXalign(x.GoPointer())

}

var xSqueezerGetYalign func(uintptr) float32

// Gets the vertical alignment, from 0 (top) to 1 (bottom).
func (x *Squeezer) GetYalign() float32 {

	return xSqueezerGetYalign(x.GoPointer())

}

var xSqueezerRemove func(uintptr, uintptr)

// Removes a child widget from @self.
func (x *Squeezer) Remove(ChildVar *gtk.Widget) {

	xSqueezerRemove(x.GoPointer(), ChildVar.GoPointer())

}

var xSqueezerSetAllowNone func(uintptr, bool)

// Sets whether to allow squeezing beyond the last child's minimum size.
//
// If set to `TRUE`, the squeezer can shrink to the point where no child can be
// shown. This is functionally equivalent to appending a widget with 0×0 minimum
// size.
func (x *Squeezer) SetAllowNone(AllowNoneVar bool) {

	xSqueezerSetAllowNone(x.GoPointer(), AllowNoneVar)

}

var xSqueezerSetHomogeneous func(uintptr, bool)

// Sets whether all children have the same size for the opposite orientation.
//
// For example, if a squeezer is horizontal and is homogeneous, it will request
// the same height for all its children. If it isn't, the squeezer may change
// size when a different child becomes visible.
func (x *Squeezer) SetHomogeneous(HomogeneousVar bool) {

	xSqueezerSetHomogeneous(x.GoPointer(), HomogeneousVar)

}

var xSqueezerSetInterpolateSize func(uintptr, bool)

// Sets whether @self interpolates its size when changing the visible child.
//
// If `TRUE`, the squeezer will interpolate its size between the one of the
// previous visible child and the one of the new visible child, according to the
// set transition duration and the orientation, e.g. if the squeezer is
// horizontal, it will interpolate the its height.
func (x *Squeezer) SetInterpolateSize(InterpolateSizeVar bool) {

	xSqueezerSetInterpolateSize(x.GoPointer(), InterpolateSizeVar)

}

var xSqueezerSetSwitchThresholdPolicy func(uintptr, FoldThresholdPolicy)

// Sets the switch threshold policy for @self.
//
// Determines when the squeezer will switch children.
//
// If set to `ADW_FOLD_THRESHOLD_POLICY_MINIMUM`, it will only switch when the
// visible child cannot fit anymore. With `ADW_FOLD_THRESHOLD_POLICY_NATURAL`,
// it will switch as soon as the visible child doesn't get their natural size.
//
// This can be useful if you have a long ellipsizing label and want to let it
// ellipsize instead of immediately switching.
func (x *Squeezer) SetSwitchThresholdPolicy(PolicyVar FoldThresholdPolicy) {

	xSqueezerSetSwitchThresholdPolicy(x.GoPointer(), PolicyVar)

}

var xSqueezerSetTransitionDuration func(uintptr, uint)

// Sets the transition animation duration for @self.
func (x *Squeezer) SetTransitionDuration(DurationVar uint) {

	xSqueezerSetTransitionDuration(x.GoPointer(), DurationVar)

}

var xSqueezerSetTransitionType func(uintptr, SqueezerTransitionType)

// Sets the type of animation used for transitions between children in @self.
func (x *Squeezer) SetTransitionType(TransitionVar SqueezerTransitionType) {

	xSqueezerSetTransitionType(x.GoPointer(), TransitionVar)

}

var xSqueezerSetXalign func(uintptr, float32)

// Sets the horizontal alignment, from 0 (start) to 1 (end).
//
// This affects the children allocation during transitions, when they exceed the
// size of the squeezer.
//
// For example, 0.5 means the child will be centered, 0 means it will keep the
// start side aligned and overflow the end side, and 1 means the opposite.
func (x *Squeezer) SetXalign(XalignVar float32) {

	xSqueezerSetXalign(x.GoPointer(), XalignVar)

}

var xSqueezerSetYalign func(uintptr, float32)

// Sets the vertical alignment, from 0 (top) to 1 (bottom).
//
// This affects the children allocation during transitions, when they exceed the
// size of the squeezer.
//
// For example, 0.5 means the child will be centered, 0 means it will keep the
// top side aligned and overflow the bottom side, and 1 means the opposite.
func (x *Squeezer) SetYalign(YalignVar float32) {

	xSqueezerSetYalign(x.GoPointer(), YalignVar)

}

func (c *Squeezer) GoPointer() uintptr {
	return c.Ptr
}

func (c *Squeezer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Squeezer) GetAccessibleRole() gtk.AccessibleRole {

	return gtk.XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Squeezer) ResetProperty(PropertyVar gtk.AccessibleProperty) {

	gtk.XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Squeezer) ResetRelation(RelationVar gtk.AccessibleRelation) {

	gtk.XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Squeezer) ResetState(StateVar gtk.AccessibleState) {

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
func (x *Squeezer) UpdateProperty(FirstPropertyVar gtk.AccessibleProperty, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Squeezer) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Squeezer) UpdateRelation(FirstRelationVar gtk.AccessibleRelation, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Squeezer) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Squeezer) UpdateState(FirstStateVar gtk.AccessibleState, varArgs ...interface{}) {

	gtk.XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Squeezer) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	gtk.XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Squeezer) GetBuildableId() string {

	return gtk.XGtkBuildableGetBuildableId(x.GoPointer())

}

// Retrieves the orientation of the @orientable.
func (x *Squeezer) GetOrientation() gtk.Orientation {

	return gtk.XGtkOrientableGetOrientation(x.GoPointer())

}

// Sets the orientation of the @orientable.
func (x *Squeezer) SetOrientation(OrientationVar gtk.Orientation) {

	gtk.XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

// An auxiliary class used by [class@Squeezer].
type SqueezerPage struct {
	gobject.Object
}

func SqueezerPageNewFromInternalPtr(ptr uintptr) *SqueezerPage {
	cls := &SqueezerPage{}
	cls.Ptr = ptr
	return cls
}

var xSqueezerPageGetChild func(uintptr) uintptr

// Returns the squeezer child to which @self belongs.
func (x *SqueezerPage) GetChild() *gtk.Widget {

	GetChildPtr := xSqueezerPageGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &gtk.Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

var xSqueezerPageGetEnabled func(uintptr) bool

// Gets whether @self is enabled.
func (x *SqueezerPage) GetEnabled() bool {

	return xSqueezerPageGetEnabled(x.GoPointer())

}

var xSqueezerPageSetEnabled func(uintptr, bool)

// Sets whether @self is enabled.
//
// If a child is disabled, it will be ignored when looking for the child
// fitting the available size best.
//
// This allows to programmatically and prematurely hide a child even if it fits
// in the available space.
//
// This can be used e.g. to ensure a certain child is hidden below a certain
// window width, or any other constraint you find suitable.
func (x *SqueezerPage) SetEnabled(EnabledVar bool) {

	xSqueezerPageSetEnabled(x.GoPointer(), EnabledVar)

}

func (c *SqueezerPage) GoPointer() uintptr {
	return c.Ptr
}

func (c *SqueezerPage) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("ADW"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSqueezer, lib, "adw_squeezer_new")

	core.PuregoSafeRegister(&xSqueezerAdd, lib, "adw_squeezer_add")
	core.PuregoSafeRegister(&xSqueezerGetAllowNone, lib, "adw_squeezer_get_allow_none")
	core.PuregoSafeRegister(&xSqueezerGetHomogeneous, lib, "adw_squeezer_get_homogeneous")
	core.PuregoSafeRegister(&xSqueezerGetInterpolateSize, lib, "adw_squeezer_get_interpolate_size")
	core.PuregoSafeRegister(&xSqueezerGetPage, lib, "adw_squeezer_get_page")
	core.PuregoSafeRegister(&xSqueezerGetPages, lib, "adw_squeezer_get_pages")
	core.PuregoSafeRegister(&xSqueezerGetSwitchThresholdPolicy, lib, "adw_squeezer_get_switch_threshold_policy")
	core.PuregoSafeRegister(&xSqueezerGetTransitionDuration, lib, "adw_squeezer_get_transition_duration")
	core.PuregoSafeRegister(&xSqueezerGetTransitionRunning, lib, "adw_squeezer_get_transition_running")
	core.PuregoSafeRegister(&xSqueezerGetTransitionType, lib, "adw_squeezer_get_transition_type")
	core.PuregoSafeRegister(&xSqueezerGetVisibleChild, lib, "adw_squeezer_get_visible_child")
	core.PuregoSafeRegister(&xSqueezerGetXalign, lib, "adw_squeezer_get_xalign")
	core.PuregoSafeRegister(&xSqueezerGetYalign, lib, "adw_squeezer_get_yalign")
	core.PuregoSafeRegister(&xSqueezerRemove, lib, "adw_squeezer_remove")
	core.PuregoSafeRegister(&xSqueezerSetAllowNone, lib, "adw_squeezer_set_allow_none")
	core.PuregoSafeRegister(&xSqueezerSetHomogeneous, lib, "adw_squeezer_set_homogeneous")
	core.PuregoSafeRegister(&xSqueezerSetInterpolateSize, lib, "adw_squeezer_set_interpolate_size")
	core.PuregoSafeRegister(&xSqueezerSetSwitchThresholdPolicy, lib, "adw_squeezer_set_switch_threshold_policy")
	core.PuregoSafeRegister(&xSqueezerSetTransitionDuration, lib, "adw_squeezer_set_transition_duration")
	core.PuregoSafeRegister(&xSqueezerSetTransitionType, lib, "adw_squeezer_set_transition_type")
	core.PuregoSafeRegister(&xSqueezerSetXalign, lib, "adw_squeezer_set_xalign")
	core.PuregoSafeRegister(&xSqueezerSetYalign, lib, "adw_squeezer_set_yalign")

	core.PuregoSafeRegister(&xSqueezerPageGetChild, lib, "adw_squeezer_page_get_child")
	core.PuregoSafeRegister(&xSqueezerPageGetEnabled, lib, "adw_squeezer_page_get_enabled")
	core.PuregoSafeRegister(&xSqueezerPageSetEnabled, lib, "adw_squeezer_page_set_enabled")

}
