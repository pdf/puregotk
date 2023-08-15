// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of function which is called when an extended
// error instance is freed. It is passed the error pointer about to be
// freed, and should free the error's private data fields.
//
// Normally, it is better to use G_DEFINE_EXTENDED_ERROR(), as it
// already takes care of getting the private data from @error.
type ErrorClearFunc func(*Error)

// Specifies the type of function which is called when an extended
// error instance is copied. It is passed the pointer to the
// destination error and source error, and should copy only the fields
// of the private data from @src_error to @dest_error.
//
// Normally, it is better to use G_DEFINE_EXTENDED_ERROR(), as it
// already takes care of getting the private data from @src_error and
// @dest_error.
type ErrorCopyFunc func(*Error, *Error)

// Specifies the type of function which is called just after an
// extended error instance is created and its fields filled. It should
// only initialize the fields in the private data, which can be
// received with the generated `*_get_private()` function.
//
// Normally, it is better to use G_DEFINE_EXTENDED_ERROR(), as it
// already takes care of getting the private data from @error.
type ErrorInitFunc func(*Error)

// The `GError` structure contains information about
// an error that has occurred.
type Error struct {
	Domain Quark

	Code int32

	Message string
}

var xClearError func()

// If @err or *@err is %NULL, does nothing. Otherwise,
// calls g_error_free() on *@err and sets *@err to %NULL.
func ClearError() {

	xClearError()

}

var xPrefixError func(**Error, string, ...interface{})

// Formats a string according to @format and prefix it to an existing
// error message. If @err is %NULL (ie: no error variable) then do
// nothing.
//
// If *@err is %NULL (ie: an error variable is present but there is no
// error condition) then also do nothing.
func PrefixError(ErrVar **Error, FormatVar string, varArgs ...interface{}) {

	xPrefixError(ErrVar, FormatVar, varArgs...)

}

var xPrefixErrorLiteral func(**Error, string)

// Prefixes @prefix to an existing error message. If @err or *@err is
// %NULL (i.e.: no error variable) then do nothing.
func PrefixErrorLiteral(ErrVar **Error, PrefixVar string) {

	xPrefixErrorLiteral(ErrVar, PrefixVar)

}

var xPropagateError func(**Error, *Error)

// If @dest is %NULL, free @src; otherwise, moves @src into *@dest.
// The error variable @dest points to must be %NULL.
//
// @src must be non-%NULL.
//
// Note that @src is no longer valid after this call. If you want
// to keep using the same GError*, you need to set it to %NULL
// after calling this function on it.
func PropagateError(DestVar **Error, SrcVar *Error) {

	xPropagateError(DestVar, SrcVar)

}

var xPropagatePrefixedError func(**Error, *Error, string, ...interface{})

// If @dest is %NULL, free @src; otherwise, moves @src into *@dest.
// *@dest must be %NULL. After the move, add a prefix as with
// g_prefix_error().
func PropagatePrefixedError(DestVar **Error, SrcVar *Error, FormatVar string, varArgs ...interface{}) {

	xPropagatePrefixedError(DestVar, SrcVar, FormatVar, varArgs...)

}

var xSetError func(**Error, Quark, int, string, ...interface{})

// Does nothing if @err is %NULL; if @err is non-%NULL, then *@err
// must be %NULL. A new #GError is created and assigned to *@err.
func SetError(ErrVar **Error, DomainVar Quark, CodeVar int, FormatVar string, varArgs ...interface{}) {

	xSetError(ErrVar, DomainVar, CodeVar, FormatVar, varArgs...)

}

var xSetErrorLiteral func(**Error, Quark, int, string)

// Does nothing if @err is %NULL; if @err is non-%NULL, then *@err
// must be %NULL. A new #GError is created and assigned to *@err.
// Unlike g_set_error(), @message is not a printf()-style format string.
// Use this function if @message contains text you don't have control over,
// that could include printf() escape sequences.
func SetErrorLiteral(ErrVar **Error, DomainVar Quark, CodeVar int, MessageVar string) {

	xSetErrorLiteral(ErrVar, DomainVar, CodeVar, MessageVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xClearError, lib, "g_clear_error")
	core.PuregoSafeRegister(&xPrefixError, lib, "g_prefix_error")
	core.PuregoSafeRegister(&xPrefixErrorLiteral, lib, "g_prefix_error_literal")
	core.PuregoSafeRegister(&xPropagateError, lib, "g_propagate_error")
	core.PuregoSafeRegister(&xPropagatePrefixedError, lib, "g_propagate_prefixed_error")
	core.PuregoSafeRegister(&xSetError, lib, "g_set_error")
	core.PuregoSafeRegister(&xSetErrorLiteral, lib, "g_set_error_literal")

}
