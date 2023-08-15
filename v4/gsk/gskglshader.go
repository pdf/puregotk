// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
	"github.com/jwijenbergh/puregotk/v4/glib"
	"github.com/jwijenbergh/puregotk/v4/gobject"
	"github.com/jwijenbergh/puregotk/v4/graphene"
)

type GLShaderClass struct {
	ParentClass uintptr
}

// An object to build the uniforms data for a `GskGLShader`.
type ShaderArgsBuilder struct {
}

// A `GskGLShader` is a snippet of GLSL that is meant to run in the
// fragment shader of the rendering pipeline.
//
// A fragment shader gets the coordinates being rendered as input and
// produces the pixel values for that particular pixel. Additionally,
// the shader can declare a set of other input arguments, called
// uniforms (as they are uniform over all the calls to your shader in
// each instance of use). A shader can also receive up to 4
// textures that it can use as input when producing the pixel data.
//
// `GskGLShader` is usually used with gtk_snapshot_push_gl_shader()
// to produce a [class@Gsk.GLShaderNode] in the rendering hierarchy,
// and then its input textures are constructed by rendering the child
// nodes to textures before rendering the shader node itself. (You can
// pass texture nodes as children if you want to directly use a texture
// as input).
//
// The actual shader code is GLSL code that gets combined with
// some other code into the fragment shader. Since the exact
// capabilities of the GPU driver differs between different OpenGL
// drivers and hardware, GTK adds some defines that you can use
// to ensure your GLSL code runs on as many drivers as it can.
//
// If the OpenGL driver is GLES, then the shader language version
// is set to 100, and GSK_GLES will be defined in the shader.
//
// Otherwise, if the OpenGL driver does not support the 3.2 core profile,
// then the shader will run with language version 110 for GL2 and 130 for GL3,
// and GSK_LEGACY will be defined in the shader.
//
// If the OpenGL driver supports the 3.2 code profile, it will be used,
// the shader language version is set to 150, and GSK_GL3 will be defined
// in the shader.
//
// The main function the shader must implement is:
//
// ```glsl
//
//	void mainImage(out vec4 fragColor,
//	               in vec2 fragCoord,
//	               in vec2 resolution,
//	               in vec2 uv)
//
// ```
//
// Where the input @fragCoord is the coordinate of the pixel we're
// currently rendering, relative to the boundary rectangle that was
// specified in the `GskGLShaderNode`, and @resolution is the width and
// height of that rectangle. This is in the typical GTK coordinate
// system with the origin in the top left. @uv contains the u and v
// coordinates that can be used to index a texture at the
// corresponding point. These coordinates are in the [0..1]x[0..1]
// region, with 0, 0 being in the lower left corder (which is typical
// for OpenGL).
//
// The output @fragColor should be a RGBA color (with
// premultiplied alpha) that will be used as the output for the
// specified pixel location. Note that this output will be
// automatically clipped to the clip region of the glshader node.
//
// In addition to the function arguments the shader can define
// up to 4 uniforms for textures which must be called u_textureN
// (i.e. u_texture1 to u_texture4) as well as any custom uniforms
// you want of types int, uint, bool, float, vec2, vec3 or vec4.
//
// All textures sources contain premultiplied alpha colors, but if some
// there are outer sources of colors there is a gsk_premultiply() helper
// to compute premultiplication when needed.
//
// Note that GTK parses the uniform declarations, so each uniform has to
// be on a line by itself with no other code, like so:
//
// ```glsl
// uniform float u_time;
// uniform vec3 u_color;
// uniform sampler2D u_texture1;
// uniform sampler2D u_texture2;
// ```
//
// GTK uses the "gsk" namespace in the symbols it uses in the
// shader, so your code should not use any symbols with the prefix gsk
// or GSK. There are some helper functions declared that you can use:
//
// ```glsl
// vec4 GskTexture(sampler2D sampler, vec2 texCoords);
// ```
//
// This samples a texture (e.g. u_texture1) at the specified
// coordinates, and containes some helper ifdefs to ensure that
// it works on all OpenGL versions.
//
// You can compile the shader yourself using [method@Gsk.GLShader.compile],
// otherwise the GSK renderer will do it when it handling the glshader
// node. If errors occurs, the returned @error will include the glsl
// sources, so you can see what GSK was passing to the compiler. You
// can also set GSK_DEBUG=shaders in the environment to see the sources
// and other relevant information about all shaders that GSK is handling.
//
// # An example shader
//
// ```glsl
// uniform float position;
// uniform sampler2D u_texture1;
// uniform sampler2D u_texture2;
//
// void mainImage(out vec4 fragColor,
//
//	               in vec2 fragCoord,
//	               in vec2 resolution,
//	               in vec2 uv) {
//	  vec4 source1 = GskTexture(u_texture1, uv);
//	  vec4 source2 = GskTexture(u_texture2, uv);
//
//	  fragColor = position * source1 + (1.0 - position) * source2;
//	}
//
// ```
type GLShader struct {
	gobject.Object
}

func GLShaderNewFromInternalPtr(ptr uintptr) *GLShader {
	cls := &GLShader{}
	cls.Ptr = ptr
	return cls
}

var xNewFromBytesGLShader func(*glib.Bytes) uintptr

// Creates a `GskGLShader` that will render pixels using the specified code.
func NewFromBytesGLShader(SourcecodeVar *glib.Bytes) *GLShader {
	NewFromBytesGLShaderPtr := xNewFromBytesGLShader(SourcecodeVar)
	if NewFromBytesGLShaderPtr == 0 {
		return nil
	}

	NewFromBytesGLShaderCls := &GLShader{}
	NewFromBytesGLShaderCls.Ptr = NewFromBytesGLShaderPtr
	return NewFromBytesGLShaderCls
}

var xNewFromResourceGLShader func(string) uintptr

// Creates a `GskGLShader` that will render pixels using the specified code.
func NewFromResourceGLShader(ResourcePathVar string) *GLShader {
	NewFromResourceGLShaderPtr := xNewFromResourceGLShader(ResourcePathVar)
	if NewFromResourceGLShaderPtr == 0 {
		return nil
	}

	NewFromResourceGLShaderCls := &GLShader{}
	NewFromResourceGLShaderCls.Ptr = NewFromResourceGLShaderPtr
	return NewFromResourceGLShaderCls
}

var xGLShaderCompile func(uintptr, uintptr) bool

// Tries to compile the @shader for the given @renderer.
//
// If there is a problem, this function returns %FALSE and reports
// an error. You should use this function before relying on the shader
// for rendering and use a fallback with a simpler shader or without
// shaders if it fails.
//
// Note that this will modify the rendering state (for example
// change the current GL context) and requires the renderer to be
// set up. This means that the widget has to be realized. Commonly you
// want to call this from the realize signal of a widget, or during
// widget snapshot.
func (x *GLShader) Compile(RendererVar *Renderer) bool {

	return xGLShaderCompile(x.GoPointer(), RendererVar.GoPointer())

}

var xGLShaderFindUniformByName func(uintptr, string) int

// Looks for a uniform by the name @name, and returns the index
// of the uniform, or -1 if it was not found.
func (x *GLShader) FindUniformByName(NameVar string) int {

	return xGLShaderFindUniformByName(x.GoPointer(), NameVar)

}

var xGLShaderFormatArgs func(uintptr, ...interface{}) *glib.Bytes

// Formats the uniform data as needed for feeding the named uniforms
// values into the shader.
//
// The argument list is a list of pairs of names, and values for the types
// that match the declared uniforms (i.e. double/int/guint/gboolean for
// primitive values and `graphene_vecN_t *` for vecN uniforms).
//
// Any uniforms of the shader that are not included in the argument list
// are zero-initialized.
func (x *GLShader) FormatArgs(varArgs ...interface{}) *glib.Bytes {

	return xGLShaderFormatArgs(x.GoPointer(), varArgs...)

}

var xGLShaderFormatArgsVa func(uintptr, []interface{}) *glib.Bytes

// Formats the uniform data as needed for feeding the named uniforms
// values into the shader.
//
// The argument list is a list of pairs of names, and values for the
// types that match the declared uniforms (i.e. double/int/guint/gboolean
// for primitive values and `graphene_vecN_t *` for vecN uniforms).
//
// It is an error to pass a uniform name that is not declared by the shader.
//
// Any uniforms of the shader that are not included in the argument list
// are zero-initialized.
func (x *GLShader) FormatArgsVa(UniformsVar []interface{}) *glib.Bytes {

	return xGLShaderFormatArgsVa(x.GoPointer(), UniformsVar)

}

var xGLShaderGetArgBool func(uintptr, *glib.Bytes, int) bool

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of bool type.
func (x *GLShader) GetArgBool(ArgsVar *glib.Bytes, IdxVar int) bool {

	return xGLShaderGetArgBool(x.GoPointer(), ArgsVar, IdxVar)

}

var xGLShaderGetArgFloat func(uintptr, *glib.Bytes, int) float32

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of float type.
func (x *GLShader) GetArgFloat(ArgsVar *glib.Bytes, IdxVar int) float32 {

	return xGLShaderGetArgFloat(x.GoPointer(), ArgsVar, IdxVar)

}

var xGLShaderGetArgInt func(uintptr, *glib.Bytes, int) int32

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of int type.
func (x *GLShader) GetArgInt(ArgsVar *glib.Bytes, IdxVar int) int32 {

	return xGLShaderGetArgInt(x.GoPointer(), ArgsVar, IdxVar)

}

var xGLShaderGetArgUint func(uintptr, *glib.Bytes, int) uint32

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of uint type.
func (x *GLShader) GetArgUint(ArgsVar *glib.Bytes, IdxVar int) uint32 {

	return xGLShaderGetArgUint(x.GoPointer(), ArgsVar, IdxVar)

}

var xGLShaderGetArgVec2 func(uintptr, *glib.Bytes, int, *graphene.Vec2)

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of vec2 type.
func (x *GLShader) GetArgVec2(ArgsVar *glib.Bytes, IdxVar int, OutValueVar *graphene.Vec2) {

	xGLShaderGetArgVec2(x.GoPointer(), ArgsVar, IdxVar, OutValueVar)

}

var xGLShaderGetArgVec3 func(uintptr, *glib.Bytes, int, *graphene.Vec3)

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of vec3 type.
func (x *GLShader) GetArgVec3(ArgsVar *glib.Bytes, IdxVar int, OutValueVar *graphene.Vec3) {

	xGLShaderGetArgVec3(x.GoPointer(), ArgsVar, IdxVar, OutValueVar)

}

var xGLShaderGetArgVec4 func(uintptr, *glib.Bytes, int, *graphene.Vec4)

// Gets the value of the uniform @idx in the @args block.
//
// The uniform must be of vec4 type.
func (x *GLShader) GetArgVec4(ArgsVar *glib.Bytes, IdxVar int, OutValueVar *graphene.Vec4) {

	xGLShaderGetArgVec4(x.GoPointer(), ArgsVar, IdxVar, OutValueVar)

}

var xGLShaderGetArgsSize func(uintptr) uint

// Get the size of the data block used to specify arguments for this shader.
func (x *GLShader) GetArgsSize() uint {

	return xGLShaderGetArgsSize(x.GoPointer())

}

var xGLShaderGetNTextures func(uintptr) int

// Returns the number of textures that the shader requires.
//
// This can be used to check that the a passed shader works
// in your usecase. It is determined by looking at the highest
// u_textureN value that the shader defines.
func (x *GLShader) GetNTextures() int {

	return xGLShaderGetNTextures(x.GoPointer())

}

var xGLShaderGetNUniforms func(uintptr) int

// Get the number of declared uniforms for this shader.
func (x *GLShader) GetNUniforms() int {

	return xGLShaderGetNUniforms(x.GoPointer())

}

var xGLShaderGetResource func(uintptr) string

// Gets the resource path for the GLSL sourcecode being used
// to render this shader.
func (x *GLShader) GetResource() string {

	return xGLShaderGetResource(x.GoPointer())

}

var xGLShaderGetSource func(uintptr) *glib.Bytes

// Gets the GLSL sourcecode being used to render this shader.
func (x *GLShader) GetSource() *glib.Bytes {

	return xGLShaderGetSource(x.GoPointer())

}

var xGLShaderGetUniformName func(uintptr, int) string

// Get the name of the declared uniform for this shader at index @idx.
func (x *GLShader) GetUniformName(IdxVar int) string {

	return xGLShaderGetUniformName(x.GoPointer(), IdxVar)

}

var xGLShaderGetUniformOffset func(uintptr, int) int

// Get the offset into the data block where data for this uniforms is stored.
func (x *GLShader) GetUniformOffset(IdxVar int) int {

	return xGLShaderGetUniformOffset(x.GoPointer(), IdxVar)

}

var xGLShaderGetUniformType func(uintptr, int) GLUniformType

// Get the type of the declared uniform for this shader at index @idx.
func (x *GLShader) GetUniformType(IdxVar int) GLUniformType {

	return xGLShaderGetUniformType(x.GoPointer(), IdxVar)

}

func (c *GLShader) GoPointer() uintptr {
	return c.Ptr
}

func (c *GLShader) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GSK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewFromBytesGLShader, lib, "gsk_gl_shader_new_from_bytes")
	core.PuregoSafeRegister(&xNewFromResourceGLShader, lib, "gsk_gl_shader_new_from_resource")

	core.PuregoSafeRegister(&xGLShaderCompile, lib, "gsk_gl_shader_compile")
	core.PuregoSafeRegister(&xGLShaderFindUniformByName, lib, "gsk_gl_shader_find_uniform_by_name")
	core.PuregoSafeRegister(&xGLShaderFormatArgs, lib, "gsk_gl_shader_format_args")
	core.PuregoSafeRegister(&xGLShaderFormatArgsVa, lib, "gsk_gl_shader_format_args_va")
	core.PuregoSafeRegister(&xGLShaderGetArgBool, lib, "gsk_gl_shader_get_arg_bool")
	core.PuregoSafeRegister(&xGLShaderGetArgFloat, lib, "gsk_gl_shader_get_arg_float")
	core.PuregoSafeRegister(&xGLShaderGetArgInt, lib, "gsk_gl_shader_get_arg_int")
	core.PuregoSafeRegister(&xGLShaderGetArgUint, lib, "gsk_gl_shader_get_arg_uint")
	core.PuregoSafeRegister(&xGLShaderGetArgVec2, lib, "gsk_gl_shader_get_arg_vec2")
	core.PuregoSafeRegister(&xGLShaderGetArgVec3, lib, "gsk_gl_shader_get_arg_vec3")
	core.PuregoSafeRegister(&xGLShaderGetArgVec4, lib, "gsk_gl_shader_get_arg_vec4")
	core.PuregoSafeRegister(&xGLShaderGetArgsSize, lib, "gsk_gl_shader_get_args_size")
	core.PuregoSafeRegister(&xGLShaderGetNTextures, lib, "gsk_gl_shader_get_n_textures")
	core.PuregoSafeRegister(&xGLShaderGetNUniforms, lib, "gsk_gl_shader_get_n_uniforms")
	core.PuregoSafeRegister(&xGLShaderGetResource, lib, "gsk_gl_shader_get_resource")
	core.PuregoSafeRegister(&xGLShaderGetSource, lib, "gsk_gl_shader_get_source")
	core.PuregoSafeRegister(&xGLShaderGetUniformName, lib, "gsk_gl_shader_get_uniform_name")
	core.PuregoSafeRegister(&xGLShaderGetUniformOffset, lib, "gsk_gl_shader_get_uniform_offset")
	core.PuregoSafeRegister(&xGLShaderGetUniformType, lib, "gsk_gl_shader_get_uniform_type")

}
