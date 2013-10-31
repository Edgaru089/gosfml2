// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <stdlib.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"

import (
	"runtime"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RenderTexture struct {
	cptr *C.sfRenderTexture
	view *View
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Construct a new render texture
//
// 	width:       Width of the render texture
// 	height:      Height of the render texture
// 	depthBuffer: Do you want a depth-buffer attached? (useful only if you're doing 3D OpenGL on the rendertexture)
func NewRenderTexture(width, height uint, depthbuffer bool) *RenderTexture {
	//create the render texture
	renderTexture := &RenderTexture{}
	cstream.ExecAndBlock(func() {
		renderTexture.cptr = C.sfRenderTexture_create(C.uint(width), C.uint(height), goBool2C(depthbuffer))
	})

	//view
	renderTexture.SetView(newViewFromPtr(C.sfRenderTexture_getView(renderTexture.cptr)))

	//GC
	runtime.SetFinalizer(renderTexture, (*RenderTexture).destroy)

	return renderTexture
}

// Destroy an existing render texture
func (this *RenderTexture) destroy() {
	C.sfRenderTexture_destroy(this.cptr)
	this.cptr = nil
}

// Get the size of the rendering region of a render texture
func (this *RenderTexture) GetSize() (size Vector2u) {
	cstream.ExecAndBlock(func() {
		size.fromC(C.sfRenderTexture_getSize(this.cptr))
	})
	return
}

// Activate or deactivate a render texture as the current target for rendering
//
// 	active: true to activate, false to deactivate
func (this *RenderTexture) SetActive(active bool) {
	cstream.Exec(func() {
		C.sfRenderTexture_setActive(this.cptr, goBool2C(active))
	})
}

// Update the contents of the target texture
func (this *RenderTexture) Display() {
	cstream.Exec(func() {
		C.sfRenderTexture_display(this.cptr)
	})
}

// Clear the rendertexture with the given color
//
// 	color: Fill color
func (this *RenderTexture) Clear(color Color) {
	cstream.Exec(func() {
		C.sfRenderTexture_clear(this.cptr, color.toC())
	})
}

// Change the current active view of a render texture
//
// 	view: Pointer to the new view
func (this *RenderTexture) SetView(view *View) {
	cstream.Exec(func() {
		this.view = view
		C.sfRenderTexture_setView(this.cptr, view.toCPtr())
	})
}

// Get the current active view of a render texture
func (this *RenderTexture) GetView() *View {
	return this.view
}

// Get the default view of a render texture
func (this *RenderTexture) GetDefaultView() *View {
	return newViewFromPtr(C.sfRenderTexture_getDefaultView(this.cptr))
}

// Get the viewport of a view applied to this target
//
// 	view: Target view
func (this *RenderTexture) GetViewport(view *View) (viewport IntRect) {
	cstream.ExecAndBlock(func() {
		viewport.fromC(C.sfRenderTexture_getViewport(this.cptr, view.toCPtr()))
	})
	return
}

// Convert a point from texture coordinates to world coordinates
//
// This function finds the 2D position that matches the
// given pixel of the render-texture. In other words, it does
// the inverse of what the graphics card does, to find the
// initial position of a rendered pixel.
//
// Initially, both coordinate systems (world units and target pixels)
// match perfectly. But if you define a custom view or resize your
// render-texture, this assertion is not true anymore, ie. a point
// located at (10, 50) in your render-texture may map to the point
// (150, 75) in your 2D world -- if the view is translated by (140, 25).
//
// This version uses a custom view for calculations, see the other
// overload of the function if you want to use the current view of the
// render-texture.
//
// 	point: Pixel to convert
// 	view:  The view to use for converting the point
func (this *RenderTexture) MapPixelToCoords(pos Vector2i, view *View) (coords Vector2f) {
	cstream.ExecAndBlock(func() {
		coords.fromC(C.sfRenderTexture_mapPixelToCoords(this.cptr, pos.toC(), view.toCPtr()))
	})
	return
}

// Convert a point from world coordinates to texture coordinates
//
// This function finds the pixel of the render-texture that matches
// the given 2D point. In other words, it goes through the same process
// as the graphics card, to compute the final position of a rendered point.
//
// Initially, both coordinate systems (world units and target pixels)
// match perfectly. But if you define a custom view or resize your
// render-texture, this assertion is not true anymore, ie. a point
// located at (150, 75) in your 2D world may map to the pixel
// (10, 50) of your render-texture -- if the view is translated by (140, 25).
//
// This version uses a custom view for calculations, see the other
// overload of the function if you want to use the current view of the
// render-texture.
//
// 	point: Point to convert
// 	view:  The view to use for converting the point
func (this *RenderTexture) MapCoordsToPixel(pos Vector2f, view *View) (coords Vector2i) {
	cstream.ExecAndBlock(func() {
		coords.fromC(C.sfRenderTexture_mapCoordsToPixel(this.cptr, pos.toC(), view.toCPtr()))
	})
	return
}

//Draws a RectangleShape on a render target
//
//renderStates: can be nil
func (this *RenderTexture) Draw(drawer Drawer, renderStates RenderStates) {
	drawer.Draw(this, renderStates)
}

// Save the current OpenGL render states and matrices
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering. Combined with popGLStates,
// it ensures that:
// SFML's internal states are not messed up by your OpenGL code
// your OpenGL states are not modified by a call to a SFML function
//
// Note that this function is quite expensive: it saves all the
// possible OpenGL states and matrices, even the ones you
// don't care about. Therefore it should be used wisely.
// It is provided for convenience, but the best results will
// be achieved if you handle OpenGL states yourself (because
// you know which states have really changed, and need to be
// saved and restored). Take a look at the resetGLStates
// function if you do so.
func (this *RenderTexture) PushGLStates() {
	cstream.ExecAndBlock(func() {
		C.sfRenderTexture_pushGLStates(this.cptr)
	})
}

// Restore the previously saved OpenGL render states and matrices
//
// See the description of pushGLStates to get a detailed
// description of these functions.
func (this *RenderTexture) PopGLStates() {
	cstream.ExecAndBlock(func() {
		C.sfRenderTexture_popGLStates(this.cptr)
	})
}

// Reset the internal OpenGL states so that the target is ready for drawing
//
// This function can be used when you mix SFML drawing
// and direct OpenGL rendering, if you choose not to use
// pushGLStates/popGLStates. It makes sure that all OpenGL
// states needed by SFML are set, so that subsequent RenderTexture.Draw
// calls will work as expected.
func (this *RenderTexture) ResetGLStates() {
	cstream.ExecAndBlock(func() {
		C.sfRenderTexture_resetGLStates(this.cptr)
	})
}

// Get the target texture of a render texture
func (this *RenderTexture) GetTexture() (texture *Texture) {
	cstream.ExecAndBlock(func() {
		texture = &Texture{}
		texture.cptr = C.sfRenderTexture_getTexture(this.cptr)
	})
	return
}

// Enable or disable the smooth filter on a render texture
//
// 	smooth: true to enable smoothing, false to disable it
func (this *RenderTexture) SetSmooth(smooth bool) {
	cstream.Exec(func() {
		C.sfRenderTexture_setSmooth(this.cptr, goBool2C(smooth))
	})
}

// Tell whether the smooth filter is enabled or not for a render texture
func (this *RenderTexture) IsSmooth() (smooth bool) {
	cstream.ExecAndBlock(func() {
		smooth = sfBool2Go(C.sfRenderTexture_isSmooth(this.cptr))
	})
	return
}

// Enable or disable texture repeating
//
// 	repeated: true to enable repeating, false to disable it
func (this *RenderTexture) SetRepeated(repeated bool) {
	cstream.Exec(func() {
		C.sfRenderTexture_setRepeated(this.cptr, goBool2C(repeated))
	})
}

// Tell whether the texture is repeated or not
func (this *RenderTexture) IsRepeated() (repeated bool) {
	cstream.ExecAndBlock(func() {
		repeated = sfBool2Go(C.sfRenderTexture_isRepeated(this.cptr))
	})
	return
}
