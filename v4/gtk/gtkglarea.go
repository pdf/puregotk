// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The `GtkGLAreaClass` structure contains only private data.
type GLAreaClass struct {
	ParentClass uintptr

	Padding uintptr
}

// `GtkGLArea` is a widget that allows drawing with OpenGL.
//
// ![An example GtkGLArea](glarea.png)
//
// `GtkGLArea` sets up its own [class@Gdk.GLContext], and creates a custom
// GL framebuffer that the widget will do GL rendering onto. It also ensures
// that this framebuffer is the default GL rendering target when rendering.
//
// In order to draw, you have to connect to the [signal@Gtk.GLArea::render]
// signal, or subclass `GtkGLArea` and override the GtkGLAreaClass.render
// virtual function.
//
// The `GtkGLArea` widget ensures that the `GdkGLContext` is associated with
// the widget's drawing area, and it is kept updated when the size and
// position of the drawing area changes.
//
// ## Drawing with GtkGLArea
//
// The simplest way to draw using OpenGL commands in a `GtkGLArea` is to
// create a widget instance and connect to the [signal@Gtk.GLArea::render] signal:
//
// The `render()` function will be called when the `GtkGLArea` is ready
// for you to draw its content:
//
// ```c
// static gboolean
// render (GtkGLArea *area, GdkGLContext *context)
//
//	{
//	  // inside this function it's safe to use GL; the given
//	  // GdkGLContext has been made current to the drawable
//	  // surface used by the `GtkGLArea` and the viewport has
//	  // already been set to be the size of the allocation
//
//	  // we can start by clearing the buffer
//	  glClearColor (0, 0, 0, 0);
//	  glClear (GL_COLOR_BUFFER_BIT);
//
//	  // draw your object
//	  // draw_an_object ();
//
//	  // we completed our drawing; the draw commands will be
//	  // flushed at the end of the signal emission chain, and
//	  // the buffers will be drawn on the window
//	  return TRUE;
//	}
//
// void setup_glarea (void)
//
//	{
//	  // create a GtkGLArea instance
//	  GtkWidget *gl_area = gtk_gl_area_new ();
//
//	  // connect to the "render" signal
//	  g_signal_connect (gl_area, "render", G_CALLBACK (render), NULL);
//	}
//
// ```
//
// If you need to initialize OpenGL state, e.g. buffer objects or
// shaders, you should use the [signal@Gtk.Widget::realize] signal;
// you can use the [signal@Gtk.Widget::unrealize] signal to clean up.
// Since the `GdkGLContext` creation and initialization may fail, you
// will need to check for errors, using [method@Gtk.GLArea.get_error].
//
// An example of how to safely initialize the GL state is:
//
// ```c
// static void
// on_realize (GtkGLarea *area)
//
//	{
//	  // We need to make the context current if we want to
//	  // call GL API
//	  gtk_gl_area_make_current (area);
//
//	  // If there were errors during the initialization or
//	  // when trying to make the context current, this
//	  // function will return a GError for you to catch
//	  if (gtk_gl_area_get_error (area) != NULL)
//	    return;
//
//	  // You can also use gtk_gl_area_set_error() in order
//	  // to show eventual initialization errors on the
//	  // GtkGLArea widget itself
//	  GError *internal_error = NULL;
//	  init_buffer_objects (&amp;error);
//	  if (error != NULL)
//	    {
//	      gtk_gl_area_set_error (area, error);
//	      g_error_free (error);
//	      return;
//	    }
//
//	  init_shaders (&amp;error);
//	  if (error != NULL)
//	    {
//	      gtk_gl_area_set_error (area, error);
//	      g_error_free (error);
//	      return;
//	    }
//	}
//
// ```
//
// If you need to change the options for creating the `GdkGLContext`
// you should use the [signal@Gtk.GLArea::create-context] signal.
type GLArea struct {
	Widget
}

func GLAreaNewFromInternalPtr(ptr uintptr) *GLArea {
	cls := &GLArea{}
	cls.Ptr = ptr
	return cls
}

var xNewGLArea func() uintptr

// Creates a new `GtkGLArea` widget.
func NewGLArea() *Widget {
	NewGLAreaPtr := xNewGLArea()
	if NewGLAreaPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(NewGLAreaPtr)

	NewGLAreaCls := &Widget{}
	NewGLAreaCls.Ptr = NewGLAreaPtr
	return NewGLAreaCls
}

var xGLAreaAttachBuffers func(uintptr)

// Binds buffers to the framebuffer.
//
// Ensures that the @area framebuffer object is made the current draw
// and read target, and that all the required buffers for the @area
// are created and bound to the framebuffer.
//
// This function is automatically called before emitting the
// [signal@Gtk.GLArea::render] signal, and doesn't normally need to be
// called by application code.
func (x *GLArea) AttachBuffers() {

	xGLAreaAttachBuffers(x.GoPointer())

}

var xGLAreaGetAutoRender func(uintptr) bool

// Returns whether the area is in auto render mode or not.
func (x *GLArea) GetAutoRender() bool {

	return xGLAreaGetAutoRender(x.GoPointer())

}

var xGLAreaGetContext func(uintptr) uintptr

// Retrieves the `GdkGLContext` used by @area.
func (x *GLArea) GetContext() *gdk.GLContext {

	GetContextPtr := xGLAreaGetContext(x.GoPointer())
	if GetContextPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetContextPtr)

	GetContextCls := &gdk.GLContext{}
	GetContextCls.Ptr = GetContextPtr
	return GetContextCls

}

var xGLAreaGetError func(uintptr) *glib.Error

// Gets the current error set on the @area.
func (x *GLArea) GetError() *glib.Error {

	return xGLAreaGetError(x.GoPointer())

}

var xGLAreaGetHasDepthBuffer func(uintptr) bool

// Returns whether the area has a depth buffer.
func (x *GLArea) GetHasDepthBuffer() bool {

	return xGLAreaGetHasDepthBuffer(x.GoPointer())

}

var xGLAreaGetHasStencilBuffer func(uintptr) bool

// Returns whether the area has a stencil buffer.
func (x *GLArea) GetHasStencilBuffer() bool {

	return xGLAreaGetHasStencilBuffer(x.GoPointer())

}

var xGLAreaGetRequiredVersion func(uintptr, int, int)

// Retrieves the required version of OpenGL.
//
// See [method@Gtk.GLArea.set_required_version].
func (x *GLArea) GetRequiredVersion(MajorVar int, MinorVar int) {

	xGLAreaGetRequiredVersion(x.GoPointer(), MajorVar, MinorVar)

}

var xGLAreaGetUseEs func(uintptr) bool

// Returns whether the `GtkGLArea` should use OpenGL ES.
//
// See [method@Gtk.GLArea.set_use_es].
func (x *GLArea) GetUseEs() bool {

	return xGLAreaGetUseEs(x.GoPointer())

}

var xGLAreaMakeCurrent func(uintptr)

// Ensures that the `GdkGLContext` used by @area is associated with
// the `GtkGLArea`.
//
// This function is automatically called before emitting the
// [signal@Gtk.GLArea::render] signal, and doesn't normally need
// to be called by application code.
func (x *GLArea) MakeCurrent() {

	xGLAreaMakeCurrent(x.GoPointer())

}

var xGLAreaQueueRender func(uintptr)

// Marks the currently rendered data (if any) as invalid, and queues
// a redraw of the widget.
//
// This ensures that the [signal@Gtk.GLArea::render] signal
// is emitted during the draw.
//
// This is only needed when [method@Gtk.GLArea.set_auto_render] has
// been called with a %FALSE value. The default behaviour is to
// emit [signal@Gtk.GLArea::render] on each draw.
func (x *GLArea) QueueRender() {

	xGLAreaQueueRender(x.GoPointer())

}

var xGLAreaSetAutoRender func(uintptr, bool)

// Sets whether the `GtkGLArea` is in auto render mode.
//
// If @auto_render is %TRUE the [signal@Gtk.GLArea::render] signal will
// be emitted every time the widget draws. This is the default and is
// useful if drawing the widget is faster.
//
// If @auto_render is %FALSE the data from previous rendering is kept
// around and will be used for drawing the widget the next time,
// unless the window is resized. In order to force a rendering
// [method@Gtk.GLArea.queue_render] must be called. This mode is
// useful when the scene changes seldom, but takes a long time to redraw.
func (x *GLArea) SetAutoRender(AutoRenderVar bool) {

	xGLAreaSetAutoRender(x.GoPointer(), AutoRenderVar)

}

var xGLAreaSetError func(uintptr, *glib.Error)

// Sets an error on the area which will be shown instead of the
// GL rendering.
//
// This is useful in the [signal@Gtk.GLArea::create-context]
// signal if GL context creation fails.
func (x *GLArea) SetError(ErrorVar *glib.Error) {

	xGLAreaSetError(x.GoPointer(), ErrorVar)

}

var xGLAreaSetHasDepthBuffer func(uintptr, bool)

// Sets whether the `GtkGLArea` should use a depth buffer.
//
// If @has_depth_buffer is %TRUE the widget will allocate and
// enable a depth buffer for the target framebuffer. Otherwise
// there will be none.
func (x *GLArea) SetHasDepthBuffer(HasDepthBufferVar bool) {

	xGLAreaSetHasDepthBuffer(x.GoPointer(), HasDepthBufferVar)

}

var xGLAreaSetHasStencilBuffer func(uintptr, bool)

// Sets whether the `GtkGLArea` should use a stencil buffer.
//
// If @has_stencil_buffer is %TRUE the widget will allocate and
// enable a stencil buffer for the target framebuffer. Otherwise
// there will be none.
func (x *GLArea) SetHasStencilBuffer(HasStencilBufferVar bool) {

	xGLAreaSetHasStencilBuffer(x.GoPointer(), HasStencilBufferVar)

}

var xGLAreaSetRequiredVersion func(uintptr, int, int)

// Sets the required version of OpenGL to be used when creating
// the context for the widget.
//
// This function must be called before the area has been realized.
func (x *GLArea) SetRequiredVersion(MajorVar int, MinorVar int) {

	xGLAreaSetRequiredVersion(x.GoPointer(), MajorVar, MinorVar)

}

var xGLAreaSetUseEs func(uintptr, bool)

// Sets whether the @area should create an OpenGL or an OpenGL ES context.
//
// You should check the capabilities of the `GdkGLContext` before drawing
// with either API.
func (x *GLArea) SetUseEs(UseEsVar bool) {

	xGLAreaSetUseEs(x.GoPointer(), UseEsVar)

}

func (c *GLArea) GoPointer() uintptr {
	return c.Ptr
}

func (c *GLArea) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when the widget is being realized.
//
// This allows you to override how the GL context is created.
// This is useful when you want to reuse an existing GL context,
// or if you want to try creating different kinds of GL options.
//
// If context creation fails then the signal handler can use
// [method@Gtk.GLArea.set_error] to register a more detailed error
// of how the construction failed.
func (x *GLArea) ConnectCreateContext(cb func(GLArea) gdk.GLContext) {
	fcb := func(clsPtr uintptr) uintptr {
		fa := GLArea{}
		fa.Ptr = clsPtr

		CreateContextCls := cb(fa)
		return CreateContextCls.Ptr

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::create-context", purego.NewCallback(fcb))
}

// Emitted every time the contents of the `GtkGLArea` should be redrawn.
//
// The @context is bound to the @area prior to emitting this function,
// and the buffers are painted to the window once the emission terminates.
func (x *GLArea) ConnectRender(cb func(GLArea, uintptr) bool) {
	fcb := func(clsPtr uintptr, ContextVarp uintptr) bool {
		fa := GLArea{}
		fa.Ptr = clsPtr

		return cb(fa, ContextVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::render", purego.NewCallback(fcb))
}

// Emitted once when the widget is realized, and then each time the widget
// is changed while realized.
//
// This is useful in order to keep GL state up to date with the widget size,
// like for instance camera properties which may depend on the width/height
// ratio.
//
// The GL context for the area is guaranteed to be current when this signal
// is emitted.
//
// The default handler sets up the GL viewport.
func (x *GLArea) ConnectResize(cb func(GLArea, int, int)) {
	fcb := func(clsPtr uintptr, WidthVarp int, HeightVarp int) {
		fa := GLArea{}
		fa.Ptr = clsPtr

		cb(fa, WidthVarp, HeightVarp)

	}
	gobject.ObjectConnect(x.GoPointer(), "signal::resize", purego.NewCallback(fcb))
}

// Retrieves the `GtkAccessibleRole` for the given `GtkAccessible`.
func (x *GLArea) GetAccessibleRole() AccessibleRole {

	return XGtkAccessibleGetAccessibleRole(x.GoPointer())

}

// Resets the accessible @property to its default value.
func (x *GLArea) ResetProperty(PropertyVar AccessibleProperty) {

	XGtkAccessibleResetProperty(x.GoPointer(), PropertyVar)

}

// Resets the accessible @relation to its default value.
func (x *GLArea) ResetRelation(RelationVar AccessibleRelation) {

	XGtkAccessibleResetRelation(x.GoPointer(), RelationVar)

}

// Resets the accessible @state to its default value.
func (x *GLArea) ResetState(StateVar AccessibleState) {

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
func (x *GLArea) UpdateProperty(FirstPropertyVar AccessibleProperty, varArgs ...interface{}) {

	XGtkAccessibleUpdateProperty(x.GoPointer(), FirstPropertyVar, varArgs...)

}

// Updates an array of accessible properties.
//
// This function should be called by `GtkWidget` types whenever an accessible
// property change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GLArea) UpdatePropertyValue(NPropertiesVar int, PropertiesVar uintptr, ValuesVar uintptr) {

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
func (x *GLArea) UpdateRelation(FirstRelationVar AccessibleRelation, varArgs ...interface{}) {

	XGtkAccessibleUpdateRelation(x.GoPointer(), FirstRelationVar, varArgs...)

}

// Updates an array of accessible relations.
//
// This function should be called by `GtkWidget` types whenever an accessible
// relation change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GLArea) UpdateRelationValue(NRelationsVar int, RelationsVar uintptr, ValuesVar uintptr) {

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
func (x *GLArea) UpdateState(FirstStateVar AccessibleState, varArgs ...interface{}) {

	XGtkAccessibleUpdateState(x.GoPointer(), FirstStateVar, varArgs...)

}

// Updates an array of accessible states.
//
// This function should be called by `GtkWidget` types whenever an accessible
// state change must be communicated to assistive technologies.
//
// This function is meant to be used by language bindings.
func (x *GLArea) UpdateStateValue(NStatesVar int, StatesVar uintptr, ValuesVar uintptr) {

	XGtkAccessibleUpdateStateValue(x.GoPointer(), NStatesVar, StatesVar, ValuesVar)

}

// Gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute
// of the &lt;object&gt; tag used to construct the @buildable.
func (x *GLArea) GetBuildableId() string {

	return XGtkBuildableGetBuildableId(x.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGLArea, lib, "gtk_gl_area_new")

	core.PuregoSafeRegister(&xGLAreaAttachBuffers, lib, "gtk_gl_area_attach_buffers")
	core.PuregoSafeRegister(&xGLAreaGetAutoRender, lib, "gtk_gl_area_get_auto_render")
	core.PuregoSafeRegister(&xGLAreaGetContext, lib, "gtk_gl_area_get_context")
	core.PuregoSafeRegister(&xGLAreaGetError, lib, "gtk_gl_area_get_error")
	core.PuregoSafeRegister(&xGLAreaGetHasDepthBuffer, lib, "gtk_gl_area_get_has_depth_buffer")
	core.PuregoSafeRegister(&xGLAreaGetHasStencilBuffer, lib, "gtk_gl_area_get_has_stencil_buffer")
	core.PuregoSafeRegister(&xGLAreaGetRequiredVersion, lib, "gtk_gl_area_get_required_version")
	core.PuregoSafeRegister(&xGLAreaGetUseEs, lib, "gtk_gl_area_get_use_es")
	core.PuregoSafeRegister(&xGLAreaMakeCurrent, lib, "gtk_gl_area_make_current")
	core.PuregoSafeRegister(&xGLAreaQueueRender, lib, "gtk_gl_area_queue_render")
	core.PuregoSafeRegister(&xGLAreaSetAutoRender, lib, "gtk_gl_area_set_auto_render")
	core.PuregoSafeRegister(&xGLAreaSetError, lib, "gtk_gl_area_set_error")
	core.PuregoSafeRegister(&xGLAreaSetHasDepthBuffer, lib, "gtk_gl_area_set_has_depth_buffer")
	core.PuregoSafeRegister(&xGLAreaSetHasStencilBuffer, lib, "gtk_gl_area_set_has_stencil_buffer")
	core.PuregoSafeRegister(&xGLAreaSetRequiredVersion, lib, "gtk_gl_area_set_required_version")
	core.PuregoSafeRegister(&xGLAreaSetUseEs, lib, "gtk_gl_area_set_use_es")

}
