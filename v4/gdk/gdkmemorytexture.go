// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type MemoryTextureClass struct {
}

// A `GdkTexture` representing image data in memory.
type MemoryTexture struct {
	Texture
}

func MemoryTextureNewFromInternalPtr(ptr uintptr) *MemoryTexture {
	cls := &MemoryTexture{}
	cls.Ptr = ptr
	return cls
}

var xNewMemoryTexture func(int, int, MemoryFormat, *glib.Bytes, uint) uintptr

// Creates a new texture for a blob of image data.
//
// The `GBytes` must contain @stride × @height pixels
// in the given format.
func NewMemoryTexture(WidthVar int, HeightVar int, FormatVar MemoryFormat, BytesVar *glib.Bytes, StrideVar uint) *MemoryTexture {
	NewMemoryTexturePtr := xNewMemoryTexture(WidthVar, HeightVar, FormatVar, BytesVar, StrideVar)
	if NewMemoryTexturePtr == 0 {
		return nil
	}

	NewMemoryTextureCls := &MemoryTexture{}
	NewMemoryTextureCls.Ptr = NewMemoryTexturePtr
	return NewMemoryTextureCls
}

func (c *MemoryTexture) GoPointer() uintptr {
	return c.Ptr
}

func (c *MemoryTexture) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Compute a concrete size for the `GdkPaintable`.
//
// Applies the sizing algorithm outlined in the
// [CSS Image spec](https://drafts.csswg.org/css-images-3/#default-sizing)
// to the given @paintable. See that link for more details.
//
// It is not necessary to call this function when both @specified_width
// and @specified_height are known, but it is useful to call this
// function in GtkWidget:measure implementations to compute the
// other dimension when only one dimension is given.
func (x *MemoryTexture) ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64) {

	XGdkPaintableComputeConcreteSize(x.GoPointer(), SpecifiedWidthVar, SpecifiedHeightVar, DefaultWidthVar, DefaultHeightVar, ConcreteWidthVar, ConcreteHeightVar)

}

// Gets an immutable paintable for the current contents displayed by @paintable.
//
// This is useful when you want to retain the current state of an animation,
// for example to take a screenshot of a running animation.
//
// If the @paintable is already immutable, it will return itself.
func (x *MemoryTexture) GetCurrentImage() *PaintableBase {

	GetCurrentImagePtr := XGdkPaintableGetCurrentImage(x.GoPointer())
	if GetCurrentImagePtr == 0 {
		return nil
	}

	GetCurrentImageCls := &PaintableBase{}
	GetCurrentImageCls.Ptr = GetCurrentImagePtr
	return GetCurrentImageCls

}

// Get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See [flags@Gdk.PaintableFlags] for the flags and what they mean.
func (x *MemoryTexture) GetFlags() PaintableFlags {

	return XGdkPaintableGetFlags(x.GoPointer())

}

// Gets the preferred aspect ratio the @paintable would like to be displayed at.
//
// The aspect ratio is the width divided by the height, so a value of 0.5
// means that the @paintable prefers to be displayed twice as high as it
// is wide. Consumers of this interface can use this to preserve aspect
// ratio when displaying the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// Usually when a @paintable returns nonzero values from
// [method@Gdk.Paintable.get_intrinsic_width] and
// [method@Gdk.Paintable.get_intrinsic_height] the aspect ratio
// should conform to those values, though that is not required.
//
// If the @paintable does not have a preferred aspect ratio,
// it returns 0. Negative values are never returned.
func (x *MemoryTexture) GetIntrinsicAspectRatio() float64 {

	return XGdkPaintableGetIntrinsicAspectRatio(x.GoPointer())

}

// Gets the preferred height the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred height, it returns 0.
// Negative values are never returned.
func (x *MemoryTexture) GetIntrinsicHeight() int {

	return XGdkPaintableGetIntrinsicHeight(x.GoPointer())

}

// Gets the preferred width the @paintable would like to be displayed at.
//
// Consumers of this interface can use this to reserve enough space to draw
// the paintable.
//
// This is a purely informational value and does not in any way limit the
// values that may be passed to [method@Gdk.Paintable.snapshot].
//
// If the @paintable does not have a preferred width, it returns 0.
// Negative values are never returned.
func (x *MemoryTexture) GetIntrinsicWidth() int {

	return XGdkPaintableGetIntrinsicWidth(x.GoPointer())

}

// Called by implementations of `GdkPaintable` to invalidate their contents.
//
// Unless the contents are invalidated, implementations must guarantee that
// multiple calls of [method@Gdk.Paintable.snapshot] produce the same output.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-contents]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_CONTENTS flag,
// it must not call this function.
func (x *MemoryTexture) InvalidateContents() {

	XGdkPaintableInvalidateContents(x.GoPointer())

}

// Called by implementations of `GdkPaintable` to invalidate their size.
//
// As long as the size is not invalidated, @paintable must return the same
// values for its intrinsic width, height and aspect ratio.
//
// This function will emit the [signal@Gdk.Paintable::invalidate-size]
// signal.
//
// If a @paintable reports the %GDK_PAINTABLE_STATIC_SIZE flag,
// it must not call this function.
func (x *MemoryTexture) InvalidateSize() {

	XGdkPaintableInvalidateSize(x.GoPointer())

}

// Snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot.
// If @width and @height are not larger than zero, this function will
// do nothing.
func (x *MemoryTexture) Snapshot(SnapshotVar *Snapshot, WidthVar float64, HeightVar float64) {

	XGdkPaintableSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar)

}

// Checks if two icons are equal.
func (x *MemoryTexture) Equal(Icon2Var gio.Icon) bool {

	return gio.XGIconEqual(x.GoPointer(), Icon2Var.GoPointer())

}

// Serializes a #GIcon into a #GVariant. An equivalent #GIcon can be retrieved
// back by calling g_icon_deserialize() on the returned value.
// As serialization will avoid using raw icon data when possible, it only
// makes sense to transfer the #GVariant between processes on the same machine,
// (as opposed to over the network), and within the same file system namespace.
func (x *MemoryTexture) Serialize() *glib.Variant {

	return gio.XGIconSerialize(x.GoPointer())

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
func (x *MemoryTexture) ToString() string {

	return gio.XGIconToString(x.GoPointer())

}

// Loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
func (x *MemoryTexture) Load(SizeVar int, TypeVar string, CancellableVar *gio.Cancellable) *gio.InputStream {

	LoadPtr := gio.XGLoadableIconLoad(x.GoPointer(), SizeVar, TypeVar, CancellableVar.GoPointer())
	if LoadPtr == 0 {
		return nil
	}

	LoadCls := &gio.InputStream{}
	LoadCls.Ptr = LoadPtr
	return LoadCls

}

// Loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking
// version of this function, see g_loadable_icon_load().
func (x *MemoryTexture) LoadAsync(SizeVar int, CancellableVar *gio.Cancellable, CallbackVar gio.AsyncReadyCallback, UserDataVar uintptr) {

	gio.XGLoadableIconLoadAsync(x.GoPointer(), SizeVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes an asynchronous icon load started in g_loadable_icon_load_async().
func (x *MemoryTexture) LoadFinish(ResVar gio.AsyncResult, TypeVar string) *gio.InputStream {

	LoadFinishPtr := gio.XGLoadableIconLoadFinish(x.GoPointer(), ResVar.GoPointer(), TypeVar)
	if LoadFinishPtr == 0 {
		return nil
	}

	LoadFinishCls := &gio.InputStream{}
	LoadFinishCls.Ptr = LoadFinishPtr
	return LoadFinishCls

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMemoryTexture, lib, "gdk_memory_texture_new")

}
