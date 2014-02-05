// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Graphics/RectangleShape.h>
// #include <SFML/Graphics/RenderWindow.h>
// #include <SFML/Graphics/RenderTexture.h>
import "C"
import "runtime"

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type RectangleShape struct {
	cptr    *C.sfRectangleShape
	texture *Texture //to prevent the GC from deleting the texture
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Create a new rectangle shape
func NewRectangleShape() *RectangleShape {
	shape := &RectangleShape{C.sfRectangleShape_create(), nil}
	runtime.SetFinalizer(shape, (*RectangleShape).destroy)
	return shape
}

// Copy an existing rectangle shape
func (this *RectangleShape) Copy() (shape *RectangleShape) {
	cstream.Exec(func() {
		shape = &RectangleShape{C.sfRectangleShape_copy(this.cptr), this.texture}
		runtime.SetFinalizer(shape, (*RectangleShape).destroy)
	})
	return shape
}

// Destroy an existing rectangle shape
func (this *RectangleShape) destroy() {
	C.sfRectangleShape_destroy(this.cptr)
	this.cptr = nil
}

// Set the position of a rectangle shape
//
// This function completely overwrites the previous position.
// See sfRectangleShape_move to apply an offset based on the previous position instead.
// The default position of a circle Shape object is (0, 0).
//
// 	position: New position
func (this *RectangleShape) SetPosition(pos Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setPosition(this.cptr, pos.toC())
	})

}

// Set the scale factors of a rectangle shape
//
// This function completely overwrites the previous scale.
// See sfRectangleShape_scale to add a factor based on the previous scale instead.
// The default scale of a circle Shape object is (1, 1).
//
// 	scale: New scale factors
func (this *RectangleShape) SetScale(scale Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setScale(this.cptr, scale.toC())
	})

}

// Set the local origin of a rectangle shape
//
// The origin of an object defines the center point for
// all transformations (position, scale, rotation).
// The coordinates of this point must be relative to the
// top-left corner of the object, and ignore all
// transformations (position, scale, rotation).
// The default origin of a circle Shape object is (0, 0).
//
// 	origin: New origin
func (this *RectangleShape) SetOrigin(orig Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setOrigin(this.cptr, orig.toC())
	})

}

// Set the orientation of a rectangle shape
//
// This function completely overwrites the previous rotation.
// See sfRectangleShape_rotate to add an angle based on the previous rotation instead.
// The default rotation of a circle Shape object is 0.
//
// 	angle: New rotation, in degrees
func (this *RectangleShape) SetRotation(rot float32) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setRotation(this.cptr, C.float(rot))
	})

}

// Get the orientation of a rectangle shape
//
// The rotation is always in the range [0, 360].
func (this *RectangleShape) GetRotation() (rotation float32) {
	cstream.Exec(func() {
		rotation = float32(C.sfRectangleShape_getRotation(this.cptr))
	})
	return
}

// Get the position of a rectangle shape
func (this *RectangleShape) GetPosition() (position Vector2f) {
	cstream.Exec(func() {
		position.fromC(C.sfRectangleShape_getPosition(this.cptr))
	})
	return
}

// Get the current scale of a rectangle shap
func (this *RectangleShape) GetScale() (scale Vector2f) {
	cstream.Exec(func() {
		scale.fromC(C.sfRectangleShape_getScale(this.cptr))
	})
	return
}

// Get the local origin of a rectangle shape
func (this *RectangleShape) GetOrigin() (origin Vector2f) {
	cstream.Exec(func() {
		origin.fromC(C.sfRectangleShape_getOrigin(this.cptr))
	})
	return
}

// Move a rectangle shape by a given offset
//
// This function adds to the current position of the object,
// unlike RectangleShape.SetPosition which overwrites it.
func (this *RectangleShape) Move(offset Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_move(this.cptr, offset.toC())
	})

}

// Scale a rectangle shape
//
// This function multiplies the current scale of the object,
// unlike RectangleShape.SetScale which overwrites it.
func (this *RectangleShape) Scale(factor Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_scale(this.cptr, factor.toC())
	})

}

// Rotate a rectangle shape
//
// This function adds to the current rotation of the object,
// unlike RectangleShape.SetRotation which overwrites it.
func (this *RectangleShape) Rotate(angle float32) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_rotate(this.cptr, C.float(angle))
	})

}

// Change the source texture of a rectangle shape
//
// texture can be nil to disable texturing.
// If resetRect is true, the TextureRect property of
// the shape is automatically adjusted to the size of the new
// texture. If it is false, the texture rect is left unchanged.
//
// 	texture:   New texture
// 	resetRect: Should the texture rect be reset to the size of the new texture?
func (this *RectangleShape) SetTexture(texture *Texture, resetRect bool) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setTexture(this.cptr, texture.cptr, goBool2C(resetRect))
	})

	this.texture = texture
}

// Set the sub-rectangle of the texture that a rectangle shape will display
//
// The texture rect is useful when you don't want to display
// the whole texture, but rather a part of it.
// By default, the texture rect covers the entire texture.
//
// 	rect:  Rectangle defining the region of the texture to display
func (this *RectangleShape) SetTextureRect(rect IntRect) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setTextureRect(this.cptr, rect.toC())
	})

}

// Get the source texture of a rectangle shape
//
// If the shape has no source texture, a nil pointer is returned.
func (this *RectangleShape) GetTexture() (texture *Texture) {
	cstream.Exec(func() {
		texture = this.texture
	})
	return
}

// Get the sub-rectangle of the texture displayed by a rectangle shape
func (this *RectangleShape) GetTextureRect() (rect IntRect) {
	cstream.Exec(func() {
		rect.fromC(C.sfRectangleShape_getTextureRect(this.cptr))
	})
	return
}

// Set the fill color of a rectangle shape
//
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
//
// 	color: New color of the shape
func (this *RectangleShape) SetFillColor(color Color) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setFillColor(this.cptr, color.toC())
	})

}

// Set the outline color of a rectangle shape
//
// You can use sfTransparent to disable the outline.
// By default, the shape's outline color is opaque white.
//
// 	color: New outline color of the shape
func (this *RectangleShape) SetOutlineColor(color Color) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setOutlineColor(this.cptr, color.toC())
	})

}

// Set the thickness of a rectangle shape's outline
//
// This number cannot be negative. Using zero disables
// the outline.
// By default, the outline thickness is 0.
//
// 	thickness: New outline thickness
func (this *RectangleShape) SetOutlineThickness(thickness float32) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setOutlineThickness(this.cptr, C.float(thickness))
	})

}

// Set the size of a rectangle shape
func (this *RectangleShape) SetSize(size Vector2f) {
	cstream.Enqueue(func() {
		C.sfRectangleShape_setSize(this.cptr, size.toC())
	})

}

// Get the size of a rectangle shape
func (this *RectangleShape) GetSize() (size Vector2f) {
	cstream.Exec(func() {
		size.fromC(C.sfRectangleShape_getSize(this.cptr))
	})
	return
}

// Get the combined transform of a rectangle shape
func (this *RectangleShape) GetTransform() (transform Transform) {
	cstream.Exec(func() {
		transform.fromC(C.sfRectangleShape_getTransform(this.cptr))
	})
	return
}

// Get the inverse of the combined transform of a rectangle shape
func (this *RectangleShape) GetInverseTransform() (transform Transform) {
	cstream.Exec(func() {
		transform.fromC(C.sfRectangleShape_getInverseTransform(this.cptr))
	})
	return
}

// Set the fill color of a rectangle shape
//
// This color is modulated (multiplied) with the shape's
// texture if any. It can be used to colorize the shape,
// or change its global opacity.
// You can use sfTransparent to make the inside of
// the shape transparent, and have the outline alone.
// By default, the shape's fill color is opaque white.
//
// 	color: New color of the shape
func (this *RectangleShape) GetFillColor() (color Color) {
	cstream.Exec(func() {
		color.fromC(C.sfRectangleShape_getFillColor(this.cptr))
	})
	return
}

// Get the outline color of a rectangle shape
func (this *RectangleShape) GetOutlineColor() (color Color) {
	cstream.Exec(func() {
		color.fromC(C.sfRectangleShape_getOutlineColor(this.cptr))
	})
	return
}

// Get the outline thickness of a rectangle shape
func (this *RectangleShape) GetOutlineThickness() (thickness float32) {
	cstream.Exec(func() {
		thickness = float32(C.sfRectangleShape_getOutlineThickness(this.cptr))
	})
	return
}

// Get the total number of points of a rectangle shape
func (this *RectangleShape) GetPointCount() (pcount uint) {
	cstream.Exec(func() {
		pcount = uint(C.sfRectangleShape_getPointCount(this.cptr))
	})
	return
}

// Get a point of a rectangle shape
//
// The result is undefined if index is out of the valid range.
//
// index: Index of the point to get, in range [0 .. GetPointCount() - 1]
func (this *RectangleShape) GetPoint(index uint) (point Vector2f) {
	cstream.Exec(func() {
		point.fromC(C.sfRectangleShape_getPoint(this.cptr, C.uint(index)))
	})
	return
}

// Get the local bounding rectangle of a rectangle shape
//
// The returned rectangle is in local coordinates, which means
// that it ignores the transformations (translation, rotation,
// scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// entity in the entity's coordinate system.
func (this *RectangleShape) GetLocalBounds() (rect FloatRect) {
	cstream.Exec(func() {
		rect.fromC(C.sfRectangleShape_getLocalBounds(this.cptr))
	})
	return
}

// Get the global bounding rectangle of a rectangle shape
//
// The returned rectangle is in global coordinates, which means
// that it takes in account the transformations (translation,
// rotation, scale, ...) that are applied to the entity.
// In other words, this function returns the bounds of the
// sprite in the global 2D world's coordinate system.
func (this *RectangleShape) GetGlobalBounds() (rect FloatRect) {
	cstream.Exec(func() {
		rect.fromC(C.sfRectangleShape_getGlobalBounds(this.cptr))
	})
	return
}

//Draws a RectangleShape on a render target
//
//	renderStates: can be nil to use the default render states
func (this *RectangleShape) Draw(target RenderTarget, renderStates RenderStates) {
	cstream.Enqueue(func() {
		rs := renderStates.toC()
		switch target.(type) {
		case *RenderWindow:
			C.sfRenderWindow_drawRectangleShape(target.(*RenderWindow).cptr, this.cptr, &rs)
		case *RenderTexture:
			C.sfRenderTexture_drawRectangleShape(target.(*RenderTexture).cptr, this.cptr, &rs)
		}
	})

}
