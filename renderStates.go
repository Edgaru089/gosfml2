package GoSFML2

/*
 #include <SFML/Graphics.h>
*/
import "C"
import "unsafe"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	Blend_Alpha    = iota ///< Pixel = Src * a + Dest * (1 - a)
	Blend_Add             ///< Pixel = Src + Dest
	Blend_Multiply        ///< Pixel = Src * Dest
	Blend_None            ///< No blending
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type BlendMode int

type RenderStates struct {
	blendMode BlendMode
	transform Transform
	texture   *C.sfTexture
	shader    *C.sfShader
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

func CreateRenderStates(blendMode BlendMode, transform Transform, texture Texture, shader Shader) (rt RenderStates) {
	rt.blendMode = blendMode
	rt.transform = transform
	rt.shader = shader.cptr
	rt.texture = texture.cptr
	return
}

/////////////////////////////////////
///		GO <-> C
/////////////////////////////////////

func (this *RenderStates) toCPtr() *C.sfRenderStates {
	return (*C.sfRenderStates)(unsafe.Pointer(this))
}
