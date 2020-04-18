# gosfml2

SFML2 Go binding, forked from bitbucket.org/krepa098/gosfml2. It compiles with CSFML 2.5

The following features are added:

 - Texture.GetNativeHandle()
 - Texture.UpdateFromPixelsUnsafe() which takes a unsafe.Pointer instead of a slice
 - Clipboard (the C API looks weird)
