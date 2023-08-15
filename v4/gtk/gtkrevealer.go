// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// These enumeration values describe the possible transitions
// when the child of a `GtkRevealer` widget is shown or hidden.
type RevealerTransitionType int

const (

	// No transition
	RevealerTransitionTypeNoneValue RevealerTransitionType = 0
	// Fade in
	RevealerTransitionTypeCrossfadeValue RevealerTransitionType = 1
	// Slide in from the left
	RevealerTransitionTypeSlideRightValue RevealerTransitionType = 2
	// Slide in from the right
	RevealerTransitionTypeSlideLeftValue RevealerTransitionType = 3
	// Slide in from the bottom
	RevealerTransitionTypeSlideUpValue RevealerTransitionType = 4
	// Slide in from the top
	RevealerTransitionTypeSlideDownValue RevealerTransitionType = 5
	// Floop in from the left
	RevealerTransitionTypeSwingRightValue RevealerTransitionType = 6
	// Floop in from the right
	RevealerTransitionTypeSwingLeftValue RevealerTransitionType = 7
	// Floop in from the bottom
	RevealerTransitionTypeSwingUpValue RevealerTransitionType = 8
	// Floop in from the top
	RevealerTransitionTypeSwingDownValue RevealerTransitionType = 9
)

// A `GtkRevealer` animates the transition of its child from invisible to visible.
//
// The style of transition can be controlled with
// [method@Gtk.Revealer.set_transition_type].
//
// These animations respect the [property@Gtk.Settings:gtk-enable-animations]
// setting.
//
// # CSS nodes
//
// `GtkRevealer` has a single CSS node with name revealer.
// When styling `GtkRevealer` using CSS, remember that it only hides its contents,
// not itself. That means applied margin, padding and borders will be visible even
// when the [property@Gtk.Revealer:reveal-child] property is set to %FALSE.
//
// # Accessibility
//
// `GtkRevealer` uses the %GTK_ACCESSIBLE_ROLE_GROUP role.
//
// The child of `GtkRevealer`, if set, is always available in the accessibility
// tree, regardless of the state of the revealer widget.
type Revealer struct {
	Widget
}

func RevealerNewFromInternalPtr(ptr uintptr) *Revealer {
	cls := &Revealer{}
	cls.Ptr = ptr
	return cls
}

var xNewRevealer func() uintptr

// Creates a new `GtkRevealer`.
func NewRevealer() *Widget {
	NewRevealerPtr := xNewRevealer()
	if NewRevealerPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewRevealerPtr)

	NewRevealerCls := &Widget{}
	NewRevealerCls.Ptr = NewRevealerPtr
	return NewRevealerCls
}

var xRevealerGetChild func(uintptr) uintptr

// Gets the child widget of @revealer.
func (x *Revealer) GetChild() *Widget {

	GetChildPtr := xRevealerGetChild(x.GoPointer())
	if GetChildPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetChildPtr)

	GetChildCls := &Widget{}
	GetChildCls.Ptr = GetChildPtr
	return GetChildCls

}

var xRevealerGetChildRevealed func(uintptr) bool

// Returns whether the child is fully revealed.
//
// In other words, this returns whether the transition
// to the revealed state is completed.
func (x *Revealer) GetChildRevealed() bool {

	return xRevealerGetChildRevealed(x.GoPointer())

}

var xRevealerGetRevealChild func(uintptr) bool

// Returns whether the child is currently revealed.
//
// This function returns %TRUE as soon as the transition
// is to the revealed state is started. To learn whether
// the child is fully revealed (ie the transition is completed),
// use [method@Gtk.Revealer.get_child_revealed].
func (x *Revealer) GetRevealChild() bool {

	return xRevealerGetRevealChild(x.GoPointer())

}

var xRevealerGetTransitionDuration func(uintptr) uint

// Returns the amount of time (in milliseconds) that
// transitions will take.
func (x *Revealer) GetTransitionDuration() uint {

	return xRevealerGetTransitionDuration(x.GoPointer())

}

var xRevealerGetTransitionType func(uintptr) RevealerTransitionType

// Gets the type of animation that will be used
// for transitions in @revealer.
func (x *Revealer) GetTransitionType() RevealerTransitionType {

	return xRevealerGetTransitionType(x.GoPointer())

}

var xRevealerSetChild func(uintptr, uintptr)

// Sets the child widget of @revealer.
func (x *Revealer) SetChild(ChildVar *Widget) {

	xRevealerSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xRevealerSetRevealChild func(uintptr, bool)

// Tells the `GtkRevealer` to reveal or conceal its child.
//
// The transition will be animated with the current
// transition type of @revealer.
func (x *Revealer) SetRevealChild(RevealChildVar bool) {

	xRevealerSetRevealChild(x.GoPointer(), RevealChildVar)

}

var xRevealerSetTransitionDuration func(uintptr, uint)

// Sets the duration that transitions will take.
func (x *Revealer) SetTransitionDuration(DurationVar uint) {

	xRevealerSetTransitionDuration(x.GoPointer(), DurationVar)

}

var xRevealerSetTransitionType func(uintptr, RevealerTransitionType)

// Sets the type of animation that will be used for
// transitions in @revealer.
//
// Available types include various kinds of fades and slides.
func (x *Revealer) SetTransitionType(TransitionVar RevealerTransitionType) {

	xRevealerSetTransitionType(x.GoPointer(), TransitionVar)

}

func (c *Revealer) GoPointer() uintptr {
	return c.Ptr
}

func (c *Revealer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Revealer) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *Revealer) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Revealer) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Revealer) ResetState(StateVar AccessibleState) {

	XGtkAccessibleResetState(x.GoPointer(), StateVar)

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
func (x *Revealer) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Revealer) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdatePropertyValue(x.GoPointer(), NPropertiesVar, PropertiesVar, ValuesVar)

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
func (x *Revealer) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Revealer) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateRelationValue(x.GoPointer(), NRelationsVar, RelationsVar, ValuesVar)

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
func (x *Revealer) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Revealer) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Revealer) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewRevealer, lib, "gtk_revealer_new")

	core.PuregoSafeRegister(&xRevealerGetChild, lib, "gtk_revealer_get_child")
	core.PuregoSafeRegister(&xRevealerGetChildRevealed, lib, "gtk_revealer_get_child_revealed")
	core.PuregoSafeRegister(&xRevealerGetRevealChild, lib, "gtk_revealer_get_reveal_child")
	core.PuregoSafeRegister(&xRevealerGetTransitionDuration, lib, "gtk_revealer_get_transition_duration")
	core.PuregoSafeRegister(&xRevealerGetTransitionType, lib, "gtk_revealer_get_transition_type")
	core.PuregoSafeRegister(&xRevealerSetChild, lib, "gtk_revealer_set_child")
	core.PuregoSafeRegister(&xRevealerSetRevealChild, lib, "gtk_revealer_set_reveal_child")
	core.PuregoSafeRegister(&xRevealerSetTransitionDuration, lib, "gtk_revealer_set_transition_duration")
	core.PuregoSafeRegister(&xRevealerSetTransitionType, lib, "gtk_revealer_set_transition_type")

}
