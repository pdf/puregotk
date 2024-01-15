// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// A function which decides whether the row indicated by @iter matches
// a given @key, and should be displayed as a possible completion for @key.
//
// Note that @key is normalized and case-folded (see g_utf8_normalize()
// and g_utf8_casefold()). If this is not appropriate, match functions
// have access to the unmodified key via
// `gtk_editable_get_text (GTK_EDITABLE (gtk_entry_completion_get_entry ()))`.
type EntryCompletionMatchFunc func(uintptr, string, *TreeIter, uintptr) bool

// `GtkEntryCompletion` is an auxiliary object to provide completion functionality
// for `GtkEntry`.
//
// It implements the [iface@Gtk.CellLayout] interface, to allow the user
// to add extra cells to the `GtkTreeView` with completion matches.
//
// “Completion functionality” means that when the user modifies the text
// in the entry, `GtkEntryCompletion` checks which rows in the model match
// the current content of the entry, and displays a list of matches.
// By default, the matching is done by comparing the entry text
// case-insensitively against the text column of the model (see
// [method@Gtk.EntryCompletion.set_text_column]), but this can be overridden
// with a custom match function (see [method@Gtk.EntryCompletion.set_match_func]).
//
// When the user selects a completion, the content of the entry is
// updated. By default, the content of the entry is replaced by the
// text column of the model, but this can be overridden by connecting
// to the [signal@Gtk.EntryCompletion::match-selected] signal and updating the
// entry in the signal handler. Note that you should return %TRUE from
// the signal handler to suppress the default behaviour.
//
// To add completion functionality to an entry, use
// [method@Gtk.Entry.set_completion].
//
// `GtkEntryCompletion` uses a [class@Gtk.TreeModelFilter] model to
// represent the subset of the entire model that is currently matching.
// While the `GtkEntryCompletion` signals
// [signal@Gtk.EntryCompletion::match-selected] and
// [signal@Gtk.EntryCompletion::cursor-on-match] take the original model
// and an iter pointing to that model as arguments, other callbacks and
// signals (such as `GtkCellLayoutDataFunc` or
// [signal@Gtk.CellArea::apply-attributes)]
// will generally take the filter model as argument. As long as you are
// only calling [method@Gtk.TreeModel.get], this will make no difference to
// you. If for some reason, you need the original model, use
// [method@Gtk.TreeModelFilter.get_model]. Don’t forget to use
// [method@Gtk.TreeModelFilter.convert_iter_to_child_iter] to obtain a
// matching iter.
type EntryCompletion struct {
	gobject.Object
}

func EntryCompletionNewFromInternalPtr(ptr uintptr) *EntryCompletion {
	cls := &EntryCompletion{}
	cls.Ptr = ptr
	return cls
}

var xNewEntryCompletion func() uintptr

// Creates a new `GtkEntryCompletion` object.
func NewEntryCompletion() *EntryCompletion {
	var cls *EntryCompletion

	cret := xNewEntryCompletion()

	if cret == 0 {
		return nil
	}
	cls = &EntryCompletion{}
	cls.Ptr = cret
	return cls
}

var xNewEntryCompletionWithArea func(uintptr) uintptr

// Creates a new `GtkEntryCompletion` object using the
// specified @area.
//
// The `GtkCellArea` is used to layout cells in the underlying
// `GtkTreeViewColumn` for the drop-down menu.
func NewEntryCompletionWithArea(AreaVar *CellArea) *EntryCompletion {
	var cls *EntryCompletion

	cret := xNewEntryCompletionWithArea(AreaVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &EntryCompletion{}
	cls.Ptr = cret
	return cls
}

var xEntryCompletionComplete func(uintptr)

// Requests a completion operation, or in other words a refiltering of the
// current list with completions, using the current key.
//
// The completion list view will be updated accordingly.
func (x *EntryCompletion) Complete() {

	xEntryCompletionComplete(x.GoPointer())

}

var xEntryCompletionComputePrefix func(uintptr, string) string

// Computes the common prefix that is shared by all rows in @completion
// that start with @key.
//
// If no row matches @key, %NULL will be returned.
// Note that a text column must have been set for this function to work,
// see [method@Gtk.EntryCompletion.set_text_column] for details.
func (x *EntryCompletion) ComputePrefix(KeyVar string) string {

	cret := xEntryCompletionComputePrefix(x.GoPointer(), KeyVar)
	return cret
}

var xEntryCompletionGetCompletionPrefix func(uintptr) string

// Get the original text entered by the user that triggered
// the completion or %NULL if there’s no completion ongoing.
func (x *EntryCompletion) GetCompletionPrefix() string {

	cret := xEntryCompletionGetCompletionPrefix(x.GoPointer())
	return cret
}

var xEntryCompletionGetEntry func(uintptr) uintptr

// Gets the entry @completion has been attached to.
func (x *EntryCompletion) GetEntry() *Widget {
	var cls *Widget

	cret := xEntryCompletionGetEntry(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xEntryCompletionGetInlineCompletion func(uintptr) bool

// Returns whether the common prefix of the possible completions should
// be automatically inserted in the entry.
func (x *EntryCompletion) GetInlineCompletion() bool {

	cret := xEntryCompletionGetInlineCompletion(x.GoPointer())
	return cret
}

var xEntryCompletionGetInlineSelection func(uintptr) bool

// Returns %TRUE if inline-selection mode is turned on.
func (x *EntryCompletion) GetInlineSelection() bool {

	cret := xEntryCompletionGetInlineSelection(x.GoPointer())
	return cret
}

var xEntryCompletionGetMinimumKeyLength func(uintptr) int

// Returns the minimum key length as set for @completion.
func (x *EntryCompletion) GetMinimumKeyLength() int {

	cret := xEntryCompletionGetMinimumKeyLength(x.GoPointer())
	return cret
}

var xEntryCompletionGetModel func(uintptr) uintptr

// Returns the model the `GtkEntryCompletion` is using as data source.
//
// Returns %NULL if the model is unset.
func (x *EntryCompletion) GetModel() *TreeModelBase {
	var cls *TreeModelBase

	cret := xEntryCompletionGetModel(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &TreeModelBase{}
	cls.Ptr = cret
	return cls
}

var xEntryCompletionGetPopupCompletion func(uintptr) bool

// Returns whether the completions should be presented in a popup window.
func (x *EntryCompletion) GetPopupCompletion() bool {

	cret := xEntryCompletionGetPopupCompletion(x.GoPointer())
	return cret
}

var xEntryCompletionGetPopupSetWidth func(uintptr) bool

// Returns whether the completion popup window will be resized to the
// width of the entry.
func (x *EntryCompletion) GetPopupSetWidth() bool {

	cret := xEntryCompletionGetPopupSetWidth(x.GoPointer())
	return cret
}

var xEntryCompletionGetPopupSingleMatch func(uintptr) bool

// Returns whether the completion popup window will appear even if there is
// only a single match.
func (x *EntryCompletion) GetPopupSingleMatch() bool {

	cret := xEntryCompletionGetPopupSingleMatch(x.GoPointer())
	return cret
}

var xEntryCompletionGetTextColumn func(uintptr) int

// Returns the column in the model of @completion to get strings from.
func (x *EntryCompletion) GetTextColumn() int {

	cret := xEntryCompletionGetTextColumn(x.GoPointer())
	return cret
}

var xEntryCompletionInsertPrefix func(uintptr)

// Requests a prefix insertion.
func (x *EntryCompletion) InsertPrefix() {

	xEntryCompletionInsertPrefix(x.GoPointer())

}

var xEntryCompletionSetInlineCompletion func(uintptr, bool)

// Sets whether the common prefix of the possible completions should
// be automatically inserted in the entry.
func (x *EntryCompletion) SetInlineCompletion(InlineCompletionVar bool) {

	xEntryCompletionSetInlineCompletion(x.GoPointer(), InlineCompletionVar)

}

var xEntryCompletionSetInlineSelection func(uintptr, bool)

// Sets whether it is possible to cycle through the possible completions
// inside the entry.
func (x *EntryCompletion) SetInlineSelection(InlineSelectionVar bool) {

	xEntryCompletionSetInlineSelection(x.GoPointer(), InlineSelectionVar)

}

var xEntryCompletionSetMatchFunc func(uintptr, uintptr, uintptr, uintptr)

// Sets the match function for @completion to be @func.
//
// The match function is used to determine if a row should or
// should not be in the completion list.
func (x *EntryCompletion) SetMatchFunc(FuncVar EntryCompletionMatchFunc, FuncDataVar uintptr, FuncNotifyVar glib.DestroyNotify) {

	xEntryCompletionSetMatchFunc(x.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(FuncNotifyVar))

}

var xEntryCompletionSetMinimumKeyLength func(uintptr, int)

// Requires the length of the search key for @completion to be at least
// @length.
//
// This is useful for long lists, where completing using a small
// key takes a lot of time and will come up with meaningless results anyway
// (ie, a too large dataset).
func (x *EntryCompletion) SetMinimumKeyLength(LengthVar int) {

	xEntryCompletionSetMinimumKeyLength(x.GoPointer(), LengthVar)

}

var xEntryCompletionSetModel func(uintptr, uintptr)

// Sets the model for a `GtkEntryCompletion`.
//
// If @completion already has a model set, it will remove it
// before setting the new model. If model is %NULL, then it
// will unset the model.
func (x *EntryCompletion) SetModel(ModelVar TreeModel) {

	xEntryCompletionSetModel(x.GoPointer(), ModelVar.GoPointer())

}

var xEntryCompletionSetPopupCompletion func(uintptr, bool)

// Sets whether the completions should be presented in a popup window.
func (x *EntryCompletion) SetPopupCompletion(PopupCompletionVar bool) {

	xEntryCompletionSetPopupCompletion(x.GoPointer(), PopupCompletionVar)

}

var xEntryCompletionSetPopupSetWidth func(uintptr, bool)

// Sets whether the completion popup window will be resized to be the same
// width as the entry.
func (x *EntryCompletion) SetPopupSetWidth(PopupSetWidthVar bool) {

	xEntryCompletionSetPopupSetWidth(x.GoPointer(), PopupSetWidthVar)

}

var xEntryCompletionSetPopupSingleMatch func(uintptr, bool)

// Sets whether the completion popup window will appear even if there is
// only a single match.
//
// You may want to set this to %FALSE if you
// are using [property@Gtk.EntryCompletion:inline-completion].
func (x *EntryCompletion) SetPopupSingleMatch(PopupSingleMatchVar bool) {

	xEntryCompletionSetPopupSingleMatch(x.GoPointer(), PopupSingleMatchVar)

}

var xEntryCompletionSetTextColumn func(uintptr, int)

// Convenience function for setting up the most used case of this code: a
// completion list with just strings.
//
// This function will set up @completion
// to have a list displaying all (and just) strings in the completion list,
// and to get those strings from @column in the model of @completion.
//
// This functions creates and adds a `GtkCellRendererText` for the selected
// column. If you need to set the text column, but don't want the cell
// renderer, use g_object_set() to set the
// [property@Gtk.EntryCompletion:text-column] property directly.
func (x *EntryCompletion) SetTextColumn(ColumnVar int) {

	xEntryCompletionSetTextColumn(x.GoPointer(), ColumnVar)

}

func (c *EntryCompletion) GoPointer() uintptr {
	return c.Ptr
}

func (c *EntryCompletion) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a match from the cursor is on a match of the list.
//
// The default behaviour is to replace the contents
// of the entry with the contents of the text column in the row
// pointed to by @iter.
//
// Note that @model is the model that was passed to
// [method@Gtk.EntryCompletion.set_model].
func (x *EntryCompletion) ConnectCursorOnMatch(cb func(EntryCompletion, uintptr, uintptr) bool) uint32 {
	fcb := func(clsPtr uintptr, ModelVarp uintptr, IterVarp uintptr) bool {
		fa := EntryCompletion{}
		fa.Ptr = clsPtr

		return cb(fa, ModelVarp, IterVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "cursor-on-match", purego.NewCallback(fcb))
}

// Emitted when the inline autocompletion is triggered.
//
// The default behaviour is to make the entry display the
// whole prefix and select the newly inserted part.
//
// Applications may connect to this signal in order to insert only a
// smaller part of the @prefix into the entry - e.g. the entry used in
// the `GtkFileChooser` inserts only the part of the prefix up to the
// next '/'.
func (x *EntryCompletion) ConnectInsertPrefix(cb func(EntryCompletion, string) bool) uint32 {
	fcb := func(clsPtr uintptr, PrefixVarp string) bool {
		fa := EntryCompletion{}
		fa.Ptr = clsPtr

		return cb(fa, PrefixVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "insert-prefix", purego.NewCallback(fcb))
}

// Emitted when a match from the list is selected.
//
// The default behaviour is to replace the contents of the
// entry with the contents of the text column in the row
// pointed to by @iter.
//
// Note that @model is the model that was passed to
// [method@Gtk.EntryCompletion.set_model].
func (x *EntryCompletion) ConnectMatchSelected(cb func(EntryCompletion, uintptr, uintptr) bool) uint32 {
	fcb := func(clsPtr uintptr, ModelVarp uintptr, IterVarp uintptr) bool {
		fa := EntryCompletion{}
		fa.Ptr = clsPtr

		return cb(fa, ModelVarp, IterVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "match-selected", purego.NewCallback(fcb))
}

// Emitted when the filter model has zero
// number of rows in completion_complete method.
//
// In other words when `GtkEntryCompletion` is out of suggestions.
func (x *EntryCompletion) ConnectNoMatches(cb func(EntryCompletion)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := EntryCompletion{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "no-matches", purego.NewCallback(fcb))
}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *EntryCompletion) GetBuildableId() string {

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
func (x *EntryCompletion) AddAttribute(CellVar *CellRenderer, AttributeVar string, ColumnVar int) {

	XGtkCellLayoutAddAttribute(x.GoPointer(), CellVar.GoPointer(), AttributeVar, ColumnVar)

}

// Unsets all the mappings on all renderers on @cell_layout and
// removes all renderers from @cell_layout.
func (x *EntryCompletion) Clear() {

	XGtkCellLayoutClear(x.GoPointer())

}

// Clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
func (x *EntryCompletion) ClearAttributes(CellVar *CellRenderer) {

	XGtkCellLayoutClearAttributes(x.GoPointer(), CellVar.GoPointer())

}

// Returns the underlying `GtkCellArea` which might be @cell_layout
// if called on a `GtkCellArea` or might be %NULL if no `GtkCellArea`
// is used by @cell_layout.
func (x *EntryCompletion) GetArea() *CellArea {
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
func (x *EntryCompletion) GetCells() *glib.List {

	cret := XGtkCellLayoutGetCells(x.GoPointer())
	return cret
}

// Adds the @cell to the end of @cell_layout. If @expand is %FALSE, then the
// @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *EntryCompletion) PackEnd(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackEnd(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Packs the @cell into the beginning of @cell_layout. If @expand is %FALSE,
// then the @cell is allocated no more space than it needs. Any unused space
// is divided evenly between cells for which @expand is %TRUE.
//
// Note that reusing the same cell renderer is not supported.
func (x *EntryCompletion) PackStart(CellVar *CellRenderer, ExpandVar bool) {

	XGtkCellLayoutPackStart(x.GoPointer(), CellVar.GoPointer(), ExpandVar)

}

// Re-inserts @cell at @position.
//
// Note that @cell has already to be packed into @cell_layout
// for this to function properly.
func (x *EntryCompletion) Reorder(CellVar *CellRenderer, PositionVar int) {

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
func (x *EntryCompletion) SetAttributes(CellVar *CellRenderer, varArgs ...interface{}) {

	XGtkCellLayoutSetAttributes(x.GoPointer(), CellVar.GoPointer(), varArgs...)

}

// Sets the `GtkCellLayout`DataFunc to use for @cell_layout.
//
// This function is used instead of the standard attributes mapping
// for setting the column value, and should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
//
// @func may be %NULL to remove a previously set function.
func (x *EntryCompletion) SetCellDataFunc(CellVar *CellRenderer, FuncVar CellLayoutDataFunc, FuncDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkCellLayoutSetCellDataFunc(x.GoPointer(), CellVar.GoPointer(), purego.NewCallback(FuncVar), FuncDataVar, purego.NewCallback(DestroyVar))

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewEntryCompletion, lib, "gtk_entry_completion_new")
	core.PuregoSafeRegister(&xNewEntryCompletionWithArea, lib, "gtk_entry_completion_new_with_area")

	core.PuregoSafeRegister(&xEntryCompletionComplete, lib, "gtk_entry_completion_complete")
	core.PuregoSafeRegister(&xEntryCompletionComputePrefix, lib, "gtk_entry_completion_compute_prefix")
	core.PuregoSafeRegister(&xEntryCompletionGetCompletionPrefix, lib, "gtk_entry_completion_get_completion_prefix")
	core.PuregoSafeRegister(&xEntryCompletionGetEntry, lib, "gtk_entry_completion_get_entry")
	core.PuregoSafeRegister(&xEntryCompletionGetInlineCompletion, lib, "gtk_entry_completion_get_inline_completion")
	core.PuregoSafeRegister(&xEntryCompletionGetInlineSelection, lib, "gtk_entry_completion_get_inline_selection")
	core.PuregoSafeRegister(&xEntryCompletionGetMinimumKeyLength, lib, "gtk_entry_completion_get_minimum_key_length")
	core.PuregoSafeRegister(&xEntryCompletionGetModel, lib, "gtk_entry_completion_get_model")
	core.PuregoSafeRegister(&xEntryCompletionGetPopupCompletion, lib, "gtk_entry_completion_get_popup_completion")
	core.PuregoSafeRegister(&xEntryCompletionGetPopupSetWidth, lib, "gtk_entry_completion_get_popup_set_width")
	core.PuregoSafeRegister(&xEntryCompletionGetPopupSingleMatch, lib, "gtk_entry_completion_get_popup_single_match")
	core.PuregoSafeRegister(&xEntryCompletionGetTextColumn, lib, "gtk_entry_completion_get_text_column")
	core.PuregoSafeRegister(&xEntryCompletionInsertPrefix, lib, "gtk_entry_completion_insert_prefix")
	core.PuregoSafeRegister(&xEntryCompletionSetInlineCompletion, lib, "gtk_entry_completion_set_inline_completion")
	core.PuregoSafeRegister(&xEntryCompletionSetInlineSelection, lib, "gtk_entry_completion_set_inline_selection")
	core.PuregoSafeRegister(&xEntryCompletionSetMatchFunc, lib, "gtk_entry_completion_set_match_func")
	core.PuregoSafeRegister(&xEntryCompletionSetMinimumKeyLength, lib, "gtk_entry_completion_set_minimum_key_length")
	core.PuregoSafeRegister(&xEntryCompletionSetModel, lib, "gtk_entry_completion_set_model")
	core.PuregoSafeRegister(&xEntryCompletionSetPopupCompletion, lib, "gtk_entry_completion_set_popup_completion")
	core.PuregoSafeRegister(&xEntryCompletionSetPopupSetWidth, lib, "gtk_entry_completion_set_popup_set_width")
	core.PuregoSafeRegister(&xEntryCompletionSetPopupSingleMatch, lib, "gtk_entry_completion_set_popup_single_match")
	core.PuregoSafeRegister(&xEntryCompletionSetTextColumn, lib, "gtk_entry_completion_set_text_column")

}
