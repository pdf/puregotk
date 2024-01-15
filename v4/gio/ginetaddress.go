// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type InetAddressClass struct {
	ParentClass uintptr
}

func (x *InetAddressClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

type InetAddressPrivate struct {
}

func (x *InetAddressPrivate) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// #GInetAddress represents an IPv4 or IPv6 internet address. Use
// g_resolver_lookup_by_name() or g_resolver_lookup_by_name_async() to
// look up the #GInetAddress for a hostname. Use
// g_resolver_lookup_by_address() or
// g_resolver_lookup_by_address_async() to look up the hostname for a
// #GInetAddress.
//
// To actually connect to a remote host, you will need a
// #GInetSocketAddress (which includes a #GInetAddress as well as a
// port number).
type InetAddress struct {
	gobject.Object
}

func InetAddressNewFromInternalPtr(ptr uintptr) *InetAddress {
	cls := &InetAddress{}
	cls.Ptr = ptr
	return cls
}

var xNewInetAddressAny func(SocketFamily) uintptr

// Creates a #GInetAddress for the "any" address (unassigned/"don't
// care") for @family.
func NewInetAddressAny(FamilyVar SocketFamily) *InetAddress {
	var cls *InetAddress

	cret := xNewInetAddressAny(FamilyVar)

	if cret == 0 {
		return nil
	}
	cls = &InetAddress{}
	cls.Ptr = cret
	return cls
}

var xNewInetAddressFromBytes func(uintptr, SocketFamily) uintptr

// Creates a new #GInetAddress from the given @family and @bytes.
// @bytes should be 4 bytes for %G_SOCKET_FAMILY_IPV4 and 16 bytes for
// %G_SOCKET_FAMILY_IPV6.
func NewInetAddressFromBytes(BytesVar uintptr, FamilyVar SocketFamily) *InetAddress {
	var cls *InetAddress

	cret := xNewInetAddressFromBytes(BytesVar, FamilyVar)

	if cret == 0 {
		return nil
	}
	cls = &InetAddress{}
	cls.Ptr = cret
	return cls
}

var xNewInetAddressFromString func(string) uintptr

// Parses @string as an IP address and creates a new #GInetAddress.
func NewInetAddressFromString(StringVar string) *InetAddress {
	var cls *InetAddress

	cret := xNewInetAddressFromString(StringVar)

	if cret == 0 {
		return nil
	}
	cls = &InetAddress{}
	cls.Ptr = cret
	return cls
}

var xNewInetAddressLoopback func(SocketFamily) uintptr

// Creates a #GInetAddress for the loopback address for @family.
func NewInetAddressLoopback(FamilyVar SocketFamily) *InetAddress {
	var cls *InetAddress

	cret := xNewInetAddressLoopback(FamilyVar)

	if cret == 0 {
		return nil
	}
	cls = &InetAddress{}
	cls.Ptr = cret
	return cls
}

var xInetAddressEqual func(uintptr, uintptr) bool

// Checks if two #GInetAddress instances are equal, e.g. the same address.
func (x *InetAddress) Equal(OtherAddressVar *InetAddress) bool {

	cret := xInetAddressEqual(x.GoPointer(), OtherAddressVar.GoPointer())
	return cret
}

var xInetAddressGetFamily func(uintptr) SocketFamily

// Gets @address's family
func (x *InetAddress) GetFamily() SocketFamily {

	cret := xInetAddressGetFamily(x.GoPointer())
	return cret
}

var xInetAddressGetIsAny func(uintptr) bool

// Tests whether @address is the "any" address for its family.
func (x *InetAddress) GetIsAny() bool {

	cret := xInetAddressGetIsAny(x.GoPointer())
	return cret
}

var xInetAddressGetIsLinkLocal func(uintptr) bool

// Tests whether @address is a link-local address (that is, if it
// identifies a host on a local network that is not connected to the
// Internet).
func (x *InetAddress) GetIsLinkLocal() bool {

	cret := xInetAddressGetIsLinkLocal(x.GoPointer())
	return cret
}

var xInetAddressGetIsLoopback func(uintptr) bool

// Tests whether @address is the loopback address for its family.
func (x *InetAddress) GetIsLoopback() bool {

	cret := xInetAddressGetIsLoopback(x.GoPointer())
	return cret
}

var xInetAddressGetIsMcGlobal func(uintptr) bool

// Tests whether @address is a global multicast address.
func (x *InetAddress) GetIsMcGlobal() bool {

	cret := xInetAddressGetIsMcGlobal(x.GoPointer())
	return cret
}

var xInetAddressGetIsMcLinkLocal func(uintptr) bool

// Tests whether @address is a link-local multicast address.
func (x *InetAddress) GetIsMcLinkLocal() bool {

	cret := xInetAddressGetIsMcLinkLocal(x.GoPointer())
	return cret
}

var xInetAddressGetIsMcNodeLocal func(uintptr) bool

// Tests whether @address is a node-local multicast address.
func (x *InetAddress) GetIsMcNodeLocal() bool {

	cret := xInetAddressGetIsMcNodeLocal(x.GoPointer())
	return cret
}

var xInetAddressGetIsMcOrgLocal func(uintptr) bool

// Tests whether @address is an organization-local multicast address.
func (x *InetAddress) GetIsMcOrgLocal() bool {

	cret := xInetAddressGetIsMcOrgLocal(x.GoPointer())
	return cret
}

var xInetAddressGetIsMcSiteLocal func(uintptr) bool

// Tests whether @address is a site-local multicast address.
func (x *InetAddress) GetIsMcSiteLocal() bool {

	cret := xInetAddressGetIsMcSiteLocal(x.GoPointer())
	return cret
}

var xInetAddressGetIsMulticast func(uintptr) bool

// Tests whether @address is a multicast address.
func (x *InetAddress) GetIsMulticast() bool {

	cret := xInetAddressGetIsMulticast(x.GoPointer())
	return cret
}

var xInetAddressGetIsSiteLocal func(uintptr) bool

// Tests whether @address is a site-local address such as 10.0.0.1
// (that is, the address identifies a host on a local network that can
// not be reached directly from the Internet, but which may have
// outgoing Internet connectivity via a NAT or firewall).
func (x *InetAddress) GetIsSiteLocal() bool {

	cret := xInetAddressGetIsSiteLocal(x.GoPointer())
	return cret
}

var xInetAddressGetNativeSize func(uintptr) uint

// Gets the size of the native raw binary address for @address. This
// is the size of the data that you get from g_inet_address_to_bytes().
func (x *InetAddress) GetNativeSize() uint {

	cret := xInetAddressGetNativeSize(x.GoPointer())
	return cret
}

var xInetAddressToBytes func(uintptr) byte

// Gets the raw binary address data from @address.
func (x *InetAddress) ToBytes() byte {

	cret := xInetAddressToBytes(x.GoPointer())
	return cret
}

var xInetAddressToString func(uintptr) string

// Converts @address to string form.
func (x *InetAddress) ToString() string {

	cret := xInetAddressToString(x.GoPointer())
	return cret
}

func (c *InetAddress) GoPointer() uintptr {
	return c.Ptr
}

func (c *InetAddress) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewInetAddressAny, lib, "g_inet_address_new_any")
	core.PuregoSafeRegister(&xNewInetAddressFromBytes, lib, "g_inet_address_new_from_bytes")
	core.PuregoSafeRegister(&xNewInetAddressFromString, lib, "g_inet_address_new_from_string")
	core.PuregoSafeRegister(&xNewInetAddressLoopback, lib, "g_inet_address_new_loopback")

	core.PuregoSafeRegister(&xInetAddressEqual, lib, "g_inet_address_equal")
	core.PuregoSafeRegister(&xInetAddressGetFamily, lib, "g_inet_address_get_family")
	core.PuregoSafeRegister(&xInetAddressGetIsAny, lib, "g_inet_address_get_is_any")
	core.PuregoSafeRegister(&xInetAddressGetIsLinkLocal, lib, "g_inet_address_get_is_link_local")
	core.PuregoSafeRegister(&xInetAddressGetIsLoopback, lib, "g_inet_address_get_is_loopback")
	core.PuregoSafeRegister(&xInetAddressGetIsMcGlobal, lib, "g_inet_address_get_is_mc_global")
	core.PuregoSafeRegister(&xInetAddressGetIsMcLinkLocal, lib, "g_inet_address_get_is_mc_link_local")
	core.PuregoSafeRegister(&xInetAddressGetIsMcNodeLocal, lib, "g_inet_address_get_is_mc_node_local")
	core.PuregoSafeRegister(&xInetAddressGetIsMcOrgLocal, lib, "g_inet_address_get_is_mc_org_local")
	core.PuregoSafeRegister(&xInetAddressGetIsMcSiteLocal, lib, "g_inet_address_get_is_mc_site_local")
	core.PuregoSafeRegister(&xInetAddressGetIsMulticast, lib, "g_inet_address_get_is_multicast")
	core.PuregoSafeRegister(&xInetAddressGetIsSiteLocal, lib, "g_inet_address_get_is_site_local")
	core.PuregoSafeRegister(&xInetAddressGetNativeSize, lib, "g_inet_address_get_native_size")
	core.PuregoSafeRegister(&xInetAddressToBytes, lib, "g_inet_address_to_bytes")
	core.PuregoSafeRegister(&xInetAddressToString, lib, "g_inet_address_to_string")

}
