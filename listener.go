// Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
// This software is provided 'as-is', without any express or implied warranty.
// In no event will the authors be held liable for any damages arising from the use of this software.
// Permission is granted to anyone to use this software for any purpose, including commercial applications, 
// and to alter it and redistribute it freely, subject to the following restrictions:
// 	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
//			If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
// 	2. Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
// 	3. This notice may not be removed or altered from any source distribution.

package gosfml2

// #include <SFML/Audio/Listener.h> 
import "C"

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

// Change the global volume of all the sounds and musics
//
// The volume is a number between 0 and 100; it is combined with
// the individual volume of each sound / music.
// The default value for the volume is 100 (maximum).
//
// 	volume: New global volume, in the range [0, 100]
func ListenerSetGlobalVolume(volume float32) {
	C.sfListener_setGlobalVolume(C.float(volume))
}

// Get the current value of the global volume
//
// return Current global volume, in the range [0, 100]
func ListenerGetGlobalVolume() float32 {
	return float32(C.sfListener_getGlobalVolume())
}

// Set the position of the listener in the scene
//
// The default listener's position is (0, 0, 0).
//
// 	position: New position of the listener
func ListenerSetPosition(pos Vector3f) {
	C.sfListener_setPosition(pos.toC())
}

// Get the current position of the listener in the scene
func ListenerGetPosition() (pos Vector3f) {
	pos.fromC(C.sfListener_getPosition())
	return
}

// Set the orientation of the listener in the scene
//
// The orientation defines the 3D axes of the listener
// (left, up, front) in the scene. The orientation vector
// doesn't have to be normalized.
// The default listener's orientation is (0, 0, -1).
//
// 	position: New direction of the listener
func ListenerSetDirection(dir Vector3f) {
	C.sfListener_setPosition(dir.toC())
}

// Get the current orientation of the listener in the scene
func ListenerGetDirection() (dir Vector3f) {
	dir.fromC(C.sfListener_getDirection())
	return
}
