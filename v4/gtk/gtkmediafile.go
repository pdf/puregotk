// Package gtk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gtk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/gdk"
	"github.com/jwijenbergh/puregotk/v4/gio"
	"github.com/jwijenbergh/puregotk/v4/gobject"
)

type MediaFileClass struct {
	ParentClass uintptr
}

// `GtkMediaFile` implements `GtkMediaStream` for files.
//
// This provides a simple way to play back video files with GTK.
//
// GTK provides a GIO extension point for `GtkMediaFile` implementations
// to allow for external implementations using various media frameworks.
//
// GTK itself includes implementations using GStreamer and ffmpeg.
type MediaFile struct {
	MediaStream
}

func MediaFileNewFromInternalPtr(ptr uintptr) *MediaFile {
	cls := &MediaFile{}
	cls.Ptr = ptr
	return cls
}

var xNewMediaFile func() uintptr

// Creates a new empty media file.
func NewMediaFile() *MediaFile {
	NewMediaFilePtr := xNewMediaFile()
	if NewMediaFilePtr == 0 {
		return nil
	}

	NewMediaFileCls := &MediaFile{}
	NewMediaFileCls.Ptr = NewMediaFilePtr
	return NewMediaFileCls
}

var xNewForFileMediaFile func(uintptr) uintptr

// Creates a new media file to play @file.
func NewForFileMediaFile(FileVar gio.File) *MediaFile {
	NewForFileMediaFilePtr := xNewForFileMediaFile(FileVar.GoPointer())
	if NewForFileMediaFilePtr == 0 {
		return nil
	}

	NewForFileMediaFileCls := &MediaFile{}
	NewForFileMediaFileCls.Ptr = NewForFileMediaFilePtr
	return NewForFileMediaFileCls
}

var xNewForFilenameMediaFile func(string) uintptr

// Creates a new media file for the given filename.
//
// This is a utility function that converts the given @filename
// to a `GFile` and calls [ctor@Gtk.MediaFile.new_for_file].
func NewForFilenameMediaFile(FilenameVar string) *MediaFile {
	NewForFilenameMediaFilePtr := xNewForFilenameMediaFile(FilenameVar)
	if NewForFilenameMediaFilePtr == 0 {
		return nil
	}

	NewForFilenameMediaFileCls := &MediaFile{}
	NewForFilenameMediaFileCls.Ptr = NewForFilenameMediaFilePtr
	return NewForFilenameMediaFileCls
}

var xNewForInputStreamMediaFile func(uintptr) uintptr

// Creates a new media file to play @stream.
//
// If you want the resulting media to be seekable,
// the stream should implement the `GSeekable` interface.
func NewForInputStreamMediaFile(StreamVar *gio.InputStream) *MediaFile {
	NewForInputStreamMediaFilePtr := xNewForInputStreamMediaFile(StreamVar.GoPointer())
	if NewForInputStreamMediaFilePtr == 0 {
		return nil
	}

	NewForInputStreamMediaFileCls := &MediaFile{}
	NewForInputStreamMediaFileCls.Ptr = NewForInputStreamMediaFilePtr
	return NewForInputStreamMediaFileCls
}

var xNewForResourceMediaFile func(string) uintptr

// Creates a new new media file for the given resource.
//
// This is a utility function that converts the given @resource
// to a `GFile` and calls [ctor@Gtk.MediaFile.new_for_file].
func NewForResourceMediaFile(ResourcePathVar string) *MediaFile {
	NewForResourceMediaFilePtr := xNewForResourceMediaFile(ResourcePathVar)
	if NewForResourceMediaFilePtr == 0 {
		return nil
	}

	NewForResourceMediaFileCls := &MediaFile{}
	NewForResourceMediaFileCls.Ptr = NewForResourceMediaFilePtr
	return NewForResourceMediaFileCls
}

var xMediaFileClear func(uintptr)

// Resets the media file to be empty.
func (x *MediaFile) Clear() {

	xMediaFileClear(x.GoPointer())

}

var xMediaFileGetFile func(uintptr) uintptr

// Returns the file that @self is currently playing from.
//
// When @self is not playing or not playing from a file,
// %NULL is returned.
func (x *MediaFile) GetFile() *gio.FileBase {

	GetFilePtr := xMediaFileGetFile(x.GoPointer())
	if GetFilePtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetFilePtr)

	GetFileCls := &gio.FileBase{}
	GetFileCls.Ptr = GetFilePtr
	return GetFileCls

}

var xMediaFileGetInputStream func(uintptr) uintptr

// Returns the stream that @self is currently playing from.
//
// When @self is not playing or not playing from a stream,
// %NULL is returned.
func (x *MediaFile) GetInputStream() *gio.InputStream {

	GetInputStreamPtr := xMediaFileGetInputStream(x.GoPointer())
	if GetInputStreamPtr == 0 {
		return nil
	}

	gobject.IncreaseRef(GetInputStreamPtr)

	GetInputStreamCls := &gio.InputStream{}
	GetInputStreamCls.Ptr = GetInputStreamPtr
	return GetInputStreamCls

}

var xMediaFileSetFile func(uintptr, uintptr)

// Sets the `GtkMediaFile` to play the given file.
//
// If any file is still playing, stop playing it.
func (x *MediaFile) SetFile(FileVar gio.File) {

	xMediaFileSetFile(x.GoPointer(), FileVar.GoPointer())

}

var xMediaFileSetFilename func(uintptr, string)

// Sets the `GtkMediaFile to play the given file.
//
// This is a utility function that converts the given @filename
// to a `GFile` and calls [method@Gtk.MediaFile.set_file].
func (x *MediaFile) SetFilename(FilenameVar string) {

	xMediaFileSetFilename(x.GoPointer(), FilenameVar)

}

var xMediaFileSetInputStream func(uintptr, uintptr)

// Sets the `GtkMediaFile` to play the given stream.
//
// If anything is still playing, stop playing it.
//
// Full control about the @stream is assumed for the duration of
// playback. The stream will not be closed.
func (x *MediaFile) SetInputStream(StreamVar *gio.InputStream) {

	xMediaFileSetInputStream(x.GoPointer(), StreamVar.GoPointer())

}

var xMediaFileSetResource func(uintptr, string)

// Sets the `GtkMediaFile to play the given resource.
//
// This is a utility function that converts the given @resource_path
// to a `GFile` and calls [method@Gtk.MediaFile.set_file].
func (x *MediaFile) SetResource(ResourcePathVar string) {

	xMediaFileSetResource(x.GoPointer(), ResourcePathVar)

}

func (c *MediaFile) GoPointer() uintptr {
	return c.Ptr
}

func (c *MediaFile) SetGoPointer(ptr uintptr) {
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
func (x *MediaFile) ComputeConcreteSize(SpecifiedWidthVar float64, SpecifiedHeightVar float64, DefaultWidthVar float64, DefaultHeightVar float64, ConcreteWidthVar float64, ConcreteHeightVar float64) {

	gdk.XGdkPaintableComputeConcreteSize(x.GoPointer(), SpecifiedWidthVar, SpecifiedHeightVar, DefaultWidthVar, DefaultHeightVar, ConcreteWidthVar, ConcreteHeightVar)

}

// Gets an immutable paintable for the current contents displayed by @paintable.
//
// This is useful when you want to retain the current state of an animation,
// for example to take a screenshot of a running animation.
//
// If the @paintable is already immutable, it will return itself.
func (x *MediaFile) GetCurrentImage() *gdk.PaintableBase {

	GetCurrentImagePtr := gdk.XGdkPaintableGetCurrentImage(x.GoPointer())
	if GetCurrentImagePtr == 0 {
		return nil
	}

	GetCurrentImageCls := &gdk.PaintableBase{}
	GetCurrentImageCls.Ptr = GetCurrentImagePtr
	return GetCurrentImageCls

}

// Get flags for the paintable.
//
// This is oftentimes useful for optimizations.
//
// See [flags@Gdk.PaintableFlags] for the flags and what they mean.
func (x *MediaFile) GetFlags() gdk.PaintableFlags {

	return gdk.XGdkPaintableGetFlags(x.GoPointer())

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
func (x *MediaFile) GetIntrinsicAspectRatio() float64 {

	return gdk.XGdkPaintableGetIntrinsicAspectRatio(x.GoPointer())

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
func (x *MediaFile) GetIntrinsicHeight() int {

	return gdk.XGdkPaintableGetIntrinsicHeight(x.GoPointer())

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
func (x *MediaFile) GetIntrinsicWidth() int {

	return gdk.XGdkPaintableGetIntrinsicWidth(x.GoPointer())

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
func (x *MediaFile) InvalidateContents() {

	gdk.XGdkPaintableInvalidateContents(x.GoPointer())

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
func (x *MediaFile) InvalidateSize() {

	gdk.XGdkPaintableInvalidateSize(x.GoPointer())

}

// Snapshots the given paintable with the given @width and @height.
//
// The paintable is drawn at the current (0,0) offset of the @snapshot.
// If @width and @height are not larger than zero, this function will
// do nothing.
func (x *MediaFile) Snapshot(SnapshotVar *gdk.Snapshot, WidthVar float64, HeightVar float64) {

	gdk.XGdkPaintableSnapshot(x.GoPointer(), SnapshotVar.GoPointer(), WidthVar, HeightVar)

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GTK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewMediaFile, lib, "gtk_media_file_new")
	core.PuregoSafeRegister(&xNewForFileMediaFile, lib, "gtk_media_file_new_for_file")
	core.PuregoSafeRegister(&xNewForFilenameMediaFile, lib, "gtk_media_file_new_for_filename")
	core.PuregoSafeRegister(&xNewForInputStreamMediaFile, lib, "gtk_media_file_new_for_input_stream")
	core.PuregoSafeRegister(&xNewForResourceMediaFile, lib, "gtk_media_file_new_for_resource")

	core.PuregoSafeRegister(&xMediaFileClear, lib, "gtk_media_file_clear")
	core.PuregoSafeRegister(&xMediaFileGetFile, lib, "gtk_media_file_get_file")
	core.PuregoSafeRegister(&xMediaFileGetInputStream, lib, "gtk_media_file_get_input_stream")
	core.PuregoSafeRegister(&xMediaFileSetFile, lib, "gtk_media_file_set_file")
	core.PuregoSafeRegister(&xMediaFileSetFilename, lib, "gtk_media_file_set_filename")
	core.PuregoSafeRegister(&xMediaFileSetInputStream, lib, "gtk_media_file_set_input_stream")
	core.PuregoSafeRegister(&xMediaFileSetResource, lib, "gtk_media_file_set_resource")

}
