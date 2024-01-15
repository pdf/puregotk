// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/pango"
)

// The `GtkLabel` widget displays a small amount of text.
//
// As the name implies, most labels are used to label another widget
// such as a [class@Button].
//
// ![An example GtkLabel](label.png)
//
// # CSS nodes
//
// ```
// label
// ├── [selection]
// ├── [link]
// ┊
// ╰── [link]
// ```
//
// `GtkLabel` has a single CSS node with the name label. A wide variety
// of style classes may be applied to labels, such as .title, .subtitle,
// .dim-label, etc. In the `GtkShortcutsWindow`, labels are used with the
// .keycap style class.
//
// If the label has a selection, it gets a subnode with name selection.
//
// If the label has links, there is one subnode per link. These subnodes
// carry the link or visited state depending on whether they have been
// visited. In this case, label node also gets a .link style class.
//
// # GtkLabel as GtkBuildable
//
// The GtkLabel implementation of the GtkBuildable interface supports a
// custom &lt;attributes&gt; element, which supports any number of &lt;attribute&gt;
// elements. The &lt;attribute&gt; element has attributes named “name“, “value“,
// “start“ and “end“ and allows you to specify [struct@Pango.Attribute]
// values for this label.
//
// An example of a UI definition fragment specifying Pango attributes:
// ```xml
// &lt;object class="GtkLabel"&gt;
//
//	&lt;attributes&gt;
//	  &lt;attribute name="weight" value="PANGO_WEIGHT_BOLD"/&gt;
//	  &lt;attribute name="background" value="red" start="5" end="10"/&gt;
//	&lt;/attributes&gt;
//
// &lt;/object&gt;
// ```
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
// # Accessibility
//
// `GtkLabel` uses the %GTK_ACCESSIBLE_ROLE_LABEL role.
//
// # Mnemonics
//
// Labels may contain “mnemonics”. Mnemonics are underlined characters in the
// label, used for keyboard navigation. Mnemonics are created by providing a
// string with an underscore before the mnemonic character, such as `"_File"`,
// to the functions [ctor@Gtk.Label.new_with_mnemonic] or
// [method@Gtk.Label.set_text_with_mnemonic].
//
// Mnemonics automatically activate any activatable widget the label is
// inside, such as a [class@Gtk.Button]; if the label is not inside the
// mnemonic’s target widget, you have to tell the label about the target
// using [class@Gtk.Label.set_mnemonic_widget]. Here’s a simple example where
// the label is inside a button:
//
// ```c
// // Pressing Alt+H will activate this button
// GtkWidget *button = gtk_button_new ();
// GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
// gtk_button_set_child (GTK_BUTTON (button), label);
// ```
//
// There’s a convenience function to create buttons with a mnemonic label
// already inside:
//
// ```c
// // Pressing Alt+H will activate this button
// GtkWidget *button = gtk_button_new_with_mnemonic ("_Hello");
// ```
//
// To create a mnemonic for a widget alongside the label, such as a
// [class@Gtk.Entry], you have to point the label at the entry with
// [method@Gtk.Label.set_mnemonic_widget]:
//
// ```c
// // Pressing Alt+H will focus the entry
// GtkWidget *entry = gtk_entry_new ();
// GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
// gtk_label_set_mnemonic_widget (GTK_LABEL (label), entry);
// ```
//
// # Markup (styled text)
//
// To make it easy to format text in a label (changing colors,
// fonts, etc.), label text can be provided in a simple
// markup format:
//
// Here’s how to create a label with a small font:
// ```c
// GtkWidget *label = gtk_label_new (NULL);
// gtk_label_set_markup (GTK_LABEL (label), "&lt;small&gt;Small text&lt;/small&gt;");
// ```
//
// (See the Pango manual for complete documentation] of available
// tags, [func@Pango.parse_markup])
//
// The markup passed to gtk_label_set_markup() must be valid; for example,
// literal &lt;, &gt; and &amp; characters must be escaped as &amp;lt;, &amp;gt;, and &amp;amp;.
// If you pass text obtained from the user, file, or a network to
// [method@Gtk.Label.set_markup], you’ll want to escape it with
// g_markup_escape_text() or g_markup_printf_escaped().
//
// Markup strings are just a convenient way to set the [struct@Pango.AttrList]
// on a label; [method@Gtk.Label.set_attributes] may be a simpler way to set
// attributes in some cases. Be careful though; [struct@Pango.AttrList] tends
// to cause internationalization problems, unless you’re applying attributes
// to the entire string (i.e. unless you set the range of each attribute
// to [0, %G_MAXINT)). The reason is that specifying the start_index and
// end_index for a [struct@Pango.Attribute] requires knowledge of the exact
// string being displayed, so translations will cause problems.
//
// # Selectable labels
//
// Labels can be made selectable with [method@Gtk.Label.set_selectable].
// Selectable labels allow the user to copy the label contents to
// the clipboard. Only labels that contain useful-to-copy information
// — such as error messages — should be made selectable.
//
// # Text layout
//
// A label can contain any number of paragraphs, but will have
// performance problems if it contains more than a small number.
// Paragraphs are separated by newlines or other paragraph separators
// understood by Pango.
//
// Labels can automatically wrap text if you call [method@Gtk.Label.set_wrap].
//
// [method@Gtk.Label.set_justify] sets how the lines in a label align
// with one another. If you want to set how the label as a whole aligns
// in its available space, see the [property@Gtk.Widget:halign] and
// [property@Gtk.Widget:valign] properties.
//
// The [property@Gtk.Label:width-chars] and [property@Gtk.Label:max-width-chars]
// properties can be used to control the size allocation of ellipsized or
// wrapped labels. For ellipsizing labels, if either is specified (and less
// than the actual text size), it is used as the minimum width, and the actual
// text size is used as the natural width of the label. For wrapping labels,
// width-chars is used as the minimum width, if specified, and max-width-chars
// is used as the natural width. Even if max-width-chars specified, wrapping
// labels will be rewrapped to use all of the available width.
//
// # Links
//
// GTK supports markup for clickable hyperlinks in addition to regular Pango
// markup. The markup for links is borrowed from HTML, using the `&lt;a&gt;` with
// “href“, “title“ and “class“ attributes. GTK renders links similar to the
// way they appear in web browsers, with colored, underlined text. The “title“
// attribute is displayed as a tooltip on the link. The “class“ attribute is
// used as style class on the CSS node for the link.
//
// An example looks like this:
//
// ```c
// const char *text =
// "Go to the"
// "&lt;a href=\"http://www.gtk.org title=\"&amp;lt;i&amp;gt;Our&amp;lt;/i&amp;gt; website\"&gt;"
// "GTK website&lt;/a&gt; for more...";
// GtkWidget *label = gtk_label_new (NULL);
// gtk_label_set_markup (GTK_LABEL (label), text);
// ```
//
// It is possible to implement custom handling for links and their tooltips
// with the [signal@Gtk.Label::activate-link] signal and the
// [method@Gtk.Label.get_current_uri] function.
type Label struct {
	Widget
}

func LabelNewFromInternalPtr(ptr uintptr) *Label {
	cls := &Label{}
	cls.Ptr = ptr
	return cls
}

var xNewLabel func(string) uintptr

// Creates a new label with the given text inside it.
//
// You can pass %NULL to get an empty label widget.
func NewLabel(StrVar string) *Label {
	var cls *Label

	cret := xNewLabel(StrVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Label{}
	cls.Ptr = cret
	return cls
}

var xNewLabelWithMnemonic func(string) uintptr

// Creates a new `GtkLabel`, containing the text in @str.
//
// If characters in @str are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use
// '__' (two underscores). The first underlined character represents a
// keyboard accelerator called a mnemonic. The mnemonic key can be used
// to activate another widget, chosen automatically, or explicitly using
// [method@Gtk.Label.set_mnemonic_widget].
//
// If [method@Gtk.Label.set_mnemonic_widget] is not called, then the first
// activatable ancestor of the `GtkLabel` will be chosen as the mnemonic
// widget. For instance, if the label is inside a button or menu item,
// the button or menu item will automatically become the mnemonic widget
// and be activated by the mnemonic.
func NewLabelWithMnemonic(StrVar string) *Label {
	var cls *Label

	cret := xNewLabelWithMnemonic(StrVar)

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Label{}
	cls.Ptr = cret
	return cls
}

var xLabelGetAttributes func(uintptr) *pango.AttrList

// Gets the labels attribute list.
//
// This is the [struct@Pango.AttrList] that was set on the label using
// [method@Gtk.Label.set_attributes], if any. This function does not
// reflect attributes that come from the labels markup (see
// [method@Gtk.Label.set_markup]). If you want to get the effective
// attributes for the label, use
// `pango_layout_get_attribute (gtk_label_get_layout (self))`.
func (x *Label) GetAttributes() *pango.AttrList {

	cret := xLabelGetAttributes(x.GoPointer())
	return cret
}

var xLabelGetCurrentUri func(uintptr) string

// Returns the URI for the currently active link in the label.
//
// The active link is the one under the mouse pointer or, in a
// selectable label, the link in which the text cursor is currently
// positioned.
//
// This function is intended for use in a [signal@Gtk.Label::activate-link]
// handler or for use in a [signal@Gtk.Widget::query-tooltip] handler.
func (x *Label) GetCurrentUri() string {

	cret := xLabelGetCurrentUri(x.GoPointer())
	return cret
}

var xLabelGetEllipsize func(uintptr) pango.EllipsizeMode

// Returns the ellipsizing position of the label.
//
// See [method@Gtk.Label.set_ellipsize].
func (x *Label) GetEllipsize() pango.EllipsizeMode {

	cret := xLabelGetEllipsize(x.GoPointer())
	return cret
}

var xLabelGetExtraMenu func(uintptr) uintptr

// Gets the extra menu model of @label.
//
// See [method@Gtk.Label.set_extra_menu].
func (x *Label) GetExtraMenu() *gio.MenuModel {
	var cls *gio.MenuModel

	cret := xLabelGetExtraMenu(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &gio.MenuModel{}
	cls.Ptr = cret
	return cls
}

var xLabelGetJustify func(uintptr) Justification

// Returns the justification of the label.
//
// See [method@Gtk.Label.set_justify].
func (x *Label) GetJustify() Justification {

	cret := xLabelGetJustify(x.GoPointer())
	return cret
}

var xLabelGetLabel func(uintptr) string

// Fetches the text from a label.
//
// The returned text includes any embedded underlines indicating
// mnemonics and Pango markup. (See [method@Gtk.Label.get_text]).
func (x *Label) GetLabel() string {

	cret := xLabelGetLabel(x.GoPointer())
	return cret
}

var xLabelGetLayout func(uintptr) uintptr

// Gets the `PangoLayout` used to display the label.
//
// The layout is useful to e.g. convert text positions to pixel
// positions, in combination with [method@Gtk.Label.get_layout_offsets].
// The returned layout is owned by the @label so need not be
// freed by the caller. The @label is free to recreate its layout
// at any time, so it should be considered read-only.
func (x *Label) GetLayout() *pango.Layout {
	var cls *pango.Layout

	cret := xLabelGetLayout(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &pango.Layout{}
	cls.Ptr = cret
	return cls
}

var xLabelGetLayoutOffsets func(uintptr, int, int)

// Obtains the coordinates where the label will draw its `PangoLayout`.
//
// The coordinates are useful to convert mouse events into coordinates
// inside the [class@Pango.Layout], e.g. to take some action if some part
// of the label is clicked. Remember when using the [class@Pango.Layout]
// functions you need to convert to and from pixels using PANGO_PIXELS()
// or [const@Pango.SCALE].
func (x *Label) GetLayoutOffsets(XVar int, YVar int) {

	xLabelGetLayoutOffsets(x.GoPointer(), XVar, YVar)

}

var xLabelGetLines func(uintptr) int

// Gets the number of lines to which an ellipsized, wrapping
// label should be limited.
//
// See [method@Gtk.Label.set_lines].
func (x *Label) GetLines() int {

	cret := xLabelGetLines(x.GoPointer())
	return cret
}

var xLabelGetMaxWidthChars func(uintptr) int

// Retrieves the desired maximum width of @label, in characters.
//
// See [method@Gtk.Label.set_width_chars].
func (x *Label) GetMaxWidthChars() int {

	cret := xLabelGetMaxWidthChars(x.GoPointer())
	return cret
}

var xLabelGetMnemonicKeyval func(uintptr) uint

// Return the mnemonic accelerator.
//
// If the label has been set so that it has a mnemonic key this function
// returns the keyval used for the mnemonic accelerator. If there is no
// mnemonic set up it returns `GDK_KEY_VoidSymbol`.
func (x *Label) GetMnemonicKeyval() uint {

	cret := xLabelGetMnemonicKeyval(x.GoPointer())
	return cret
}

var xLabelGetMnemonicWidget func(uintptr) uintptr

// Retrieves the target of the mnemonic (keyboard shortcut) of this
// label.
//
// See [method@Gtk.Label.set_mnemonic_widget].
func (x *Label) GetMnemonicWidget() *Widget {
	var cls *Widget

	cret := xLabelGetMnemonicWidget(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &Widget{}
	cls.Ptr = cret
	return cls
}

var xLabelGetNaturalWrapMode func(uintptr) NaturalWrapMode

// Returns line wrap mode used by the label.
//
// See [method@Gtk.Label.set_natural_wrap_mode].
func (x *Label) GetNaturalWrapMode() NaturalWrapMode {

	cret := xLabelGetNaturalWrapMode(x.GoPointer())
	return cret
}

var xLabelGetSelectable func(uintptr) bool

// Returns whether the label is selectable.
func (x *Label) GetSelectable() bool {

	cret := xLabelGetSelectable(x.GoPointer())
	return cret
}

var xLabelGetSelectionBounds func(uintptr, int, int) bool

// Gets the selected range of characters in the label.
func (x *Label) GetSelectionBounds(StartVar int, EndVar int) bool {

	cret := xLabelGetSelectionBounds(x.GoPointer(), StartVar, EndVar)
	return cret
}

var xLabelGetSingleLineMode func(uintptr) bool

// Returns whether the label is in single line mode.
func (x *Label) GetSingleLineMode() bool {

	cret := xLabelGetSingleLineMode(x.GoPointer())
	return cret
}

var xLabelGetTabs func(uintptr) *pango.TabArray

// Gets the tabs for @self.
//
// The returned array will be %NULL if “standard” (8-space) tabs are used.
// Free the return value with [method@Pango.TabArray.free].
func (x *Label) GetTabs() *pango.TabArray {

	cret := xLabelGetTabs(x.GoPointer())
	return cret
}

var xLabelGetText func(uintptr) string

// Fetches the text from a label.
//
// The returned text is as it appears on screen. This does not include
// any embedded underlines indicating mnemonics or Pango markup. (See
// [method@Gtk.Label.get_label])
func (x *Label) GetText() string {

	cret := xLabelGetText(x.GoPointer())
	return cret
}

var xLabelGetUseMarkup func(uintptr) bool

// Returns whether the label’s text is interpreted as Pango markup.
//
// See [method@Gtk.Label.set_use_markup].
func (x *Label) GetUseMarkup() bool {

	cret := xLabelGetUseMarkup(x.GoPointer())
	return cret
}

var xLabelGetUseUnderline func(uintptr) bool

// Returns whether an embedded underlines in the label indicate mnemonics.
//
// See [method@Gtk.Label.set_use_underline].
func (x *Label) GetUseUnderline() bool {

	cret := xLabelGetUseUnderline(x.GoPointer())
	return cret
}

var xLabelGetWidthChars func(uintptr) int

// Retrieves the desired width of @label, in characters.
//
// See [method@Gtk.Label.set_width_chars].
func (x *Label) GetWidthChars() int {

	cret := xLabelGetWidthChars(x.GoPointer())
	return cret
}

var xLabelGetWrap func(uintptr) bool

// Returns whether lines in the label are automatically wrapped.
//
// See [method@Gtk.Label.set_wrap].
func (x *Label) GetWrap() bool {

	cret := xLabelGetWrap(x.GoPointer())
	return cret
}

var xLabelGetWrapMode func(uintptr) pango.WrapMode

// Returns line wrap mode used by the label.
//
// See [method@Gtk.Label.set_wrap_mode].
func (x *Label) GetWrapMode() pango.WrapMode {

	cret := xLabelGetWrapMode(x.GoPointer())
	return cret
}

var xLabelGetXalign func(uintptr) float32

// Gets the `xalign` of the label.
//
// See the [property@Gtk.Label:xalign] property.
func (x *Label) GetXalign() float32 {

	cret := xLabelGetXalign(x.GoPointer())
	return cret
}

var xLabelGetYalign func(uintptr) float32

// Gets the `yalign` of the label.
//
// See the [property@Gtk.Label:yalign] property.
func (x *Label) GetYalign() float32 {

	cret := xLabelGetYalign(x.GoPointer())
	return cret
}

var xLabelSelectRegion func(uintptr, int, int)

// Selects a range of characters in the label, if the label is selectable.
//
// See [method@Gtk.Label.set_selectable]. If the label is not selectable,
// this function has no effect. If @start_offset or
// @end_offset are -1, then the end of the label will be substituted.
func (x *Label) SelectRegion(StartOffsetVar int, EndOffsetVar int) {

	xLabelSelectRegion(x.GoPointer(), StartOffsetVar, EndOffsetVar)

}

var xLabelSetAttributes func(uintptr, *pango.AttrList)

// Apply attributes to the label text.
//
// The attributes set with this function will be applied and merged with
// any other attributes previously effected by way of the
// [property@Gtk.Label:use-underline] or [property@Gtk.Label:use-markup]
// properties. While it is not recommended to mix markup strings with
// manually set attributes, if you must; know that the attributes will
// be applied to the label after the markup string is parsed.
func (x *Label) SetAttributes(AttrsVar *pango.AttrList) {

	xLabelSetAttributes(x.GoPointer(), AttrsVar)

}

var xLabelSetEllipsize func(uintptr, pango.EllipsizeMode)

// Sets the mode used to ellipsizei the text.
//
// The text will be ellipsized if there is not enough space
// to render the entire string.
func (x *Label) SetEllipsize(ModeVar pango.EllipsizeMode) {

	xLabelSetEllipsize(x.GoPointer(), ModeVar)

}

var xLabelSetExtraMenu func(uintptr, uintptr)

// Sets a menu model to add when constructing
// the context menu for @label.
func (x *Label) SetExtraMenu(ModelVar *gio.MenuModel) {

	xLabelSetExtraMenu(x.GoPointer(), ModelVar.GoPointer())

}

var xLabelSetJustify func(uintptr, Justification)

// Sets the alignment of the lines in the text of the label relative to
// each other.
//
// %GTK_JUSTIFY_LEFT is the default value when the widget is first created
// with [ctor@Gtk.Label.new]. If you instead want to set the alignment of
// the label as a whole, use [method@Gtk.Widget.set_halign] instead.
// [method@Gtk.Label.set_justify] has no effect on labels containing
// only a single line.
func (x *Label) SetJustify(JtypeVar Justification) {

	xLabelSetJustify(x.GoPointer(), JtypeVar)

}

var xLabelSetLabel func(uintptr, string)

// Sets the text of the label.
//
// The label is interpreted as including embedded underlines and/or Pango
// markup depending on the values of the [property@Gtk.Label:use-underline]
// and [property@Gtk.Label:use-markup] properties.
func (x *Label) SetLabel(StrVar string) {

	xLabelSetLabel(x.GoPointer(), StrVar)

}

var xLabelSetLines func(uintptr, int)

// Sets the number of lines to which an ellipsized, wrapping label
// should be limited.
//
// This has no effect if the label is not wrapping or ellipsized.
// Set this to -1 if you don’t want to limit the number of lines.
func (x *Label) SetLines(LinesVar int) {

	xLabelSetLines(x.GoPointer(), LinesVar)

}

var xLabelSetMarkup func(uintptr, string)

// Sets the labels text and attributes from markup.
//
// The string must be marked up with Pango markup
// (see [func@Pango.parse_markup]).
//
// If the @str is external data, you may need to escape it
// with g_markup_escape_text() or g_markup_printf_escaped():
//
// ```c
// GtkWidget *self = gtk_label_new (NULL);
// const char *str = "...";
// const char *format = "&lt;span style=\"italic\"&gt;\%s&lt;/span&gt;";
// char *markup;
//
// markup = g_markup_printf_escaped (format, str);
// gtk_label_set_markup (GTK_LABEL (self), markup);
// g_free (markup);
// ```
//
// This function will set the [property@Gtk.Label:use-markup] property
// to %TRUE as a side effect.
//
// If you set the label contents using the [property@Gtk.Label:label]
// property you should also ensure that you set the
// [property@Gtk.Label:use-markup] property accordingly.
//
// See also: [method@Gtk.Label.set_text]
func (x *Label) SetMarkup(StrVar string) {

	xLabelSetMarkup(x.GoPointer(), StrVar)

}

var xLabelSetMarkupWithMnemonic func(uintptr, string)

// Sets the labels text, attributes and mnemonic from markup.
//
// Parses @str which is marked up with Pango markup (see [func@Pango.parse_markup]),
// setting the label’s text and attribute list based on the parse results.
// If characters in @str are preceded by an underscore, they are underlined
// indicating that they represent a keyboard accelerator called a mnemonic.
//
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using [method@Gtk.Label.set_mnemonic_widget].
func (x *Label) SetMarkupWithMnemonic(StrVar string) {

	xLabelSetMarkupWithMnemonic(x.GoPointer(), StrVar)

}

var xLabelSetMaxWidthChars func(uintptr, int)

// Sets the desired maximum width in characters of @label to @n_chars.
func (x *Label) SetMaxWidthChars(NCharsVar int) {

	xLabelSetMaxWidthChars(x.GoPointer(), NCharsVar)

}

var xLabelSetMnemonicWidget func(uintptr, uintptr)

// Associate the label with its mnemonic target.
//
// If the label has been set so that it has a mnemonic key (using
// i.e. [method@Gtk.Label.set_markup_with_mnemonic],
// [method@Gtk.Label.set_text_with_mnemonic],
// [ctor@Gtk.Label.new_with_mnemonic]
// or the [property@Gtk.Label:use_underline] property) the label can be
// associated with a widget that is the target of the mnemonic. When the
// label is inside a widget (like a [class@Gtk.Button] or a
// [class@Gtk.Notebook] tab) it is automatically associated with the correct
// widget, but sometimes (i.e. when the target is a [class@Gtk.Entry] next to
// the label) you need to set it explicitly using this function.
//
// The target widget will be accelerated by emitting the
// [signal@GtkWidget::mnemonic-activate] signal on it. The default handler for
// this signal will activate the widget if there are no mnemonic collisions
// and toggle focus between the colliding widgets otherwise.
func (x *Label) SetMnemonicWidget(WidgetVar *Widget) {

	xLabelSetMnemonicWidget(x.GoPointer(), WidgetVar.GoPointer())

}

var xLabelSetNaturalWrapMode func(uintptr, NaturalWrapMode)

// Select the line wrapping for the natural size request.
//
// This only affects the natural size requested, for the actual wrapping used,
// see the [property@Gtk.Label:wrap-mode] property.
func (x *Label) SetNaturalWrapMode(WrapModeVar NaturalWrapMode) {

	xLabelSetNaturalWrapMode(x.GoPointer(), WrapModeVar)

}

var xLabelSetSelectable func(uintptr, bool)

// Makes text in the label selectable.
//
// Selectable labels allow the user to select text from the label,
// for copy-and-paste.
func (x *Label) SetSelectable(SettingVar bool) {

	xLabelSetSelectable(x.GoPointer(), SettingVar)

}

var xLabelSetSingleLineMode func(uintptr, bool)

// Sets whether the label is in single line mode.
func (x *Label) SetSingleLineMode(SingleLineModeVar bool) {

	xLabelSetSingleLineMode(x.GoPointer(), SingleLineModeVar)

}

var xLabelSetTabs func(uintptr, *pango.TabArray)

// Sets the default tab stops for paragraphs in @self.
func (x *Label) SetTabs(TabsVar *pango.TabArray) {

	xLabelSetTabs(x.GoPointer(), TabsVar)

}

var xLabelSetText func(uintptr, string)

// Sets the text within the `GtkLabel` widget.
//
// It overwrites any text that was there before.
//
// This function will clear any previously set mnemonic accelerators,
// and set the [property@Gtk.Label:use-underline property] to %FALSE as
// a side effect.
//
// This function will set the [property@Gtk.Label:use-markup] property
// to %FALSE as a side effect.
//
// See also: [method@Gtk.Label.set_markup]
func (x *Label) SetText(StrVar string) {

	xLabelSetText(x.GoPointer(), StrVar)

}

var xLabelSetTextWithMnemonic func(uintptr, string)

// Sets the label’s text from the string @str.
//
// If characters in @str are preceded by an underscore, they are underlined
// indicating that they represent a keyboard accelerator called a mnemonic.
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using [method@Gtk.Label.set_mnemonic_widget].
func (x *Label) SetTextWithMnemonic(StrVar string) {

	xLabelSetTextWithMnemonic(x.GoPointer(), StrVar)

}

var xLabelSetUseMarkup func(uintptr, bool)

// Sets whether the text of the label contains markup.
//
// See [method@Gtk.Label.set_markup].
func (x *Label) SetUseMarkup(SettingVar bool) {

	xLabelSetUseMarkup(x.GoPointer(), SettingVar)

}

var xLabelSetUseUnderline func(uintptr, bool)

// Sets whether underlines in the text indicate mnemonics.
func (x *Label) SetUseUnderline(SettingVar bool) {

	xLabelSetUseUnderline(x.GoPointer(), SettingVar)

}

var xLabelSetWidthChars func(uintptr, int)

// Sets the desired width in characters of @label to @n_chars.
func (x *Label) SetWidthChars(NCharsVar int) {

	xLabelSetWidthChars(x.GoPointer(), NCharsVar)

}

var xLabelSetWrap func(uintptr, bool)

// Toggles line wrapping within the `GtkLabel` widget.
//
// %TRUE makes it break lines if text exceeds the widget’s size.
// %FALSE lets the text get cut off by the edge of the widget if
// it exceeds the widget size.
//
// Note that setting line wrapping to %TRUE does not make the label
// wrap at its parent container’s width, because GTK widgets
// conceptually can’t make their requisition depend on the parent
// container’s size. For a label that wraps at a specific position,
// set the label’s width using [method@Gtk.Widget.set_size_request].
func (x *Label) SetWrap(WrapVar bool) {

	xLabelSetWrap(x.GoPointer(), WrapVar)

}

var xLabelSetWrapMode func(uintptr, pango.WrapMode)

// Controls how line wrapping is done.
//
// This only affects the label if line wrapping is on. (See
// [method@Gtk.Label.set_wrap]) The default is %PANGO_WRAP_WORD
// which means wrap on word boundaries.
//
// For sizing behavior, also consider the [property@Gtk.Label:natural-wrap-mode]
// property.
func (x *Label) SetWrapMode(WrapModeVar pango.WrapMode) {

	xLabelSetWrapMode(x.GoPointer(), WrapModeVar)

}

var xLabelSetXalign func(uintptr, float32)

// Sets the `xalign` of the label.
//
// See the [property@Gtk.Label:xalign] property.
func (x *Label) SetXalign(XalignVar float32) {

	xLabelSetXalign(x.GoPointer(), XalignVar)

}

var xLabelSetYalign func(uintptr, float32)

// Sets the `yalign` of the label.
//
// See the [property@Gtk.Label:yalign] property.
func (x *Label) SetYalign(YalignVar float32) {

	xLabelSetYalign(x.GoPointer(), YalignVar)

}

func (c *Label) GoPointer() uintptr {
	return c.Ptr
}

func (c *Label) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Gets emitted when the user activates a link in the label.
//
// The ::activate-current-link is a [keybinding signal](class.SignalAction.html).
//
// Applications may also emit the signal with g_signal_emit_by_name()
// if they need to control activation of URIs programmatically.
//
// The default bindings for this signal are all forms of the Enter key.
func (x *Label) ConnectActivateCurrentLink(cb func(Label)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Label{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate-current-link", purego.NewCallback(fcb))
}

// Gets emitted to activate a URI.
//
// Applications may connect to it to override the default behaviour,
// which is to call gtk_show_uri().
func (x *Label) ConnectActivateLink(cb func(Label, string) bool) uint32 {
	fcb := func(clsPtr uintptr, UriVarp string) bool {
		fa := Label{}
		fa.Ptr = clsPtr

		return cb(fa, UriVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "activate-link", purego.NewCallback(fcb))
}

// Gets emitted to copy the slection to the clipboard.
//
// The ::copy-clipboard signal is a [keybinding signal](class.SignalAction.html).
//
// The default binding for this signal is Ctrl-c.
func (x *Label) ConnectCopyClipboard(cb func(Label)) uint32 {
	fcb := func(clsPtr uintptr) {
		fa := Label{}
		fa.Ptr = clsPtr

		cb(fa)

	}
	return gobject.SignalConnect(x.GoPointer(), "copy-clipboard", purego.NewCallback(fcb))
}

// Gets emitted when the user initiates a cursor movement.
//
// The ::move-cursor signal is a [keybinding signal](class.SignalAction.html).
// If the cursor is not visible in @entry, this signal causes the viewport to
// be moved instead.
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor
// programmatically.
//
// The default bindings for this signal come in two variants,
// the variant with the Shift modifier extends the selection,
// the variant without the Shift modifier does not.
// There are too many key combinations to list them all here.
// - Arrow keys move by individual characters/lines
// - Ctrl-arrow key combinations move by words/paragraphs
// - Home/End keys move to the ends of the buffer
func (x *Label) ConnectMoveCursor(cb func(Label, MovementStep, int, bool)) uint32 {
	fcb := func(clsPtr uintptr, StepVarp MovementStep, CountVarp int, ExtendSelectionVarp bool) {
		fa := Label{}
		fa.Ptr = clsPtr

		cb(fa, StepVarp, CountVarp, ExtendSelectionVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "move-cursor", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *Label) GetAccessibleRole() AccessibleRole {

	cret := XGtkAccessibleGetAccessibleRole(x.GoPointer())
	return cret
}

// Resets the accessible @property to its default value.
func (x *Label) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *Label) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *Label) ResetState(StateVar AccessibleState) {

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
func (x *Label) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Label) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *Label) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Label) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *Label) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *Label) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *Label) GetBuildableId() string {

	cret := XGtkBuildableGetBuildableId(x.GoPointer())
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewLabel, lib, "gtk_label_new")
	core.PuregoSafeRegister(&xNewLabelWithMnemonic, lib, "gtk_label_new_with_mnemonic")

	core.PuregoSafeRegister(&xLabelGetAttributes, lib, "gtk_label_get_attributes")
	core.PuregoSafeRegister(&xLabelGetCurrentUri, lib, "gtk_label_get_current_uri")
	core.PuregoSafeRegister(&xLabelGetEllipsize, lib, "gtk_label_get_ellipsize")
	core.PuregoSafeRegister(&xLabelGetExtraMenu, lib, "gtk_label_get_extra_menu")
	core.PuregoSafeRegister(&xLabelGetJustify, lib, "gtk_label_get_justify")
	core.PuregoSafeRegister(&xLabelGetLabel, lib, "gtk_label_get_label")
	core.PuregoSafeRegister(&xLabelGetLayout, lib, "gtk_label_get_layout")
	core.PuregoSafeRegister(&xLabelGetLayoutOffsets, lib, "gtk_label_get_layout_offsets")
	core.PuregoSafeRegister(&xLabelGetLines, lib, "gtk_label_get_lines")
	core.PuregoSafeRegister(&xLabelGetMaxWidthChars, lib, "gtk_label_get_max_width_chars")
	core.PuregoSafeRegister(&xLabelGetMnemonicKeyval, lib, "gtk_label_get_mnemonic_keyval")
	core.PuregoSafeRegister(&xLabelGetMnemonicWidget, lib, "gtk_label_get_mnemonic_widget")
	core.PuregoSafeRegister(&xLabelGetNaturalWrapMode, lib, "gtk_label_get_natural_wrap_mode")
	core.PuregoSafeRegister(&xLabelGetSelectable, lib, "gtk_label_get_selectable")
	core.PuregoSafeRegister(&xLabelGetSelectionBounds, lib, "gtk_label_get_selection_bounds")
	core.PuregoSafeRegister(&xLabelGetSingleLineMode, lib, "gtk_label_get_single_line_mode")
	core.PuregoSafeRegister(&xLabelGetTabs, lib, "gtk_label_get_tabs")
	core.PuregoSafeRegister(&xLabelGetText, lib, "gtk_label_get_text")
	core.PuregoSafeRegister(&xLabelGetUseMarkup, lib, "gtk_label_get_use_markup")
	core.PuregoSafeRegister(&xLabelGetUseUnderline, lib, "gtk_label_get_use_underline")
	core.PuregoSafeRegister(&xLabelGetWidthChars, lib, "gtk_label_get_width_chars")
	core.PuregoSafeRegister(&xLabelGetWrap, lib, "gtk_label_get_wrap")
	core.PuregoSafeRegister(&xLabelGetWrapMode, lib, "gtk_label_get_wrap_mode")
	core.PuregoSafeRegister(&xLabelGetXalign, lib, "gtk_label_get_xalign")
	core.PuregoSafeRegister(&xLabelGetYalign, lib, "gtk_label_get_yalign")
	core.PuregoSafeRegister(&xLabelSelectRegion, lib, "gtk_label_select_region")
	core.PuregoSafeRegister(&xLabelSetAttributes, lib, "gtk_label_set_attributes")
	core.PuregoSafeRegister(&xLabelSetEllipsize, lib, "gtk_label_set_ellipsize")
	core.PuregoSafeRegister(&xLabelSetExtraMenu, lib, "gtk_label_set_extra_menu")
	core.PuregoSafeRegister(&xLabelSetJustify, lib, "gtk_label_set_justify")
	core.PuregoSafeRegister(&xLabelSetLabel, lib, "gtk_label_set_label")
	core.PuregoSafeRegister(&xLabelSetLines, lib, "gtk_label_set_lines")
	core.PuregoSafeRegister(&xLabelSetMarkup, lib, "gtk_label_set_markup")
	core.PuregoSafeRegister(&xLabelSetMarkupWithMnemonic, lib, "gtk_label_set_markup_with_mnemonic")
	core.PuregoSafeRegister(&xLabelSetMaxWidthChars, lib, "gtk_label_set_max_width_chars")
	core.PuregoSafeRegister(&xLabelSetMnemonicWidget, lib, "gtk_label_set_mnemonic_widget")
	core.PuregoSafeRegister(&xLabelSetNaturalWrapMode, lib, "gtk_label_set_natural_wrap_mode")
	core.PuregoSafeRegister(&xLabelSetSelectable, lib, "gtk_label_set_selectable")
	core.PuregoSafeRegister(&xLabelSetSingleLineMode, lib, "gtk_label_set_single_line_mode")
	core.PuregoSafeRegister(&xLabelSetTabs, lib, "gtk_label_set_tabs")
	core.PuregoSafeRegister(&xLabelSetText, lib, "gtk_label_set_text")
	core.PuregoSafeRegister(&xLabelSetTextWithMnemonic, lib, "gtk_label_set_text_with_mnemonic")
	core.PuregoSafeRegister(&xLabelSetUseMarkup, lib, "gtk_label_set_use_markup")
	core.PuregoSafeRegister(&xLabelSetUseUnderline, lib, "gtk_label_set_use_underline")
	core.PuregoSafeRegister(&xLabelSetWidthChars, lib, "gtk_label_set_width_chars")
	core.PuregoSafeRegister(&xLabelSetWrap, lib, "gtk_label_set_wrap")
	core.PuregoSafeRegister(&xLabelSetWrapMode, lib, "gtk_label_set_wrap_mode")
	core.PuregoSafeRegister(&xLabelSetXalign, lib, "gtk_label_set_xalign")
	core.PuregoSafeRegister(&xLabelSetYalign, lib, "gtk_label_set_yalign")

}
