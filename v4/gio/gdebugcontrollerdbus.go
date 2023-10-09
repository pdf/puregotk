// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// The virtual function table for #GDebugControllerDBus.
type DebugControllerDBusClass struct {
	ParentClass uintptr

	Padding uintptr
}

// #GDebugControllerDBus is an implementation of #GDebugController which exposes
// debug settings as a D-Bus object.
//
// It is a #GInitable object, and will register an object at
// `/org/gtk/Debugging` on the bus given as
// #GDebugControllerDBus:connection once it’s initialized. The object will be
// unregistered when the last reference to the #GDebugControllerDBus is dropped.
//
// This D-Bus object can be used by remote processes to enable or disable debug
// output in this process. Remote processes calling
// `org.gtk.Debugging.SetDebugEnabled()` will affect the value of
// #GDebugController:debug-enabled and, by default, g_log_get_debug_enabled().
// default.
//
// By default, all processes will be able to call `SetDebugEnabled()`. If this
// process is privileged, or might expose sensitive information in its debug
// output, you may want to restrict the ability to enable debug output to
// privileged users or processes.
//
// One option is to install a D-Bus security policy which restricts access to
// `SetDebugEnabled()`, installing something like the following in
// `$datadir/dbus-1/system.d/`:
// |[&lt;!-- language="XML" --&gt;
// &lt;?xml version="1.0"?&gt; &lt;!--*-nxml-*--&gt;
// &lt;!DOCTYPE busconfig PUBLIC "-//freedesktop//DTD D-BUS Bus Configuration 1.0//EN"
//
//	"http://www.freedesktop.org/standards/dbus/1.0/busconfig.dtd"&gt;
//
// &lt;busconfig&gt;
//
//	&lt;policy user="root"&gt;
//	  &lt;allow send_destination="com.example.MyService" send_interface="org.gtk.Debugging"/&gt;
//	&lt;/policy&gt;
//	&lt;policy context="default"&gt;
//	  &lt;deny send_destination="com.example.MyService" send_interface="org.gtk.Debugging"/&gt;
//	&lt;/policy&gt;
//
// &lt;/busconfig&gt;
// ]|
//
// This will prevent the `SetDebugEnabled()` method from being called by all
// except root. It will not prevent the `DebugEnabled` property from being read,
// as it’s accessed through the `org.freedesktop.DBus.Properties` interface.
//
// Another option is to use polkit to allow or deny requests on a case-by-case
// basis, allowing for the possibility of dynamic authorisation. To do this,
// connect to the #GDebugControllerDBus::authorize signal and query polkit in
// it:
// |[&lt;!-- language="C" --&gt;
//
//	g_autoptr(GError) child_error = NULL;
//	g_autoptr(GDBusConnection) connection = g_bus_get_sync (G_BUS_TYPE_SYSTEM, NULL, NULL);
//	gulong debug_controller_authorize_id = 0;
//
//	// Set up the debug controller.
//	debug_controller = G_DEBUG_CONTROLLER (g_debug_controller_dbus_new (priv-&gt;connection, NULL, &amp;child_error));
//	if (debug_controller == NULL)
//	  {
//	    g_error ("Could not register debug controller on bus: %s"),
//	             child_error-&gt;message);
//	  }
//
//	debug_controller_authorize_id = g_signal_connect (debug_controller,
//	                                                  "authorize",
//	                                                  G_CALLBACK (debug_controller_authorize_cb),
//	                                                  self);
//
//	static gboolean
//	debug_controller_authorize_cb (GDebugControllerDBus  *debug_controller,
//	                               GDBusMethodInvocation *invocation,
//	                               gpointer               user_data)
//	{
//	  g_autoptr(PolkitAuthority) authority = NULL;
//	  g_autoptr(PolkitSubject) subject = NULL;
//	  g_autoptr(PolkitAuthorizationResult) auth_result = NULL;
//	  g_autoptr(GError) local_error = NULL;
//	  GDBusMessage *message;
//	  GDBusMessageFlags message_flags;
//	  PolkitCheckAuthorizationFlags flags = POLKIT_CHECK_AUTHORIZATION_FLAGS_NONE;
//
//	  message = g_dbus_method_invocation_get_message (invocation);
//	  message_flags = g_dbus_message_get_flags (message);
//
//	  authority = polkit_authority_get_sync (NULL, &amp;local_error);
//	  if (authority == NULL)
//	    {
//	      g_warning ("Failed to get polkit authority: %s", local_error-&gt;message);
//	      return FALSE;
//	    }
//
//	  if (message_flags &amp; G_DBUS_MESSAGE_FLAGS_ALLOW_INTERACTIVE_AUTHORIZATION)
//	    flags |= POLKIT_CHECK_AUTHORIZATION_FLAGS_ALLOW_USER_INTERACTION;
//
//	  subject = polkit_system_bus_name_new (g_dbus_method_invocation_get_sender (invocation));
//
//	  auth_result = polkit_authority_check_authorization_sync (authority,
//	                                                           subject,
//	                                                           "com.example.MyService.set-debug-enabled",
//	                                                           NULL,
//	                                                           flags,
//	                                                           NULL,
//	                                                           &amp;local_error);
//	  if (auth_result == NULL)
//	    {
//	      g_warning ("Failed to get check polkit authorization: %s", local_error-&gt;message);
//	      return FALSE;
//	    }
//
//	  return polkit_authorization_result_get_is_authorized (auth_result);
//	}
//
// ]|
type DebugControllerDBus struct {
	gobject.Object
}

func DebugControllerDBusNewFromInternalPtr(ptr uintptr) *DebugControllerDBus {
	cls := &DebugControllerDBus{}
	cls.Ptr = ptr
	return cls
}

var xNewDebugControllerDBus func(uintptr, uintptr, **glib.Error) uintptr

// Create a new #GDebugControllerDBus and synchronously initialize it.
//
// Initializing the object will export the debug object on @connection. The
// object will remain registered until the last reference to the
// #GDebugControllerDBus is dropped.
//
// Initialization may fail if registering the object on @connection fails.
func NewDebugControllerDBus(ConnectionVar *DBusConnection, CancellableVar *Cancellable) (*DebugControllerDBus, error) {
	var cls *DebugControllerDBus
	var cerr *glib.Error

	cret := xNewDebugControllerDBus(ConnectionVar.GoPointer(), CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DebugControllerDBus{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xDebugControllerDBusStop func(uintptr)

// Stop the debug controller, unregistering its object from the bus.
//
// Any pending method calls to the object will complete successfully, but new
// ones will return an error. This method will block until all pending
// #GDebugControllerDBus::authorize signals have been handled. This is expected
// to not take long, as it will just be waiting for threads to join. If any
// #GDebugControllerDBus::authorize signal handlers are still executing in other
// threads, this will block until after they have returned.
//
// This method will be called automatically when the final reference to the
// #GDebugControllerDBus is dropped. You may want to call it explicitly to know
// when the controller has been fully removed from the bus, or to break
// reference count cycles.
//
// Calling this method from within a #GDebugControllerDBus::authorize signal
// handler will cause a deadlock and must not be done.
func (x *DebugControllerDBus) Stop() {

	xDebugControllerDBusStop(x.GoPointer())

}

func (c *DebugControllerDBus) GoPointer() uintptr {
	return c.Ptr
}

func (c *DebugControllerDBus) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when a D-Bus peer is trying to change the debug settings and used
// to determine if that is authorized.
//
// This signal is emitted in a dedicated worker thread, so handlers are
// allowed to perform blocking I/O. This means that, for example, it is
// appropriate to call `polkit_authority_check_authorization_sync()` to check
// authorization using polkit.
//
// If %FALSE is returned then no further handlers are run and the request to
// change the debug settings is rejected.
//
// Otherwise, if %TRUE is returned, signal emission continues. If no handlers
// return %FALSE, then the debug settings are allowed to be changed.
//
// Signal handlers must not modify @invocation, or cause it to return a value.
//
// The default class handler just returns %TRUE.
func (x *DebugControllerDBus) ConnectAuthorize(cb func(DebugControllerDBus, uintptr) bool) uint32 {
	fcb := func(clsPtr uintptr, InvocationVarp uintptr) bool {
		fa := DebugControllerDBus{}
		fa.Ptr = clsPtr

		return cb(fa, InvocationVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "authorize", purego.NewCallback(fcb))
}

// Get the value of #GDebugController:debug-enabled.
func (x *DebugControllerDBus) GetDebugEnabled() bool {

	cret := XGDebugControllerGetDebugEnabled(x.GoPointer())
	return cret
}

// Set the value of #GDebugController:debug-enabled.
func (x *DebugControllerDBus) SetDebugEnabled(DebugEnabledVar bool) {

	XGDebugControllerSetDebugEnabled(x.GoPointer(), DebugEnabledVar)

}

// Initializes the object implementing the interface.
//
// This method is intended for language bindings. If writing in C,
// g_initable_new() should typically be used instead.
//
// The object must be initialized before any real use after initial
// construction, either with this function or g_async_initable_init_async().
//
// Implementations may also support cancellation. If @cancellable is not %NULL,
// then initialization can be cancelled by triggering the cancellable object
// from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. If @cancellable is not %NULL and
// the object doesn't support cancellable initialization the error
// %G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// If the object is not initialized, or initialization returns with an
// error, then all operations on the object except g_object_ref() and
// g_object_unref() are considered to be invalid, and have undefined
// behaviour. See the [introduction][ginitable] for more details.
//
// Callers should not assume that a class which implements #GInitable can be
// initialized multiple times, unless the class explicitly documents itself as
// supporting this. Generally, a class’ implementation of init() can assume
// (and assert) that it will only be called once. Previously, this documentation
// recommended all #GInitable implementations should be idempotent; that
// recommendation was relaxed in GLib 2.54.
//
// If a class explicitly supports being initialized multiple times, it is
// recommended that the method is idempotent: multiple calls with the same
// arguments should return the same results. Only the first call initializes
// the object; further calls return the result of the first call.
//
// One reason why a class might need to support idempotent initialization is if
// it is designed to be used via the singleton pattern, with a
// #GObjectClass.constructor that sometimes returns an existing instance.
// In this pattern, a caller would expect to be able to call g_initable_init()
// on the result of g_object_new(), regardless of whether it is in fact a new
// instance.
func (x *DebugControllerDBus) Init(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGInitableInit(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDebugControllerDBus, lib, "g_debug_controller_dbus_new")

	core.PuregoSafeRegister(&xDebugControllerDBusStop, lib, "g_debug_controller_dbus_stop")

}
