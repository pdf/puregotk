// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ListViewClass struct {
}

// `GtkListView` presents a large dynamic list of items.
//
// `GtkListView` uses its factory to generate one row widget for each visible
// item and shows them in a linear display, either vertically or horizontally.
//
// The [property@Gtk.ListView:show-separators] property offers a simple way to
// display separators between the rows.
//
// `GtkListView` allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// [property@Gtk.ListView:enable-rubberband].
//
// If you need multiple columns with headers, see [class@Gtk.ColumnView].
//
// To learn more about the list widget framework, see the
// [overview](section-list-widget.html).
//
// An example of using `GtkListView`:
// ```c
// static void
// setup_listitem_cb (GtkListItemFactory *factory,
//
//	GtkListItem        *list_item)
//
//	{
//	  GtkWidget *image;
//
//	  image = gtk_image_new ();
//	  gtk_image_set_icon_size (GTK_IMAGE (image), GTK_ICON_SIZE_LARGE);
//	  gtk_list_item_set_child (list_item, image);
//	}
//
// static void
// bind_listitem_cb (GtkListItemFactory *factory,
//
//	GtkListItem        *list_item)
//
//	{
//	  GtkWidget *image;
//	  GAppInfo *app_info;
//
//	  image = gtk_list_item_get_child (list_item);
//	  app_info = gtk_list_item_get_item (list_item);
//	  gtk_image_set_from_gicon (GTK_IMAGE (image), g_app_info_get_icon (app_info));
//	}
//
// static void
// activate_cb (GtkListView  *list,
//
//	guint         position,
//	gpointer      unused)
//
//	{
//	  GAppInfo *app_info;
//
//	  app_info = g_list_model_get_item (G_LIST_MODEL (gtk_list_view_get_model (list)), position);
//	  g_app_info_launch (app_info, NULL, NULL, NULL);
//	  g_object_unref (app_info);
//	}
//
// ...
//
//	model = create_application_list ();
//
//	factory = gtk_signal_list_item_factory_new ();
//	g_signal_connect (factory, "setup", G_CALLBACK (setup_listitem_cb), NULL);
//	g_signal_connect (factory, "bind", G_CALLBACK (bind_listitem_cb), NULL);
//
//	list = gtk_list_view_new (GTK_SELECTION_MODEL (gtk_single_selection_new (model)), factory);
//
//	g_signal_connect (list, "activate", G_CALLBACK (activate_cb), NULL);
//
//	gtk_scrolled_window_set_child (GTK_SCROLLED_WINDOW (sw), list);
//
// ```
//
// # CSS nodes
//
// ```
// listview[.separators][.rich-list][.navigation-sidebar][.data-table]
// ├── row[.activatable]
// │
// ├── row[.activatable]
// │
// ┊
// ╰── [rubberband]
// ```
//
// `GtkListView` uses a single CSS node named `listview`. It may carry the
// `.separators` style class, when [property@Gtk.ListView:show-separators]
// property is set. Each child widget uses a single CSS node named `row`.
// If the [property@Gtk.ListItem:activatable] property is set, the
// corresponding row will have the `.activatable` style class. For
// rubberband selection, a node with name `rubberband` is used.
//
// The main listview node may also carry style classes to select
// the style of [list presentation](ListContainers.html#list-styles):
// .rich-list, .navigation-sidebar or .data-table.
//
// # Accessibility
//
// `GtkListView` uses the %GTK_ACCESSIBLE_ROLE_LIST role, and the list
// items use the %GTK_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListView struct {
	ListBase
}

func ListViewNewFromInternalPtr(ptr uintptr) *ListView {
	cls := &ListView{}
	cls.Ptr = ptr
	return cls
}

var xNewListView func(uintptr, uintptr) uintptr

// Creates a new `GtkListView` that uses the given @factory for
// mapping items to widgets.
//
// The function takes ownership of the
// arguments, so you can write code like
// ```c
// list_view = gtk_list_view_new (create_model (),
//
//	gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));
//
// ```
func NewListView(ModelVar SelectionModel, FactoryVar *ListItemFactory) *Widget {
	var cls *Widget

	cret := xNewListView(ModelVar.GoPointer(), FactoryVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xListViewGetEnableRubberband func(uintptr) bool

// Returns whether rows can be selected by dragging with the mouse.
func (x *ListView) GetEnableRubberband() bool {

	cret := xListViewGetEnableRubberband(x.GoPointer())
	return cret
}

var xListViewGetFactory func(uintptr) uintptr

// Gets the factory that's currently used to populate list items.
func (x *ListView) GetFactory() *ListItemFactory {
	var cls *ListItemFactory

	cret := xListViewGetFactory(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &ListItemFactory{}
	cls.Ptr = cret
	return cls
}

var xListViewGetModel func(uintptr) uintptr

// Gets the model that's currently used to read the items displayed.
func (x *ListView) GetModel() *SelectionModelBase {
	var cls *SelectionModelBase

	cret := xListViewGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &SelectionModelBase{}
	cls.Ptr = cret
	return cls
}

var xListViewGetShowSeparators func(uintptr) bool

// Returns whether the list box should show separators
// between rows.
func (x *ListView) GetShowSeparators() bool {

	cret := xListViewGetShowSeparators(x.GoPointer())
	return cret
}

var xListViewGetSingleClickActivate func(uintptr) bool

// Returns whether rows will be activated on single click and
// selected on hover.
func (x *ListView) GetSingleClickActivate() bool {

	cret := xListViewGetSingleClickActivate(x.GoPointer())
	return cret
}

var xListViewSetEnableRubberband func(uintptr, bool)

// Sets whether selections can be changed by dragging with the mouse.
func (x *ListView) SetEnableRubberband(EnableRubberbandVar bool) {

	xListViewSetEnableRubberband(x.GoPointer(), EnableRubberbandVar)

}

var xListViewSetFactory func(uintptr, uintptr)

// Sets the `GtkListItemFactory` to use for populating list items.
func (x *ListView) SetFactory(FactoryVar *ListItemFactory) {

	xListViewSetFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xListViewSetModel func(uintptr, uintptr)

// Sets the model to use.
//
// This must be a [iface@Gtk.SelectionModel] to use.
func (x *ListView) SetModel(ModelVar SelectionModel) {

	xListViewSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xListViewSetShowSeparators func(uintptr, bool)

// Sets whether the list box should show separators
// between rows.
func (x *ListView) SetShowSeparators(ShowSeparatorsVar bool) {

	xListViewSetShowSeparators(x.GoPointer(), ShowSeparatorsVar)

}

var xListViewSetSingleClickActivate func(uintptr, bool)

// Sets whether rows should be activated on single click and
// selected on hover.
func (x *ListView) SetSingleClickActivate(SingleClickActivateVar bool) {

	xListViewSetSingleClickActivate(x.GoPointer(), SingleClickActivateVar)

}

func (c *ListView) GoPointer() uintptr {
	return c.Ptr
}

func (c *ListView) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a row has been activated by the user,
// usually via activating the GtkListView|list.activate-item action.
//
// This allows for a convenient way to handle activation in a listview.
// See [method@Gtk.ListItem.set_activatable] for details on how to use
// this signal.
func (x *ListView) ConnectActivate(cb func(ListView, uint)) uint32 {
	fcb := func(clsPtr uintptr, PositionVarp uint) {
		fa := ListView{}
		fa.Ptr = clsPtr

		cb(fa, PositionVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *ListView) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *ListView) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *ListView) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *ListView) ResetState(StateVar AccessibleState) {

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
func (x *ListView) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ListView) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *ListView) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ListView) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *ListView) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *ListView) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *ListView) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Retrieves the orientation of the @orientable.
func (x *ListView) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *ListView) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

// Returns the size of a non-scrolling border around the
// outside of the scrollable.
//
// An example for this would be treeview headers. GTK can use
// this information to display overlaid graphics, like the
// overshoot indication, at the right position.
func (x *ListView) GetBorder(BorderVar *Border) bool {

	cret := XGtkScrollableGetBorder(x.GoPointer(), BorderVar)
	return cret
}

// Retrieves the `GtkAdjustment` used for horizontal scrolling.
func (x *ListView) GetHadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetHadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the horizontal `GtkScrollablePolicy`.
func (x *ListView) GetHscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetHscrollPolicy(x.GoPointer())
	return cret
}

// Retrieves the `GtkAdjustment` used for vertical scrolling.
func (x *ListView) GetVadjustment() *Adjustment {
	var cls *Adjustment

	cret := XGtkScrollableGetVadjustment(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Adjustment{}
	cls.Ptr = cret
	return cls
}

// Gets the vertical `GtkScrollablePolicy`.
func (x *ListView) GetVscrollPolicy() ScrollablePolicy {

	cret := XGtkScrollableGetVscrollPolicy(x.GoPointer())
	return cret
}

// Sets the horizontal adjustment of the `GtkScrollable`.
func (x *ListView) SetHadjustment(HadjustmentVar *Adjustment) {

	XGtkScrollableSetHadjustment(x.GoPointer(), HadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether horizontal scrolling should start
// below the minimum width or below the natural width.
func (x *ListView) SetHscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetHscrollPolicy(x.GoPointer(), PolicyVar)

}

// Sets the vertical adjustment of the `GtkScrollable`.
func (x *ListView) SetVadjustment(VadjustmentVar *Adjustment) {

	XGtkScrollableSetVadjustment(x.GoPointer(), VadjustmentVar.GoPointer())

}

// Sets the `GtkScrollablePolicy`.
//
// The policy determines whether vertical scrolling should start
// below the minimum height or below the natural height.
func (x *ListView) SetVscrollPolicy(PolicyVar ScrollablePolicy) {

	XGtkScrollableSetVscrollPolicy(x.GoPointer(), PolicyVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewListView, lib, "gtk_list_view_new")

	core.PuregoSafeRegister(&xListViewGetEnableRubberband, lib, "gtk_list_view_get_enable_rubberband")
	core.PuregoSafeRegister(&xListViewGetFactory, lib, "gtk_list_view_get_factory")
	core.PuregoSafeRegister(&xListViewGetModel, lib, "gtk_list_view_get_model")
	core.PuregoSafeRegister(&xListViewGetShowSeparators, lib, "gtk_list_view_get_show_separators")
	core.PuregoSafeRegister(&xListViewGetSingleClickActivate, lib, "gtk_list_view_get_single_click_activate")
	core.PuregoSafeRegister(&xListViewSetEnableRubberband, lib, "gtk_list_view_set_enable_rubberband")
	core.PuregoSafeRegister(&xListViewSetFactory, lib, "gtk_list_view_set_factory")
	core.PuregoSafeRegister(&xListViewSetModel, lib, "gtk_list_view_set_model")
	core.PuregoSafeRegister(&xListViewSetShowSeparators, lib, "gtk_list_view_set_show_separators")
	core.PuregoSafeRegister(&xListViewSetSingleClickActivate, lib, "gtk_list_view_set_single_click_activate")

}
