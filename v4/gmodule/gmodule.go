// Package gmodule was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gmodule

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

// Specifies the type of the module initialization function.
// If a module contains a function named g_module_check_init() it is called
// automatically when the module is loaded. It is passed the #GModule structure
// and should return %NULL on success or a string describing the initialization
// error.
type ModuleCheckInit func(*Module) string

// Specifies the type of the module function called when it is unloaded.
// If a module contains a function named g_module_unload() it is called
// automatically when the module is unloaded.
// It is passed the #GModule structure.
type ModuleUnload func(*Module)

// The #GModule struct is an opaque data structure to represent a
// [dynamically-loaded module][glib-Dynamic-Loading-of-Modules].
// It should only be accessed via the following functions.
type Module struct {
}

// Flags passed to g_module_open().
// Note that these flags are not supported on all platforms.
type ModuleFlags int

const (

	// specifies that symbols are only resolved when
	//     needed. The default action is to bind all symbols when the module
	//     is loaded.
	GModuleBindLazyValue ModuleFlags = 1
	// specifies that symbols in the module should
	//     not be added to the global name space. The default action on most
	//     platforms is to place symbols in the module in the global name space,
	//     which may cause conflicts with existing symbols.
	GModuleBindLocalValue ModuleFlags = 2
	// mask for all flags.
	GModuleBindMaskValue ModuleFlags = 3
)

// Errors returned by g_module_open_full().
type ModuleError int

const (

	// there was an error loading or opening a module file
	GModuleErrorFailedValue ModuleError = 0
	// a module returned an error from its `g_module_check_init()` function
	GModuleErrorCheckFailedValue ModuleError = 1
)

var xModuleBuildPath func(string, string) string

// A portable way to build the filename of a module. The platform-specific
// prefix and suffix are added to the filename, if needed, and the result
// is added to the directory, using the correct separator character.
//
// The directory should specify the directory where the module can be found.
// It can be %NULL or an empty string to indicate that the module is in a
// standard platform-specific directory, though this is not recommended
// since the wrong module may be found.
//
// For example, calling g_module_build_path() on a Linux system with a
// @directory of `/lib` and a @module_name of "mylibrary" will return
// `/lib/libmylibrary.so`. On a Windows system, using `\Windows` as the
// directory it will return `\Windows\mylibrary.dll`.
func ModuleBuildPath(DirectoryVar string, ModuleNameVar string) string {

	cret := xModuleBuildPath(DirectoryVar, ModuleNameVar)
	return cret
}

var xNewModuleError func() string

// Gets a string describing the last module error.
func NewModuleError() string {

	cret := xNewModuleError()
	return cret
}

var xModuleSupported func() bool

// Checks if modules are supported on the current platform.
func ModuleSupported() bool {

	cret := xModuleSupported()
	return cret
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GMODULE"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	core.PuregoSafeRegister(&xModuleBuildPath, lib, "g_module_build_path")
	core.PuregoSafeRegister(&xNewModuleError, lib, "g_module_error")
	core.PuregoSafeRegister(&xModuleSupported, lib, "g_module_supported")

}
