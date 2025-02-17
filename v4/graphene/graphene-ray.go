// Package graphene was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package graphene

// A ray emitted from an origin in a given direction.
//
// The contents of the `graphene_ray_t` structure are private, and should not
// be modified directly.
type Ray struct {
	Origin uintptr

	Direction uintptr
}

// The type of intersection.
type RayIntersectionKind int

const (

	// No intersection
	RayIntersectionKindNoneValue RayIntersectionKind = 0
	// The ray is entering the intersected
	//   object
	RayIntersectionKindEnterValue RayIntersectionKind = 1
	// The ray is leaving the intersected
	//   object
	RayIntersectionKindLeaveValue RayIntersectionKind = 2
)
