// Added by Edgaru089

package gosfml2

// #include <SFML/Window/Clipboard.h>
import "C"
import (
	"unicode/utf8"
	"unsafe"
)

// ClipboardGetString gets the content of the clipboard as a Unicode string
func ClipboardGetString() string {
	strUtf32 := (*[1 << 30]int32)(unsafe.Pointer(C.sfClipboard_getUnicodeString()))[:]
	str := make([]byte, 0)
	buffer := make([]byte, 3)
	for i := 0; strUtf32[i] != 0; i++ {
		str = append(str, buffer[:utf8.EncodeRune(buffer, strUtf32[i])]...)
	}
	return string(str)
}

// ClipboardSetString sets the content of the clipboard
func ClipboardSetString(text string) {
	buffer := make([]rune, 0)
	for _, c := range text {
		buffer = append(buffer, c)
	}
	buffer = append(buffer, 0)
	C.sfClipboard_setUnicodeString((*C.sfUint32)(unsafe.Pointer(&buffer[0])))
}
