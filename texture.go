// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/Texture.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Texture struct {
	cptr *C.sfTexture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new texture
//
// 	width:  Texture width
// 	height: Texture height
func NewTexture(width, height uint) (texture *Texture, err error) {
	texture = &Texture{}
	cstream.ExecAndBlock(func() {
		texture.cptr = C.sfTexture_create(C.uint(width), C.uint(height))
	})
	runtime.SetFinalizer(texture, (*Texture).destroy)

	if texture.cptr == nil {
		err = errors.New("NewTexture: Cannot create texture")
	}

	return
}

// Create a new texture from an image
//
// 	file: Path of the image file to load
// 	area: Area of the source image to load (nil to load the entire image)
func NewTextureFromFile(file string, area *IntRect) (texture *Texture, err error) {
	cFile := C.CString(file)
	defer C.free(unsafe.Pointer(cFile))
	texture = &Texture{}
	cstream.ExecAndBlock(func() {
		texture.cptr = C.sfTexture_createFromFile(cFile, area.toCPtr())
	})
	runtime.SetFinalizer(texture, (*Texture).destroy)

	if texture.cptr == nil {
		err = errors.New("NewTextureFromFile: Cannot load texture " + file)
	}

	return
}

// Create a new texture from a file in memory
//
// 	data: Slice containing the file data
// 	area: Area of the source image to load (nil to load the entire image)
func NewTextureFromMemory(data []byte, area *IntRect) (texture *Texture, err error) {
	if len(data) > 0 {
		texture = &Texture{}
		cstream.ExecAndBlock(func() {
			texture.cptr = C.sfTexture_createFromMemory(unsafe.Pointer(&data[0]), C.size_t(len(data)), area.toCPtr())
		})
		runtime.SetFinalizer(texture, (*Texture).destroy)
	}
	err = errors.New("NewTextureFromMemory: no data")
	return
}

// Create a new texture from an image
//
// 	image: Image to upload to the texture
// 	area:  Area of the source image to load (NULL to load the entire image)
func NewTextureFromImage(image *Image, area *IntRect) (texture *Texture, err error) {
	texture = &Texture{}
	cstream.ExecAndBlock(func() {
		texture.cptr = C.sfTexture_createFromImage(image.toCPtr(), area.toCPtr())
	})
	runtime.SetFinalizer(texture, (*Texture).destroy)

	if texture.cptr == nil {
		err = errors.New("NewTextureFromFile: Cannot create texture from image")
	}

	return
}

// Copy an existing texture
func (this *Texture) Copy() (texture *Texture) {
	cstream.ExecAndBlock(func() {
		texture = &Texture{C.sfTexture_copy(this.cptr)}
		runtime.SetFinalizer(texture, (*Texture).destroy)
	})

	return
}

// Destroy an existing texture
func (this *Texture) destroy() {
	cstream.ExecAndBlock(func() {
		C.sfTexture_destroy(this.cptr)
		this.cptr = nil
	})
}

// Return the size of the texture
func (this *Texture) GetSize() (size Vector2u) {
	cstream.ExecAndBlock(func() {
		size.fromC(C.sfTexture_getSize(this.cptr))
	})
	return
}

// Update a texture from the contents of a window
//
// 	window:  Window to copy to the texture
// 	x:       X offset in the texture where to copy the source pixels
// 	y:       Y offset in the texture where to copy the source pixels
func (this *Texture) UpdateFromWindow(window *Window, x, y uint) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_updateFromWindow(this.cptr, window.cptr, C.uint(x), C.uint(y))
	})
}

// Update a texture from the contents of a render-window
//
// 	renderWindow: Render-window to copy to the texture
// 	x:            X offset in the texture where to copy the source pixels
// 	y:            Y offset in the texture where to copy the source pixels
func (this *Texture) UpdateFromRenderWindow(window *RenderWindow, x, y uint) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_updateFromRenderWindow(this.cptr, window.cptr, C.uint(x), C.uint(y))
	})
}

// Update a texture from an image
//
// 	image:   Image to copy to the texture
// 	x:       X offset in the texture where to copy the source pixels
// 	y:       Y offset in the texture where to copy the source pixels
func (this *Texture) UpdateFromImage(image *Image, x, y uint) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_updateFromImage(this.cptr, image.toCPtr(), C.uint(x), C.uint(y))
	})
}

// Update a texture from an array of pixels
//
// 	pixels:  Slice of pixels to copy to the texture
// 	width:   Width of the pixel region contained in pixels
// 	height:  Height of the pixel region contained in pixels
// 	x:       X offset in the texture where to copy the source pixels
// 	y:       Y offset in the texture where to copy the source pixels
func (this *Texture) UpdateFromPixels(pixels []byte, width, height, x, y uint) {
	if len(pixels) > 0 {
		cstream.ExecAndBlock(func() {
			C.sfTexture_updateFromPixels(this.cptr, (*C.sfUint8)(unsafe.Pointer(&pixels[0])), C.uint(width), C.uint(height), C.uint(x), C.uint(y))
		})
	}
}

// Enable or disable the smooth filter on a texture
func (this *Texture) SetSmooth(smooth bool) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_setSmooth(this.cptr, goBool2C(smooth))
	})
}

// Tell whether the smooth filter is enabled or not for a texture
func (this *Texture) IsSmooth() (smooth bool) {
	cstream.ExecAndBlock(func() {
		smooth = sfBool2Go(C.sfTexture_isSmooth(this.cptr))
	})
	return
}

// Enable or disable repeating for a texture
//
// Repeating is involved when using texture coordinates
// outside the texture rectangle [0, 0, width, height].
// In this case, if repeat mode is enabled, the whole texture
// will be repeated as many times as needed to reach the
// coordinate (for example, if the X texture coordinate is
// 3 * width, the texture will be repeated 3 times).
// If repeat mode is disabled, the "extra space" will instead
// be filled with border pixels.
// Warning: on very old graphics cards, white pixels may appear
// when the texture is repeated. With such cards, repeat mode
// can be used reliably only if the texture has power-of-two
// dimensions (such as 256x128).
// Repeating is disabled by default.
func (this *Texture) SetRepeated(repeated bool) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_setRepeated(this.cptr, goBool2C(repeated))
	})
}

// Tell whether a texture is repeated or not
func (this *Texture) IsRepeated() (repeated bool) {
	cstream.ExecAndBlock(func() {
		repeated = sfBool2Go(C.sfTexture_isRepeated(this.cptr))
	})
	return
}

// Get the maximum texture size allowed
func GetMaximumTextureSize() (maxSize uint) {
	cstream.ExecAndBlock(func() {
		maxSize = uint(C.sfTexture_getMaximumSize())
	})
	return
}

// Bind a texture for rendering
//
// This function is not part of the graphics API, it mustn't be
// used when drawing SFML entities. It must be used only if you
// mix sfTexture with OpenGL code.
//
// 	texture: Pointer to the texture to bind, can be nil to use no texture
func BindTexture(texture *Texture) {
	cstream.ExecAndBlock(func() {
		C.sfTexture_bind(texture.toCPtr())
	})
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *Texture) toCPtr() *C.sfTexture {
	if this != nil {
		return this.cptr
	}
	return nil
}
