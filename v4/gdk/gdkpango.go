// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/cairo"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

var xPangoLayoutGetClipRegion func(uintptr, int, int, int, int) *cairo.Region

// Obtains a clip region which contains the areas where the given ranges
// of text would be drawn.
//
// @x_origin and @y_origin are the top left point to center the layout.
// @index_ranges should contain ranges of bytes in the layout’s text.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn layout may in fact touch areas out of
// the clip region.  The clip region is mainly useful for highlightling parts
// of text, such as when text is selected.
func PangoLayoutGetClipRegion(LayoutVar *pango.Layout, XOriginVar int, YOriginVar int, IndexRangesVar int, NRangesVar int) *cairo.Region {

	return xPangoLayoutGetClipRegion(LayoutVar.GoPointer(), XOriginVar, YOriginVar, IndexRangesVar, NRangesVar)

}

var xPangoLayoutLineGetClipRegion func(*pango.LayoutLine, int, int, uintptr, int) *cairo.Region

// Obtains a clip region which contains the areas where the given
// ranges of text would be drawn.
//
// @x_origin and @y_origin are the top left position of the layout.
// @index_ranges should contain ranges of bytes in the layout’s text.
// The clip region will include space to the left or right of the line
// (to the layout bounding box) if you have indexes above or below the
// indexes contained inside the line. This is to draw the selection all
// the way to the side of the layout. However, the clip region is in line
// coordinates, not layout coordinates.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn line may in fact touch areas out of
// the clip region.  The clip region is mainly useful for highlightling parts
// of text, such as when text is selected.
func PangoLayoutLineGetClipRegion(LineVar *pango.LayoutLine, XOriginVar int, YOriginVar int, IndexRangesVar uintptr, NRangesVar int) *cairo.Region {

	return xPangoLayoutLineGetClipRegion(LineVar, XOriginVar, YOriginVar, IndexRangesVar, NRangesVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xPangoLayoutGetClipRegion, lib, "gdk_pango_layout_get_clip_region")
	core.PuregoSafeRegister(&xPangoLayoutLineGetClipRegion, lib, "gdk_pango_layout_line_get_clip_region")

}
