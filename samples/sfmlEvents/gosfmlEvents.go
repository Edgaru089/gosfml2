/*
#############################################
#	GOSFML2
#	Events
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"runtime"
	"unicode"
)

func init() {
	runtime.LockOSThread()
}

/////////////////////////////////////
///		LOGGER
/////////////////////////////////////

type Logger []*sf.Text

func newLogger(nbOfItems int, font *sf.Font) Logger {
	logger := make(Logger, nbOfItems)
	for i := 0; i < len(logger); i++ {
		logger[i], _ = sf.NewText(font)
		logger[i].SetColor(sf.ColorBlack())
		logger[i].SetPosition(sf.Vector2f{100, 150 + float32(i)*20})
		logger[i].SetCharacterSize(12)
	}
	return logger
}

//implement io.Writer
func (logger Logger) Write(p []byte) (n int, err error) {
	for i := 0; i < len(logger)-1; i++ {
		oldMsg := logger[i+1].GetString()
		logger[i].SetString(oldMsg)
	}
	logger[len(logger)-1].SetString(string(p))

	return len(p), nil
}

/////////////////////////////////////
///		Event Handler
/////////////////////////////////////

type MyEventHandler struct {
	logger Logger
	wnd    *sf.RenderWindow
}

func (this MyEventHandler) OnEvent(event sf.Event) {
	switch ev := event.(type) {
	case sf.EventKeyPressed:
		fmt.Fprint(this.logger, "Key pressed: ", int(ev.Code), " Shift: ", ev.Shift, " Control: ", ev.Control, " Alt: ", ev.Alt, " System: ", ev.System)

		//exit on ESC
		if ev.Code == sf.KeyEscape {
			this.wnd.Close()
		}
	case sf.EventKeyReleased:
		fmt.Fprint(this.logger, "Key released: ", int(ev.Code), " Shift: ", ev.Shift, " Control: ", ev.Control, " Alt: ", ev.Alt, " System: ", ev.System)
	case sf.EventGainedFocus:
		fmt.Fprint(this.logger, "Gained Focus")
	case sf.EventLostFocus:
		fmt.Fprint(this.logger, "Lost Focus")
	case sf.EventResized:
		fmt.Fprint(this.logger, "Resized width: ", ev.Width, " height: ", ev.Height)
	case sf.EventTextEntered:
		if unicode.IsPrint(ev.Char) {
			fmt.Fprint(this.logger, "Text entered: ", string(ev.Char))
		}
	case sf.EventMouseButtonPressed:
		fmt.Fprint(this.logger, "Mouse pressed: ", int(ev.Button), " [X: ", ev.X, " Y: ", ev.Y, "]")
	case sf.EventMouseButtonReleased:
		fmt.Fprint(this.logger, "Mouse released: ", int(ev.Button), " [X: ", ev.X, " Y: ", ev.Y, "]")
	case sf.EventMouseLeft:
		fmt.Fprint(this.logger, "Mouse left")
	case sf.EventMouseEntered:
		fmt.Fprint(this.logger, "Mouse entered")
	case sf.EventMouseWheelMoved:
		fmt.Fprint(this.logger, "Mouse wheel moved: ", ev.Delta)
	case sf.EventMouseMoved:
		fmt.Fprint(this.logger, "Mouse moved: [X: ", ev.X, " Y: ", ev.Y, "]")
	case sf.EventClosed:
		this.wnd.Close()
	case sf.EventJoystickConnected:
		fmt.Fprint(this.logger, "Joystick connected id: ", ev.JoystickId)
	case sf.EventJoystickDisconnected:
		fmt.Fprint(this.logger, "Joystick disconnected id: ", ev.JoystickId)
	case sf.EventJoystickButtonPressed:
		fmt.Fprint(this.logger, "Joystick Button pressed: ", ev.Button)
	case sf.EventJoystickButtonReleased:
		fmt.Fprint(this.logger, "Joystick Button released: ", ev.Button)
	case sf.EventJoystickMoved:
		fmt.Fprint(this.logger, "Joystick moved: [Axis", int(ev.Axis), "Value: ", ev.Position, "]")
	}
}

/////////////////////////////////////
///		MAIN
/////////////////////////////////////

func main() {
	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Events (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())

	//load font
	font, _ := sf.NewFontFromFile("resources/Vera.ttf")

	text, _ := sf.NewText(font)
	text.SetColor(sf.ColorBlack())
	text.SetPosition(sf.Vector2f{40, 50})
	text.SetString("Generate input!\nLog:")

	handler := MyEventHandler{
		logger: newLogger(20, font),
		wnd:    renderWindow,
	}

	for renderWindow.IsOpen() {
		renderWindow.DispatchEvents(handler.OnEvent)
		renderWindow.Clear(sf.ColorWhite())
		renderWindow.Draw(text, sf.DefaultRenderStates())
		for _, t := range handler.logger {
			renderWindow.Draw(t, sf.DefaultRenderStates())
		}
		renderWindow.Display()
	}
}
