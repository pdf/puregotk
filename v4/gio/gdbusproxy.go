// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

// Class structure for #GDBusProxy.
type DBusProxyClass struct {
	ParentClass uintptr

	Padding uintptr
}

func (x *DBusProxyClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type DBusProxyPrivate struct {
}

func (x *DBusProxyPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GDBusProxy is a base class used for proxies to access a D-Bus
// interface on a remote object. A #GDBusProxy can be constructed for
// both well-known and unique names.
//
// By default, #GDBusProxy will cache all properties (and listen to
// changes) of the remote object, and proxy all signals that get
// emitted. This behaviour can be changed by passing suitable
// #GDBusProxyFlags when the proxy is created. If the proxy is for a
// well-known name, the property cache is flushed when the name owner
// vanishes and reloaded when a name owner appears.
//
// The unique name owner of the proxy's name is tracked and can be read from
// #GDBusProxy:g-name-owner. Connect to the #GObject::notify signal to
// get notified of changes. Additionally, only signals and property
// changes emitted from the current name owner are considered and
// calls are always sent to the current name owner. This avoids a
// number of race conditions when the name is lost by one owner and
// claimed by another. However, if no name owner currently exists,
// then calls will be sent to the well-known name which may result in
// the message bus launching an owner (unless
// %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START is set).
//
// If the proxy is for a stateless D-Bus service, where the name owner may
// be started and stopped between calls, the #GDBusProxy:g-name-owner tracking
// of #GDBusProxy will cause the proxy to drop signal and property changes from
// the service after it has restarted for the first time. When interacting
// with a stateless D-Bus service, do not use #GDBusProxy — use direct D-Bus
// method calls and signal connections.
//
// The generic #GDBusProxy::g-properties-changed and
// #GDBusProxy::g-signal signals are not very convenient to work with.
// Therefore, the recommended way of working with proxies is to subclass
// #GDBusProxy, and have more natural properties and signals in your derived
// class. This [example][gdbus-example-gdbus-codegen] shows how this can
// easily be done using the [gdbus-codegen][gdbus-codegen] tool.
//
// A #GDBusProxy instance can be used from multiple threads but note
// that all signals (e.g. #GDBusProxy::g-signal, #GDBusProxy::g-properties-changed
// and #GObject::notify) are emitted in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread where the instance was constructed.
//
// An example using a proxy for a well-known name can be found in
// [gdbus-example-watch-proxy.c](https://gitlab.gnome.org/GNOME/glib/-/blob/HEAD/gio/tests/gdbus-example-watch-proxy.c)
type DBusProxy struct {
	gobject.Object
}

func DBusProxyNewFromInternalPtr(ptr uintptr) *DBusProxy {
	cls := &DBusProxy{}
	cls.Ptr = ptr
	return cls
}

var xNewDBusProxyFinish func(uintptr, **glib.Error) uintptr

// Finishes creating a #GDBusProxy.
func NewDBusProxyFinish(ResVar AsyncResult) (*DBusProxy, error) {
	var cls *DBusProxy
	var cerr *glib.Error

	cret := xNewDBusProxyFinish(ResVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DBusProxy{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewDBusProxyForBusFinish func(uintptr, **glib.Error) uintptr

// Finishes creating a #GDBusProxy.
func NewDBusProxyForBusFinish(ResVar AsyncResult) (*DBusProxy, error) {
	var cls *DBusProxy
	var cerr *glib.Error

	cret := xNewDBusProxyForBusFinish(ResVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DBusProxy{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewDBusProxyForBusSync func(BusType, DBusProxyFlags, *DBusInterfaceInfo, string, string, string, uintptr, **glib.Error) uintptr

// Like g_dbus_proxy_new_sync() but takes a #GBusType instead of a #GDBusConnection.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxyForBusSync(BusTypeVar BusType, FlagsVar DBusProxyFlags, InfoVar *DBusInterfaceInfo, NameVar string, ObjectPathVar string, InterfaceNameVar string, CancellableVar *Cancellable) (*DBusProxy, error) {
	var cls *DBusProxy
	var cerr *glib.Error

	cret := xNewDBusProxyForBusSync(BusTypeVar, FlagsVar, InfoVar, NameVar, ObjectPathVar, InterfaceNameVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DBusProxy{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewDBusProxySync func(uintptr, DBusProxyFlags, *DBusInterfaceInfo, string, string, string, uintptr, **glib.Error) uintptr

// Creates a proxy for accessing @interface_name on the remote object
// at @object_path owned by @name at @connection and synchronously
// loads D-Bus properties unless the
// %G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is used.
//
// If the %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets up
// match rules for signals. Connect to the #GDBusProxy::g-signal signal
// to handle signals from the remote object.
//
// If both %G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES and
// %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS are set, this constructor is
// guaranteed to return immediately without blocking.
//
// If @name is a well-known name and the
// %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION
// flags aren't set and no name owner currently exists, the message bus
// will be requested to launch a name owner for the name.
//
// This is a synchronous failable constructor. See g_dbus_proxy_new()
// and g_dbus_proxy_new_finish() for the asynchronous version.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
func NewDBusProxySync(ConnectionVar *DBusConnection, FlagsVar DBusProxyFlags, InfoVar *DBusInterfaceInfo, NameVar string, ObjectPathVar string, InterfaceNameVar string, CancellableVar *Cancellable) (*DBusProxy, error) {
	var cls *DBusProxy
	var cerr *glib.Error

	cret := xNewDBusProxySync(ConnectionVar.GoPointer(), FlagsVar, InfoVar, NameVar, ObjectPathVar, InterfaceNameVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &DBusProxy{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xDBusProxyCall func(uintptr, string, *glib.Variant, DBusCallFlags, int, uintptr, uintptr, uintptr)

// Asynchronously invokes the @method_name method on @proxy.
//
// If @method_name contains any dots, then @name is split into interface and
// method name parts. This allows using @proxy for invoking methods on
// other interfaces.
//
// If the #GDBusConnection associated with @proxy is closed then
// the operation will fail with %G_IO_ERROR_CLOSED. If
// @cancellable is canceled, the operation will fail with
// %G_IO_ERROR_CANCELLED. If @parameters contains a value not
// compatible with the D-Bus protocol, the operation fails with
// %G_IO_ERROR_INVALID_ARGUMENT.
//
// If the @parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
// |[&lt;!-- language="C" --&gt;
//
//	g_dbus_proxy_call (proxy,
//	                   "TwoStrings",
//	                   g_variant_new ("(ss)",
//	                                  "Thing One",
//	                                  "Thing Two"),
//	                   G_DBUS_CALL_FLAGS_NONE,
//	                   -1,
//	                   NULL,
//	                   (GAsyncReadyCallback) two_strings_done,
//	                   &amp;data);
//
// ]|
//
// If @proxy has an expected interface (see
// #GDBusProxy:g-interface-info) and @method_name is referenced by it,
// then the return value is checked against the return type.
//
// This is an asynchronous method. When the operation is finished,
// @callback will be invoked in the
// [thread-default main context][g-main-context-push-thread-default]
// of the thread you are calling this method from.
// You can then call g_dbus_proxy_call_finish() to get the result of
// the operation. See g_dbus_proxy_call_sync() for the synchronous
// version of this method.
//
// If @callback is %NULL then the D-Bus method call message will be sent with
// the %G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
func (x *DBusProxy) Call(MethodNameVar string, ParametersVar *glib.Variant, FlagsVar DBusCallFlags, TimeoutMsecVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xDBusProxyCall(x.GoPointer(), MethodNameVar, ParametersVar, FlagsVar, TimeoutMsecVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xDBusProxyCallFinish func(uintptr, uintptr, **glib.Error) *glib.Variant

// Finishes an operation started with g_dbus_proxy_call().
func (x *DBusProxy) CallFinish(ResVar AsyncResult) (*glib.Variant, error) {
	var cerr *glib.Error

	cret := xDBusProxyCallFinish(x.GoPointer(), ResVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusProxyCallSync func(uintptr, string, *glib.Variant, DBusCallFlags, int, uintptr, **glib.Error) *glib.Variant

// Synchronously invokes the @method_name method on @proxy.
//
// If @method_name contains any dots, then @name is split into interface and
// method name parts. This allows using @proxy for invoking methods on
// other interfaces.
//
// If the #GDBusConnection associated with @proxy is disconnected then
// the operation will fail with %G_IO_ERROR_CLOSED. If
// @cancellable is canceled, the operation will fail with
// %G_IO_ERROR_CANCELLED. If @parameters contains a value not
// compatible with the D-Bus protocol, the operation fails with
// %G_IO_ERROR_INVALID_ARGUMENT.
//
// If the @parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
// |[&lt;!-- language="C" --&gt;
//
//	g_dbus_proxy_call_sync (proxy,
//	                        "TwoStrings",
//	                        g_variant_new ("(ss)",
//	                                       "Thing One",
//	                                       "Thing Two"),
//	                        G_DBUS_CALL_FLAGS_NONE,
//	                        -1,
//	                        NULL,
//	                        &amp;error);
//
// ]|
//
// The calling thread is blocked until a reply is received. See
// g_dbus_proxy_call() for the asynchronous version of this
// method.
//
// If @proxy has an expected interface (see
// #GDBusProxy:g-interface-info) and @method_name is referenced by it,
// then the return value is checked against the return type.
func (x *DBusProxy) CallSync(MethodNameVar string, ParametersVar *glib.Variant, FlagsVar DBusCallFlags, TimeoutMsecVar int, CancellableVar *Cancellable) (*glib.Variant, error) {
	var cerr *glib.Error

	cret := xDBusProxyCallSync(x.GoPointer(), MethodNameVar, ParametersVar, FlagsVar, TimeoutMsecVar, CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusProxyCallWithUnixFdList func(uintptr, string, *glib.Variant, DBusCallFlags, int, uintptr, uintptr, uintptr, uintptr)

// Like g_dbus_proxy_call() but also takes a #GUnixFDList object.
//
// This method is only available on UNIX.
func (x *DBusProxy) CallWithUnixFdList(MethodNameVar string, ParametersVar *glib.Variant, FlagsVar DBusCallFlags, TimeoutMsecVar int, FdListVar *UnixFDList, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xDBusProxyCallWithUnixFdList(x.GoPointer(), MethodNameVar, ParametersVar, FlagsVar, TimeoutMsecVar, FdListVar.GoPointer(), CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xDBusProxyCallWithUnixFdListFinish func(uintptr, *uintptr, uintptr, **glib.Error) *glib.Variant

// Finishes an operation started with g_dbus_proxy_call_with_unix_fd_list().
func (x *DBusProxy) CallWithUnixFdListFinish(OutFdListVar **UnixFDList, ResVar AsyncResult) (*glib.Variant, error) {
	var cerr *glib.Error

	cret := xDBusProxyCallWithUnixFdListFinish(x.GoPointer(), gobject.ConvertPtr(OutFdListVar), ResVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusProxyCallWithUnixFdListSync func(uintptr, string, *glib.Variant, DBusCallFlags, int, uintptr, *uintptr, uintptr, **glib.Error) *glib.Variant

// Like g_dbus_proxy_call_sync() but also takes and returns #GUnixFDList objects.
//
// This method is only available on UNIX.
func (x *DBusProxy) CallWithUnixFdListSync(MethodNameVar string, ParametersVar *glib.Variant, FlagsVar DBusCallFlags, TimeoutMsecVar int, FdListVar *UnixFDList, OutFdListVar **UnixFDList, CancellableVar *Cancellable) (*glib.Variant, error) {
	var cerr *glib.Error

	cret := xDBusProxyCallWithUnixFdListSync(x.GoPointer(), MethodNameVar, ParametersVar, FlagsVar, TimeoutMsecVar, FdListVar.GoPointer(), gobject.ConvertPtr(OutFdListVar), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusProxyGetCachedProperty func(uintptr, string) *glib.Variant

// Looks up the value for a property from the cache. This call does no
// blocking IO.
//
// If @proxy has an expected interface (see
// #GDBusProxy:g-interface-info) and @property_name is referenced by
// it, then @value is checked against the type of the property.
func (x *DBusProxy) GetCachedProperty(PropertyNameVar string) *glib.Variant {

	cret := xDBusProxyGetCachedProperty(x.GoPointer(), PropertyNameVar)
	return cret
}

var xDBusProxyGetCachedPropertyNames func(uintptr) uintptr

// Gets the names of all cached properties on @proxy.
func (x *DBusProxy) GetCachedPropertyNames() uintptr {

	cret := xDBusProxyGetCachedPropertyNames(x.GoPointer())
	return cret
}

var xDBusProxyGetConnection func(uintptr) uintptr

// Gets the connection @proxy is for.
func (x *DBusProxy) GetConnection() *DBusConnection {
	var cls *DBusConnection

	cret := xDBusProxyGetConnection(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &DBusConnection{}
	cls.Ptr = cret
	return cls
}

var xDBusProxyGetDefaultTimeout func(uintptr) int

// Gets the timeout to use if -1 (specifying default timeout) is
// passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the #GDBusProxy:g-default-timeout property for more details.
func (x *DBusProxy) GetDefaultTimeout() int {

	cret := xDBusProxyGetDefaultTimeout(x.GoPointer())
	return cret
}

var xDBusProxyGetFlags func(uintptr) DBusProxyFlags

// Gets the flags that @proxy was constructed with.
func (x *DBusProxy) GetFlags() DBusProxyFlags {

	cret := xDBusProxyGetFlags(x.GoPointer())
	return cret
}

var xDBusProxyGetInterfaceInfo func(uintptr) *DBusInterfaceInfo

// Returns the #GDBusInterfaceInfo, if any, specifying the interface
// that @proxy conforms to. See the #GDBusProxy:g-interface-info
// property for more details.
func (x *DBusProxy) GetInterfaceInfo() *DBusInterfaceInfo {

	cret := xDBusProxyGetInterfaceInfo(x.GoPointer())
	return cret
}

var xDBusProxyGetInterfaceName func(uintptr) string

// Gets the D-Bus interface name @proxy is for.
func (x *DBusProxy) GetInterfaceName() string {

	cret := xDBusProxyGetInterfaceName(x.GoPointer())
	return cret
}

var xDBusProxyGetName func(uintptr) string

// Gets the name that @proxy was constructed for.
//
// When connected to a message bus, this will usually be non-%NULL.
// However, it may be %NULL for a proxy that communicates using a peer-to-peer
// pattern.
func (x *DBusProxy) GetName() string {

	cret := xDBusProxyGetName(x.GoPointer())
	return cret
}

var xDBusProxyGetNameOwner func(uintptr) string

// The unique name that owns the name that @proxy is for or %NULL if
// no-one currently owns that name. You may connect to the
// #GObject::notify signal to track changes to the
// #GDBusProxy:g-name-owner property.
func (x *DBusProxy) GetNameOwner() string {

	cret := xDBusProxyGetNameOwner(x.GoPointer())
	return cret
}

var xDBusProxyGetObjectPath func(uintptr) string

// Gets the object path @proxy is for.
func (x *DBusProxy) GetObjectPath() string {

	cret := xDBusProxyGetObjectPath(x.GoPointer())
	return cret
}

var xDBusProxySetCachedProperty func(uintptr, string, *glib.Variant)

// If @value is not %NULL, sets the cached value for the property with
// name @property_name to the value in @value.
//
// If @value is %NULL, then the cached value is removed from the
// property cache.
//
// If @proxy has an expected interface (see
// #GDBusProxy:g-interface-info) and @property_name is referenced by
// it, then @value is checked against the type of the property.
//
// If the @value #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.
// |[&lt;!-- language="C" --&gt;
//
//	g_dbus_proxy_set_cached_property (proxy,
//	                                  "SomeProperty",
//	                                  g_variant_new ("(si)",
//	                                                "A String",
//	                                                42));
//
// ]|
//
// Normally you will not need to use this method since @proxy
// is tracking changes using the
// `org.freedesktop.DBus.Properties.PropertiesChanged`
// D-Bus signal. However, for performance reasons an object may
// decide to not use this signal for some properties and instead
// use a proprietary out-of-band mechanism to transmit changes.
//
// As a concrete example, consider an object with a property
// `ChatroomParticipants` which is an array of strings. Instead of
// transmitting the same (long) array every time the property changes,
// it is more efficient to only transmit the delta using e.g. signals
// `ChatroomParticipantJoined(String name)` and
// `ChatroomParticipantParted(String name)`.
func (x *DBusProxy) SetCachedProperty(PropertyNameVar string, ValueVar *glib.Variant) {

	xDBusProxySetCachedProperty(x.GoPointer(), PropertyNameVar, ValueVar)

}

var xDBusProxySetDefaultTimeout func(uintptr, int)

// Sets the timeout to use if -1 (specifying default timeout) is
// passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the #GDBusProxy:g-default-timeout property for more details.
func (x *DBusProxy) SetDefaultTimeout(TimeoutMsecVar int) {

	xDBusProxySetDefaultTimeout(x.GoPointer(), TimeoutMsecVar)

}

var xDBusProxySetInterfaceInfo func(uintptr, *DBusInterfaceInfo)

// Ensure that interactions with @proxy conform to the given
// interface. See the #GDBusProxy:g-interface-info property for more
// details.
func (x *DBusProxy) SetInterfaceInfo(InfoVar *DBusInterfaceInfo) {

	xDBusProxySetInterfaceInfo(x.GoPointer(), InfoVar)

}

func (c *DBusProxy) GoPointer() uintptr {
	return c.Ptr
}

func (c *DBusProxy) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Emitted when one or more D-Bus properties on @proxy changes. The
// local cache has already been updated when this signal fires. Note
// that both @changed_properties and @invalidated_properties are
// guaranteed to never be %NULL (either may be empty though).
//
// If the proxy has the flag
// %G_DBUS_PROXY_FLAGS_GET_INVALIDATED_PROPERTIES set, then
// @invalidated_properties will always be empty.
//
// This signal corresponds to the
// `PropertiesChanged` D-Bus signal on the
// `org.freedesktop.DBus.Properties` interface.
func (x *DBusProxy) ConnectGPropertiesChanged(cb func(DBusProxy, uintptr, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, ChangedPropertiesVarp uintptr, InvalidatedPropertiesVarp uintptr) {
		fa := DBusProxy{}
		fa.Ptr = clsPtr

		cb(fa, ChangedPropertiesVarp, InvalidatedPropertiesVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "g-properties-changed", purego.NewCallback(fcb))
}

// Emitted when a signal from the remote object and interface that @proxy is for, has been received.
//
// Since 2.72 this signal supports detailed connections. You can connect to
// the detailed signal `g-signal::x` in order to receive callbacks only when
// signal `x` is received from the remote object.
func (x *DBusProxy) ConnectGSignal(cb func(DBusProxy, string, string, uintptr)) uint32 {
	fcb := func(clsPtr uintptr, SenderNameVarp string, SignalNameVarp string, ParametersVarp uintptr) {
		fa := DBusProxy{}
		fa.Ptr = clsPtr

		cb(fa, SenderNameVarp, SignalNameVarp, ParametersVarp)

	}
	return gobject.SignalConnect(x.GoPointer(), "g-signal", purego.NewCallback(fcb))
}

// Starts asynchronous initialization of the object implementing the
// interface. This must be done before any real use of the object after
// initial construction. If the object also implements #GInitable you can
// optionally call g_initable_init() instead.
//
// This method is intended for language bindings. If writing in C,
// g_async_initable_new_async() should typically be used instead.
//
// When the initialization is finished, @callback will be called. You can
// then call g_async_initable_init_finish() to get the result of the
// initialization.
//
// Implementations may also support cancellation. If @cancellable is not
// %NULL, then initialization can be cancelled by triggering the cancellable
// object from another thread. If the operation was cancelled, the error
// %G_IO_ERROR_CANCELLED will be returned. If @cancellable is not %NULL, and
// the object doesn't support cancellable initialization, the error
// %G_IO_ERROR_NOT_SUPPORTED will be returned.
//
// As with #GInitable, if the object is not initialized, or initialization
// returns with an error, then all operations on the object except
// g_object_ref() and g_object_unref() are considered to be invalid, and
// have undefined behaviour. They will often fail with g_critical() or
// g_warning(), but this must not be relied on.
//
// Callers should not assume that a class which implements #GAsyncInitable can
// be initialized multiple times; for more information, see g_initable_init().
// If a class explicitly supports being initialized multiple times,
// implementation requires yielding all subsequent calls to init_async() on the
// results of the first call.
//
// For classes that also support the #GInitable interface, the default
// implementation of this method will run the g_initable_init() function
// in a thread, so if you want to support asynchronous initialization via
// threads, just implement the #GAsyncInitable interface without overriding
// any interface methods.
func (x *DBusProxy) InitAsync(IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGAsyncInitableInitAsync(x.GoPointer(), IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes asynchronous initialization and returns the result.
// See g_async_initable_init_async().
func (x *DBusProxy) InitFinish(ResVar AsyncResult) (bool, error) {
	var cerr *glib.Error

	cret := XGAsyncInitableInitFinish(x.GoPointer(), ResVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

// Finishes the async construction for the various g_async_initable_new
// calls, returning the created object or %NULL on error.
func (x *DBusProxy) NewFinish(ResVar AsyncResult) (*gobject.Object, error) {
	var cls *gobject.Object
	var cerr *glib.Error

	cret := XGAsyncInitableNewFinish(x.GoPointer(), ResVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gobject.Object{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

// Gets the #GDBusObject that @interface_ belongs to, if any.
func (x *DBusProxy) DupObject() *DBusObjectBase {
	var cls *DBusObjectBase

	cret := XGDbusInterfaceDupObject(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &DBusObjectBase{}
	cls.Ptr = cret
	return cls
}

// Gets D-Bus introspection information for the D-Bus interface
// implemented by @interface_.
func (x *DBusProxy) GetInfo() *DBusInterfaceInfo {

	cret := XGDbusInterfaceGetInfo(x.GoPointer())
	return cret
}

// Gets the #GDBusObject that @interface_ belongs to, if any.
//
// It is not safe to use the returned object if @interface_ or
// the returned object is being used from other threads. See
// g_dbus_interface_dup_object() for a thread-safe alternative.
func (x *DBusProxy) GetObject() *DBusObjectBase {
	var cls *DBusObjectBase

	cret := XGDbusInterfaceGetObject(x.GoPointer())

	if cret == 0 {
		return nil
	}
	gobject.IncreaseRef(cret)
	cls = &DBusObjectBase{}
	cls.Ptr = cret
	return cls
}

// Sets the #GDBusObject for @interface_ to @object.
//
// Note that @interface_ will hold a weak reference to @object.
func (x *DBusProxy) SetObject(ObjectVar DBusObject) {

	XGDbusInterfaceSetObject(x.GoPointer(), ObjectVar.GoPointer())

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
func (x *DBusProxy) Init(CancellableVar *Cancellable) (bool, error) {
	var cerr *glib.Error

	cret := XGInitableInit(x.GoPointer(), CancellableVar.GoPointer(), &cerr)
	if cerr == nil {
		return cret, nil
	}
	return cret, cerr

}

var xDBusProxyNew func(uintptr, DBusProxyFlags, *DBusInterfaceInfo, string, string, string, uintptr, uintptr, uintptr)

// Creates a proxy for accessing @interface_name on the remote object
// at @object_path owned by @name at @connection and asynchronously
// loads D-Bus properties unless the
// %G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES flag is used. Connect to
// the #GDBusProxy::g-properties-changed signal to get notified about
// property changes.
//
// If the %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS flag is not set, also sets up
// match rules for signals. Connect to the #GDBusProxy::g-signal signal
// to handle signals from the remote object.
//
// If both %G_DBUS_PROXY_FLAGS_DO_NOT_LOAD_PROPERTIES and
// %G_DBUS_PROXY_FLAGS_DO_NOT_CONNECT_SIGNALS are set, this constructor is
// guaranteed to complete immediately without blocking.
//
// If @name is a well-known name and the
// %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START and %G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START_AT_CONSTRUCTION
// flags aren't set and no name owner currently exists, the message bus
// will be requested to launch a name owner for the name.
//
// This is a failable asynchronous constructor - when the proxy is
// ready, @callback will be invoked and you can use
// g_dbus_proxy_new_finish() to get the result.
//
// See g_dbus_proxy_new_sync() and for a synchronous version of this constructor.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
func DBusProxyNew(ConnectionVar *DBusConnection, FlagsVar DBusProxyFlags, InfoVar *DBusInterfaceInfo, NameVar string, ObjectPathVar string, InterfaceNameVar string, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xDBusProxyNew(ConnectionVar.GoPointer(), FlagsVar, InfoVar, NameVar, ObjectPathVar, InterfaceNameVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xDBusProxyNewForBus func(BusType, DBusProxyFlags, *DBusInterfaceInfo, string, string, string, uintptr, uintptr, uintptr)

// Like g_dbus_proxy_new() but takes a #GBusType instead of a #GDBusConnection.
//
// #GDBusProxy is used in this [example][gdbus-wellknown-proxy].
func DBusProxyNewForBus(BusTypeVar BusType, FlagsVar DBusProxyFlags, InfoVar *DBusInterfaceInfo, NameVar string, ObjectPathVar string, InterfaceNameVar string, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xDBusProxyNewForBus(BusTypeVar, FlagsVar, InfoVar, NameVar, ObjectPathVar, InterfaceNameVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewDBusProxyFinish, lib, "g_dbus_proxy_new_finish")
	core.PuregoSafeRegister(&xNewDBusProxyForBusFinish, lib, "g_dbus_proxy_new_for_bus_finish")
	core.PuregoSafeRegister(&xNewDBusProxyForBusSync, lib, "g_dbus_proxy_new_for_bus_sync")
	core.PuregoSafeRegister(&xNewDBusProxySync, lib, "g_dbus_proxy_new_sync")

	core.PuregoSafeRegister(&xDBusProxyCall, lib, "g_dbus_proxy_call")
	core.PuregoSafeRegister(&xDBusProxyCallFinish, lib, "g_dbus_proxy_call_finish")
	core.PuregoSafeRegister(&xDBusProxyCallSync, lib, "g_dbus_proxy_call_sync")
	core.PuregoSafeRegister(&xDBusProxyCallWithUnixFdList, lib, "g_dbus_proxy_call_with_unix_fd_list")
	core.PuregoSafeRegister(&xDBusProxyCallWithUnixFdListFinish, lib, "g_dbus_proxy_call_with_unix_fd_list_finish")
	core.PuregoSafeRegister(&xDBusProxyCallWithUnixFdListSync, lib, "g_dbus_proxy_call_with_unix_fd_list_sync")
	core.PuregoSafeRegister(&xDBusProxyGetCachedProperty, lib, "g_dbus_proxy_get_cached_property")
	core.PuregoSafeRegister(&xDBusProxyGetCachedPropertyNames, lib, "g_dbus_proxy_get_cached_property_names")
	core.PuregoSafeRegister(&xDBusProxyGetConnection, lib, "g_dbus_proxy_get_connection")
	core.PuregoSafeRegister(&xDBusProxyGetDefaultTimeout, lib, "g_dbus_proxy_get_default_timeout")
	core.PuregoSafeRegister(&xDBusProxyGetFlags, lib, "g_dbus_proxy_get_flags")
	core.PuregoSafeRegister(&xDBusProxyGetInterfaceInfo, lib, "g_dbus_proxy_get_interface_info")
	core.PuregoSafeRegister(&xDBusProxyGetInterfaceName, lib, "g_dbus_proxy_get_interface_name")
	core.PuregoSafeRegister(&xDBusProxyGetName, lib, "g_dbus_proxy_get_name")
	core.PuregoSafeRegister(&xDBusProxyGetNameOwner, lib, "g_dbus_proxy_get_name_owner")
	core.PuregoSafeRegister(&xDBusProxyGetObjectPath, lib, "g_dbus_proxy_get_object_path")
	core.PuregoSafeRegister(&xDBusProxySetCachedProperty, lib, "g_dbus_proxy_set_cached_property")
	core.PuregoSafeRegister(&xDBusProxySetDefaultTimeout, lib, "g_dbus_proxy_set_default_timeout")
	core.PuregoSafeRegister(&xDBusProxySetInterfaceInfo, lib, "g_dbus_proxy_set_interface_info")

	core.PuregoSafeRegister(&xDBusProxyNew, lib, "g_dbus_proxy_new")
	core.PuregoSafeRegister(&xDBusProxyNewForBus, lib, "g_dbus_proxy_new_for_bus")

}
