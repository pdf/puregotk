// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type DropDownClass struct {
	ParentClass uintptr
}

// `GtkDropDown` is a widget that allows the user to choose an item
// from a list of options.
//
// ![An example GtkDropDown](drop-down.png)
//
// The `GtkDropDown` displays the selected choice.
//
// The options are given to `GtkDropDown` in the form of `GListModel`
// and how the individual options are represented is determined by
// a [class@Gtk.ListItemFactory]. The default factory displays simple strings.
//
// `GtkDropDown` knows how to obtain strings from the items in a
// [class@Gtk.StringList]; for other models, you have to provide an expression
// to find the strings via [method@Gtk.DropDown.set_expression].
//
// `GtkDropDown` can optionally allow search in the popup, which is
// useful if the list of options is long. To enable the search entry,
// use [method@Gtk.DropDown.set_enable_search].
//
// Here is a UI definition example for `GtkDropDown` with a simple model:
// ```xml
// &lt;object class="GtkDropDown"&gt;
//
//	&lt;property name="model"&gt;
//	  &lt;object class="GtkStringList"&gt;
//	    &lt;items&gt;
//	      &lt;item translatable="yes"&gt;Factory&lt;/item&gt;
//	      &lt;item translatable="yes"&gt;Home&lt;/item&gt;
//	      &lt;item translatable="yes"&gt;Subway&lt;/item&gt;
//	    &lt;/items&gt;
//	  &lt;/object&gt;
//	&lt;/property&gt;
//
// &lt;/object&gt;
// ```
//
// # CSS nodes
//
// `GtkDropDown` has a single CSS node with name dropdown,
// with the button and popover nodes as children.
//
// # Accessibility
//
// `GtkDropDown` uses the %GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type DropDown struct {
	Widget
}

func DropDownNewFromInternalPtr(ptr uintptr) *DropDown {
	cls := &DropDown{}
	cls.Ptr = ptr
	return cls
}

var xNewDropDown func(uintptr, uintptr) uintptr

// Creates a new `GtkDropDown`.
//
// You may want to call [method@Gtk.DropDown.set_factory]
// to set up a way to map its items to widgets.
func NewDropDown(ModelVar gio.ListModel, ExpressionVar *Expression) *Widget {
	var cls *Widget

	cret := xNewDropDown(ModelVar.GoPointer(), ExpressionVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xNewFromStringsDropDown func(uintptr) uintptr

// Creates a new `GtkDropDown` that is populated with
// the strings.
func NewFromStringsDropDown(StringsVar uintptr) *Widget {
	var cls *Widget

	cret := xNewFromStringsDropDown(StringsVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetEnableSearch func(uintptr) bool

// Returns whether search is enabled.
func (x *DropDown) GetEnableSearch() bool {

	cret := xDropDownGetEnableSearch(x.GoPointer())
	return cret
}

var xDropDownGetExpression func(uintptr) uintptr

// Gets the expression set that is used to obtain strings from items.
//
// See [method@Gtk.DropDown.set_expression].
func (x *DropDown) GetExpression() *Expression {
	var cls *Expression

	cret := xDropDownGetExpression(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Expression{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetFactory func(uintptr) uintptr

// Gets the factory that's currently used to populate list items.
//
// The factory returned by this function is always used for the
// item in the button. It is also used for items in the popup
// if [property@Gtk.DropDown:list-factory] is not set.
func (x *DropDown) GetFactory() *ListItemFactory {
	var cls *ListItemFactory

	cret := xDropDownGetFactory(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetListFactory func(uintptr) uintptr

// Gets the factory that's currently used to populate list items in the popup.
func (x *DropDown) GetListFactory() *ListItemFactory {
	var cls *ListItemFactory

	cret := xDropDownGetListFactory(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetModel func(uintptr) uintptr

// Gets the model that provides the displayed items.
func (x *DropDown) GetModel() *gio.ListModelBase {
	var cls *gio.ListModelBase

	cret := xDropDownGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.ListModelBase{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetSelected func(uintptr) uint

// Gets the position of the selected item.
func (x *DropDown) GetSelected() uint {

	cret := xDropDownGetSelected(x.GoPointer())
	return cret
}

var xDropDownGetSelectedItem func(uintptr) uintptr

// Gets the selected item. If no item is selected, %NULL is returned.
func (x *DropDown) GetSelectedItem() *gobject.Object {
	var cls *gobject.Object

	cret := xDropDownGetSelectedItem(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gobject.Object{}
	cls.Ptr = cret
	return cls
}

var xDropDownGetShowArrow func(uintptr) bool

// Returns whether to show an arrow within the widget.
func (x *DropDown) GetShowArrow() bool {

	cret := xDropDownGetShowArrow(x.GoPointer())
	return cret
}

var xDropDownSetEnableSearch func(uintptr, bool)

// Sets whether a search entry will be shown in the popup that
// allows to search for items in the list.
//
// Note that [property@Gtk.DropDown:expression] must be set for
// search to work.
func (x *DropDown) SetEnableSearch(EnableSearchVar bool) {

	xDropDownSetEnableSearch(x.GoPointer(), EnableSearchVar)

}

var xDropDownSetExpression func(uintptr, uintptr)

// Sets the expression that gets evaluated to obtain strings from items.
//
// This is used for search in the popup. The expression must have
// a value type of %G_TYPE_STRING.
func (x *DropDown) SetExpression(ExpressionVar *Expression) {

	xDropDownSetExpression(x.GoPointer(), ExpressionVar.GoPointer())

}

var xDropDownSetFactory func(uintptr, uintptr)

// Sets the `GtkListItemFactory` to use for populating list items.
func (x *DropDown) SetFactory(FactoryVar *ListItemFactory) {

	xDropDownSetFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xDropDownSetListFactory func(uintptr, uintptr)

// Sets the `GtkListItemFactory` to use for populating list items in the popup.
func (x *DropDown) SetListFactory(FactoryVar *ListItemFactory) {

	xDropDownSetListFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xDropDownSetModel func(uintptr, uintptr)

// Sets the `GListModel` to use.
func (x *DropDown) SetModel(ModelVar gio.ListModel) {

	xDropDownSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xDropDownSetSelected func(uintptr, uint)

// Selects the item at the given position.
func (x *DropDown) SetSelected(PositionVar uint) {

	xDropDownSetSelected(x.GoPointer(), PositionVar)

}

var xDropDownSetShowArrow func(uintptr, bool)

// Sets whether an arrow will be displayed within the widget.
func (x *DropDown) SetShowArrow(ShowArrowVar bool) {

	xDropDownSetShowArrow(x.GoPointer(), ShowArrowVar)

}

func (c *DropDown) GoPointer() uintptr {
	return c.Ptr
}

func (c *DropDown) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted to when the drop down is activated.
//
// The `::activate` signal on `GtkDropDown` is an action signal and
// emitting it causes the drop down to pop up its dropdown.
func (x *DropDown) ConnectActivate(cb func(DropDown)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := DropDown{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *DropDown) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *DropDown) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *DropDown) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *DropDown) ResetState(StateVar AccessibleState) {

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
func (x *DropDown) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DropDown) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *DropDown) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DropDown) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *DropDown) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *DropDown) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *DropDown) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDropDown, lib, "gtk_drop_down_new")
	core.PuregoSafeRegister(&xNewFromStringsDropDown, lib, "gtk_drop_down_new_from_strings")

	core.PuregoSafeRegister(&xDropDownGetEnableSearch, lib, "gtk_drop_down_get_enable_search")
	core.PuregoSafeRegister(&xDropDownGetExpression, lib, "gtk_drop_down_get_expression")
	core.PuregoSafeRegister(&xDropDownGetFactory, lib, "gtk_drop_down_get_factory")
	core.PuregoSafeRegister(&xDropDownGetListFactory, lib, "gtk_drop_down_get_list_factory")
	core.PuregoSafeRegister(&xDropDownGetModel, lib, "gtk_drop_down_get_model")
	core.PuregoSafeRegister(&xDropDownGetSelected, lib, "gtk_drop_down_get_selected")
	core.PuregoSafeRegister(&xDropDownGetSelectedItem, lib, "gtk_drop_down_get_selected_item")
	core.PuregoSafeRegister(&xDropDownGetShowArrow, lib, "gtk_drop_down_get_show_arrow")
	core.PuregoSafeRegister(&xDropDownSetEnableSearch, lib, "gtk_drop_down_set_enable_search")
	core.PuregoSafeRegister(&xDropDownSetExpression, lib, "gtk_drop_down_set_expression")
	core.PuregoSafeRegister(&xDropDownSetFactory, lib, "gtk_drop_down_set_factory")
	core.PuregoSafeRegister(&xDropDownSetListFactory, lib, "gtk_drop_down_set_list_factory")
	core.PuregoSafeRegister(&xDropDownSetModel, lib, "gtk_drop_down_set_model")
	core.PuregoSafeRegister(&xDropDownSetSelected, lib, "gtk_drop_down_set_selected")
	core.PuregoSafeRegister(&xDropDownSetShowArrow, lib, "gtk_drop_down_set_show_arrow")

}
