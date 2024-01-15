// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of function passed to g_tree_traverse(). It is
// passed the key and value of each node, together with the @user_data
// parameter passed to g_tree_traverse(). If the function returns
// %TRUE, the traversal is stopped.
type TraverseFunc func(uintptr, uintptr, uintptr) bool

// Specifies the type of function passed to g_tree_foreach_node(). It is
// passed each node, together with the @user_data parameter passed to
// g_tree_foreach_node(). If the function returns %TRUE, the traversal is
// stopped.
type TraverseNodeFunc func(*TreeNode, uintptr) bool

// The GTree struct is an opaque data structure representing a
// [balanced binary tree][glib-Balanced-Binary-Trees]. It should be
// accessed only by using the following functions.
type Tree struct {
}

func (x *Tree) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xNewTree func(uintptr) *Tree

// Creates a new #GTree.
func NewTree(KeyCompareFuncVar CompareFunc) *Tree {

	cret := xNewTree(purego.NewCallback(KeyCompareFuncVar))
	return cret
}

var xNewTreeFull func(uintptr, uintptr, uintptr, uintptr) *Tree

// Creates a new #GTree like g_tree_new() and allows to specify functions
// to free the memory allocated for the key and value that get called when
// removing the entry from the #GTree.
func NewTreeFull(KeyCompareFuncVar CompareDataFunc, KeyCompareDataVar uintptr, KeyDestroyFuncVar DestroyNotify, ValueDestroyFuncVar DestroyNotify) *Tree {

	cret := xNewTreeFull(purego.NewCallback(KeyCompareFuncVar), KeyCompareDataVar, purego.NewCallback(KeyDestroyFuncVar), purego.NewCallback(ValueDestroyFuncVar))
	return cret
}

var xNewTreeWithData func(uintptr, uintptr) *Tree

// Creates a new #GTree with a comparison function that accepts user data.
// See g_tree_new() for more details.
func NewTreeWithData(KeyCompareFuncVar CompareDataFunc, KeyCompareDataVar uintptr) *Tree {

	cret := xNewTreeWithData(purego.NewCallback(KeyCompareFuncVar), KeyCompareDataVar)
	return cret
}

var xTreeDestroy func(uintptr)

// Removes all keys and values from the #GTree and decreases its
// reference count by one. If keys and/or values are dynamically
// allocated, you should either free them first or create the #GTree
// using g_tree_new_full(). In the latter case the destroy functions
// you supplied will be called on all keys and values before destroying
// the #GTree.
func (x *Tree) Destroy() {

	xTreeDestroy(x.GoPointer())

}

var xTreeForeach func(uintptr, uintptr, uintptr)

// Calls the given function for each of the key/value pairs in the #GTree.
// The function is passed the key and value of each pair, and the given
// @data parameter. The tree is traversed in sorted order.
//
// The tree may not be modified while iterating over it (you can't
// add/remove items). To remove all items matching a predicate, you need
// to add each item to a list in your #GTraverseFunc as you walk over
// the tree, then walk the list and remove each item.
func (x *Tree) Foreach(FuncVar TraverseFunc, UserDataVar uintptr) {

	xTreeForeach(x.GoPointer(), purego.NewCallback(FuncVar), UserDataVar)

}

var xTreeForeachNode func(uintptr, uintptr, uintptr)

// Calls the given function for each of the nodes in the #GTree.
// The function is passed the pointer to the particular node, and the given
// @data parameter. The tree traversal happens in-order.
//
// The tree may not be modified while iterating over it (you can't
// add/remove items). To remove all items matching a predicate, you need
// to add each item to a list in your #GTraverseFunc as you walk over
// the tree, then walk the list and remove each item.
func (x *Tree) ForeachNode(FuncVar TraverseNodeFunc, UserDataVar uintptr) {

	xTreeForeachNode(x.GoPointer(), purego.NewCallback(FuncVar), UserDataVar)

}

var xTreeHeight func(uintptr) int

// Gets the height of a #GTree.
//
// If the #GTree contains no nodes, the height is 0.
// If the #GTree contains only one root node the height is 1.
// If the root node has children the height is 2, etc.
func (x *Tree) Height() int {

	cret := xTreeHeight(x.GoPointer())
	return cret
}

var xTreeInsert func(uintptr, uintptr, uintptr)

// Inserts a key/value pair into a #GTree.
//
// Inserts a new key and value into a #GTree as g_tree_insert_node() does,
// only this function does not return the inserted or set node.
func (x *Tree) Insert(KeyVar uintptr, ValueVar uintptr) {

	xTreeInsert(x.GoPointer(), KeyVar, ValueVar)

}

var xTreeInsertNode func(uintptr, uintptr, uintptr) *TreeNode

// Inserts a key/value pair into a #GTree.
//
// If the given key already exists in the #GTree its corresponding value
// is set to the new value. If you supplied a @value_destroy_func when
// creating the #GTree, the old value is freed using that function. If
// you supplied a @key_destroy_func when creating the #GTree, the passed
// key is freed using that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added,
// so that the distance from the root to every leaf is as small as possible.
// The cost of maintaining a balanced tree while inserting new key/value
// result in a O(n log(n)) operation where most of the other operations
// are O(log(n)).
func (x *Tree) InsertNode(KeyVar uintptr, ValueVar uintptr) *TreeNode {

	cret := xTreeInsertNode(x.GoPointer(), KeyVar, ValueVar)
	return cret
}

var xTreeLookup func(uintptr, uintptr) uintptr

// Gets the value corresponding to the given key. Since a #GTree is
// automatically balanced as key/value pairs are added, key lookup
// is O(log n) (where n is the number of key/value pairs in the tree).
func (x *Tree) Lookup(KeyVar uintptr) uintptr {

	cret := xTreeLookup(x.GoPointer(), KeyVar)
	return cret
}

var xTreeLookupExtended func(uintptr, uintptr, uintptr, uintptr) bool

// Looks up a key in the #GTree, returning the original key and the
// associated value. This is useful if you need to free the memory
// allocated for the original key, for example before calling
// g_tree_remove().
func (x *Tree) LookupExtended(LookupKeyVar uintptr, OrigKeyVar uintptr, ValueVar uintptr) bool {

	cret := xTreeLookupExtended(x.GoPointer(), LookupKeyVar, OrigKeyVar, ValueVar)
	return cret
}

var xTreeLookupNode func(uintptr, uintptr) *TreeNode

// Gets the tree node corresponding to the given key. Since a #GTree is
// automatically balanced as key/value pairs are added, key lookup
// is O(log n) (where n is the number of key/value pairs in the tree).
func (x *Tree) LookupNode(KeyVar uintptr) *TreeNode {

	cret := xTreeLookupNode(x.GoPointer(), KeyVar)
	return cret
}

var xTreeLowerBound func(uintptr, uintptr) *TreeNode

// Gets the lower bound node corresponding to the given key,
// or %NULL if the tree is empty or all the nodes in the tree
// have keys that are strictly lower than the searched key.
//
// The lower bound is the first node that has its key greater
// than or equal to the searched key.
func (x *Tree) LowerBound(KeyVar uintptr) *TreeNode {

	cret := xTreeLowerBound(x.GoPointer(), KeyVar)
	return cret
}

var xTreeNnodes func(uintptr) int

// Gets the number of nodes in a #GTree.
func (x *Tree) Nnodes() int {

	cret := xTreeNnodes(x.GoPointer())
	return cret
}

var xTreeNodeFirst func(uintptr) *TreeNode

// Returns the first in-order node of the tree, or %NULL
// for an empty tree.
func (x *Tree) NodeFirst() *TreeNode {

	cret := xTreeNodeFirst(x.GoPointer())
	return cret
}

var xTreeNodeLast func(uintptr) *TreeNode

// Returns the last in-order node of the tree, or %NULL
// for an empty tree.
func (x *Tree) NodeLast() *TreeNode {

	cret := xTreeNodeLast(x.GoPointer())
	return cret
}

var xTreeRef func(uintptr) *Tree

// Increments the reference count of @tree by one.
//
// It is safe to call this function from any thread.
func (x *Tree) Ref() *Tree {

	cret := xTreeRef(x.GoPointer())
	return cret
}

var xTreeRemove func(uintptr, uintptr) bool

// Removes a key/value pair from a #GTree.
//
// If the #GTree was created using g_tree_new_full(), the key and value
// are freed using the supplied destroy functions, otherwise you have to
// make sure that any dynamically allocated values are freed yourself.
// If the key does not exist in the #GTree, the function does nothing.
//
// The cost of maintaining a balanced tree while removing a key/value
// result in a O(n log(n)) operation where most of the other operations
// are O(log(n)).
func (x *Tree) Remove(KeyVar uintptr) bool {

	cret := xTreeRemove(x.GoPointer(), KeyVar)
	return cret
}

var xTreeRemoveAll func(uintptr)

// Removes all nodes from a #GTree and destroys their keys and values,
// then resets the #GTree’s root to %NULL.
func (x *Tree) RemoveAll() {

	xTreeRemoveAll(x.GoPointer())

}

var xTreeReplace func(uintptr, uintptr, uintptr)

// Inserts a new key and value into a #GTree as g_tree_replace_node() does,
// only this function does not return the inserted or set node.
func (x *Tree) Replace(KeyVar uintptr, ValueVar uintptr) {

	xTreeReplace(x.GoPointer(), KeyVar, ValueVar)

}

var xTreeReplaceNode func(uintptr, uintptr, uintptr) *TreeNode

// Inserts a new key and value into a #GTree similar to g_tree_insert_node().
// The difference is that if the key already exists in the #GTree, it gets
// replaced by the new key. If you supplied a @value_destroy_func when
// creating the #GTree, the old value is freed using that function. If you
// supplied a @key_destroy_func when creating the #GTree, the old key is
// freed using that function.
//
// The tree is automatically 'balanced' as new key/value pairs are added,
// so that the distance from the root to every leaf is as small as possible.
func (x *Tree) ReplaceNode(KeyVar uintptr, ValueVar uintptr) *TreeNode {

	cret := xTreeReplaceNode(x.GoPointer(), KeyVar, ValueVar)
	return cret
}

var xTreeSearch func(uintptr, uintptr, uintptr) uintptr

// Searches a #GTree using @search_func.
//
// The @search_func is called with a pointer to the key of a key/value
// pair in the tree, and the passed in @user_data. If @search_func returns
// 0 for a key/value pair, then the corresponding value is returned as
// the result of g_tree_search(). If @search_func returns -1, searching
// will proceed among the key/value pairs that have a smaller key; if
// @search_func returns 1, searching will proceed among the key/value
// pairs that have a larger key.
func (x *Tree) Search(SearchFuncVar CompareFunc, UserDataVar uintptr) uintptr {

	cret := xTreeSearch(x.GoPointer(), purego.NewCallback(SearchFuncVar), UserDataVar)
	return cret
}

var xTreeSearchNode func(uintptr, uintptr, uintptr) *TreeNode

// Searches a #GTree using @search_func.
//
// The @search_func is called with a pointer to the key of a key/value
// pair in the tree, and the passed in @user_data. If @search_func returns
// 0 for a key/value pair, then the corresponding node is returned as
// the result of g_tree_search(). If @search_func returns -1, searching
// will proceed among the key/value pairs that have a smaller key; if
// @search_func returns 1, searching will proceed among the key/value
// pairs that have a larger key.
func (x *Tree) SearchNode(SearchFuncVar CompareFunc, UserDataVar uintptr) *TreeNode {

	cret := xTreeSearchNode(x.GoPointer(), purego.NewCallback(SearchFuncVar), UserDataVar)
	return cret
}

var xTreeSteal func(uintptr, uintptr) bool

// Removes a key and its associated value from a #GTree without calling
// the key and value destroy functions.
//
// If the key does not exist in the #GTree, the function does nothing.
func (x *Tree) Steal(KeyVar uintptr) bool {

	cret := xTreeSteal(x.GoPointer(), KeyVar)
	return cret
}

var xTreeTraverse func(uintptr, uintptr, TraverseType, uintptr)

// Calls the given function for each node in the #GTree.
func (x *Tree) Traverse(TraverseFuncVar TraverseFunc, TraverseTypeVar TraverseType, UserDataVar uintptr) {

	xTreeTraverse(x.GoPointer(), purego.NewCallback(TraverseFuncVar), TraverseTypeVar, UserDataVar)

}

var xTreeUnref func(uintptr)

// Decrements the reference count of @tree by one.
// If the reference count drops to 0, all keys and values will
// be destroyed (if destroy functions were specified) and all
// memory allocated by @tree will be released.
//
// It is safe to call this function from any thread.
func (x *Tree) Unref() {

	xTreeUnref(x.GoPointer())

}

var xTreeUpperBound func(uintptr, uintptr) *TreeNode

// Gets the upper bound node corresponding to the given key,
// or %NULL if the tree is empty or all the nodes in the tree
// have keys that are lower than or equal to the searched key.
//
// The upper bound is the first node that has its key strictly greater
// than the searched key.
func (x *Tree) UpperBound(KeyVar uintptr) *TreeNode {

	cret := xTreeUpperBound(x.GoPointer(), KeyVar)
	return cret
}

// An opaque type which identifies a specific node in a #GTree.
type TreeNode struct {
}

func (x *TreeNode) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

var xTreeNodeKey func(uintptr) uintptr

// Gets the key stored at a particular tree node.
func (x *TreeNode) Key() uintptr {

	cret := xTreeNodeKey(x.GoPointer())
	return cret
}

var xTreeNodeNext func(uintptr) *TreeNode

// Returns the next in-order node of the tree, or %NULL
// if the passed node was already the last one.
func (x *TreeNode) Next() *TreeNode {

	cret := xTreeNodeNext(x.GoPointer())
	return cret
}

var xTreeNodePrevious func(uintptr) *TreeNode

// Returns the previous in-order node of the tree, or %NULL
// if the passed node was already the first one.
func (x *TreeNode) Previous() *TreeNode {

	cret := xTreeNodePrevious(x.GoPointer())
	return cret
}

var xTreeNodeValue func(uintptr) uintptr

// Gets the value stored at a particular tree node.
func (x *TreeNode) Value() uintptr {

	cret := xTreeNodeValue(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTree, lib, "g_tree_new")
	core.PuregoSafeRegister(&xNewTreeFull, lib, "g_tree_new_full")
	core.PuregoSafeRegister(&xNewTreeWithData, lib, "g_tree_new_with_data")

	core.PuregoSafeRegister(&xTreeDestroy, lib, "g_tree_destroy")
	core.PuregoSafeRegister(&xTreeForeach, lib, "g_tree_foreach")
	core.PuregoSafeRegister(&xTreeForeachNode, lib, "g_tree_foreach_node")
	core.PuregoSafeRegister(&xTreeHeight, lib, "g_tree_height")
	core.PuregoSafeRegister(&xTreeInsert, lib, "g_tree_insert")
	core.PuregoSafeRegister(&xTreeInsertNode, lib, "g_tree_insert_node")
	core.PuregoSafeRegister(&xTreeLookup, lib, "g_tree_lookup")
	core.PuregoSafeRegister(&xTreeLookupExtended, lib, "g_tree_lookup_extended")
	core.PuregoSafeRegister(&xTreeLookupNode, lib, "g_tree_lookup_node")
	core.PuregoSafeRegister(&xTreeLowerBound, lib, "g_tree_lower_bound")
	core.PuregoSafeRegister(&xTreeNnodes, lib, "g_tree_nnodes")
	core.PuregoSafeRegister(&xTreeNodeFirst, lib, "g_tree_node_first")
	core.PuregoSafeRegister(&xTreeNodeLast, lib, "g_tree_node_last")
	core.PuregoSafeRegister(&xTreeRef, lib, "g_tree_ref")
	core.PuregoSafeRegister(&xTreeRemove, lib, "g_tree_remove")
	core.PuregoSafeRegister(&xTreeRemoveAll, lib, "g_tree_remove_all")
	core.PuregoSafeRegister(&xTreeReplace, lib, "g_tree_replace")
	core.PuregoSafeRegister(&xTreeReplaceNode, lib, "g_tree_replace_node")
	core.PuregoSafeRegister(&xTreeSearch, lib, "g_tree_search")
	core.PuregoSafeRegister(&xTreeSearchNode, lib, "g_tree_search_node")
	core.PuregoSafeRegister(&xTreeSteal, lib, "g_tree_steal")
	core.PuregoSafeRegister(&xTreeTraverse, lib, "g_tree_traverse")
	core.PuregoSafeRegister(&xTreeUnref, lib, "g_tree_unref")
	core.PuregoSafeRegister(&xTreeUpperBound, lib, "g_tree_upper_bound")

	core.PuregoSafeRegister(&xTreeNodeKey, lib, "g_tree_node_key")
	core.PuregoSafeRegister(&xTreeNodeNext, lib, "g_tree_node_next")
	core.PuregoSafeRegister(&xTreeNodePrevious, lib, "g_tree_node_previous")
	core.PuregoSafeRegister(&xTreeNodeValue, lib, "g_tree_node_value")

}
