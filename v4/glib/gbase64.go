// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xBase64Decode func(string, uint) uintptr

// Decode a sequence of Base-64 encoded text into binary data.  Note
// that the returned binary data is not necessarily zero-terminated,
// so it should not be used as a character string.
func Base64Decode(TextVar string, OutLenVar uint) uintptr {

	return xBase64Decode(TextVar, OutLenVar)

}

var xBase64DecodeInplace func(uintptr, uint) byte

// Decode a sequence of Base-64 encoded text into binary data
// by overwriting the input data.
func Base64DecodeInplace(TextVar uintptr, OutLenVar uint) byte {

	return xBase64DecodeInplace(TextVar, OutLenVar)

}

var xBase64DecodeStep func(uintptr, uint, uintptr, int, uint) uint

// Incrementally decode a sequence of binary data from its Base-64 stringified
// representation. By calling this function multiple times you can convert
// data in chunks to avoid having to have the full encoded data in memory.
//
// The output buffer must be large enough to fit all the data that will
// be written to it. Since base64 encodes 3 bytes in 4 chars you need
// at least: (@len / 4) * 3 + 3 bytes (+ 3 may be needed in case of non-zero
// state).
func Base64DecodeStep(InVar uintptr, LenVar uint, OutVar uintptr, StateVar int, SaveVar uint) uint {

	return xBase64DecodeStep(InVar, LenVar, OutVar, StateVar, SaveVar)

}

var xBase64Encode func(uintptr, uint) string

// Encode a sequence of binary data into its Base-64 stringified
// representation.
func Base64Encode(DataVar uintptr, LenVar uint) string {

	return xBase64Encode(DataVar, LenVar)

}

var xBase64EncodeClose func(bool, uintptr, int, int) uint

// Flush the status from a sequence of calls to g_base64_encode_step().
//
// The output buffer must be large enough to fit all the data that will
// be written to it. It will need up to 4 bytes, or up to 5 bytes if
// line-breaking is enabled.
//
// The @out array will not be automatically nul-terminated.
func Base64EncodeClose(BreakLinesVar bool, OutVar uintptr, StateVar int, SaveVar int) uint {

	return xBase64EncodeClose(BreakLinesVar, OutVar, StateVar, SaveVar)

}

var xBase64EncodeStep func(uintptr, uint, bool, uintptr, int, int) uint

// Incrementally encode a sequence of binary data into its Base-64 stringified
// representation. By calling this function multiple times you can convert
// data in chunks to avoid having to have the full encoded data in memory.
//
// When all of the data has been converted you must call
// g_base64_encode_close() to flush the saved state.
//
// The output buffer must be large enough to fit all the data that will
// be written to it. Due to the way base64 encodes you will need
// at least: (@len / 3 + 1) * 4 + 4 bytes (+ 4 may be needed in case of
// non-zero state). If you enable line-breaking you will need at least:
// ((@len / 3 + 1) * 4 + 4) / 76 + 1 bytes of extra space.
//
// @break_lines is typically used when putting base64-encoded data in emails.
// It breaks the lines at 76 columns instead of putting all of the text on
// the same line. This avoids problems with long lines in the email system.
// Note however that it breaks the lines with `LF` characters, not
// `CR LF` sequences, so the result cannot be passed directly to SMTP
// or certain other protocols.
func Base64EncodeStep(InVar uintptr, LenVar uint, BreakLinesVar bool, OutVar uintptr, StateVar int, SaveVar int) uint {

	return xBase64EncodeStep(InVar, LenVar, BreakLinesVar, OutVar, StateVar, SaveVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBase64Decode, lib, "g_base64_decode")
	core.PuregoSafeRegister(&xBase64DecodeInplace, lib, "g_base64_decode_inplace")
	core.PuregoSafeRegister(&xBase64DecodeStep, lib, "g_base64_decode_step")
	core.PuregoSafeRegister(&xBase64Encode, lib, "g_base64_encode")
	core.PuregoSafeRegister(&xBase64EncodeClose, lib, "g_base64_encode_close")
	core.PuregoSafeRegister(&xBase64EncodeStep, lib, "g_base64_encode_step")

}
