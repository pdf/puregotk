// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// A parse context is used to parse a stream of bytes that
// you expect to contain marked-up text.
//
// See g_markup_parse_context_new(), #GMarkupParser, and so
// on for more details.
type MarkupParseContext struct {
}

// Any of the fields in #GMarkupParser can be %NULL, in which case they
// will be ignored. Except for the @error function, any of these callbacks
// can set an error; in particular the %G_MARKUP_ERROR_UNKNOWN_ELEMENT,
// %G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE, and %G_MARKUP_ERROR_INVALID_CONTENT
// errors are intended to be set from these callbacks. If you set an error
// from a callback, g_markup_parse_context_parse() will report that error
// back to its caller.
type MarkupParser struct {
}

// A mixed enumerated type and flags field. You must specify one type
// (string, strdup, boolean, tristate).  Additionally, you may  optionally
// bitwise OR the type with the flag %G_MARKUP_COLLECT_OPTIONAL.
//
// It is likely that this enum will be extended in the future to
// support other types.
type MarkupCollectType int

const (

	// used to terminate the list of attributes
	//     to collect
	GMarkupCollectInvalidValue MarkupCollectType = 0
	// collect the string pointer directly from
	//     the attribute_values[] array. Expects a parameter of type (const
	//     char **). If %G_MARKUP_COLLECT_OPTIONAL is specified and the
	//     attribute isn't present then the pointer will be set to %NULL
	GMarkupCollectStringValue MarkupCollectType = 1
	// as with %G_MARKUP_COLLECT_STRING, but
	//     expects a parameter of type (char **) and g_strdup()s the
	//     returned pointer. The pointer must be freed with g_free()
	GMarkupCollectStrdupValue MarkupCollectType = 2
	// expects a parameter of type (gboolean *)
	//     and parses the attribute value as a boolean. Sets %FALSE if the
	//     attribute isn't present. Valid boolean values consist of
	//     (case-insensitive) "false", "f", "no", "n", "0" and "true", "t",
	//     "yes", "y", "1"
	GMarkupCollectBooleanValue MarkupCollectType = 3
	// as with %G_MARKUP_COLLECT_BOOLEAN, but
	//     in the case of a missing attribute a value is set that compares
	//     equal to neither %FALSE nor %TRUE G_MARKUP_COLLECT_OPTIONAL is
	//     implied
	GMarkupCollectTristateValue MarkupCollectType = 4
	// can be bitwise ORed with the other fields.
	//     If present, allows the attribute not to appear. A default value
	//     is set depending on what value type is used
	GMarkupCollectOptionalValue MarkupCollectType = 65536
)

// Flags that affect the behaviour of the parser.
type MarkupParseFlags int

const (

	// flag you should not use
	GMarkupDoNotUseThisUnsupportedFlagValue MarkupParseFlags = 1
	// When this flag is set, CDATA marked
	//     sections are not passed literally to the @passthrough function of
	//     the parser. Instead, the content of the section (without the
	//     `&lt;![CDATA[` and `]]&gt;`) is
	//     passed to the @text function. This flag was added in GLib 2.12
	GMarkupTreatCdataAsTextValue MarkupParseFlags = 2
	// Normally errors caught by GMarkup
	//     itself have line/column information prefixed to them to let the
	//     caller know the location of the error. When this flag is set the
	//     location information is also prefixed to errors generated by the
	//     #GMarkupParser implementation functions
	GMarkupPrefixErrorPositionValue MarkupParseFlags = 4
	// Ignore (don't report) qualified
	//     attributes and tags, along with their contents.  A qualified
	//     attribute or tag is one that contains ':' in its name (ie: is in
	//     another namespace).  Since: 2.40.
	GMarkupIgnoreQualifiedValue MarkupParseFlags = 8
)

// Error codes returned by markup parsing.
type MarkupError int

const (

	// text being parsed was not valid UTF-8
	GMarkupErrorBadUtf8Value MarkupError = 0
	// document contained nothing, or only whitespace
	GMarkupErrorEmptyValue MarkupError = 1
	// document was ill-formed
	GMarkupErrorParseValue MarkupError = 2
	// error should be set by #GMarkupParser
	//     functions; element wasn't known
	GMarkupErrorUnknownElementValue MarkupError = 3
	// error should be set by #GMarkupParser
	//     functions; attribute wasn't known
	GMarkupErrorUnknownAttributeValue MarkupError = 4
	// error should be set by #GMarkupParser
	//     functions; content was invalid
	GMarkupErrorInvalidContentValue MarkupError = 5
	// error should be set by #GMarkupParser
	//     functions; a required attribute was missing
	GMarkupErrorMissingAttributeValue MarkupError = 6
)

var xMarkupCollectAttributes func(string, string, string, **Error, MarkupCollectType, string, ...interface{}) bool

// Collects the attributes of the element from the data passed to the
// #GMarkupParser start_element function, dealing with common error
// conditions and supporting boolean values.
//
// This utility function is not required to write a parser but can save
// a lot of typing.
//
// The @element_name, @attribute_names, @attribute_values and @error
// parameters passed to the start_element callback should be passed
// unmodified to this function.
//
// Following these arguments is a list of "supported" attributes to collect.
// It is an error to specify multiple attributes with the same name. If any
// attribute not in the list appears in the @attribute_names array then an
// unknown attribute error will result.
//
// The #GMarkupCollectType field allows specifying the type of collection
// to perform and if a given attribute must appear or is optional.
//
// The attribute name is simply the name of the attribute to collect.
//
// The pointer should be of the appropriate type (see the descriptions
// under #GMarkupCollectType) and may be %NULL in case a particular
// attribute is to be allowed but ignored.
//
// This function deals with issuing errors for missing attributes
// (of type %G_MARKUP_ERROR_MISSING_ATTRIBUTE), unknown attributes
// (of type %G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE) and duplicate
// attributes (of type %G_MARKUP_ERROR_INVALID_CONTENT) as well
// as parse errors for boolean-valued attributes (again of type
// %G_MARKUP_ERROR_INVALID_CONTENT). In all of these cases %FALSE
// will be returned and @error will be set as appropriate.
func MarkupCollectAttributes(ElementNameVar string, AttributeNamesVar string, AttributeValuesVar string, ErrorVar **Error, FirstTypeVar MarkupCollectType, FirstAttrVar string, varArgs ...interface{}) bool {

	cret := xMarkupCollectAttributes(ElementNameVar, AttributeNamesVar, AttributeValuesVar, ErrorVar, FirstTypeVar, FirstAttrVar, varArgs...)
	return cret
}

var xMarkupEscapeText func(string, int) string

// Escapes text so that the markup parser will parse it verbatim.
// Less than, greater than, ampersand, etc. are replaced with the
// corresponding entities. This function would typically be used
// when writing out a file to be parsed with the markup parser.
//
// Note that this function doesn't protect whitespace and line endings
// from being processed according to the XML rules for normalization
// of line endings and attribute values.
//
// Note also that this function will produce character references in
// the range of &amp;#x1; ... &amp;#x1f; for all control sequences
// except for tabstop, newline and carriage return.  The character
// references in this range are not valid XML 1.0, but they are
// valid XML 1.1 and will be accepted by the GMarkup parser.
func MarkupEscapeText(TextVar string, LengthVar int) string {

	cret := xMarkupEscapeText(TextVar, LengthVar)
	return cret
}

var xMarkupPrintfEscaped func(string, ...interface{}) string

// Formats arguments according to @format, escaping
// all string and character arguments in the fashion
// of g_markup_escape_text(). This is useful when you
// want to insert literal strings into XML-style markup
// output, without having to worry that the strings
// might themselves contain markup.
//
// |[&lt;!-- language="C" --&gt;
// const char *store = "Fortnum &amp; Mason";
// const char *item = "Tea";
// char *output;
//
// output = g_markup_printf_escaped ("&lt;purchase&gt;"
//
//	"&lt;store&gt;%s&lt;/store&gt;"
//	"&lt;item&gt;%s&lt;/item&gt;"
//	"&lt;/purchase&gt;",
//	store, item);
//
// ]|
func MarkupPrintfEscaped(FormatVar string, varArgs ...interface{}) string {

	cret := xMarkupPrintfEscaped(FormatVar, varArgs...)
	return cret
}

var xMarkupVprintfEscaped func(string, []interface{}) string

// Formats the data in @args according to @format, escaping
// all string and character arguments in the fashion
// of g_markup_escape_text(). See g_markup_printf_escaped().
func MarkupVprintfEscaped(FormatVar string, ArgsVar []interface{}) string {

	cret := xMarkupVprintfEscaped(FormatVar, ArgsVar)
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xMarkupCollectAttributes, lib, "g_markup_collect_attributes")
	core.PuregoSafeRegister(&xMarkupEscapeText, lib, "g_markup_escape_text")
	core.PuregoSafeRegister(&xMarkupPrintfEscaped, lib, "g_markup_printf_escaped")
	core.PuregoSafeRegister(&xMarkupVprintfEscaped, lib, "g_markup_vprintf_escaped")

}
