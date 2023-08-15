// Package glib was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package glib

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

var xBitLock func(int, int)

// Sets the indicated @lock_bit in @address.  If the bit is already
// set, this call will block until g_bit_unlock() unsets the
// corresponding bit.
//
// Attempting to lock on two different bits within the same integer is
// not supported and will very probably cause deadlocks.
//
// The value of the bit that is set is (1u &lt;&lt; @bit).  If @bit is not
// between 0 and 31 then the result is undefined.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably. While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func BitLock(AddressVar int, LockBitVar int) {

	xBitLock(AddressVar, LockBitVar)

}

var xBitTrylock func(int, int) bool

// Sets the indicated @lock_bit in @address, returning %TRUE if
// successful.  If the bit is already set, returns %FALSE immediately.
//
// Attempting to lock on two different bits within the same integer is
// not supported.
//
// The value of the bit that is set is (1u &lt;&lt; @bit).  If @bit is not
// between 0 and 31 then the result is undefined.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably. While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func BitTrylock(AddressVar int, LockBitVar int) bool {

	return xBitTrylock(AddressVar, LockBitVar)

}

var xBitUnlock func(int, int)

// Clears the indicated @lock_bit in @address.  If another thread is
// currently blocked in g_bit_lock() on this same bit then it will be
// woken up.
//
// This function accesses @address atomically.  All other accesses to
// @address must be atomic in order for this function to work
// reliably. While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func BitUnlock(AddressVar int, LockBitVar int) {

	xBitUnlock(AddressVar, LockBitVar)

}

var xPointerBitLock func(uintptr, int)

// This is equivalent to g_bit_lock, but working on pointers (or other
// pointer-sized values).
//
// For portability reasons, you may only lock on the bottom 32 bits of
// the pointer.
//
// While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func PointerBitLock(AddressVar uintptr, LockBitVar int) {

	xPointerBitLock(AddressVar, LockBitVar)

}

var xPointerBitTrylock func(uintptr, int) bool

// This is equivalent to g_bit_trylock(), but working on pointers (or
// other pointer-sized values).
//
// For portability reasons, you may only lock on the bottom 32 bits of
// the pointer.
//
// While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func PointerBitTrylock(AddressVar uintptr, LockBitVar int) bool {

	return xPointerBitTrylock(AddressVar, LockBitVar)

}

var xPointerBitUnlock func(uintptr, int)

// This is equivalent to g_bit_unlock, but working on pointers (or other
// pointer-sized values).
//
// For portability reasons, you may only lock on the bottom 32 bits of
// the pointer.
//
// While @address has a `volatile` qualifier, this is a historical
// artifact and the argument passed to it should not be `volatile`.
func PointerBitUnlock(AddressVar uintptr, LockBitVar int) {

	xPointerBitUnlock(AddressVar, LockBitVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GLIB"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xBitLock, lib, "g_bit_lock")
	core.PuregoSafeRegister(&xBitTrylock, lib, "g_bit_trylock")
	core.PuregoSafeRegister(&xBitUnlock, lib, "g_bit_unlock")
	core.PuregoSafeRegister(&xPointerBitLock, lib, "g_pointer_bit_lock")
	core.PuregoSafeRegister(&xPointerBitTrylock, lib, "g_pointer_bit_trylock")
	core.PuregoSafeRegister(&xPointerBitUnlock, lib, "g_pointer_bit_unlock")

}
