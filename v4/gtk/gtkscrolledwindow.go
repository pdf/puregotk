// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Specifies which corner a child widget should be placed in when packed into
// a `GtkScrolledWindow.`
//
// This is effectively the opposite of where the scroll bars are placed.
type CornerType int

const (

	// Place the scrollbars on the right and bottom of the
	//   widget (default behaviour).
	CornerTopLeftValue CornerType = 0
	// Place the scrollbars on the top and right of the
	//   widget.
	CornerBottomLeftValue CornerType = 1
	// Place the scrollbars on the left and bottom of the
	//   widget.
	CornerTopRightValue CornerType = 2
	// Place the scrollbars on the top and left of the
	//   widget.
	CornerBottomRightValue CornerType = 3
)

// Determines how the size should be computed to achieve the one of the
// visibility mode for the scrollbars.
type PolicyType int

const (

	// The scrollbar is always visible. The view size is
	//   independent of the content.
	PolicyAlwaysValue PolicyType = 0
	// The scrollbar will appear and disappear as necessary.
	//   For example, when all of a `GtkTreeView` can not be seen.
	PolicyAutomaticValue PolicyType = 1
	// The scrollbar should never appear. In this mode the
	//   content determines the size.
	PolicyNeverValue PolicyType = 2
	// Don't show a scrollbar, but don't force the
	//   size to follow the content. This can be used e.g. to make multiple
	//   scrolled windows share a scrollbar.
	PolicyExternalValue PolicyType = 3
)

// `GtkScrolledWindow` is a container that makes its child scrollable.
//
// It does so using either internally added scrollbars or externally
// associated adjustments, and optionally draws a frame around the child.
//
// Widgets with native scrolling support, i.e. those whose classes implement
// the [iface@Gtk.Scrollable] interface, are added directly. For other types
// of widget, the class [class@Gtk.Viewport] acts as an adaptor, giving
// scrollability to other widgets. [method@Gtk.ScrolledWindow.set_child]
// intelligently accounts for whether or not the added child is a `GtkScrollable`.
// If it isn’t, then it wraps the child in a `GtkViewport`. Therefore, you can
// just add any child widget and not worry about the details.
//
// If [method@Gtk.ScrolledWindow.set_child] has added a `GtkViewport` for you,
// you can remove both your added child widget from the `GtkViewport`, and the
// `GtkViewport` from the `GtkScrolledWindow`, like this:
//
// ```c
// GtkWidget *scrolled_window = gtk_scrolled_window_new ();
// GtkWidget *child_widget = gtk_button_new ();
//
// // GtkButton is not a GtkScrollable, so GtkScrolledWindow will automatically
// // add a GtkViewport.
// gtk_box_append (GTK_BOX (scrolled_window), child_widget);
//
// // Either of these will result in child_widget being unparented:
// gtk_box_remove (GTK_BOX (scrolled_window), child_widget);
// // or
// gtk_box_remove (GTK_BOX (scrolled_window),
//
//	gtk_bin_get_child (GTK_BIN (scrolled_window)));
//
// ```
//
// Unless [property@Gtk.ScrolledWindow:hscrollbar-policy] and
// [property@Gtk.ScrolledWindow:vscrollbar-policy] are %GTK_POLICY_NEVER or
// %GTK_POLICY_EXTERNAL, `GtkScrolledWindow` adds internal `GtkScrollbar` widgets
// around its child. The scroll position of the child, and if applicable the
// scrollbars, is controlled by the [property@Gtk.ScrolledWindow:hadjustment]
// and [property@Gtk.ScrolledWindow:vadjustment] that are associated with the
// `GtkScrolledWindow`. See the docs on [class@Gtk.Scrollbar] for the details,
// but note that the “step_increment” and “page_increment” fields are only
// effective if the policy causes scrollbars to be present.
//
// If a `GtkScrolledWindow` doesn’t behave quite as you would like, or
// doesn’t have exactly the right layout, it’s very possible to set up
// your own scrolling with `GtkScrollbar` and for example a `GtkGrid`.
//
// # Touch support
//
// `GtkScrolledWindow` has built-in support for touch devices. When a
// touchscreen is used, swiping will move the scrolled window, and will
// expose 'kinetic' behavior. This can be turned off with the
// [property@Gtk.ScrolledWindow:kinetic-scrolling] property if it is undesired.
//
// `GtkScrolledWindow` also displays visual 'overshoot' indication when
// the content is pulled beyond the end, and this situation can be
// captured with the [signal@Gtk.ScrolledWindow::edge-overshot] signal.
//
// If no mouse device is present, the scrollbars will overlaid as
// narrow, auto-hiding indicators over the content. If traditional
// scrollbars are desired although no mouse is present, this behaviour
// can be turned off with the [property@Gtk.ScrolledWindow:overlay-scrolling]
// property.
//
// # CSS nodes
//
// `GtkScrolledWindow` has a main CSS node with name scrolledwindow.
// It gets a .frame style class added when [property@Gtk.ScrolledWindow:has-frame]
// is %TRUE.
//
// It uses subnodes with names overshoot and undershoot to draw the overflow
// and underflow indications. These nodes get the .left, .right, .top or .bottom
// style class added depending on where the indication is drawn.
//
// `GtkScrolledWindow` also sets the positional style classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering) on its scrollbars.
//
// If both scrollbars are visible, the area where they meet is drawn
// with a subnode named junction.
//
// # Accessibility
//
// `GtkScrolledWindow` uses the %GTK_ACCESSIBLE_ROLE_GROUP role.
type ScrolledWindow struct {
	Widget
}

func ScrolledWindowNewFromInternalPtr(ptr uintptr) *ScrolledWindow {
	cls := &ScrolledWindow{}
	cls.Ptr = ptr
	return cls
}

var xNewScrolledWindow func() uintptr

// Creates a new scrolled window.
func NewScrolledWindow() *Widget {
	var cls *Widget

	cret := xNewScrolledWindow()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowGetChild func(uintptr) uintptr

// Gets the child widget of @scrolled_window.
func (x *ScrolledWindow) GetChild() *Widget {
	var cls *Widget

	cret := xScrolledWindowGetChild(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowGetHadjustment func(uintptr) uintptr

// Returns the horizontal scrollbar’s adjustment.
//
// This is the adjustment used to connect the horizontal scrollbar
// to the child widget’s horizontal scroll functionality.
func (x *ScrolledWindow) GetHadjustment() *Adjustment {
	var cls *Adjustment

	cret := xScrolledWindowGetHadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowGetHasFrame func(uintptr) bool

// Gets whether the scrolled window draws a frame.
func (x *ScrolledWindow) GetHasFrame() bool {

	cret := xScrolledWindowGetHasFrame(x.GoPointer())
	return cret
}

var xScrolledWindowGetHscrollbar func(uintptr) uintptr

// Returns the horizontal scrollbar of @scrolled_window.
func (x *ScrolledWindow) GetHscrollbar() *Widget {
	var cls *Widget

	cret := xScrolledWindowGetHscrollbar(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowGetKineticScrolling func(uintptr) bool

// Returns the specified kinetic scrolling behavior.
func (x *ScrolledWindow) GetKineticScrolling() bool {

	cret := xScrolledWindowGetKineticScrolling(x.GoPointer())
	return cret
}

var xScrolledWindowGetMaxContentHeight func(uintptr) int

// Returns the maximum content height set.
func (x *ScrolledWindow) GetMaxContentHeight() int {

	cret := xScrolledWindowGetMaxContentHeight(x.GoPointer())
	return cret
}

var xScrolledWindowGetMaxContentWidth func(uintptr) int

// Returns the maximum content width set.
func (x *ScrolledWindow) GetMaxContentWidth() int {

	cret := xScrolledWindowGetMaxContentWidth(x.GoPointer())
	return cret
}

var xScrolledWindowGetMinContentHeight func(uintptr) int

// Gets the minimal content height of @scrolled_window.
func (x *ScrolledWindow) GetMinContentHeight() int {

	cret := xScrolledWindowGetMinContentHeight(x.GoPointer())
	return cret
}

var xScrolledWindowGetMinContentWidth func(uintptr) int

// Gets the minimum content width of @scrolled_window.
func (x *ScrolledWindow) GetMinContentWidth() int {

	cret := xScrolledWindowGetMinContentWidth(x.GoPointer())
	return cret
}

var xScrolledWindowGetOverlayScrolling func(uintptr) bool

// Returns whether overlay scrolling is enabled for this scrolled window.
func (x *ScrolledWindow) GetOverlayScrolling() bool {

	cret := xScrolledWindowGetOverlayScrolling(x.GoPointer())
	return cret
}

var xScrolledWindowGetPlacement func(uintptr) CornerType

// Gets the placement of the contents with respect to the scrollbars.
func (x *ScrolledWindow) GetPlacement() CornerType {

	cret := xScrolledWindowGetPlacement(x.GoPointer())
	return cret
}

var xScrolledWindowGetPolicy func(uintptr, *PolicyType, *PolicyType)

// Retrieves the current policy values for the horizontal and vertical
// scrollbars.
//
// See [method@Gtk.ScrolledWindow.set_policy].
func (x *ScrolledWindow) GetPolicy(HscrollbarPolicyVar *PolicyType, VscrollbarPolicyVar *PolicyType) {

	xScrolledWindowGetPolicy(x.GoPointer(), HscrollbarPolicyVar, VscrollbarPolicyVar)

}

var xScrolledWindowGetPropagateNaturalHeight func(uintptr) bool

// Reports whether the natural height of the child will be calculated
// and propagated through the scrolled window’s requested natural height.
func (x *ScrolledWindow) GetPropagateNaturalHeight() bool {

	cret := xScrolledWindowGetPropagateNaturalHeight(x.GoPointer())
	return cret
}

var xScrolledWindowGetPropagateNaturalWidth func(uintptr) bool

// Reports whether the natural width of the child will be calculated
// and propagated through the scrolled window’s requested natural width.
func (x *ScrolledWindow) GetPropagateNaturalWidth() bool {

	cret := xScrolledWindowGetPropagateNaturalWidth(x.GoPointer())
	return cret
}

var xScrolledWindowGetVadjustment func(uintptr) uintptr

// Returns the vertical scrollbar’s adjustment.
//
// This is the adjustment used to connect the vertical
// scrollbar to the child widget’s vertical scroll functionality.
func (x *ScrolledWindow) GetVadjustment() *Adjustment {
	var cls *Adjustment

	cret := xScrolledWindowGetVadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowGetVscrollbar func(uintptr) uintptr

// Returns the vertical scrollbar of @scrolled_window.
func (x *ScrolledWindow) GetVscrollbar() *Widget {
	var cls *Widget

	cret := xScrolledWindowGetVscrollbar(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xScrolledWindowSetChild func(uintptr, uintptr)

// Sets the child widget of @scrolled_window.
func (x *ScrolledWindow) SetChild(ChildVar *Widget) {

	xScrolledWindowSetChild(x.GoPointer(), ChildVar.GoPointer())

}

var xScrolledWindowSetHadjustment func(uintptr, uintptr)

// Sets the `GtkAdjustment` for the horizontal scrollbar.
func (x *ScrolledWindow) SetHadjustment(HadjustmentVar *Adjustment) {

	xScrolledWindowSetHadjustment(x.GoPointer(), HadjustmentVar.GoPointer())

}

var xScrolledWindowSetHasFrame func(uintptr, bool)

// Changes the frame drawn around the contents of @scrolled_window.
func (x *ScrolledWindow) SetHasFrame(HasFrameVar bool) {

	xScrolledWindowSetHasFrame(x.GoPointer(), HasFrameVar)

}

var xScrolledWindowSetKineticScrolling func(uintptr, bool)

// Turns kinetic scrolling on or off.
//
// Kinetic scrolling only applies to devices with source
// %GDK_SOURCE_TOUCHSCREEN.
func (x *ScrolledWindow) SetKineticScrolling(KineticScrollingVar bool) {

	xScrolledWindowSetKineticScrolling(x.GoPointer(), KineticScrollingVar)

}

var xScrolledWindowSetMaxContentHeight func(uintptr, int)

// Sets the maximum height that @scrolled_window should keep visible.
//
// The @scrolled_window will grow up to this height before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than [property@Gtk.ScrolledWindow:min-content-height].
func (x *ScrolledWindow) SetMaxContentHeight(HeightVar int) {

	xScrolledWindowSetMaxContentHeight(x.GoPointer(), HeightVar)

}

var xScrolledWindowSetMaxContentWidth func(uintptr, int)

// Sets the maximum width that @scrolled_window should keep visible.
//
// The @scrolled_window will grow up to this width before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content width to a
// value smaller than [property@Gtk.ScrolledWindow:min-content-width].
func (x *ScrolledWindow) SetMaxContentWidth(WidthVar int) {

	xScrolledWindowSetMaxContentWidth(x.GoPointer(), WidthVar)

}

var xScrolledWindowSetMinContentHeight func(uintptr, int)

// Sets the minimum height that @scrolled_window should keep visible.
//
// Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content height to a
// value greater than [property@Gtk.ScrolledWindow:max-content-height].
func (x *ScrolledWindow) SetMinContentHeight(HeightVar int) {

	xScrolledWindowSetMinContentHeight(x.GoPointer(), HeightVar)

}

var xScrolledWindowSetMinContentWidth func(uintptr, int)

// Sets the minimum width that @scrolled_window should keep visible.
//
// Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content width to a
// value greater than [property@Gtk.ScrolledWindow:max-content-width].
func (x *ScrolledWindow) SetMinContentWidth(WidthVar int) {

	xScrolledWindowSetMinContentWidth(x.GoPointer(), WidthVar)

}

var xScrolledWindowSetOverlayScrolling func(uintptr, bool)

// Enables or disables overlay scrolling for this scrolled window.
func (x *ScrolledWindow) SetOverlayScrolling(OverlayScrollingVar bool) {

	xScrolledWindowSetOverlayScrolling(x.GoPointer(), OverlayScrollingVar)

}

var xScrolledWindowSetPlacement func(uintptr, CornerType)

// Sets the placement of the contents with respect to the scrollbars
// for the scrolled window.
//
// The default is %GTK_CORNER_TOP_LEFT, meaning the child is
// in the top left, with the scrollbars underneath and to the right.
// Other values in [enum@Gtk.CornerType] are %GTK_CORNER_TOP_RIGHT,
// %GTK_CORNER_BOTTOM_LEFT, and %GTK_CORNER_BOTTOM_RIGHT.
//
// See also [method@Gtk.ScrolledWindow.get_placement] and
// [method@Gtk.ScrolledWindow.unset_placement].
func (x *ScrolledWindow) SetPlacement(WindowPlacementVar CornerType) {

	xScrolledWindowSetPlacement(x.GoPointer(), WindowPlacementVar)

}

var xScrolledWindowSetPolicy func(uintptr, PolicyType, PolicyType)

// Sets the scrollbar policy for the horizontal and vertical scrollbars.
//
// The policy determines when the scrollbar should appear; it is a value
// from the [enum@Gtk.PolicyType] enumeration. If %GTK_POLICY_ALWAYS, the
// scrollbar is always present; if %GTK_POLICY_NEVER, the scrollbar is
// never present; if %GTK_POLICY_AUTOMATIC, the scrollbar is present only
// if needed (that is, if the slider part of the bar would be smaller
// than the trough — the display is larger than the page size).
func (x *ScrolledWindow) SetPolicy(HscrollbarPolicyVar PolicyType, VscrollbarPolicyVar PolicyType) {

	xScrolledWindowSetPolicy(x.GoPointer(), HscrollbarPolicyVar, VscrollbarPolicyVar)

}

var xScrolledWindowSetPropagateNaturalHeight func(uintptr, bool)

// Sets whether the natural height of the child should be calculated
// and propagated through the scrolled window’s requested natural height.
func (x *ScrolledWindow) SetPropagateNaturalHeight(PropagateVar bool) {

	xScrolledWindowSetPropagateNaturalHeight(x.GoPointer(), PropagateVar)

}

var xScrolledWindowSetPropagateNaturalWidth func(uintptr, bool)

// Sets whether the natural width of the child should be calculated
// and propagated through the scrolled window’s requested natural width.
func (x *ScrolledWindow) SetPropagateNaturalWidth(PropagateVar bool) {

	xScrolledWindowSetPropagateNaturalWidth(x.GoPointer(), PropagateVar)

}

var xScrolledWindowSetVadjustment func(uintptr, uintptr)

// Sets the `GtkAdjustment` for the vertical scrollbar.
func (x *ScrolledWindow) SetVadjustment(VadjustmentVar *Adjustment) {

	xScrolledWindowSetVadjustment(x.GoPointer(), VadjustmentVar.GoPointer())

}

var xScrolledWindowUnsetPlacement func(uintptr)

// Unsets the placement of the contents with respect to the scrollbars.
//
// If no window placement is set for a scrolled window,
// it defaults to %GTK_CORNER_TOP_LEFT.
func (x *ScrolledWindow) UnsetPlacement() {

	xScrolledWindowUnsetPlacement(x.GoPointer())

}

func (c *ScrolledWindow) GoPointer() uintptr {
	return c.Ptr
}

func (c *ScrolledWindow) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted whenever user initiated scrolling makes the scrolled
// window firmly surpass the limits defined by the adjustment
// in that orientation.
//
// A similar behavior without edge resistance is provided by the
// [signal@Gtk.ScrolledWindow::edge-reached] signal.
//
// Note: The @pos argument is LTR/RTL aware, so callers should be
// aware too if intending to provide behavior on horizontal edges.
func (x *ScrolledWindow) ConnectEdgeOvershot(cb func(ScrolledWindow, PositionType)) uint32 {
	fcb := func(clsPtr uintptr, PosVarp PositionType) {
		fa := ScrolledWindow{}
		fa.Ptr = clsPtr

		cb(fa, PosVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "edge-overshot", purego.NewCallback(fcb))
}

// Emitted whenever user-initiated scrolling makes the scrolled
// window exactly reach the lower or upper limits defined by the
// adjustment in that orientation.
//
// A similar behavior with edge resistance is provided by the
// [signal@Gtk.ScrolledWindow::edge-overshot] signal.
//
// Note: The @pos argument is LTR/RTL aware, so callers should be
// aware too if intending to provide behavior on horizontal edges.
func (x *ScrolledWindow) ConnectEdgeReached(cb func(ScrolledWindow, PositionType)) uint32 {
	fcb := func(clsPtr uintptr, PosVarp PositionType) {
		fa := ScrolledWindow{}
		fa.Ptr = clsPtr

		cb(fa, PosVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "edge-reached", purego.NewCallback(fcb))
}

// Emitted when focus is moved away from the scrolled window by a
// keybinding.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The default bindings for this signal are
// `Ctrl + Tab` to move forward and `Ctrl + Shift + Tab` to
// move backward.
func (x *ScrolledWindow) ConnectMoveFocusOut(cb func(ScrolledWindow, DirectionType)) uint32 {
	fcb := func(clsPtr uintptr, DirectionTypeVarp DirectionType) {
		fa := ScrolledWindow{}
		fa.Ptr = clsPtr

		cb(fa, DirectionTypeVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "move-focus-out", purego.NewCallback(fcb))
}

// Emitted when a keybinding that scrolls is pressed.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// The horizontal or vertical adjustment is updated which triggers a
// signal that the scrolled window’s child may listen to and scroll itself.
func (x *ScrolledWindow) ConnectScrollChild(cb func(ScrolledWindow, ScrollType, bool) bool) uint32 {
	fcb := func(clsPtr uintptr, ScrollVarp ScrollType, HorizontalVarp bool) bool {
		fa := ScrolledWindow{}
		fa.Ptr = clsPtr

		return cb(fa, ScrollVarp, HorizontalVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "scroll-child", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ScrolledWindow) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ScrolledWindow) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ScrolledWindow) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ScrolledWindow) ResetState(StateVar AccessibleState) {

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
func (x *ScrolledWindow) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScrolledWindow) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ScrolledWindow) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScrolledWindow) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ScrolledWindow) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ScrolledWindow) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ScrolledWindow) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewScrolledWindow, lib, "gtk_scrolled_window_new")

	core.PuregoSafeRegister(&xScrolledWindowGetChild, lib, "gtk_scrolled_window_get_child")
	core.PuregoSafeRegister(&xScrolledWindowGetHadjustment, lib, "gtk_scrolled_window_get_hadjustment")
	core.PuregoSafeRegister(&xScrolledWindowGetHasFrame, lib, "gtk_scrolled_window_get_has_frame")
	core.PuregoSafeRegister(&xScrolledWindowGetHscrollbar, lib, "gtk_scrolled_window_get_hscrollbar")
	core.PuregoSafeRegister(&xScrolledWindowGetKineticScrolling, lib, "gtk_scrolled_window_get_kinetic_scrolling")
	core.PuregoSafeRegister(&xScrolledWindowGetMaxContentHeight, lib, "gtk_scrolled_window_get_max_content_height")
	core.PuregoSafeRegister(&xScrolledWindowGetMaxContentWidth, lib, "gtk_scrolled_window_get_max_content_width")
	core.PuregoSafeRegister(&xScrolledWindowGetMinContentHeight, lib, "gtk_scrolled_window_get_min_content_height")
	core.PuregoSafeRegister(&xScrolledWindowGetMinContentWidth, lib, "gtk_scrolled_window_get_min_content_width")
	core.PuregoSafeRegister(&xScrolledWindowGetOverlayScrolling, lib, "gtk_scrolled_window_get_overlay_scrolling")
	core.PuregoSafeRegister(&xScrolledWindowGetPlacement, lib, "gtk_scrolled_window_get_placement")
	core.PuregoSafeRegister(&xScrolledWindowGetPolicy, lib, "gtk_scrolled_window_get_policy")
	core.PuregoSafeRegister(&xScrolledWindowGetPropagateNaturalHeight, lib, "gtk_scrolled_window_get_propagate_natural_height")
	core.PuregoSafeRegister(&xScrolledWindowGetPropagateNaturalWidth, lib, "gtk_scrolled_window_get_propagate_natural_width")
	core.PuregoSafeRegister(&xScrolledWindowGetVadjustment, lib, "gtk_scrolled_window_get_vadjustment")
	core.PuregoSafeRegister(&xScrolledWindowGetVscrollbar, lib, "gtk_scrolled_window_get_vscrollbar")
	core.PuregoSafeRegister(&xScrolledWindowSetChild, lib, "gtk_scrolled_window_set_child")
	core.PuregoSafeRegister(&xScrolledWindowSetHadjustment, lib, "gtk_scrolled_window_set_hadjustment")
	core.PuregoSafeRegister(&xScrolledWindowSetHasFrame, lib, "gtk_scrolled_window_set_has_frame")
	core.PuregoSafeRegister(&xScrolledWindowSetKineticScrolling, lib, "gtk_scrolled_window_set_kinetic_scrolling")
	core.PuregoSafeRegister(&xScrolledWindowSetMaxContentHeight, lib, "gtk_scrolled_window_set_max_content_height")
	core.PuregoSafeRegister(&xScrolledWindowSetMaxContentWidth, lib, "gtk_scrolled_window_set_max_content_width")
	core.PuregoSafeRegister(&xScrolledWindowSetMinContentHeight, lib, "gtk_scrolled_window_set_min_content_height")
	core.PuregoSafeRegister(&xScrolledWindowSetMinContentWidth, lib, "gtk_scrolled_window_set_min_content_width")
	core.PuregoSafeRegister(&xScrolledWindowSetOverlayScrolling, lib, "gtk_scrolled_window_set_overlay_scrolling")
	core.PuregoSafeRegister(&xScrolledWindowSetPlacement, lib, "gtk_scrolled_window_set_placement")
	core.PuregoSafeRegister(&xScrolledWindowSetPolicy, lib, "gtk_scrolled_window_set_policy")
	core.PuregoSafeRegister(&xScrolledWindowSetPropagateNaturalHeight, lib, "gtk_scrolled_window_set_propagate_natural_height")
	core.PuregoSafeRegister(&xScrolledWindowSetPropagateNaturalWidth, lib, "gtk_scrolled_window_set_propagate_natural_width")
	core.PuregoSafeRegister(&xScrolledWindowSetVadjustment, lib, "gtk_scrolled_window_set_vadjustment")
	core.PuregoSafeRegister(&xScrolledWindowUnsetPlacement, lib, "gtk_scrolled_window_unset_placement")

}
