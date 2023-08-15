// Package pango was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package pango

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// The `PangoLogAttr` structure stores information about the attributes of a
// single character.
type LogAttr struct {
	IsLineBreak uint

	IsMandatoryBreak uint

	IsCharBreak uint

	IsWhite uint

	IsCursorPosition uint

	IsWordStart uint

	IsWordEnd uint

	IsSentenceBoundary uint

	IsSentenceStart uint

	IsSentenceEnd uint

	BackspaceDeletesCharacter uint

	IsExpandableSpace uint

	IsWordBoundary uint

	BreakInsertsHyphen uint

	BreakRemovesPreceding uint

	Reserved uint
}

var xAttrBreak func(string, int, *AttrList, int, uintptr, int)

// Apply customization from attributes to the breaks in @attrs.
//
// The line breaks are assumed to have been produced
// by [func@Pango.default_break] and [func@Pango.tailor_break].
func AttrBreak(TextVar string, LengthVar int, AttrListVar *AttrList, OffsetVar int, AttrsVar uintptr, AttrsLenVar int) {

	xAttrBreak(TextVar, LengthVar, AttrListVar, OffsetVar, AttrsVar, AttrsLenVar)

}

var xBreak func(string, int, *Analysis, uintptr, int)

// Determines possible line, word, and character breaks
// for a string of Unicode text with a single analysis.
//
// For most purposes you may want to use [func@Pango.get_log_attrs].
func Break(TextVar string, LengthVar int, AnalysisVar *Analysis, AttrsVar uintptr, AttrsLenVar int) {

	xBreak(TextVar, LengthVar, AnalysisVar, AttrsVar, AttrsLenVar)

}

var xDefaultBreak func(string, int, *Analysis, *LogAttr, int)

// This is the default break algorithm.
//
// It applies rules from the [Unicode Line Breaking Algorithm](http://www.unicode.org/unicode/reports/tr14/)
// without language-specific tailoring, therefore the @analyis argument is unused
// and can be %NULL.
//
// See [func@Pango.tailor_break] for language-specific breaks.
//
// See [func@Pango.attr_break] for attribute-based customization.
func DefaultBreak(TextVar string, LengthVar int, AnalysisVar *Analysis, AttrsVar *LogAttr, AttrsLenVar int) {

	xDefaultBreak(TextVar, LengthVar, AnalysisVar, AttrsVar, AttrsLenVar)

}

var xGetLogAttrs func(string, int, int, *Language, uintptr, int)

// Computes a `PangoLogAttr` for each character in @text.
//
// The @attrs array must have one `PangoLogAttr` for
// each position in @text; if @text contains N characters,
// it has N+1 positions, including the last position at the
// end of the text. @text should be an entire paragraph;
// logical attributes can't be computed without context
// (for example you need to see spaces on either side of
// a word to know the word is a word).
func GetLogAttrs(TextVar string, LengthVar int, LevelVar int, LanguageVar *Language, AttrsVar uintptr, AttrsLenVar int) {

	xGetLogAttrs(TextVar, LengthVar, LevelVar, LanguageVar, AttrsVar, AttrsLenVar)

}

var xTailorBreak func(string, int, *Analysis, int, uintptr, int)

// Apply language-specific tailoring to the breaks in @attrs.
//
// The line breaks are assumed to have been produced by [func@Pango.default_break].
//
// If @offset is not -1, it is used to apply attributes from @analysis that are
// relevant to line breaking.
//
// Note that it is better to pass -1 for @offset and use [func@Pango.attr_break]
// to apply attributes to the whole paragraph.
func TailorBreak(TextVar string, LengthVar int, AnalysisVar *Analysis, OffsetVar int, AttrsVar uintptr, AttrsLenVar int) {

	xTailorBreak(TextVar, LengthVar, AnalysisVar, OffsetVar, AttrsVar, AttrsLenVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("PANGO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xAttrBreak, lib, "pango_attr_break")
	core.PuregoSafeRegister(&xBreak, lib, "pango_break")
	core.PuregoSafeRegister(&xDefaultBreak, lib, "pango_default_break")
	core.PuregoSafeRegister(&xGetLogAttrs, lib, "pango_get_log_attrs")
	core.PuregoSafeRegister(&xTailorBreak, lib, "pango_tailor_break")

}
