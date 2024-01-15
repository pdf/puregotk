// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A widget displaying a single row of a GtkTreeModel
//
// A `GtkCellView` displays a single row of a `GtkTreeModel` using a `GtkCellArea`
// and `GtkCellAreaContext`. A `GtkCellAreaContext` can be provided to the
// `GtkCellView` at construction time in order to keep the cellview in context
// of a group of cell views, this ensures that the renderers displayed will
// be properly aligned with each other (like the aligned cells in the menus
// of `GtkComboBox`).
//
// `GtkCellView` is `GtkOrientable` in order to decide in which orientation
// the underlying `GtkCellAreaContext` should be allocated. Taking the `GtkComboBox`
// menu as an example, cellviews should be oriented horizontally if the menus are
// listed top-to-bottom and thus all share the same width but may have separate
// individual heights (left-to-right menus should be allocated vertically since
// they all share the same height but may have variable widths).
//
// # CSS nodes
//
// GtkCellView has a single CSS node with name cellview.
type CellView struct {
	Widget
}

func CellViewNewFromInternalPtr(ptr uintptr) *CellView {
	cls := &CellView{}
	cls.Ptr = ptr
	return cls
}

var xNewCellView func() uintptr

// Creates a new `GtkCellView` widget.
func NewCellView() *CellView {
	var cls *CellView

	cret := xNewCellView()

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellView{}
	cls.Ptr = cret
	return cls
}

var xNewCellViewWithContext func(uintptr, uintptr) uintptr

// Creates a new `GtkCellView` widget with a specific `GtkCellArea`
// to layout cells and a specific `GtkCellAreaContext`.
//
// Specifying the same context for a handful of cells lets
// the underlying area synchronize the geometry for those cells,
// in this way alignments with cellviews for other rows are
// possible.
func NewCellViewWithContext(AreaVar *CellArea, ContextVar *CellAreaContext) *CellView {
	var cls *CellView

	cret := xNewCellViewWithContext(AreaVar.GoPointer(), ContextVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellView{}
	cls.Ptr = cret
	return cls
}

var xNewCellViewWithMarkup func(string) uintptr

// Creates a new `GtkCellView` widget, adds a `GtkCellRendererText`
// to it, and makes it show @markup. The text can be marked up with
// the [Pango text markup language](https://docs.gtk.org/Pango/pango_markup.html).
func NewCellViewWithMarkup(MarkupVar string) *CellView {
	var cls *CellView

	cret := xNewCellViewWithMarkup(MarkupVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellView{}
	cls.Ptr = cret
	return cls
}

var xNewCellViewWithText func(string) uintptr

// Creates a new `GtkCellView` widget, adds a `GtkCellRendererText`
// to it, and makes it show @text.
func NewCellViewWithText(TextVar string) *CellView {
	var cls *CellView

	cret := xNewCellViewWithText(TextVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellView{}
	cls.Ptr = cret
	return cls
}

var xNewCellViewWithTexture func(uintptr) uintptr

// Creates a new `GtkCellView` widget, adds a `GtkCellRendererPixbuf`
// to it, and makes it show @texture.
func NewCellViewWithTexture(TextureVar *gdk.Texture) *CellView {
	var cls *CellView

	cret := xNewCellViewWithTexture(TextureVar.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellView{}
	cls.Ptr = cret
	return cls
}

var xCellViewGetDisplayedRow func(uintptr) *TreePath

// Returns a `GtkTreePath` referring to the currently
// displayed row. If no row is currently displayed,
// %NULL is returned.
func (x *CellView) GetDisplayedRow() *TreePath {

	cret := xCellViewGetDisplayedRow(x.GoPointer())
	return cret
}

var xCellViewGetDrawSensitive func(uintptr) bool

// Gets whether @cell_view is configured to draw all of its
// cells in a sensitive state.
func (x *CellView) GetDrawSensitive() bool {

	cret := xCellViewGetDrawSensitive(x.GoPointer())
	return cret
}

var xCellViewGetFitModel func(uintptr) bool

// Gets whether @cell_view is configured to request space
// to fit the entire `GtkTreeModel`.
func (x *CellView) GetFitModel() bool {

	cret := xCellViewGetFitModel(x.GoPointer())
	return cret
}

var xCellViewGetModel func(uintptr) uintptr

// Returns the model for @cell_view. If no model is used %NULL is
// returned.
func (x *CellView) GetModel() *TreeModelBase {
	var cls *TreeModelBase

	cret := xCellViewGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

var xCellViewSetDisplayedRow func(uintptr, *TreePath)

// Sets the row of the model that is currently displayed
// by the `GtkCellView`. If the path is unset, then the
// contents of the cellview “stick” at their last value;
// this is not normally a desired result, but may be
// a needed intermediate state if say, the model for
// the `GtkCellView` becomes temporarily empty.
func (x *CellView) SetDisplayedRow(PathVar *TreePath) {

	xCellViewSetDisplayedRow(x.GoPointer(), PathVar)

}

var xCellViewSetDrawSensitive func(uintptr, bool)

// Sets whether @cell_view should draw all of its
// cells in a sensitive state, this is used by `GtkComboBox` menus
// to ensure that rows with insensitive cells that contain
// children appear sensitive in the parent menu item.
func (x *CellView) SetDrawSensitive(DrawSensitiveVar bool) {

	xCellViewSetDrawSensitive(x.GoPointer(), DrawSensitiveVar)

}

var xCellViewSetFitModel func(uintptr, bool)

// Sets whether @cell_view should request space to fit the entire `GtkTreeModel`.
//
// This is used by `GtkComboBox` to ensure that the cell view displayed on
// the combo box’s button always gets enough space and does not resize
// when selection changes.
func (x *CellView) SetFitModel(FitModelVar bool) {

	xCellViewSetFitModel(x.GoPointer(), FitModelVar)

}

var xCellViewSetModel func(uintptr, uintptr)

// Sets the model for @cell_view.  If @cell_view already has a model
// set, it will remove it before setting the new model.  If @model is
// %NULL, then it will unset the old model.
func (x *CellView) SetModel(ModelVar TreeModel) {

	xCellViewSetModel(x.GoPointer(), ModelVar.GoPointer())

}

func (c *CellView) GoPointer() uintptr {
	return c.Ptr
}

func (c *CellView) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *CellView) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *CellView) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *CellView) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *CellView) ResetState(StateVar AccessibleState) {

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
func (x *CellView) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CellView) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *CellView) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CellView) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *CellView) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *CellView) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *CellView) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

// Adds an attribute mapping to the list in @cell_layout.
//
// The @column is the column of the model to get a value from, and the
// @attribute is the property on @cell to be set from that value. So for
// example if column 2 of the model contains strings, you could have the
// “text” attribute of a `GtkCellRendererText` get its values from column 2.
// In this context "attribute" and "property" are used interchangeably.
func (x *CellView) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *CellView) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *CellView) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *CellView) GetArea() *CellArea {
	var cls *CellArea

	cret := XGtkCellLayoutGetArea(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &CellArea{}
	cls.Ptr = cret
	return cls
}

// Returns the cell renderers which have been added to @cell_layout.
func (x *CellView) GetCells() *glib.List {

	cret := XGtkCellLayoutGetCells(x.GoPointer())
	return cret
}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *CellView) PackEnd(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackEnd(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *CellView) PackStart(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackStart(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *CellView) Reorder(CellVar *CellRenderer, PositionVar int) {

	XGtkCellLayoutReorder(x.GoPointer(), CellVar.GoPointer(), PositionVar)

}

// Sets the attributes in the parameter list as the attributes
// of @cell_layout.
//
// See [method@Gtk.CellLayout.add_attribute] for more details.
//
// The attributes should be in attribute/column order, as in
// gtk_cell_layout_add_attribute(). All existing attributes are
// removed, and replaced with the new attributes.
func (x *CellView) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *CellView) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

// Retrieves the orientation of the @orientable.
func (x *CellView) GetOrientation() Orientation {

	cret := XGtkOrientableGetOrientation(x.GoPointer())
	return cret
}

// Sets the orientation of the @orientable.
func (x *CellView) SetOrientation(OrientationVar Orientation) {

	XGtkOrientableSetOrientation(x.GoPointer(), OrientationVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewCellView, lib, "gtk_cell_view_new")
	core.PuregoSafeRegister(&xNewCellViewWithContext, lib, "gtk_cell_view_new_with_context")
	core.PuregoSafeRegister(&xNewCellViewWithMarkup, lib, "gtk_cell_view_new_with_markup")
	core.PuregoSafeRegister(&xNewCellViewWithText, lib, "gtk_cell_view_new_with_text")
	core.PuregoSafeRegister(&xNewCellViewWithTexture, lib, "gtk_cell_view_new_with_texture")

	core.PuregoSafeRegister(&xCellViewGetDisplayedRow, lib, "gtk_cell_view_get_displayed_row")
	core.PuregoSafeRegister(&xCellViewGetDrawSensitive, lib, "gtk_cell_view_get_draw_sensitive")
	core.PuregoSafeRegister(&xCellViewGetFitModel, lib, "gtk_cell_view_get_fit_model")
	core.PuregoSafeRegister(&xCellViewGetModel, lib, "gtk_cell_view_get_model")
	core.PuregoSafeRegister(&xCellViewSetDisplayedRow, lib, "gtk_cell_view_set_displayed_row")
	core.PuregoSafeRegister(&xCellViewSetDrawSensitive, lib, "gtk_cell_view_set_draw_sensitive")
	core.PuregoSafeRegister(&xCellViewSetFitModel, lib, "gtk_cell_view_set_fit_model")
	core.PuregoSafeRegister(&xCellViewSetModel, lib, "gtk_cell_view_set_model")

}
