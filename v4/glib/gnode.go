// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

// Specifies the type of function passed to g_node_children_foreach().
// The function is called with each child node, together with the user
// data passed to g_node_children_foreach().
type NodeForeachFunc func(*Node, uintptr)

// Specifies the type of function passed to g_node_traverse(). The
// function is called with each of the nodes visited, together with the
// user data passed to g_node_traverse(). If the function returns
// %TRUE, then the traversal is stopped.
type NodeTraverseFunc func(*Node, uintptr) bool

// The #GNode struct represents one node in a [n-ary tree][glib-N-ary-Trees].
type Node struct {
	Data uintptr

	Next *Node

	Prev *Node

	Parent *Node

	Children *Node
}

// Specifies which nodes are visited during several of the tree
// functions, including g_node_traverse() and g_node_find().
type TraverseFlags int

const (

	// only leaf nodes should be visited. This name has
	//                     been introduced in 2.6, for older version use
	//                     %G_TRAVERSE_LEAFS.
	GTraverseLeavesValue TraverseFlags = 1
	// only non-leaf nodes should be visited. This
	//                         name has been introduced in 2.6, for older
	//                         version use %G_TRAVERSE_NON_LEAFS.
	GTraverseNonLeavesValue TraverseFlags = 2
	// all nodes should be visited.
	GTraverseAllValue TraverseFlags = 3
	// a mask of all traverse flags.
	GTraverseMaskValue TraverseFlags = 3
	// identical to %G_TRAVERSE_LEAVES.
	GTraverseLeafsValue TraverseFlags = 1
	// identical to %G_TRAVERSE_NON_LEAVES.
	GTraverseNonLeafsValue TraverseFlags = 2
)

// Specifies the type of traversal performed by g_tree_traverse(),
// g_node_traverse() and g_node_find(). The different orders are
// illustrated here:
//   - In order: A, B, C, D, E, F, G, H, I
//     ![](Sorted_binary_tree_inorder.svg)
//   - Pre order: F, B, A, D, C, E, G, I, H
//     ![](Sorted_binary_tree_preorder.svg)
//   - Post order: A, C, E, D, B, H, I, G, F
//     ![](Sorted_binary_tree_postorder.svg)
//   - Level order: F, B, G, A, D, I, C, E, H
//     ![](Sorted_binary_tree_breadth-first_traversal.svg)
type TraverseType int

const (

	// vists a node's left child first, then the node itself,
	//              then its right child. This is the one to use if you
	//              want the output sorted according to the compare
	//              function.
	GInOrderValue TraverseType = 0
	// visits a node, then its children.
	GPreOrderValue TraverseType = 1
	// visits the node's children, then the node itself.
	GPostOrderValue TraverseType = 2
	// is not implemented for
	//              [balanced binary trees][glib-Balanced-Binary-Trees].
	//              For [n-ary trees][glib-N-ary-Trees], it
	//              vists the root node first, then its children, then
	//              its grandchildren, and so on. Note that this is less
	//              efficient than the other orders.
	GLevelOrderValue TraverseType = 3
)
