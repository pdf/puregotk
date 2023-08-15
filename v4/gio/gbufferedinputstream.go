// Package gio was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gio

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
)

type BufferedInputStreamClass struct {
	ParentClass uintptr
}

type BufferedInputStreamPrivate struct {
}

// Buffered input stream implements #GFilterInputStream and provides
// for buffered reads.
//
// By default, #GBufferedInputStream's buffer size is set at 4 kilobytes.
//
// To create a buffered input stream, use g_buffered_input_stream_new(),
// or g_buffered_input_stream_new_sized() to specify the buffer's size at
// construction.
//
// To get the size of a buffer within a buffered input stream, use
// g_buffered_input_stream_get_buffer_size(). To change the size of a
// buffered input stream's buffer, use
// g_buffered_input_stream_set_buffer_size(). Note that the buffer's size
// cannot be reduced below the size of the data within the buffer.
type BufferedInputStream struct {
	FilterInputStream
}

func BufferedInputStreamNewFromInternalPtr(ptr uintptr) *BufferedInputStream {
	cls := &BufferedInputStream{}
	cls.Ptr = ptr
	return cls
}

var xNewBufferedInputStream func(uintptr) uintptr

// Creates a new #GInputStream from the given @base_stream, with
// a buffer set to the default size (4 kilobytes).
func NewBufferedInputStream(BaseStreamVar *InputStream) *InputStream {
	NewBufferedInputStreamPtr := xNewBufferedInputStream(BaseStreamVar.GoPointer())
	if NewBufferedInputStreamPtr == 0 {
		return nil
	}

	NewBufferedInputStreamCls := &InputStream{}
	NewBufferedInputStreamCls.Ptr = NewBufferedInputStreamPtr
	return NewBufferedInputStreamCls
}

var xNewSizedBufferedInputStream func(uintptr, uint) uintptr

// Creates a new #GBufferedInputStream from the given @base_stream,
// with a buffer set to @size.
func NewSizedBufferedInputStream(BaseStreamVar *InputStream, SizeVar uint) *InputStream {
	NewSizedBufferedInputStreamPtr := xNewSizedBufferedInputStream(BaseStreamVar.GoPointer(), SizeVar)
	if NewSizedBufferedInputStreamPtr == 0 {
		return nil
	}

	NewSizedBufferedInputStreamCls := &InputStream{}
	NewSizedBufferedInputStreamCls.Ptr = NewSizedBufferedInputStreamPtr
	return NewSizedBufferedInputStreamCls
}

var xBufferedInputStreamFill func(uintptr, int, uintptr) int

// Tries to read @count bytes from the stream into the buffer.
// Will block during this read.
//
// If @count is zero, returns zero and does nothing. A value of @count
// larger than %G_MAXSSIZE will cause a %G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes read into the buffer is returned.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. near the end of a file. Zero is returned on end of file
// (or if @count is zero),  but never otherwise.
//
// If @count is -1 then the attempted read size is equal to the number of
// bytes that are required to fill the buffer.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
//
// For the asynchronous, non-blocking, version of this function, see
// g_buffered_input_stream_fill_async().
func (x *BufferedInputStream) Fill(CountVar int, CancellableVar *Cancellable) int {

	return xBufferedInputStreamFill(x.GoPointer(), CountVar, CancellableVar.GoPointer())

}

var xBufferedInputStreamFillAsync func(uintptr, int, int, uintptr, uintptr, uintptr)

// Reads data into @stream's buffer asynchronously, up to @count size.
// @io_priority can be used to prioritize reads. For the synchronous
// version of this function, see g_buffered_input_stream_fill().
//
// If @count is -1 then the attempted read size is equal to the number
// of bytes that are required to fill the buffer.
func (x *BufferedInputStream) FillAsync(CountVar int, IoPriorityVar int, CancellableVar *Cancellable, CallbackVar AsyncReadyCallback, UserDataVar uintptr) {

	xBufferedInputStreamFillAsync(x.GoPointer(), CountVar, IoPriorityVar, CancellableVar.GoPointer(), purego.NewCallback(CallbackVar), UserDataVar)

}

var xBufferedInputStreamFillFinish func(uintptr, uintptr) int

// Finishes an asynchronous read.
func (x *BufferedInputStream) FillFinish(ResultVar AsyncResult) int {

	return xBufferedInputStreamFillFinish(x.GoPointer(), ResultVar.GoPointer())

}

var xBufferedInputStreamGetAvailable func(uintptr) uint

// Gets the size of the available data within the stream.
func (x *BufferedInputStream) GetAvailable() uint {

	return xBufferedInputStreamGetAvailable(x.GoPointer())

}

var xBufferedInputStreamGetBufferSize func(uintptr) uint

// Gets the size of the input buffer.
func (x *BufferedInputStream) GetBufferSize() uint {

	return xBufferedInputStreamGetBufferSize(x.GoPointer())

}

var xBufferedInputStreamPeek func(uintptr, uintptr, uint, uint) uint

// Peeks in the buffer, copying data of size @count into @buffer,
// offset @offset bytes.
func (x *BufferedInputStream) Peek(BufferVar uintptr, OffsetVar uint, CountVar uint) uint {

	return xBufferedInputStreamPeek(x.GoPointer(), BufferVar, OffsetVar, CountVar)

}

var xBufferedInputStreamPeekBuffer func(uintptr, uint) uintptr

// Returns the buffer with the currently available bytes. The returned
// buffer must not be modified and will become invalid when reading from
// the stream or filling the buffer.
func (x *BufferedInputStream) PeekBuffer(CountVar uint) uintptr {

	return xBufferedInputStreamPeekBuffer(x.GoPointer(), CountVar)

}

var xBufferedInputStreamReadByte func(uintptr, uintptr) int

// Tries to read a single byte from the stream or the buffer. Will block
// during this read.
//
// On success, the byte read from the stream is returned. On end of stream
// -1 is returned but it's not an exceptional error and @error is not set.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
//
// On error -1 is returned and @error is set accordingly.
func (x *BufferedInputStream) ReadByte(CancellableVar *Cancellable) int {

	return xBufferedInputStreamReadByte(x.GoPointer(), CancellableVar.GoPointer())

}

var xBufferedInputStreamSetBufferSize func(uintptr, uint)

// Sets the size of the internal buffer of @stream to @size, or to the
// size of the contents of the buffer. The buffer can never be resized
// smaller than its current contents.
func (x *BufferedInputStream) SetBufferSize(SizeVar uint) {

	xBufferedInputStreamSetBufferSize(x.GoPointer(), SizeVar)

}

func (c *BufferedInputStream) GoPointer() uintptr {
	return c.Ptr
}

func (c *BufferedInputStream) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

// Tests if the stream supports the #GSeekableIface.
func (x *BufferedInputStream) CanSeek() bool {

	return XGSeekableCanSeek(x.GoPointer())

}

// Tests if the length of the stream can be adjusted with
// g_seekable_truncate().
func (x *BufferedInputStream) CanTruncate() bool {

	return XGSeekableCanTruncate(x.GoPointer())

}

// Seeks in the stream by the given @offset, modified by @type.
//
// Attempting to seek past the end of the stream will have different
// results depending on if the stream is fixed-sized or resizable.  If
// the stream is resizable then seeking past the end and then writing
// will result in zeros filling the empty space.  Seeking past the end
// of a resizable stream and reading will result in EOF.  Seeking past
// the end of a fixed-sized stream will fail.
//
// Any operation that would result in a negative offset will fail.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned.
func (x *BufferedInputStream) Seek(OffsetVar int64, TypeVar glib.SeekType, CancellableVar *Cancellable) bool {

	return XGSeekableSeek(x.GoPointer(), OffsetVar, TypeVar, CancellableVar.GoPointer())

}

// Tells the current position within the stream.
func (x *BufferedInputStream) Tell() int64 {

	return XGSeekableTell(x.GoPointer())

}

// Sets the length of the stream to @offset. If the stream was previously
// larger than @offset, the extra data is discarded. If the stream was
// previously shorter than @offset, it is extended with NUL ('\0') bytes.
//
// If @cancellable is not %NULL, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error %G_IO_ERROR_CANCELLED will be returned. If an
// operation was partially finished when the operation was cancelled the
// partial result will be returned, without an error.
func (x *BufferedInputStream) Truncate(OffsetVar int64, CancellableVar *Cancellable) bool {

	return XGSeekableTruncate(x.GoPointer(), OffsetVar, CancellableVar.GoPointer())

}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GIO"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewBufferedInputStream, lib, "g_buffered_input_stream_new")
	core.PuregoSafeRegister(&xNewSizedBufferedInputStream, lib, "g_buffered_input_stream_new_sized")

	core.PuregoSafeRegister(&xBufferedInputStreamFill, lib, "g_buffered_input_stream_fill")
	core.PuregoSafeRegister(&xBufferedInputStreamFillAsync, lib, "g_buffered_input_stream_fill_async")
	core.PuregoSafeRegister(&xBufferedInputStreamFillFinish, lib, "g_buffered_input_stream_fill_finish")
	core.PuregoSafeRegister(&xBufferedInputStreamGetAvailable, lib, "g_buffered_input_stream_get_available")
	core.PuregoSafeRegister(&xBufferedInputStreamGetBufferSize, lib, "g_buffered_input_stream_get_buffer_size")
	core.PuregoSafeRegister(&xBufferedInputStreamPeek, lib, "g_buffered_input_stream_peek")
	core.PuregoSafeRegister(&xBufferedInputStreamPeekBuffer, lib, "g_buffered_input_stream_peek_buffer")
	core.PuregoSafeRegister(&xBufferedInputStreamReadByte, lib, "g_buffered_input_stream_read_byte")
	core.PuregoSafeRegister(&xBufferedInputStreamSetBufferSize, lib, "g_buffered_input_stream_set_buffer_size")

}
