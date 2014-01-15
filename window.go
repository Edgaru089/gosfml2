// Copyright (C) 2012 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

// #include <SFML/Window/Window.h>
// #include <stdlib.h>
import "C"

import (
	"errors"
	"runtime"
	"unsafe"
)

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	StyleNone       = C.sfNone         ///< No border / title bar (this flag and all others are mutually exclusive)
	StyleTitlebar   = C.sfTitlebar     ///< Title bar + fixed border
	StyleResize     = C.sfResize       ///< Titlebar + resizable border + maximize button
	StyleClose      = C.sfClose        ///< Titlebar + close button
	StyleFullscreen = C.sfFullscreen   ///< Fullscreen mode (this flag and all others are mutually exclusive)
	StyleDefault    = C.sfDefaultStyle ///< Default window style
)

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

type Window struct {
	cptr *C.sfWindow
}

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

// Window and RenderWindow are SystemWindows
type SystemWindow interface {
	SetVSyncEnabled(bool)
	SetFramerateLimit(uint)
	SetJoystickThreshold(float32)
	SetKeyRepeatEnabled(bool)
	Display()
	IsOpen() bool
	Close()
	SetTitle(string)
	SetIcon(uint, uint, []byte) error
	SetMouseCursorVisible(bool)
	SetActive(bool) bool
}

//TEST
var _ SystemWindow = &RenderWindow{}
var _ SystemWindow = &Window{}

/////////////////////////////////////
///		FUNCTIONS
/////////////////////////////////////

// Construct a new window
//
// 	videoMode:       Video mode to use
// 	title:           Title of the window
// 	style:           Window style
// 	contextSettings: Creation settings (pass nil to use default values)
func NewWindow(videoMode VideoMode, title string, style int, contextSettings ContextSettings) (window *Window) {
	//string conversion
	utf32 := strToRunes(title)

	//convert contextSettings to C
	cs := contextSettings.toC()

	//create the window
	window = &Window{C.sfWindow_createUnicode(videoMode.toC(), (*C.sfUint32)(unsafe.Pointer(&utf32[0])), C.sfUint32(style), &cs)}

	//GC cleanup
	runtime.SetFinalizer(window, (*Window).destroy)

	return window
}

// Get the creation settings of a window
func (this *Window) GetSettings() (settings ContextSettings) {
	cstream.ExecAndBlock(func() {
		settings.fromC(C.sfWindow_getSettings(this.cptr))
	})

	return
}

// Change the size of the rendering region of a window
//
// 	size: New size, in pixels
func (this *Window) SetSize(size Vector2u) {
	cstream.Exec(func() {
		C.sfWindow_setSize(this.cptr, size.toC())
	})
}

// Get the size of the rendering region of a window
func (this *Window) GetSize() (size Vector2u) {
	cstream.ExecAndBlock(func() {
		csize := C.sfWindow_getSize(this.cptr)
		size = Vector2u{uint(csize.x), uint(csize.y)}
	})
	return
}

// Change the position of a window on screen
//
// Only works for top-level windows
//
// 	pos: New position, in pixels
func (this *Window) SetPosition(pos Vector2i) {
	C.sfWindow_setPosition(this.cptr, pos.toC())
}

// Get the position of a render window
func (this *Window) GetPosition() (pos Vector2i) {
	cstream.ExecAndBlock(func() {
		pos.fromC(C.sfWindow_getPosition(this.cptr))
	})
	return
}

// Tell whether or not a window is opened
func (this *Window) IsOpen() (open bool) {
	cstream.ExecAndBlock(func() {
		open = sfBool2Go(C.sfWindow_isOpen(this.cptr))
	})
	return
}

// Close a window (but doesn't destroy the internal data)
func (this *Window) Close() {
	cstream.ExecAndBlock(func() {
		C.sfWindow_close(this.cptr)
	})
}

// Destroy an existing window
func (this *Window) destroy() {
	cstream.ExecAndBlock(func() {
		C.sfWindow_destroy(this.cptr)
		this.cptr = nil
	})
}

// Get the event on top of event queue of a window, if any, and pop it
//
// returns nil if there are no events left.
func (this *Window) PollEvent() Event {
	cEvent := C.sfEvent{}
	var hasEvent C.sfBool
	cstream.ExecAndBlock(func() {
		hasEvent = C.sfWindow_pollEvent(this.cptr, &cEvent)
	})

	if hasEvent != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Wait for an event and return it
func (this *Window) WaitEvent() Event {
	cEvent := C.sfEvent{}
	var hasError C.sfBool

	cstream.ExecAndBlock(func() {
		hasError = C.sfWindow_waitEvent(this.cptr, &cEvent)
	})

	if hasError != 0 {
		return handleEvent(&cEvent)
	}
	return nil
}

// Change the title of a window
//
// 	title: New title
func (this *Window) SetTitle(title string) {
	cstream.Exec(func() {
		utf32 := strToRunes(title)
		C.sfWindow_setUnicodeTitle(this.cptr, (*C.sfUint32)(unsafe.Pointer(&utf32[0])))
	})
}

// Change a window's icon
//
// 	width:  Icon's width, in pixels
// 	height: Icon's height, in pixels
// 	pixels: Slice of pixels, format must be RGBA 32 bits
func (this *Window) SetIcon(width, height uint, data []byte) error {
	if len(data) >= int(width*height*4) {
		cstream.Exec(func() {
			C.sfWindow_setIcon(this.cptr, C.uint(width), C.uint(height), (*C.sfUint8)(&data[0]))
		})
		return nil
	}
	return errors.New("SetIcon: Slice length does not match specified dimensions")
}

// Limit the framerate to a maximum fixed frequency for a window
//
// 	limit: Framerate limit, in frames per seconds (use 0 to disable limit)
func (this *Window) SetFramerateLimit(limit uint) {
	cstream.Exec(func() {
		C.sfWindow_setFramerateLimit(this.cptr, C.uint(limit))
	})
}

///Change the joystick threshold, ie. the value below which no move event will be generated
//
// threshold: New threshold, in range [0, 100]
func (this *Window) SetJoystickThreshold(threshold float32) {
	cstream.Exec(func() {
		C.sfWindow_setJoystickThreshold(this.cptr, C.float(threshold))
	})
}

// Enable or disable automatic key-repeat
//
// If key repeat is enabled, you will receive repeated
// KeyPress events while keeping a key pressed. If it is disabled,
// you will only get a single event when the key is pressed.
//
// Key repeat is enabled by default.
func (this *Window) SetKeyRepeatEnabled(enabled bool) {
	cstream.Exec(func() {
		C.sfWindow_setKeyRepeatEnabled(this.cptr, goBool2C(enabled))
	})
}

// Display a window on screen
func (this *Window) Display() {
	cstream.ExecAndBlock(func() {
		C.sfWindow_display(this.cptr)
	})
}

// Enable / disable vertical synchronization on a window
//
// 	enabled: true to enable v-sync, false to deactivate
func (this *Window) SetVSyncEnabled(enabled bool) {
	cstream.Exec(func() {
		C.sfWindow_setVerticalSyncEnabled(this.cptr, goBool2C(enabled))
	})
}

// Activate or deactivate a window as the current target for rendering
//
// 	active: true to activate, false to deactivate
//
// return True if operation was successful, false otherwise
func (this *Window) SetActive(active bool) (success bool) {
	cstream.ExecAndBlock(func() {
		success = sfBool2Go(C.sfWindow_setActive(this.cptr, goBool2C(active)))
	})
	return
}

// Show or hide the mouse cursor on a render window
//
// 	visible: true to show, false to hide
func (this *Window) SetMouseCursorVisible(visible bool) {
	cstream.Exec(func() {
		C.sfWindow_setMouseCursorVisible(this.cptr, goBool2C(visible))
	})
}
