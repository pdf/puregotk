// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Determines whether the spin button displays values outside the adjustment
// bounds.
//
// See [method@Gtk.SpinButton.set_update_policy].
type SpinButtonUpdatePolicy int

const (

	// When refreshing your `GtkSpinButton`, the value is
	//   always displayed
	UpdateAlwaysValue SpinButtonUpdatePolicy = 0
	// When refreshing your `GtkSpinButton`, the value is
	//   only displayed if it is valid within the bounds of the spin button's
	//   adjustment
	UpdateIfValidValue SpinButtonUpdatePolicy = 1
)

// The values of the GtkSpinType enumeration are used to specify the
// change to make in gtk_spin_button_spin().
type SpinType int

const (

	// Increment by the adjustments step increment.
	SpinStepForwardValue SpinType = 0
	// Decrement by the adjustments step increment.
	SpinStepBackwardValue SpinType = 1
	// Increment by the adjustments page increment.
	SpinPageForwardValue SpinType = 2
	// Decrement by the adjustments page increment.
	SpinPageBackwardValue SpinType = 3
	// Go to the adjustments lower bound.
	SpinHomeValue SpinType = 4
	// Go to the adjustments upper bound.
	SpinEndValue SpinType = 5
	// Change by a specified amount.
	SpinUserDefinedValue SpinType = 6
)

// A `GtkSpinButton` is an ideal way to allow the user to set the
// value of some attribute.
//
// ![An example GtkSpinButton](spinbutton.png)
//
// Rather than having to directly type a number into a `GtkEntry`,
// `GtkSpinButton` allows the user to click on one of two arrows
// to increment or decrement the displayed value. A value can still be
// typed in, with the bonus that it can be checked to ensure it is in a
// given range.
//
// The main properties of a `GtkSpinButton` are through an adjustment.
// See the [class@Gtk.Adjustment] documentation for more details about
// an adjustment's properties.
//
// Note that `GtkSpinButton` will by default make its entry large enough
// to accommodate the lower and upper bounds of the adjustment. If this
// is not desired, the automatic sizing can be turned off by explicitly
// setting [property@Gtk.Editable:width-chars] to a value != -1.
//
// ## Using a GtkSpinButton to get an integer
//
// ```c
// // Provides a function to retrieve an integer value from a GtkSpinButton
// // and creates a spin button to model percentage values.
//
// int
// grab_int_value (GtkSpinButton *button,
//
//	gpointer       user_data)
//
//	{
//	  return gtk_spin_button_get_value_as_int (button);
//	}
//
// void
// create_integer_spin_button (void)
// {
//
//	  GtkWidget *window, *button;
//	  GtkAdjustment *adjustment;
//
//	  adjustment = gtk_adjustment_new (50.0, 0.0, 100.0, 1.0, 5.0, 0.0);
//
//	  window = gtk_window_new ();
//
//	  // creates the spinbutton, with no decimal places
//	  button = gtk_spin_button_new (adjustment, 1.0, 0);
//	  gtk_window_set_child (GTK_WINDOW (window), button);
//
//	  gtk_widget_show (window);
//	}
//
// ```
//
// ## Using a GtkSpinButton to get a floating point value
//
// ```c
// // Provides a function to retrieve a floating point value from a
// // GtkSpinButton, and creates a high precision spin button.
//
// float
// grab_float_value (GtkSpinButton *button,
//
//	gpointer       user_data)
//
//	{
//	  return gtk_spin_button_get_value (button);
//	}
//
// void
// create_floating_spin_button (void)
//
//	{
//	  GtkWidget *window, *button;
//	  GtkAdjustment *adjustment;
//
//	  adjustment = gtk_adjustment_new (2.500, 0.0, 5.0, 0.001, 0.1, 0.0);
//
//	  window = gtk_window_new ();
//
//	  // creates the spinbutton, with three decimal places
//	  button = gtk_spin_button_new (adjustment, 0.001, 3);
//	  gtk_window_set_child (GTK_WINDOW (window), button);
//
//	  gtk_widget_show (window);
//	}
//
// ```
//
// # CSS nodes
//
// ```
// spinbutton.horizontal
// ├── text
// │    ├── undershoot.left
// │    ╰── undershoot.right
// ├── button.down
// ╰── button.up
// ```
//
// ```
// spinbutton.vertical
// ├── button.up
// ├── text
// │    ├── undershoot.left
// │    ╰── undershoot.right
// ╰── button.down
// ```
//
// `GtkSpinButton`s main CSS node has the name spinbutton. It creates subnodes
// for the entry and the two buttons, with these names. The button nodes have
// the style classes .up and .down. The `GtkText` subnodes (if present) are put
// below the text node. The orientation of the spin button is reflected in
// the .vertical or .horizontal style class on the main node.
//
// # Accessiblity
//
// `GtkSpinButton` uses the %GTK_ACCESSIBLE_ROLE_SPIN_BUTTON role.
type SpinButton struct {
	Widget
}

func SpinButtonNewFromInternalPtr(ptr uintptr) *SpinButton {
	cls := &SpinButton{}
	cls.Ptr = ptr
	return cls
}

var xNewSpinButton func(uintptr, float64, uint) uintptr

// Creates a new `GtkSpinButton`.
func NewSpinButton(AdjustmentVar *Adjustment, ClimbRateVar float64, DigitsVar uint) *Widget {
	var cls *Widget

	cret := xNewSpinButton(AdjustmentVar.GoPointer(), ClimbRateVar, DigitsVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xNewWithRangeSpinButton func(float64, float64, float64) uintptr

// Creates a new `GtkSpinButton` with the given properties.
//
// This is a convenience constructor that allows creation
// of a numeric `GtkSpinButton` without manually creating
// an adjustment. The value is initially set to the minimum
// value and a page increment of 10 * @step is the default.
// The precision of the spin button is equivalent to the
// precision of @step.
//
// Note that the way in which the precision is derived works
// best if @step is a power of ten. If the resulting precision
// is not suitable for your needs, use
// [method@Gtk.SpinButton.set_digits] to correct it.
func NewWithRangeSpinButton(MinVar float64, MaxVar float64, StepVar float64) *Widget {
	var cls *Widget

	cret := xNewWithRangeSpinButton(MinVar, MaxVar, StepVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xSpinButtonConfigure func(uintptr, uintptr, float64, uint)

// Changes the properties of an existing spin button.
//
// The adjustment, climb rate, and number of decimal places
// are updated accordingly.
func (x *SpinButton) Configure(AdjustmentVar *Adjustment, ClimbRateVar float64, DigitsVar uint) {

	xSpinButtonConfigure(x.GoPointer(), AdjustmentVar.GoPointer(), ClimbRateVar, DigitsVar)

}

var xSpinButtonGetAdjustment func(uintptr) uintptr

// Get the adjustment associated with a `GtkSpinButton`.
func (x *SpinButton) GetAdjustment() *Adjustment {
	var cls *Adjustment

	cret := xSpinButtonGetAdjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

var xSpinButtonGetClimbRate func(uintptr) float64

// Returns the acceleration rate for repeated changes.
func (x *SpinButton) GetClimbRate() float64 {

	cret := xSpinButtonGetClimbRate(x.GoPointer())
	return cret
}

var xSpinButtonGetDigits func(uintptr) uint

// Fetches the precision of @spin_button.
func (x *SpinButton) GetDigits() uint {

	cret := xSpinButtonGetDigits(x.GoPointer())
	return cret
}

var xSpinButtonGetIncrements func(uintptr, float64, float64)

// Gets the current step and page the increments
// used by @spin_button.
//
// See [method@Gtk.SpinButton.set_increments].
func (x *SpinButton) GetIncrements(StepVar float64, PageVar float64) {

	xSpinButtonGetIncrements(x.GoPointer(), StepVar, PageVar)

}

var xSpinButtonGetNumeric func(uintptr) bool

// Returns whether non-numeric text can be typed into the spin button.
func (x *SpinButton) GetNumeric() bool {

	cret := xSpinButtonGetNumeric(x.GoPointer())
	return cret
}

var xSpinButtonGetRange func(uintptr, float64, float64)

// Gets the range allowed for @spin_button.
//
// See [method@Gtk.SpinButton.set_range].
func (x *SpinButton) GetRange(MinVar float64, MaxVar float64) {

	xSpinButtonGetRange(x.GoPointer(), MinVar, MaxVar)

}

var xSpinButtonGetSnapToTicks func(uintptr) bool

// Returns whether the values are corrected to the nearest step.
func (x *SpinButton) GetSnapToTicks() bool {

	cret := xSpinButtonGetSnapToTicks(x.GoPointer())
	return cret
}

var xSpinButtonGetUpdatePolicy func(uintptr) SpinButtonUpdatePolicy

// Gets the update behavior of a spin button.
//
// See [method@Gtk.SpinButton.set_update_policy].
func (x *SpinButton) GetUpdatePolicy() SpinButtonUpdatePolicy {

	cret := xSpinButtonGetUpdatePolicy(x.GoPointer())
	return cret
}

var xSpinButtonGetValue func(uintptr) float64

// Get the value in the @spin_button.
func (x *SpinButton) GetValue() float64 {

	cret := xSpinButtonGetValue(x.GoPointer())
	return cret
}

var xSpinButtonGetValueAsInt func(uintptr) int

// Get the value @spin_button represented as an integer.
func (x *SpinButton) GetValueAsInt() int {

	cret := xSpinButtonGetValueAsInt(x.GoPointer())
	return cret
}

var xSpinButtonGetWrap func(uintptr) bool

// Returns whether the spin button’s value wraps around to the
// opposite limit when the upper or lower limit of the range is
// exceeded.
func (x *SpinButton) GetWrap() bool {

	cret := xSpinButtonGetWrap(x.GoPointer())
	return cret
}

var xSpinButtonSetAdjustment func(uintptr, uintptr)

// Replaces the `GtkAdjustment` associated with @spin_button.
func (x *SpinButton) SetAdjustment(AdjustmentVar *Adjustment) {

	xSpinButtonSetAdjustment(x.GoPointer(), AdjustmentVar.GoPointer())

}

var xSpinButtonSetClimbRate func(uintptr, float64)

// Sets the acceleration rate for repeated changes when you
// hold down a button or key.
func (x *SpinButton) SetClimbRate(ClimbRateVar float64) {

	xSpinButtonSetClimbRate(x.GoPointer(), ClimbRateVar)

}

var xSpinButtonSetDigits func(uintptr, uint)

// Set the precision to be displayed by @spin_button.
//
// Up to 20 digit precision is allowed.
func (x *SpinButton) SetDigits(DigitsVar uint) {

	xSpinButtonSetDigits(x.GoPointer(), DigitsVar)

}

var xSpinButtonSetIncrements func(uintptr, float64, float64)

// Sets the step and page increments for spin_button.
//
// This affects how quickly the value changes when
// the spin button’s arrows are activated.
func (x *SpinButton) SetIncrements(StepVar float64, PageVar float64) {

	xSpinButtonSetIncrements(x.GoPointer(), StepVar, PageVar)

}

var xSpinButtonSetNumeric func(uintptr, bool)

// Sets the flag that determines if non-numeric text can be typed
// into the spin button.
func (x *SpinButton) SetNumeric(NumericVar bool) {

	xSpinButtonSetNumeric(x.GoPointer(), NumericVar)

}

var xSpinButtonSetRange func(uintptr, float64, float64)

// Sets the minimum and maximum allowable values for @spin_button.
//
// If the current value is outside this range, it will be adjusted
// to fit within the range, otherwise it will remain unchanged.
func (x *SpinButton) SetRange(MinVar float64, MaxVar float64) {

	xSpinButtonSetRange(x.GoPointer(), MinVar, MaxVar)

}

var xSpinButtonSetSnapToTicks func(uintptr, bool)

// Sets the policy as to whether values are corrected to the
// nearest step increment when a spin button is activated after
// providing an invalid value.
func (x *SpinButton) SetSnapToTicks(SnapToTicksVar bool) {

	xSpinButtonSetSnapToTicks(x.GoPointer(), SnapToTicksVar)

}

var xSpinButtonSetUpdatePolicy func(uintptr, SpinButtonUpdatePolicy)

// Sets the update behavior of a spin button.
//
// This determines whether the spin button is always
// updated or only when a valid value is set.
func (x *SpinButton) SetUpdatePolicy(PolicyVar SpinButtonUpdatePolicy) {

	xSpinButtonSetUpdatePolicy(x.GoPointer(), PolicyVar)

}

var xSpinButtonSetValue func(uintptr, float64)

// Sets the value of @spin_button.
func (x *SpinButton) SetValue(ValueVar float64) {

	xSpinButtonSetValue(x.GoPointer(), ValueVar)

}

var xSpinButtonSetWrap func(uintptr, bool)

// Sets the flag that determines if a spin button value wraps
// around to the opposite limit when the upper or lower limit
// of the range is exceeded.
func (x *SpinButton) SetWrap(WrapVar bool) {

	xSpinButtonSetWrap(x.GoPointer(), WrapVar)

}

var xSpinButtonSpin func(uintptr, SpinType, float64)

// Increment or decrement a spin button’s value in a specified
// direction by a specified amount.
func (x *SpinButton) Spin(DirectionVar SpinType, IncrementVar float64) {

	xSpinButtonSpin(x.GoPointer(), DirectionVar, IncrementVar)

}

var xSpinButtonUpdate func(uintptr)

// Manually force an update of the spin button.
func (x *SpinButton) Update() {

	xSpinButtonUpdate(x.GoPointer())

}

func (c *SpinButton) GoPointer() uintptr {
	return c.Ptr
}

func (c *SpinButton) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the user initiates a value change.
//
// This is a [keybinding signal](class.SignalAction.html).
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor
// programmatically.
//
// The default bindings for this signal are Up/Down and PageUp/PageDown.
func (x *SpinButton) ConnectChangeValue(cb func(SpinButton, ScrollType)) uint32 {
	fcb := func(clsPtr uintptr, ScrollVarp ScrollType) {
		fa := SpinButton{}
		fa.Ptr = clsPtr

		cb(fa, ScrollVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "change-value", purego.NewCallback(fcb))
}

// Emitted to convert the users input into a double value.
//
// The signal handler is expected to use [method@Gtk.Editable.get_text]
// to retrieve the text of the spinbutton and set @new_value to the
// new value.
//
// The default conversion uses g_strtod().
func (x *SpinButton) ConnectInput(cb func(SpinButton, float64) int) uint32 {
	fcb := func(clsPtr uintptr, NewValueVarp float64) int {
		fa := SpinButton{}
		fa.Ptr = clsPtr

		return cb(fa, NewValueVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "input", purego.NewCallback(fcb))
}

// Emitted to tweak the formatting of the value for display.
//
// ```c
// // show leading zeros
// static gboolean
// on_output (GtkSpinButton *spin,
//
//	gpointer       data)
//
//	{
//	   GtkAdjustment *adjustment;
//	   char *text;
//	   int value;
//
//	   adjustment = gtk_spin_button_get_adjustment (spin);
//	   value = (int)gtk_adjustment_get_value (adjustment);
//	   text = g_strdup_printf ("%02d", value);
//	   gtk_editable_set_text (GTK_EDITABLE (spin), text):
//	   g_free (text);
//
//	   return TRUE;
//	}
//
// ```
func (x *SpinButton) ConnectOutput(cb func(SpinButton) bool) uint32 {
	fcb := func(clsPtr uintptr) bool {
		fa := SpinButton{}
		fa.Ptr = clsPtr

		return cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "output", purego.NewCallback(fcb))
}

// Emitted when the value is changed.
//
// Also see the [signal@Gtk.SpinButton::output] signal.
func (x *SpinButton) ConnectValueChanged(cb func(SpinButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := SpinButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "value-changed", purego.NewCallback(fcb))
}

// Emitted right after the spinbutton wraps from its maximum
// to its minimum value or vice-versa.
func (x *SpinButton) ConnectWrapped(cb func(SpinButton)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := SpinButton{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "wrapped", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *SpinButton) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *SpinButton) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *SpinButton) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *SpinButton) ResetState(StateVar AccessibleState) {

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
func (x *SpinButton) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SpinButton) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *SpinButton) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SpinButton) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *SpinButton) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *SpinButton) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *SpinButton) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Emits the `GtkCellEditable::editing-done` signal.
func (x *SpinButton) EditingDone() {

	XGtkCellEditableEditingDone(x.GoPointer())

}

// Emits the `GtkCellEditable::remove-widget` signal.
func (x *SpinButton) RemoveWidget() {

	XGtkCellEditableRemoveWidget(x.GoPointer())

}

// Begins editing on a @cell_editable.
//
// The `GtkCellRenderer` for the cell creates and returns a `GtkCellEditable` from
// gtk_cell_renderer_start_editing(), configured for the `GtkCellRenderer` type.
//
// gtk_cell_editable_start_editing() can then set up @cell_editable suitably for
// editing a cell, e.g. making the Esc key emit `GtkCellEditable::editing-done`.
//
// Note that the @cell_editable is created on-demand for the current edit; its
// lifetime is temporary and does not persist across other edits and/or cells.
func (x *SpinButton) StartEditing(EventVar *gdk.Event) {

	XGtkCellEditableStartEditing(x.GoPointer(), EventVar.GoPointer())

}

// Deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (x *SpinButton) DeleteSelection() {

	XGtkEditableDeleteSelection(x.GoPointer())

}

// Deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters deleted are those from @start_pos to
// the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (x *SpinButton) DeleteText(StartPosVar int, EndPosVar int) {

	XGtkEditableDeleteText(x.GoPointer(), StartPosVar, EndPosVar)

}

// Undoes the setup done by [method@Gtk.Editable.init_delegate].
//
// This is a helper function that should be called from dispose,
// before removing the delegate object.
func (x *SpinButton) FinishDelegate() {

	XGtkEditableFinishDelegate(x.GoPointer())

}

// Gets the alignment of the editable.
func (x *SpinButton) GetAlignment() float32 {

	cret := XGtkEditableGetAlignment(x.GoPointer())
	return cret
}

// Retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is negative,
// then the characters retrieved are those characters from @start_pos to
// the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *SpinButton) GetChars(StartPosVar int, EndPosVar int) string {

	cret := XGtkEditableGetChars(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Gets the `GtkEditable` that @editable is delegating its
// implementation to.
//
// Typically, the delegate is a [class@Gtk.Text] widget.
func (x *SpinButton) GetDelegate() *EditableBase {
	var cls *EditableBase

	cret := XGtkEditableGetDelegate(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &EditableBase{}
	cls.Ptr = cret
	return cls
}

// Retrieves whether @editable is editable.
func (x *SpinButton) GetEditable() bool {

	cret := XGtkEditableGetEditable(x.GoPointer())
	return cret
}

// Gets if undo/redo actions are enabled for @editable
func (x *SpinButton) GetEnableUndo() bool {

	cret := XGtkEditableGetEnableUndo(x.GoPointer())
	return cret
}

// Retrieves the desired maximum width of @editable, in characters.
func (x *SpinButton) GetMaxWidthChars() int {

	cret := XGtkEditableGetMaxWidthChars(x.GoPointer())
	return cret
}

// Retrieves the current position of the cursor relative
// to the start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (x *SpinButton) GetPosition() int {

	cret := XGtkEditableGetPosition(x.GoPointer())
	return cret
}

// Retrieves the selection bound of the editable.
//
// @start_pos will be filled with the start of the selection and
// @end_pos with end. If no text was selected both will be identical
// and %FALSE will be returned.
//
// Note that positions are specified in characters, not bytes.
func (x *SpinButton) GetSelectionBounds(StartPosVar int, EndPosVar int) bool {

	cret := XGtkEditableGetSelectionBounds(x.GoPointer(), StartPosVar, EndPosVar)
	return cret
}

// Retrieves the contents of @editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (x *SpinButton) GetText() string {

	cret := XGtkEditableGetText(x.GoPointer())
	return cret
}

// Gets the number of characters of space reserved
// for the contents of the editable.
func (x *SpinButton) GetWidthChars() int {

	cret := XGtkEditableGetWidthChars(x.GoPointer())
	return cret
}

// Sets up a delegate for `GtkEditable`.
//
// This is assuming that the get_delegate vfunc in the `GtkEditable`
// interface has been set up for the @editable's type.
//
// This is a helper function that should be called in instance init,
// after creating the delegate object.
func (x *SpinButton) InitDelegate() {

	XGtkEditableInitDelegate(x.GoPointer())

}

// Inserts @length bytes of @text into the contents of the
// widget, at position @position.
//
// Note that the position is in characters, not in bytes.
// The function updates @position to point after the newly
// inserted text.
func (x *SpinButton) InsertText(TextVar string, LengthVar int, PositionVar int) {

	XGtkEditableInsertText(x.GoPointer(), TextVar, LengthVar, PositionVar)

}

// Selects a region of text.
//
// The characters that are selected are those characters at positions
// from @start_pos up to, but not including @end_pos. If @end_pos is
// negative, then the characters selected are those characters from
// @start_pos to  the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (x *SpinButton) SelectRegion(StartPosVar int, EndPosVar int) {

	XGtkEditableSelectRegion(x.GoPointer(), StartPosVar, EndPosVar)

}

// Sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when
// the displayed text is shorter than the width of the editable.
func (x *SpinButton) SetAlignment(XalignVar float32) {

	XGtkEditableSetAlignment(x.GoPointer(), XalignVar)

}

// Determines if the user can edit the text in the editable widget.
func (x *SpinButton) SetEditable(IsEditableVar bool) {

	XGtkEditableSetEditable(x.GoPointer(), IsEditableVar)

}

// If enabled, changes to @editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not
// stored in secure memory. As such, undo is forcefully disabled
// when [property@Gtk.Text:visibility] is set to %FALSE.
func (x *SpinButton) SetEnableUndo(EnableUndoVar bool) {

	XGtkEditableSetEnableUndo(x.GoPointer(), EnableUndoVar)

}

// Sets the desired maximum width in characters of @editable.
func (x *SpinButton) SetMaxWidthChars(NCharsVar int) {

	XGtkEditableSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

// Sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than
// or equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character
// of the editable. Note that @position is in characters, not in bytes.
func (x *SpinButton) SetPosition(PositionVar int) {

	XGtkEditableSetPosition(x.GoPointer(), PositionVar)

}

// Sets the text in the editable to the given value.
//
// This is replacing the current contents.
func (x *SpinButton) SetText(TextVar string) {

	XGtkEditableSetText(x.GoPointer(), TextVar)

}

// Changes the size request of the editable to be about the
// right size for @n_chars characters.
//
// Note that it changes the size request, the size can still
// be affected by how you pack the widget into containers.
// If @n_chars is -1, the size reverts to the default size.
func (x *SpinButton) SetWidthChars(NCharsVar int) {

	XGtkEditableSetWidthChars(x.GoPointer(), NCharsVar)

}

// Retrieves the orientation of the @orientable.
func (x *SpinButton) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *SpinButton) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewSpinButton, lib, "gtk_spin_button_new")
	core.PuregoSafeRegister(&xNewWithRangeSpinButton, lib, "gtk_spin_button_new_with_range")

	core.PuregoSafeRegister(&xSpinButtonConfigure, lib, "gtk_spin_button_configure")
	core.PuregoSafeRegister(&xSpinButtonGetAdjustment, lib, "gtk_spin_button_get_adjustment")
	core.PuregoSafeRegister(&xSpinButtonGetClimbRate, lib, "gtk_spin_button_get_climb_rate")
	core.PuregoSafeRegister(&xSpinButtonGetDigits, lib, "gtk_spin_button_get_digits")
	core.PuregoSafeRegister(&xSpinButtonGetIncrements, lib, "gtk_spin_button_get_increments")
	core.PuregoSafeRegister(&xSpinButtonGetNumeric, lib, "gtk_spin_button_get_numeric")
	core.PuregoSafeRegister(&xSpinButtonGetRange, lib, "gtk_spin_button_get_range")
	core.PuregoSafeRegister(&xSpinButtonGetSnapToTicks, lib, "gtk_spin_button_get_snap_to_ticks")
	core.PuregoSafeRegister(&xSpinButtonGetUpdatePolicy, lib, "gtk_spin_button_get_update_policy")
	core.PuregoSafeRegister(&xSpinButtonGetValue, lib, "gtk_spin_button_get_value")
	core.PuregoSafeRegister(&xSpinButtonGetValueAsInt, lib, "gtk_spin_button_get_value_as_int")
	core.PuregoSafeRegister(&xSpinButtonGetWrap, lib, "gtk_spin_button_get_wrap")
	core.PuregoSafeRegister(&xSpinButtonSetAdjustment, lib, "gtk_spin_button_set_adjustment")
	core.PuregoSafeRegister(&xSpinButtonSetClimbRate, lib, "gtk_spin_button_set_climb_rate")
	core.PuregoSafeRegister(&xSpinButtonSetDigits, lib, "gtk_spin_button_set_digits")
	core.PuregoSafeRegister(&xSpinButtonSetIncrements, lib, "gtk_spin_button_set_increments")
	core.PuregoSafeRegister(&xSpinButtonSetNumeric, lib, "gtk_spin_button_set_numeric")
	core.PuregoSafeRegister(&xSpinButtonSetRange, lib, "gtk_spin_button_set_range")
	core.PuregoSafeRegister(&xSpinButtonSetSnapToTicks, lib, "gtk_spin_button_set_snap_to_ticks")
	core.PuregoSafeRegister(&xSpinButtonSetUpdatePolicy, lib, "gtk_spin_button_set_update_policy")
	core.PuregoSafeRegister(&xSpinButtonSetValue, lib, "gtk_spin_button_set_value")
	core.PuregoSafeRegister(&xSpinButtonSetWrap, lib, "gtk_spin_button_set_wrap")
	core.PuregoSafeRegister(&xSpinButtonSpin, lib, "gtk_spin_button_spin")
	core.PuregoSafeRegister(&xSpinButtonUpdate, lib, "gtk_spin_button_update")

}
