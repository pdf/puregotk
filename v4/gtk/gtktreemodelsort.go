// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TreeModelSortClass struct {
	ParentClass uintptr

	Padding uintptr
}

type TreeModelSortPrivate struct {
}

// A GtkTreeModel which makes an underlying tree model sortable
//
// The `GtkTreeModelSort` is a model which implements the `GtkTreeSortable`
// interface.  It does not hold any data itself, but rather is created with
// a child model and proxies its data.  It has identical column types to
// this child model, and the changes in the child are propagated.  The
// primary purpose of this model is to provide a way to sort a different
// model without modifying it. Note that the sort function used by
// `GtkTreeModelSort` is not guaranteed to be stable.
//
// The use of this is best demonstrated through an example.  In the
// following sample code we create two `GtkTreeView` widgets each with a
// view of the same data.  As the model is wrapped here by a
// `GtkTreeModelSort`, the two `GtkTreeView`s can each sort their
// view of the data without affecting the other.  By contrast, if we
// simply put the same model in each widget, then sorting the first would
// sort the second.
//
// ## Using a `GtkTreeModelSort`
//
// |[&lt;!-- language="C" --&gt;
//
//	{
//	  GtkTreeView *tree_view1;
//	  GtkTreeView *tree_view2;
//	  GtkTreeModel *sort_model1;
//	  GtkTreeModel *sort_model2;
//	  GtkTreeModel *child_model;
//
//	  // get the child model
//	  child_model = get_my_model ();
//
//	  // Create the first tree
//	  sort_model1 = gtk_tree_model_sort_new_with_model (child_model);
//	  tree_view1 = gtk_tree_view_new_with_model (sort_model1);
//
//	  // Create the second tree
//	  sort_model2 = gtk_tree_model_sort_new_with_model (child_model);
//	  tree_view2 = gtk_tree_view_new_with_model (sort_model2);
//
//	  // Now we can sort the two models independently
//	  gtk_tree_sortable_set_sort_column_id (GTK_TREE_SORTABLE (sort_model1),
//	                                        COLUMN_1, GTK_SORT_ASCENDING);
//	  gtk_tree_sortable_set_sort_column_id (GTK_TREE_SORTABLE (sort_model2),
//	                                        COLUMN_1, GTK_SORT_DESCENDING);
//	}
//
// ]|
//
// To demonstrate how to access the underlying child model from the sort
// model, the next example will be a callback for the `GtkTreeSelection`
// `GtkTreeSelection::changed` signal.  In this callback, we get a string
// from COLUMN_1 of the model.  We then modify the string, find the same
// selected row on the child model, and change the row there.
//
// ## Accessing the child model of in a selection changed callback
//
// |[&lt;!-- language="C" --&gt;
// void
// selection_changed (GtkTreeSelection *selection, gpointer data)
//
//	{
//	  GtkTreeModel *sort_model = NULL;
//	  GtkTreeModel *child_model;
//	  GtkTreeIter sort_iter;
//	  GtkTreeIter child_iter;
//	  char *some_data = NULL;
//	  char *modified_data;
//
//	  // Get the current selected row and the model.
//	  if (! gtk_tree_selection_get_selected (selection,
//	                                         &amp;sort_model,
//	                                         &amp;sort_iter))
//	    return;
//
//	  // Look up the current value on the selected row and get
//	  // a new value to change it to.
//	  gtk_tree_model_get (GTK_TREE_MODEL (sort_model), &amp;sort_iter,
//	                      COLUMN_1, &amp;some_data,
//	                      -1);
//
//	  modified_data = change_the_data (some_data);
//	  g_free (some_data);
//
//	  // Get an iterator on the child model, instead of the sort model.
//	  gtk_tree_model_sort_convert_iter_to_child_iter (GTK_TREE_MODEL_SORT (sort_model),
//	                                                  &amp;child_iter,
//	                                                  &amp;sort_iter);
//
//	  // Get the child model and change the value of the row. In this
//	  // example, the child model is a GtkListStore. It could be any other
//	  // type of model, though.
//	  child_model = gtk_tree_model_sort_get_model (GTK_TREE_MODEL_SORT (sort_model));
//	  gtk_list_store_set (GTK_LIST_STORE (child_model), &amp;child_iter,
//	                      COLUMN_1, &amp;modified_data,
//	                      -1);
//	  g_free (modified_data);
//	}
//
// ]|
type TreeModelSort struct {
	gobject.Object
}

func TreeModelSortNewFromInternalPtr(ptr uintptr) *TreeModelSort {
	cls := &TreeModelSort{}
	cls.Ptr = ptr
	return cls
}

var xNewWithModelTreeModelSort func(uintptr) uintptr

// Creates a new `GtkTreeModelSort`, with @child_model as the child model.
func NewWithModelTreeModelSort(ChildModelVar TreeModel) *TreeModelSort {
	NewWithModelTreeModelSortPtr := xNewWithModelTreeModelSort(ChildModelVar.GoPointer())
	if NewWithModelTreeModelSortPtr == 0 {
		return nil
	}

	NewWithModelTreeModelSortCls := &TreeModelSort{}
	NewWithModelTreeModelSortCls.Ptr = NewWithModelTreeModelSortPtr
	return NewWithModelTreeModelSortCls
}

var xTreeModelSortClearCache func(uintptr)

// This function should almost never be called.  It clears the @tree_model_sort
// of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node().  This might be useful if the child model being
// sorted is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes.  As a side effect of this function, all unreffed
// iters will be invalid.
func (x *TreeModelSort) ClearCache() {

	xTreeModelSortClearCache(x.GoPointer())

}

var xTreeModelSortConvertChildIterToIter func(uintptr, *TreeIter, *TreeIter) bool

// Sets @sort_iter to point to the row in @tree_model_sort that corresponds to
// the row pointed at by @child_iter.  If @sort_iter was not set, %FALSE
// is returned.  Note: a boolean is only returned since 2.14.
func (x *TreeModelSort) ConvertChildIterToIter(SortIterVar *TreeIter, ChildIterVar *TreeIter) bool {

	return xTreeModelSortConvertChildIterToIter(x.GoPointer(), SortIterVar, ChildIterVar)

}

var xTreeModelSortConvertChildPathToPath func(uintptr, *TreePath) *TreePath

// Converts @child_path to a path relative to @tree_model_sort.  That is,
// @child_path points to a path in the child model.  The returned path will
// point to the same row in the sorted model.  If @child_path isn’t a valid
// path on the child model, then %NULL is returned.
func (x *TreeModelSort) ConvertChildPathToPath(ChildPathVar *TreePath) *TreePath {

	return xTreeModelSortConvertChildPathToPath(x.GoPointer(), ChildPathVar)

}

var xTreeModelSortConvertIterToChildIter func(uintptr, *TreeIter, *TreeIter)

// Sets @child_iter to point to the row pointed to by @sorted_iter.
func (x *TreeModelSort) ConvertIterToChildIter(ChildIterVar *TreeIter, SortedIterVar *TreeIter) {

	xTreeModelSortConvertIterToChildIter(x.GoPointer(), ChildIterVar, SortedIterVar)

}

var xTreeModelSortConvertPathToChildPath func(uintptr, *TreePath) *TreePath

// Converts @sorted_path to a path on the child model of @tree_model_sort.
// That is, @sorted_path points to a location in @tree_model_sort.  The
// returned path will point to the same location in the model not being
// sorted.  If @sorted_path does not point to a location in the child model,
// %NULL is returned.
func (x *TreeModelSort) ConvertPathToChildPath(SortedPathVar *TreePath) *TreePath {

	return xTreeModelSortConvertPathToChildPath(x.GoPointer(), SortedPathVar)

}

var xTreeModelSortGetModel func(uintptr) uintptr

// Returns the model the `GtkTreeModelSort` is sorting.
func (x *TreeModelSort) GetModel() *TreeModelBase {

	GetModelPtr := xTreeModelSortGetModel(x.GoPointer())
	if GetModelPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetModelPtr)

	GetModelCls := &TreeModelBase{}
	GetModelCls.Ptr = GetModelPtr
	return GetModelCls

}

var xTreeModelSortIterIsValid func(uintptr, *TreeIter) bool

// &gt; This function is slow. Only use it for debugging and/or testing
// &gt; purposes.
//
// Checks if the given iter is a valid iter for this `GtkTreeModelSort`.
func (x *TreeModelSort) IterIsValid(IterVar *TreeIter) bool {

	return xTreeModelSortIterIsValid(x.GoPointer(), IterVar)

}

var xTreeModelSortResetDefaultSortFunc func(uintptr)

// This resets the default sort function to be in the “unsorted” state.  That
// is, it is in the same order as the child model. It will re-sort the model
// to be in the same order as the child model only if the `GtkTreeModelSort`
// is in “unsorted” state.
func (x *TreeModelSort) ResetDefaultSortFunc() {

	xTreeModelSortResetDefaultSortFunc(x.GoPointer())

}

func (c *TreeModelSort) GoPointer() uintptr {
	return c.Ptr
}

func (c *TreeModelSort) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Asks the `GtkTreeDragSource` to delete the row at @path, because
// it was moved somewhere else via drag-and-drop. Returns %FALSE
// if the deletion fails because @path no longer exists, or for
// some model-specific reason. Should robustly handle a @path no
// longer found in the model!
func (x *TreeModelSort) DragDataDelete(PathVar *TreePath) bool {

	return XGtkTreeDragSourceDragDataDelete(x.GoPointer(), PathVar)

}

// Asks the `GtkTreeDragSource` to return a `GdkContentProvider` representing
// the row at @path. Should robustly handle a @path no
// longer found in the model!
func (x *TreeModelSort) DragDataGet(PathVar *TreePath) *gdk.ContentProvider {

	DragDataGetPtr := XGtkTreeDragSourceDragDataGet(x.GoPointer(), PathVar)
	if DragDataGetPtr == 0 {
		return nil
	}

	DragDataGetCls := &gdk.ContentProvider{}
	DragDataGetCls.Ptr = DragDataGetPtr
	return DragDataGetCls

}

// Asks the `GtkTreeDragSource` whether a particular row can be used as
// the source of a DND operation. If the source doesn’t implement
// this interface, the row is assumed draggable.
func (x *TreeModelSort) RowDraggable(PathVar *TreePath) bool {

	return XGtkTreeDragSourceRowDraggable(x.GoPointer(), PathVar)

}

// Creates a new `GtkTreeModel`, with @child_model as the child_model
// and @root as the virtual root.
func (x *TreeModelSort) FilterNew(RootVar *TreePath) *TreeModelBase {

	FilterNewPtr := XGtkTreeModelFilterNew(x.GoPointer(), RootVar)
	if FilterNewPtr == 0 {
		return nil
	}

	FilterNewCls := &TreeModelBase{}
	FilterNewCls.Ptr = FilterNewPtr
	return FilterNewCls

}

// Calls @func on each node in model in a depth-first fashion.
//
// If @func returns %TRUE, then the tree ceases to be walked,
// and gtk_tree_model_foreach() returns.
func (x *TreeModelSort) Foreach(FuncVar TreeModelForeachFunc, UserDataVar uintptr) {

	XGtkTreeModelForeach(x.GoPointer(), purego.NewCallback(FuncVar), UserDataVar)

}

// Gets the value of one or more cells in the row referenced by @iter.
//
// The variable argument list should contain integer column numbers,
// each column number followed by a place to store the value being
// retrieved.  The list is terminated by a -1. For example, to get a
// value from column 0 with type %G_TYPE_STRING, you would
// write: `gtk_tree_model_get (model, iter, 0, &amp;place_string_here, -1)`,
// where `place_string_here` is a #gchararray
// to be filled with the string.
//
// Returned values with type %G_TYPE_OBJECT have to be unreferenced,
// values with type %G_TYPE_STRING or %G_TYPE_BOXED have to be freed.
// Other values are passed by value.
func (x *TreeModelSort) Get(IterVar *TreeIter, varArgs ...interface{}) {

	XGtkTreeModelGet(x.GoPointer(), IterVar, varArgs...)

}

// Returns the type of the column.
func (x *TreeModelSort) GetColumnType(IndexVar int) []interface{} {

	return XGtkTreeModelGetColumnType(x.GoPointer(), IndexVar)

}

// Returns a set of flags supported by this interface.
//
// The flags are a bitwise combination of `GtkTreeModel`Flags.
// The flags supported should not change during the lifetime
// of the @tree_model.
func (x *TreeModelSort) GetFlags() TreeModelFlags {

	return XGtkTreeModelGetFlags(x.GoPointer())

}

// Sets @iter to a valid iterator pointing to @path.
//
// If @path does not exist, @iter is set to an invalid
// iterator and %FALSE is returned.
func (x *TreeModelSort) GetIter(IterVar *TreeIter, PathVar *TreePath) bool {

	return XGtkTreeModelGetIter(x.GoPointer(), IterVar, PathVar)

}

// Initializes @iter with the first iterator in the tree
// (the one at the path "0").
//
// Returns %FALSE if the tree is empty, %TRUE otherwise.
func (x *TreeModelSort) GetIterFirst(IterVar *TreeIter) bool {

	return XGtkTreeModelGetIterFirst(x.GoPointer(), IterVar)

}

// Sets @iter to a valid iterator pointing to @path_string, if it
// exists.
//
// Otherwise, @iter is left invalid and %FALSE is returned.
func (x *TreeModelSort) GetIterFromString(IterVar *TreeIter, PathStringVar string) bool {

	return XGtkTreeModelGetIterFromString(x.GoPointer(), IterVar, PathStringVar)

}

// Returns the number of columns supported by @tree_model.
func (x *TreeModelSort) GetNColumns() int {

	return XGtkTreeModelGetNColumns(x.GoPointer())

}

// Returns a newly-created `GtkTreePath` referenced by @iter.
//
// This path should be freed with gtk_tree_path_free().
func (x *TreeModelSort) GetPath(IterVar *TreeIter) *TreePath {

	return XGtkTreeModelGetPath(x.GoPointer(), IterVar)

}

// Generates a string representation of the iter.
//
// This string is a “:” separated list of numbers.
// For example, “4:10:0:3” would be an acceptable
// return value for this string.
func (x *TreeModelSort) GetStringFromIter(IterVar *TreeIter) string {

	return XGtkTreeModelGetStringFromIter(x.GoPointer(), IterVar)

}

// Gets the value of one or more cells in the row referenced by @iter.
//
// See [method@Gtk.TreeModel.get], this version takes a va_list
// for language bindings to use.
func (x *TreeModelSort) GetValist(IterVar *TreeIter, VarArgsVar []interface{}) {

	XGtkTreeModelGetValist(x.GoPointer(), IterVar, VarArgsVar)

}

// Initializes and sets @value to that at @column.
//
// When done with @value, g_value_unset() needs to be called
// to free any allocated memory.
func (x *TreeModelSort) GetValue(IterVar *TreeIter, ColumnVar int, ValueVar *gobject.Value) {

	XGtkTreeModelGetValue(x.GoPointer(), IterVar, ColumnVar, ValueVar)

}

// Sets @iter to point to the first child of @parent.
//
// If @parent has no children, %FALSE is returned and @iter is
// set to be invalid. @parent will remain a valid node after this
// function has been called.
//
// If @parent is %NULL returns the first node, equivalent to
// `gtk_tree_model_get_iter_first (tree_model, iter);`
func (x *TreeModelSort) IterChildren(IterVar *TreeIter, ParentVar *TreeIter) bool {

	return XGtkTreeModelIterChildren(x.GoPointer(), IterVar, ParentVar)

}

// Returns %TRUE if @iter has children, %FALSE otherwise.
func (x *TreeModelSort) IterHasChild(IterVar *TreeIter) bool {

	return XGtkTreeModelIterHasChild(x.GoPointer(), IterVar)

}

// Returns the number of children that @iter has.
//
// As a special case, if @iter is %NULL, then the number
// of toplevel nodes is returned.
func (x *TreeModelSort) IterNChildren(IterVar *TreeIter) int {

	return XGtkTreeModelIterNChildren(x.GoPointer(), IterVar)

}

// Sets @iter to point to the node following it at the current level.
//
// If there is no next @iter, %FALSE is returned and @iter is set
// to be invalid.
func (x *TreeModelSort) IterNext(IterVar *TreeIter) bool {

	return XGtkTreeModelIterNext(x.GoPointer(), IterVar)

}

// Sets @iter to be the child of @parent, using the given index.
//
// The first index is 0. If @n is too big, or @parent has no children,
// @iter is set to an invalid iterator and %FALSE is returned. @parent
// will remain a valid node after this function has been called. As a
// special case, if @parent is %NULL, then the @n-th root node
// is set.
func (x *TreeModelSort) IterNthChild(IterVar *TreeIter, ParentVar *TreeIter, NVar int) bool {

	return XGtkTreeModelIterNthChild(x.GoPointer(), IterVar, ParentVar, NVar)

}

// Sets @iter to be the parent of @child.
//
// If @child is at the toplevel, and doesn’t have a parent, then
// @iter is set to an invalid iterator and %FALSE is returned.
// @child will remain a valid node after this function has been
// called.
//
// @iter will be initialized before the lookup is performed, so @child
// and @iter cannot point to the same memory location.
func (x *TreeModelSort) IterParent(IterVar *TreeIter, ChildVar *TreeIter) bool {

	return XGtkTreeModelIterParent(x.GoPointer(), IterVar, ChildVar)

}

// Sets @iter to point to the previous node at the current level.
//
// If there is no previous @iter, %FALSE is returned and @iter is
// set to be invalid.
func (x *TreeModelSort) IterPrevious(IterVar *TreeIter) bool {

	return XGtkTreeModelIterPrevious(x.GoPointer(), IterVar)

}

// Lets the tree ref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons.
//
// This function is primarily meant as a way for views to let
// caching models know when nodes are being displayed (and hence,
// whether or not to cache that node). Being displayed means a node
// is in an expanded branch, regardless of whether the node is currently
// visible in the viewport. For example, a file-system based model
// would not want to keep the entire file-hierarchy in memory,
// just the sections that are currently being displayed by
// every current view.
//
// A model should be expected to be able to get an iter independent
// of its reffed state.
func (x *TreeModelSort) RefNode(IterVar *TreeIter) {

	XGtkTreeModelRefNode(x.GoPointer(), IterVar)

}

// Emits the ::row-changed signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-changed].
func (x *TreeModelSort) RowChanged(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowChanged(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-deleted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-deleted].
//
// This should be called by models after a row has been removed.
// The location pointed to by @path should be the location that
// the row previously was at. It may not be a valid location anymore.
//
// Nodes that are deleted are not unreffed, this means that any
// outstanding references on the deleted node should not be released.
func (x *TreeModelSort) RowDeleted(PathVar *TreePath) {

	XGtkTreeModelRowDeleted(x.GoPointer(), PathVar)

}

// Emits the ::row-has-child-toggled signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-has-child-toggled].
//
// This should be called by models after the child
// state of a node changes.
func (x *TreeModelSort) RowHasChildToggled(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowHasChildToggled(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::row-inserted signal on @tree_model.
//
// See [signal@Gtk.TreeModel::row-inserted].
func (x *TreeModelSort) RowInserted(PathVar *TreePath, IterVar *TreeIter) {

	XGtkTreeModelRowInserted(x.GoPointer(), PathVar, IterVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelSort) RowsReordered(PathVar *TreePath, IterVar *TreeIter, NewOrderVar int) {

	XGtkTreeModelRowsReordered(x.GoPointer(), PathVar, IterVar, NewOrderVar)

}

// Emits the ::rows-reordered signal on @tree_model.
//
// See [signal@Gtk.TreeModel::rows-reordered].
//
// This should be called by models when their rows have been
// reordered.
func (x *TreeModelSort) RowsReorderedWithLength(PathVar *TreePath, IterVar *TreeIter, NewOrderVar uintptr, LengthVar int) {

	XGtkTreeModelRowsReorderedWithLength(x.GoPointer(), PathVar, IterVar, NewOrderVar, LengthVar)

}

// Lets the tree unref the node.
//
// This is an optional method for models to implement.
// To be more specific, models may ignore this call as it exists
// primarily for performance reasons. For more information on what
// this means, see gtk_tree_model_ref_node().
//
// Please note that nodes that are deleted are not unreffed.
func (x *TreeModelSort) UnrefNode(IterVar *TreeIter) {

	XGtkTreeModelUnrefNode(x.GoPointer(), IterVar)

}

// Fills in @sort_column_id and @order with the current sort column and the
// order. It returns %TRUE unless the @sort_column_id is
// %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// %GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (x *TreeModelSort) GetSortColumnId(SortColumnIdVar int, OrderVar *SortType) bool {

	return XGtkTreeSortableGetSortColumnId(x.GoPointer(), SortColumnIdVar, OrderVar)

}

// Returns %TRUE if the model has a default sort function. This is used
// primarily by GtkTreeViewColumns in order to determine if a model can
// go back to the default state, or not.
func (x *TreeModelSort) HasDefaultSortFunc() bool {

	return XGtkTreeSortableHasDefaultSortFunc(x.GoPointer())

}

// Sets the default comparison function used when sorting to be @sort_func.
// If the current sort column id of @sortable is
// %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using
// this function.
//
// If @sort_func is %NULL, then there will be no default comparison function.
// This means that once the model  has been sorted, it can’t go back to the
// default state. In this case, when the current sort column id of @sortable
// is %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
func (x *TreeModelSort) SetDefaultSortFunc(SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkTreeSortableSetDefaultSortFunc(x.GoPointer(), purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(DestroyVar))

}

// Sets the current sort column to be @sort_column_id. The @sortable will
// resort itself to reflect this change, after emitting a
// `GtkTreeSortable::sort-column-changed` signal. @sort_column_id may either be
// a regular column id, or one of the following special values:
//
//   - %GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
//     will be used, if it is set
//
// - %GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
func (x *TreeModelSort) SetSortColumnId(SortColumnIdVar int, OrderVar SortType) {

	XGtkTreeSortableSetSortColumnId(x.GoPointer(), SortColumnIdVar, OrderVar)

}

// Sets the comparison function used when sorting to be @sort_func. If the
// current sort column id of @sortable is the same as @sort_column_id, then
// the model will sort using this function.
func (x *TreeModelSort) SetSortFunc(SortColumnIdVar int, SortFuncVar TreeIterCompareFunc, UserDataVar uintptr, DestroyVar glib.DestroyNotify) {

	XGtkTreeSortableSetSortFunc(x.GoPointer(), SortColumnIdVar, purego.NewCallback(SortFuncVar), UserDataVar, purego.NewCallback(DestroyVar))

}

// Emits a `GtkTreeSortable::sort-column-changed` signal on @sortable.
func (x *TreeModelSort) SortColumnChanged() {

	XGtkTreeSortableSortColumnChanged(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewWithModelTreeModelSort, lib, "gtk_tree_model_sort_new_with_model")

	core.PuregoSafeRegister(&xTreeModelSortClearCache, lib, "gtk_tree_model_sort_clear_cache")
	core.PuregoSafeRegister(&xTreeModelSortConvertChildIterToIter, lib, "gtk_tree_model_sort_convert_child_iter_to_iter")
	core.PuregoSafeRegister(&xTreeModelSortConvertChildPathToPath, lib, "gtk_tree_model_sort_convert_child_path_to_path")
	core.PuregoSafeRegister(&xTreeModelSortConvertIterToChildIter, lib, "gtk_tree_model_sort_convert_iter_to_child_iter")
	core.PuregoSafeRegister(&xTreeModelSortConvertPathToChildPath, lib, "gtk_tree_model_sort_convert_path_to_child_path")
	core.PuregoSafeRegister(&xTreeModelSortGetModel, lib, "gtk_tree_model_sort_get_model")
	core.PuregoSafeRegister(&xTreeModelSortIterIsValid, lib, "gtk_tree_model_sort_iter_is_valid")
	core.PuregoSafeRegister(&xTreeModelSortResetDefaultSortFunc, lib, "gtk_tree_model_sort_reset_default_sort_func")

}
