// Package gdk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gdk

import (
	"unsafe"

	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdkpixbuf"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type TextureClass struct {
}

func (x *TextureClass) GoPointer() uintptr {
	return uintptr(unsafe.Pointer(x))
}

// Possible errors that can be returned by `GdkTexture` constructors.
type TextureError int

const (

	// Not enough memory to handle this image
	TextureErrorTooLargeValue TextureError = 0
	// The image data appears corrupted
	TextureErrorCorruptImageValue TextureError = 1
	// The image contains features
	//   that cannot be loaded
	TextureErrorUnsupportedContentValue TextureError = 2
	// The image format is not supported
	TextureErrorUnsupportedFormatValue TextureError = 3
)

// `GdkTexture` is the basic element used to refer to pixel data.
//
// It is primarily meant for pixel data that will not change over
// multiple frames, and will be used for a long time.
//
// There are various ways to create `GdkTexture` objects from a
// [class@GdkPixbuf.Pixbuf], or a Cairo surface, or other pixel data.
//
// The ownership of the pixel data is transferred to the `GdkTexture`
// instance; you can only make a copy of it, via [method@Gdk.Texture.download].
//
// `GdkTexture` is an immutable object: That means you cannot change
// anything about it other than increasing the reference count via
// [method@GObject.Object.ref], and consequently, it is a thread-safe object.
type Texture struct {
	gobject.Object
}

func TextureNewFromInternalPtr(ptr uintptr) *Texture {
	cls := &Texture{}
	cls.Ptr = ptr
	return cls
}

var xNewTextureForPixbuf func(uintptr) uintptr

// Creates a new texture object representing the `GdkPixbuf`.
//
// This function is threadsafe, so that you can e.g. use GTask
// and [method@Gio.Task.run_in_thread] to avoid blocking the main thread
// while loading a big image.
func NewTextureForPixbuf(PixbufVar *gdkpixbuf.Pixbuf) *Texture {
	var cls *Texture

	cret := xNewTextureForPixbuf(PixbufVar.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &Texture{}
	cls.Ptr = cret
	return cls
}

var xNewTextureFromBytes func(*glib.Bytes, **glib.Error) uintptr

// Creates a new texture by loading an image from memory,
//
// The file format is detected automatically. The supported formats
// are PNG and JPEG, though more formats might be available.
//
// If %NULL is returned, then @error will be set.
//
// This function is threadsafe, so that you can e.g. use GTask
// and [method@Gio.Task.run_in_thread] to avoid blocking the main thread
// while loading a big image.
func NewTextureFromBytes(BytesVar *glib.Bytes) (*Texture, error) {
	var cls *Texture
	var cerr *glib.Error

	cret := xNewTextureFromBytes(BytesVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &Texture{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewTextureFromFile func(uintptr, **glib.Error) uintptr

// Creates a new texture by loading an image from a file.
//
// The file format is detected automatically. The supported formats
// are PNG and JPEG, though more formats might be available.
//
// If %NULL is returned, then @error will be set.
//
// This function is threadsafe, so that you can e.g. use GTask
// and [method@Gio.Task.run_in_thread] to avoid blocking the main thread
// while loading a big image.
func NewTextureFromFile(FileVar gio.File) (*Texture, error) {
	var cls *Texture
	var cerr *glib.Error

	cret := xNewTextureFromFile(FileVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &Texture{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewTextureFromFilename func(string, **glib.Error) uintptr

// Creates a new texture by loading an image from a file.
//
// The file format is detected automatically. The supported formats
// are PNG and JPEG, though more formats might be available.
//
// If %NULL is returned, then @error will be set.
//
// This function is threadsafe, so that you can e.g. use GTask
// and [method@Gio.Task.run_in_thread] to avoid blocking the main thread
// while loading a big image.
func NewTextureFromFilename(PathVar string) (*Texture, error) {
	var cls *Texture
	var cerr *glib.Error

	cret := xNewTextureFromFilename(PathVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &Texture{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

var xNewTextureFromResource func(string) uintptr

// Creates a new texture by loading an image from a resource.
//
// The file format is detected automatically. The supported formats
// are PNG and JPEG, though more formats might be available.
//
// It is a fatal error if @resource_path does not specify a valid
// image resource and the program will abort if that happens.
// If you are unsure about the validity of a resource, use
// [ctor@Gdk.Texture.new_from_file] to load it.
//
// This function is threadsafe, so that you can e.g. use GTask
// and [method@Gio.Task.run_in_thread] to avoid blocking the main thread
// while loading a big image.
func NewTextureFromResource(ResourcePathVar string) *Texture {
	var cls *Texture

	cret := xNewTextureFromResource(ResourcePathVar)

	if cret == 0 {
		return nil
	}
	cls = &Texture{}
	cls.Ptr = cret
	return cls
}

var xTextureDownload func(uintptr, uintptr, uint)

// Downloads the @texture into local memory.
//
// This may be an expensive operation, as the actual texture data
// may reside on a GPU or on a remote display server.
//
// The data format of the downloaded data is equivalent to
// %CAIRO_FORMAT_ARGB32, so every downloaded pixel requires
// 4 bytes of memory.
//
// Downloading a texture into a Cairo image surface:
// ```c
// surface = cairo_image_surface_create (CAIRO_FORMAT_ARGB32,
//
//	gdk_texture_get_width (texture),
//	gdk_texture_get_height (texture));
//
// gdk_texture_download (texture,
//
//	cairo_image_surface_get_data (surface),
//	cairo_image_surface_get_stride (surface));
//
// cairo_surface_mark_dirty (surface);
// ```
func (x *Texture) Download(DataVar uintptr, StrideVar uint) {

	xTextureDownload(x.GoPointer(), DataVar, StrideVar)

}

var xTextureGetHeight func(uintptr) int

// Returns the height of the @texture, in pixels.
func (x *Texture) GetHeight() int {

	cret := xTextureGetHeight(x.GoPointer())
	return cret
}

var xTextureGetWidth func(uintptr) int

// Returns the width of @texture, in pixels.
func (x *Texture) GetWidth() int {

	cret := xTextureGetWidth(x.GoPointer())
	return cret
}

var xTextureSaveToPng func(uintptr, string) bool

// Store the given @texture to the @filename as a PNG file.
//
// This is a utility function intended for debugging and testing.
// If you want more control over formats, proper error handling or
// want to store to a [iface@Gio.File] or other location, you might want to
// use [method@Gdk.Texture.save_to_png_bytes] or look into the
// gdk-pixbuf library.
func (x *Texture) SaveToPng(FilenameVar string) bool {

	cret := xTextureSaveToPng(x.GoPointer(), FilenameVar)
	return cret
}

var xTextureSaveToPngBytes func(uintptr) *glib.Bytes

// Store the given @texture in memory as a PNG file.
//
// Use [ctor@Gdk.Texture.new_from_bytes] to read it back.
//
// If you want to serialize a texture, this is a convenient and
// portable way to do that.
//
// If you need more control over the generated image, such as
// attaching metadata, you should look into an image handling
// library such as the gdk-pixbuf library.
//
// If you are dealing with high dynamic range float data, you
// might also want to consider [method@Gdk.Texture.save_to_tiff_bytes]
// instead.
func (x *Texture) SaveToPngBytes() *glib.Bytes {

	cret := xTextureSaveToPngBytes(x.GoPointer())
	return cret
}

var xTextureSaveToTiff func(uintptr, string) bool

// Store the given @texture to the @filename as a TIFF file.
//
// GTK will attempt to store data without loss.
func (x *Texture) SaveToTiff(FilenameVar string) bool {

	cret := xTextureSaveToTiff(x.GoPointer(), FilenameVar)
	return cret
}

var xTextureSaveToTiffBytes func(uintptr) *glib.Bytes

// Store the given @texture in memory as a TIFF file.
//
// Use [ctor@Gdk.Texture.new_from_bytes] to read it back.
//
// This function is intended to store a representation of the
// texture's data that is as accurate as possible. This is
// particularly relevant when working with high dynamic range
// images and floating-point texture data.
//
// If that is not your concern and you are interested in a
// smaller size and a more portable format, you might want to
// use [method@Gdk.Texture.save_to_png_bytes].
func (x *Texture) SaveToTiffBytes() *glib.Bytes {

	cret := xTextureSaveToTiffBytes(x.GoPointer())
	return cret
}

func (c *Texture) GoPointer() uintptr {
	return c.Ptr
}

func (c *Texture) SetGoPointer(ptr uintptr) {
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
func (x *Texture) ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64) {

	XGdkPaintableComputeConcreteSize(x.GoPointer(), SpecifiedWidthVar, SpecifiedHeightVar, DefaultWidthVar, DefaultHeightVar, ConcreteWidthVar, ConcreteHeightVar)

}

// Gets an immutable paintable for the current contents displayed by @paintable.
//
// This is useful when you want to retain the current state of an animation,
// for example to take a screenshot of a running animation.
//
// If the @paintable is already immutable, it will return itself.
func (x *Texture) GetCurrentImage() *PaintableBase {
	var cls *PaintableBase

	cret := XGdkPaintableGetCurrentImage(x.GoPointer())

	if cret == 0 {
		return nil
	}
	cls = &PaintableBase{}
	cls.Ptr = cret
	return cls
}

// Get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See [flags@Gdk.PaintableFlags] for the flags and what they mean.
func (x *Texture) GetFlags() PaintableFlags {

	cret := XGdkPaintableGetFlags(x.GoPointer())
	return cret
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
func (x *Texture) GetIntrinsicAspectRatio() float64 {

	cret := XGdkPaintableGetIntrinsicAspectRatio(x.GoPointer())
	return cret
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
func (x *Texture) GetIntrinsicHeight() int {

	cret := XGdkPaintableGetIntrinsicHeight(x.GoPointer())
	return cret
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
func (x *Texture) GetIntrinsicWidth() int {

	cret := XGdkPaintableGetIntrinsicWidth(x.GoPointer())
	return cret
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
func (x *Texture) InvalidateContents() {

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
func (x *Texture) InvalidateSize() {

	XGdkPaintableInvalidateSize(x.GoPointer())

}

// Snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot.
// If @width and @height are not larger than zero, this function will
// do nothing.
func (x *Texture) Snapshot(SnapshotVar *Snapshot, WidthVar float64, HeightVar float64) {

	XGdkPaintableSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar)

}

// Checks if two icons are equal.
func (x *Texture) Equal(Icon2Var gio.Icon) bool {

	cret := gio.XGIconEqual(x.GoPointer(), Icon2Var.GoPointer())
	return cret
}

// Serializes a #GIcon into a #GVariant. An equivalent #GIcon can be retrieved
// back by calling g_icon_deserialize() on the returned value.
// As serialization will avoid using raw icon data when possible, it only
// makes sense to transfer the #GVariant between processes on the same machine,
// (as opposed to over the network), and within the same file system namespace.
func (x *Texture) Serialize() *glib.Variant {

	cret := gio.XGIconSerialize(x.GoPointer())
	return cret
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
func (x *Texture) ToString() string {

	cret := gio.XGIconToString(x.GoPointer())
	return cret
}

// Loads a loadable icon. For the asynchronous version of this function,
// see g_loadable_icon_load_async().
func (x *Texture) Load(SizeVar int, TypeVar string, CancellableVar *gio.Cancellable) (*gio.InputStream, error) {
	var cls *gio.InputStream
	var cerr *glib.Error

	cret := gio.XGLoadableIconLoad(x.GoPointer(), SizeVar, TypeVar, CancellableVar.GoPointer(), &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gio.InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

// Loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking
// version of this function, see g_loadable_icon_load().
func (x *Texture) LoadAsync(SizeVar int, CancellableVar *gio.Cancellable, CallbackVar gio.AsyncReadyCallback, UserDataVar uintptr) {

	gio.XGLoadableIconLoadAsync(x.GoPointer(), SizeVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

// Finishes an asynchronous icon load started in g_loadable_icon_load_async().
func (x *Texture) LoadFinish(ResVar gio.AsyncResult, TypeVar string) (*gio.InputStream, error) {
	var cls *gio.InputStream
	var cerr *glib.Error

	cret := gio.XGLoadableIconLoadFinish(x.GoPointer(), ResVar.GoPointer(), TypeVar, &cerr)

	if cret == 0 {
		return nil, cerr
	}
	cls = &gio.InputStream{}
	cls.Ptr = cret
	if cerr == nil {
		return cls, nil
	}
	return cls, cerr

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GDK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewTextureForPixbuf, lib, "gdk_texture_new_for_pixbuf")
	core.PuregoSafeRegister(&xNewTextureFromBytes, lib, "gdk_texture_new_from_bytes")
	core.PuregoSafeRegister(&xNewTextureFromFile, lib, "gdk_texture_new_from_file")
	core.PuregoSafeRegister(&xNewTextureFromFilename, lib, "gdk_texture_new_from_filename")
	core.PuregoSafeRegister(&xNewTextureFromResource, lib, "gdk_texture_new_from_resource")

	core.PuregoSafeRegister(&xTextureDownload, lib, "gdk_texture_download")
	core.PuregoSafeRegister(&xTextureGetHeight, lib, "gdk_texture_get_height")
	core.PuregoSafeRegister(&xTextureGetWidth, lib, "gdk_texture_get_width")
	core.PuregoSafeRegister(&xTextureSaveToPng, lib, "gdk_texture_save_to_png")
	core.PuregoSafeRegister(&xTextureSaveToPngBytes, lib, "gdk_texture_save_to_png_bytes")
	core.PuregoSafeRegister(&xTextureSaveToTiff, lib, "gdk_texture_save_to_tiff")
	core.PuregoSafeRegister(&xTextureSaveToTiffBytes, lib, "gdk_texture_save_to_tiff_bytes")

}
