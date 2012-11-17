/*
Copyright (c) 2012 krepa098 (krepa098 at gmail dot com)
This software is provided 'as-is', without any express or implied warranty.
In no event will the authors be held liable for any damages arising from the use of this software.
Permission is granted to anyone to use this software for any purpose, including commercial applications, 
and to alter it and redistribute it freely, subject to the following restrictions:
	1.	The origin of this software must not be misrepresented; you must not claim that you wrote the original software. 
		If you use this software in a product, an acknowledgment in the product documentation would be appreciated but is not required.
	2. 	Altered source versions must be plainly marked as such, and must not be misrepresented as being the original software.
	3. 	This notice may not be removed or altered from any source distribution.
*/

package gosfml2

// #include <SFML/Window.h>
// int getEventType(sfEvent* ev) { return ev->type; }
// sfSizeEvent* getSizeEvent(sfEvent* ev) { return &ev->size; }
// sfKeyEvent* getKeyEvent(sfEvent* ev) { return &ev->key; }
// sfTextEvent* getTextEvent(sfEvent* ev) { return &ev->text; }
// sfMouseMoveEvent* getMouseMoveEvent(sfEvent* ev) { return &ev->mouseMove; }
// sfMouseButtonEvent* getMouseButtonEvent(sfEvent* ev) { return &ev->mouseButton; }
// sfMouseWheelEvent* getMouseWheelEvent(sfEvent* ev) { return &ev->mouseWheel; }
// sfJoystickMoveEvent* getJoystickMoveEvent(sfEvent* ev) { return &ev->joystickMove; }
// sfJoystickButtonEvent* getJoystickButtonEvent(sfEvent* ev) { return &ev->joystickButton; }
// sfJoystickConnectEvent* getJoystickConnectEvent(sfEvent* ev) { return &ev->joystickConnect; }
import "C"

/////////////////////////////////////
///		CONSTS
/////////////////////////////////////

const (
	event_Closed = iota
	event_Resized
	event_LostFocus
	event_GainedFocus
	event_TextEntered
	event_KeyPressed
	event_KeyReleased
	event_MouseWheelMoved
	event_MouseButtonPressed
	event_MouseButtonReleased
	event_MouseMoved
	event_MouseEntered
	event_MouseLeft
	event_JoystickButtonPressed
	event_JoystickButtonReleased
	event_JoystickMoved
	event_JoystickConnected
	event_JoystickDisconnected
	event_None
)

type eventType int

/////////////////////////////////////
///		INTERFACES
/////////////////////////////////////

type Event interface{}

///////////////////////////////////////////////////////////////
//	EmptyEvents

type EventLostFocus struct{}
type EventGainedFocus struct{}
type EventMouseEntered struct{}
type EventMouseLeft struct{}
type EventClosed struct{}

///////////////////////////////////////////////////////////////
//	KeyEvent

type eventKey struct {
	Code    KeyCode
	Alt     int
	Control int
	Shift   int
	System  int
}

type EventKeyPressed eventKey
type EventKeyReleased eventKey

func newKeyEventFromC(ev *C.sfKeyEvent) eventKey {
	return eventKey{Code: KeyCode(ev.code), Alt: int(ev.alt), Control: int(ev.control), Shift: int(ev.shift), System: int(ev.system)}
}

///////////////////////////////////////////////////////////////
//	SizeEvent

type EventResized struct {
	Width  uint
	Height uint
}

func newSizeEventFromC(ev *C.sfSizeEvent) EventResized {
	return EventResized{Width: uint(ev.width), Height: uint(ev.height)}
}

///////////////////////////////////////////////////////////////
//	TextEvent

type EventTextEntered struct {
	Char rune //32bits
}

func newTextEventFromC(ev *C.sfTextEvent) EventTextEntered {
	return EventTextEntered{Char: rune(ev.unicode)}
}

///////////////////////////////////////////////////////////////
//	MouseMoveEvent

type EventMouseMoved struct {
	X int
	Y int
}

func newMouseMoveEventFromC(ev *C.sfMouseMoveEvent) EventMouseMoved {
	return EventMouseMoved{X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseButtonEvent

type eventMouseButton struct {
	Button MouseButton
	X      int
	Y      int
}

type EventMouseButtonPressed eventMouseButton
type EventMouseButtonReleased eventMouseButton

func newMouseButtonEventFromC(ev *C.sfMouseButtonEvent) eventMouseButton {
	return eventMouseButton{Button: MouseButton(ev.button), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	MouseWheelEvent

type EventMouseWheelMoved struct {
	Delta int
	X     int
	Y     int
}

func newMouseWheelEventFromC(ev *C.sfMouseWheelEvent) EventMouseWheelMoved {
	return EventMouseWheelMoved{Delta: int(ev.delta), X: int(ev.x), Y: int(ev.y)}
}

///////////////////////////////////////////////////////////////
//	JoystickMoveEvent

type EventJoystickMoved struct {
	JoystickId uint
	Axis       JoystickAxis
	position   float32
}

func newJoystickMoveEventFromC(ev *C.sfJoystickMoveEvent) EventJoystickMoved {
	return EventJoystickMoved{JoystickId: uint(ev.joystickId), Axis: JoystickAxis(ev.axis), position: float32(ev.position)}
}

///////////////////////////////////////////////////////////////
//	JoystickButtonEvent

type eventJoystickButton struct {
	JoystickId uint
	Button     uint
}

type EventJoystickButtonPressed eventJoystickButton
type EventJoystickButtonReleased eventJoystickButton

func newJoystickButtonEventFromC(ev *C.sfJoystickButtonEvent) eventJoystickButton {
	return eventJoystickButton{JoystickId: uint(ev.joystickId), Button: uint(ev.button)}
}

///////////////////////////////////////////////////////////////
//	JoystickConnectEvent

type eventJoystickConnection struct {
	JoystickId uint
}

type EventJoystickConnected eventJoystickConnection
type EventJoystickDisconnected eventJoystickConnection

func newJoystickConnectEventFromC(ev *C.sfJoystickConnectEvent) eventJoystickConnection {
	return eventJoystickConnection{JoystickId: uint(ev.joystickId)}
}

///////////////////////////////////////////////////////////////
//standard event handling method used by Window & RenderWindow

func handleEvent(cEvent *C.sfEvent) (ev Event) {
	switch eventType(C.getEventType(cEvent)) {
	case event_Resized:
		ev = newSizeEventFromC(C.getSizeEvent(cEvent))
	case event_Closed:
		ev = EventClosed{}
	case event_LostFocus:
		ev = EventLostFocus{}
	case event_GainedFocus:
		ev = EventGainedFocus{}
	case event_TextEntered:
		ev = newTextEventFromC(C.getTextEvent(cEvent))
	case event_KeyReleased:
		ev = (EventKeyReleased)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case event_KeyPressed:
		ev = (EventKeyPressed)(newKeyEventFromC(C.getKeyEvent(cEvent)))
	case event_MouseWheelMoved:
		ev = newMouseWheelEventFromC(C.getMouseWheelEvent(cEvent))
	case event_MouseButtonReleased:
		ev = (EventMouseButtonReleased)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case event_MouseButtonPressed:
		ev = (EventMouseButtonPressed)(newMouseButtonEventFromC(C.getMouseButtonEvent(cEvent)))
	case event_MouseMoved:
		ev = newMouseMoveEventFromC(C.getMouseMoveEvent(cEvent))
	case event_MouseLeft:
		ev = EventMouseLeft{}
	case event_MouseEntered:
		ev = EventMouseEntered{}
	case event_JoystickButtonReleased:
		ev = (EventJoystickButtonReleased)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case event_JoystickButtonPressed:
		ev = (EventJoystickButtonPressed)(newJoystickButtonEventFromC(C.getJoystickButtonEvent(cEvent)))
	case event_JoystickMoved:
		ev = newJoystickMoveEventFromC(C.getJoystickMoveEvent(cEvent))
	case event_JoystickDisconnected:
		ev = (EventJoystickDisconnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	case event_JoystickConnected:
		ev = (EventJoystickConnected)(newJoystickConnectEventFromC(C.getJoystickConnectEvent(cEvent)))
	}
	return
}
