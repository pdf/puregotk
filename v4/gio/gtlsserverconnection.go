// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

// vtable for a #GTlsServerConnection implementation.
type TlsServerConnectionInterface struct {
	GIface uintptr
}

// #GTlsServerConnection is the server-side subclass of #GTlsConnection,
// representing a server-side TLS connection.
type TlsServerConnection interface {
	GoPointer() uintptr
	SetGoPointer(uintptr)
}
type TlsServerConnectionBase struct {
	Ptr uintptr
}

func (x *TlsServerConnectionBase) GoPointer() uintptr {
	return x.Ptr
}

func (x *TlsServerConnectionBase) SetGoPointer(ptr uintptr) {
	x.Ptr = ptr
}

var xTlsServerConnectionNew func(uintptr, uintptr, **glib.Error) uintptr

// Creates a new #GTlsServerConnection wrapping @base_io_stream (which
// must have pollable input and output streams).
//
// See the documentation for #GTlsConnection:base-io-stream for restrictions
// on when application code can run operations on the @base_io_stream after
// this function has returned.
func TlsServerConnectionNew(BaseIoStreamVar *IOStream, CertificateVar *TlsCertificate) (*TlsServerConnectionBase, error) {
	var cls *TlsServerConnectionBase
	var cerr *glib.Error

	cret := xTlsServerConnectionNew(BaseIoStreamVar.GoPointer(), CertificateVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &TlsServerConnectionBase{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xTlsServerConnectionNew, lib, "g_tls_server_connection_new")

}
