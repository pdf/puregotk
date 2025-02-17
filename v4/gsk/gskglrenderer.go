// Package gsk was automatically generated by github.com/jwijenbergh/puregotk DO NOT EDIT
package gsk

import (
	"github.com/jwijenbergh/purego"
	"github.com/jwijenbergh/puregotk/internal/core"
)

type GLRendererClass struct {
}

type GLRenderer struct {
	Renderer
}

func GLRendererNewFromInternalPtr(ptr uintptr) *GLRenderer {
	cls := &GLRenderer{}
	cls.Ptr = ptr
	return cls
}

var xNewGLRenderer func() uintptr

// Creates a new `GskRenderer` using the new OpenGL renderer.
func NewGLRenderer() *Renderer {
	var cls *Renderer

	cret := xNewGLRenderer()

	if cret == 0 {
		return nil
	}
	cls = &Renderer{}
	cls.Ptr = cret
	return cls
}

func (c *GLRenderer) GoPointer() uintptr {
	return c.Ptr
}

func (c *GLRenderer) SetGoPointer(ptr uintptr) {
	c.Ptr = ptr
}

func init() {
	lib, err := purego.Dlopen(core.GetPath("GSK"), purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	core.PuregoSafeRegister(&xNewGLRenderer, lib, "gsk_gl_renderer_new")

}
