// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type ColumnViewColumnClass struct {
}

// `GtkColumnViewColumn` represents the columns being added to `GtkColumnView`.
//
// The main ingredient for a `GtkColumnViewColumn` is the `GtkListItemFactory`
// that tells the columnview how to create cells for this column from items in
// the model.
//
// Columns have a title, and can optionally have a header menu set
// with [method@Gtk.ColumnViewColumn.set_header_menu].
//
// A sorter can be associated with a column using
// [method@Gtk.ColumnViewColumn.set_sorter], to let users influence sorting
// by clicking on the column header.
type ColumnViewColumn struct {
	gobject.Object
}

func ColumnViewColumnNewFromInternalPtr(ptr uintptr) *ColumnViewColumn {
	cls := &ColumnViewColumn{}
	cls.Ptr = ptr
	return cls
}

var xNewColumnViewColumn func(string, uintptr) uintptr

// Creates a new `GtkColumnViewColumn` that uses the given @factory for
// mapping items to widgets.
//
// You most likely want to call [method@Gtk.ColumnView.append_column] next.
//
// The function takes ownership of the argument, so you can write code like:
//
// ```c
// column = gtk_column_view_column_new (_("Name"),
//
//	gtk_builder_list_item_factory_new_from_resource ("/name.ui"));
//
// ```
func NewColumnViewColumn(TitleVar string, FactoryVar *ListItemFactory) *ColumnViewColumn {
	NewColumnViewColumnPtr := xNewColumnViewColumn(TitleVar, FactoryVar.GoPointer())
	if NewColumnViewColumnPtr == 0 {
		return nil
	}

	NewColumnViewColumnCls := &ColumnViewColumn{}
	NewColumnViewColumnCls.Ptr = NewColumnViewColumnPtr
	return NewColumnViewColumnCls
}

var xColumnViewColumnGetColumnView func(uintptr) uintptr

// Gets the column view that's currently displaying this column.
//
// If @self has not been added to a column view yet, %NULL is returned.
func (x *ColumnViewColumn) GetColumnView() *ColumnView {

	GetColumnViewPtr := xColumnViewColumnGetColumnView(x.GoPointer())
	if GetColumnViewPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetColumnViewPtr)

	GetColumnViewCls := &ColumnView{}
	GetColumnViewCls.Ptr = GetColumnViewPtr
	return GetColumnViewCls

}

var xColumnViewColumnGetExpand func(uintptr) bool

// Returns whether this column should expand.
func (x *ColumnViewColumn) GetExpand() bool {

	return xColumnViewColumnGetExpand(x.GoPointer())

}

var xColumnViewColumnGetFactory func(uintptr) uintptr

// Gets the factory that's currently used to populate list items for
// this column.
func (x *ColumnViewColumn) GetFactory() *ListItemFactory {

	GetFactoryPtr := xColumnViewColumnGetFactory(x.GoPointer())
	if GetFactoryPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFactoryPtr)

	GetFactoryCls := &ListItemFactory{}
	GetFactoryCls.Ptr = GetFactoryPtr
	return GetFactoryCls

}

var xColumnViewColumnGetFixedWidth func(uintptr) int

// Gets the fixed width of the column.
func (x *ColumnViewColumn) GetFixedWidth() int {

	return xColumnViewColumnGetFixedWidth(x.GoPointer())

}

var xColumnViewColumnGetHeaderMenu func(uintptr) uintptr

// Gets the menu model that is used to create the context menu
// for the column header.
func (x *ColumnViewColumn) GetHeaderMenu() *gio.MenuModel {

	GetHeaderMenuPtr := xColumnViewColumnGetHeaderMenu(x.GoPointer())
	if GetHeaderMenuPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetHeaderMenuPtr)

	GetHeaderMenuCls := &gio.MenuModel{}
	GetHeaderMenuCls.Ptr = GetHeaderMenuPtr
	return GetHeaderMenuCls

}

var xColumnViewColumnGetResizable func(uintptr) bool

// Returns whether this column is resizable.
func (x *ColumnViewColumn) GetResizable() bool {

	return xColumnViewColumnGetResizable(x.GoPointer())

}

var xColumnViewColumnGetSorter func(uintptr) uintptr

// Returns the sorter that is associated with the column.
func (x *ColumnViewColumn) GetSorter() *Sorter {

	GetSorterPtr := xColumnViewColumnGetSorter(x.GoPointer())
	if GetSorterPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetSorterPtr)

	GetSorterCls := &Sorter{}
	GetSorterCls.Ptr = GetSorterPtr
	return GetSorterCls

}

var xColumnViewColumnGetTitle func(uintptr) string

// Returns the title set with gtk_column_view_column_set_title().
func (x *ColumnViewColumn) GetTitle() string {

	return xColumnViewColumnGetTitle(x.GoPointer())

}

var xColumnViewColumnGetVisible func(uintptr) bool

// Returns whether this column is visible.
func (x *ColumnViewColumn) GetVisible() bool {

	return xColumnViewColumnGetVisible(x.GoPointer())

}

var xColumnViewColumnSetExpand func(uintptr, bool)

// Sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that
// have the expand set to %TRUE.
func (x *ColumnViewColumn) SetExpand(ExpandVar bool) {

	xColumnViewColumnSetExpand(x.GoPointer(), ExpandVar)

}

var xColumnViewColumnSetFactory func(uintptr, uintptr)

// Sets the `GtkListItemFactory` to use for populating list items for this
// column.
func (x *ColumnViewColumn) SetFactory(FactoryVar *ListItemFactory) {

	xColumnViewColumnSetFactory(x.GoPointer(), FactoryVar.GoPointer())

}

var xColumnViewColumnSetFixedWidth func(uintptr, int)

// If @fixed_width is not -1, sets the fixed width of @column;
// otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated
// width. Interactive resizing also sets the “fixed-width” property.
func (x *ColumnViewColumn) SetFixedWidth(FixedWidthVar int) {

	xColumnViewColumnSetFixedWidth(x.GoPointer(), FixedWidthVar)

}

var xColumnViewColumnSetHeaderMenu func(uintptr, uintptr)

// Sets the menu model that is used to create the context menu
// for the column header.
func (x *ColumnViewColumn) SetHeaderMenu(MenuVar *gio.MenuModel) {

	xColumnViewColumnSetHeaderMenu(x.GoPointer(), MenuVar.GoPointer())

}

var xColumnViewColumnSetResizable func(uintptr, bool)

// Sets whether this column should be resizable by dragging.
func (x *ColumnViewColumn) SetResizable(ResizableVar bool) {

	xColumnViewColumnSetResizable(x.GoPointer(), ResizableVar)

}

var xColumnViewColumnSetSorter func(uintptr, uintptr)

// Associates a sorter with the column.
//
// If @sorter is %NULL, the column will not let users change
// the sorting by clicking on its header.
//
// This sorter can be made active by clicking on the column
// header, or by calling [method@Gtk.ColumnView.sort_by_column].
//
// See [method@Gtk.ColumnView.get_sorter] for the necessary steps
// for setting up customizable sorting for [class@Gtk.ColumnView].
func (x *ColumnViewColumn) SetSorter(SorterVar *Sorter) {

	xColumnViewColumnSetSorter(x.GoPointer(), SorterVar.GoPointer())

}

var xColumnViewColumnSetTitle func(uintptr, string)

// Sets the title of this column.
//
// The title is displayed in the header of a `GtkColumnView`
// for this column and is therefore user-facing text that should
// be translated.
func (x *ColumnViewColumn) SetTitle(TitleVar string) {

	xColumnViewColumnSetTitle(x.GoPointer(), TitleVar)

}

var xColumnViewColumnSetVisible func(uintptr, bool)

// Sets whether this column should be visible in views.
func (x *ColumnViewColumn) SetVisible(VisibleVar bool) {

	xColumnViewColumnSetVisible(x.GoPointer(), VisibleVar)

}

func (c *ColumnViewColumn) GoPointer() uintptr {
	return c.Ptr
}

func (c *ColumnViewColumn) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewColumnViewColumn, lib, "gtk_column_view_column_new")

	core.PuregoSafeRegister(&xColumnViewColumnGetColumnView, lib, "gtk_column_view_column_get_column_view")
	core.PuregoSafeRegister(&xColumnViewColumnGetExpand, lib, "gtk_column_view_column_get_expand")
	core.PuregoSafeRegister(&xColumnViewColumnGetFactory, lib, "gtk_column_view_column_get_factory")
	core.PuregoSafeRegister(&xColumnViewColumnGetFixedWidth, lib, "gtk_column_view_column_get_fixed_width")
	core.PuregoSafeRegister(&xColumnViewColumnGetHeaderMenu, lib, "gtk_column_view_column_get_header_menu")
	core.PuregoSafeRegister(&xColumnViewColumnGetResizable, lib, "gtk_column_view_column_get_resizable")
	core.PuregoSafeRegister(&xColumnViewColumnGetSorter, lib, "gtk_column_view_column_get_sorter")
	core.PuregoSafeRegister(&xColumnViewColumnGetTitle, lib, "gtk_column_view_column_get_title")
	core.PuregoSafeRegister(&xColumnViewColumnGetVisible, lib, "gtk_column_view_column_get_visible")
	core.PuregoSafeRegister(&xColumnViewColumnSetExpand, lib, "gtk_column_view_column_set_expand")
	core.PuregoSafeRegister(&xColumnViewColumnSetFactory, lib, "gtk_column_view_column_set_factory")
	core.PuregoSafeRegister(&xColumnViewColumnSetFixedWidth, lib, "gtk_column_view_column_set_fixed_width")
	core.PuregoSafeRegister(&xColumnViewColumnSetHeaderMenu, lib, "gtk_column_view_column_set_header_menu")
	core.PuregoSafeRegister(&xColumnViewColumnSetResizable, lib, "gtk_column_view_column_set_resizable")
	core.PuregoSafeRegister(&xColumnViewColumnSetSorter, lib, "gtk_column_view_column_set_sorter")
	core.PuregoSafeRegister(&xColumnViewColumnSetTitle, lib, "gtk_column_view_column_set_title")
	core.PuregoSafeRegister(&xColumnViewColumnSetVisible, lib, "gtk_column_view_column_set_visible")

}
