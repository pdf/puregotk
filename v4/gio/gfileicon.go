// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type FileIconClass struct {
}

// #GFileIcon specifies an icon by pointing to an image file
// to be used as icon.
type FileIcon struct {
	gobject.Object
}

func FileIconNewFromInternalPtr(ptr uintptr) *FileIcon {
	cls := &FileIcon{}
	cls.Ptr = ptr
	return cls
}

var xNewFileIcon func(uintptr) uintptr

// Creates a new icon for a file.
func NewFileIcon(FileVar File) *FileIcon {
	NewFileIconPtr := xNewFileIcon(FileVar.GoPointer())
	if NewFileIconPtr == 0 {
		return nil
	}

	NewFileIconCls := &FileIcon{}
	NewFileIconCls.Ptr = NewFileIconPtr
	return NewFileIconCls
}

var xFileIconGetFile func(uintptr) uintptr

// Gets the #GFile associated with the given @icon.
func (x *FileIcon) GetFile() *FileBase {

	GetFilePtr := xFileIconGetFile(x.GoPointer())
	if GetFilePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFilePtr)

	GetFileCls := &FileBase{}
	GetFileCls.Ptr = GetFilePtr
	return GetFileCls

}

func (c *FileIcon) GoPointer() uintptr {
	return c.Ptr
}

func (c *FileIcon) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Checks if two icons are equal.
func (x *FileIcon) Equal(Icon2Var Icon) bool {

	return XGIconEqual(x.GoPointer(), Icon2Var.GoPointer())

}

// Serializes a #GIcon into a #GVariant. An equivalent #GIcon can be retrieved
// back by calling g_icon_deserialize() on the returned value.
// As serialization will avoid using raw icon data when possible, it only
// makes sense to transfer the #GVariant between processes on the same machine,
// (as opposed to over the network), and within the same file system namespace.
func (x *FileIcon) Serialize() *glib.Variant {

	return XGIconSerialize(x.GoPointer())

}

// Generates a textual representation of @icon that can be used for
// serialization such as when passing @icon to a different process or
// saving it to persistent storage. Use g_icon_new_for_string() to
// get @icon back from the returned string.
//
// The encoding of the returned string is proprietary to #GIcon except
// in the following two cases
//
//   - If @icon is a #GFileIcon, the returned string is a native path
//     (such as `/path/to/my icon.png`) without escaping
//     if the #GFile for @icon is a native file.  If the file is not
//     native, the returned string is the result of g_file_get_uri()
//     (such as `sftp://path/to/my%20icon.png`).
//
//   - If @icon is a #GThemedIcon with exactly one name and no fallbacks,
//     the encoding is simply the name (such as `network-server`).
func (x *FileIcon) ToString() string {

	return XGIconToString(x.GoPointer())

}

// Loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
func (x *FileIcon) Load(SizeVar int, TypeVar string, CancellableVar *Cancellable) *InputStream {

	LoadPtr := XGLoadableIconLoad(x.GoPointer(), SizeVar, TypeVar, CancellableVar.GoPointer())
	if LoadPtr == 0 {
		return nil
	}

	LoadCls := &InputStream{}
	LoadCls.Ptr = LoadPtr
	return LoadCls

}

// Loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking
// version of this function, see g_loadable_icon_load().
func (x *FileIcon) LoadAsync(SizeVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	XGLoadableIconLoadAsync(x.GoPointer(), SizeVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes an asynchronous icon load started in g_loadable_icon_load_async().
func (x *FileIcon) LoadFinish(ResVar AsyncResult, TypeVar string) *InputStream {

	LoadFinishPtr := XGLoadableIconLoadFinish(x.GoPointer(), ResVar.GoPointer(), TypeVar)
	if LoadFinishPtr == 0 {
		return nil
	}

	LoadFinishCls := &InputStream{}
	LoadFinishCls.Ptr = LoadFinishPtr
	return LoadFinishCls

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFileIcon, lib, "g_file_icon_new")

	core.PuregoSafeRegister(&xFileIconGetFile, lib, "g_file_icon_get_file")

}
