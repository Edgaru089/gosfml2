// Added by Edgaru089

package gosfml2

// #include <SFML/Window/Cursor.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// Cursor defines the appearance of a system cursor
type Cursor struct {
	cptr *C.sfCursor
}

// Enumeration of the native system cursor types
//
// Refer to the following table to determine which cursor
// is available on which platform.
//
//    Type                         | Linux | Mac OS X | Windows |
//   ------------------------------|:-----:|:--------:|:--------:
//    CursorArrow                  |  yes  |    yes   |   yes   |
//    CursorArrowWait              |  no   |    no    |   yes   |
//    CursorWait                   |  yes  |    no    |   yes   |
//    CursorText                   |  yes  |    yes   |   yes   |
//    CursorHand                   |  yes  |    yes   |   yes   |
//    CursorSizeHorizontal         |  yes  |    yes   |   yes   |
//    CursorSizeVertical           |  yes  |    yes   |   yes   |
//    CursorSizeTopLeftBottomRight |  no   |    no    |   yes   |
//    CursorSizeBottomLeftTopRight |  no   |    no    |   yes   |
//    CursorSizeAll                |  yes  |    no    |   yes   |
//    CursorCross                  |  yes  |    yes   |   yes   |
//    CursorHelp                   |  yes  |    no    |   yes   |
//    CursorNotAllowed             |  yes  |    yes   |   yes   |
const (
	CursorArrow                  = iota // Arrow cursor (default)
	CursorArrowWait                     // Busy arrow cursor
	CursorWait                          // Busy cursor
	CursorText                          // I-beam, cursor when hovering over a field allowing text entry
	CursorHand                          // Pointing hand cursor
	CursorSizeHorizontal                // Horizontal double arrow cursor
	CursorSizeVertical                  // Vertical double arrow cursor
	CursorSizeTopLeftBottomRight        // Double arrow cursor going from top-left to bottom-right
	CursorSizeBottomLeftTopRight        // Double arrow cursor going from bottom-left to top-right
	CursorSizeAll                       // Combination of SizeHorizontal and SizeVertical
	CursorCross                         // Crosshair cursor
	CursorHelp                          // Help cursor
	CursorNotAllowed                    // Action not allowed cursor
)

// NewCursorFromPixels creates a new Cursor with a provided image
//
// pixels must be an array of width by height pixels
// in 32-bit RGBA format. If not, this function returns nil.
//
// If pixels is nil or either width or height are 0,
// the function returns nil.
//
// In addition to specifying the pixel data, you can also
// specify the location of the hotspot of the cursor. The
// hotspot is the pixel coordinate within the cursor image
// which will be located exactly where the mouse pointer
// position is. Any mouse actions that are performed will
// return the window/screen location of the hotspot.
func NewCursorFromPixels(pixels []byte, size, hotspot Vector2u) *Cursor {
	if pixels == nil || size.X == 0 || size.Y == 0 || len(pixels) != int(size.X*size.Y*4) {
		return nil
	}
	c := C.sfCursor_createFromPixels(
		(*C.sfUint8)((unsafe.Pointer)(&pixels[0])),
		C.sfVector2u{x: C.uint(size.X), y: C.uint(size.Y)},
		C.sfVector2u{x: C.uint(hotspot.X), y: C.uint(hotspot.Y)},
	)
	if c == nil {
		return nil
	}
	cr := &Cursor{cptr: c}
	runtime.SetFinalizer(cr, (*Cursor).destroy)
	return cr
}

// NewCursorFromPixelsUnsafe creates a new Cursor with a provided image, pointed to by unsafe.Pointer
func NewCursorFromPixelsUnsafe(pixels unsafe.Pointer, size, hotspot Vector2u) *Cursor {
	if pixels == nil || size.X == 0 || size.Y == 0 {
		return nil
	}
	c := C.sfCursor_createFromPixels(
		(*C.sfUint8)(pixels),
		C.sfVector2u{x: C.uint(size.X), y: C.uint(size.Y)},
		C.sfVector2u{x: C.uint(hotspot.X), y: C.uint(hotspot.Y)},
	)
	if c == nil {
		return nil
	}
	cr := &Cursor{cptr: c}
	runtime.SetFinalizer(cr, (*Cursor).destroy)
	return cr
}

// NewCursorFromSystem creates a new system cursor, returning nil if the requested
// cursor is not supported
func NewCursorFromSystem(cursorType int) *Cursor {
	c := C.sfCursor_createFromSystem(C.sfCursorType(cursorType))
	if c == nil {
		return nil
	}
	cr := &Cursor{cptr: c}
	runtime.SetFinalizer(cr, (*Cursor).destroy)
	return cr
}

func (c *Cursor) destroy() {
	C.sfCursor_destroy(c.cptr)
}
