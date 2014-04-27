/*
#############################################
#	GOSFML2
#	SFML Examples:	 Pong
#	Ported from C++ to Go
#############################################
*/

package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

const (
	paddleSpeed = 400
	ballSpeed   = 400
	ballRadius  = 10
	width       = 800
	height      = 600
)

type Game struct {
	leftPaddle   *sf.RectangleShape
	rightPaddle  *sf.RectangleShape
	ball         *sf.CircleShape
	ballSound    *sf.Sound
	renderWindow *sf.RenderWindow

	paddleSize       sf.Vector2f
	rightPaddleSpeed float32
	ballAngle        float32
	isPlaying        bool
}

func (this *Game) play() {
	this.paddleSize = sf.Vector2f{25, 100}

	ticker := time.NewTicker(time.Second / 60)
	AITicker := time.NewTicker(time.Second / 10)
	rand.Seed(time.Now().UnixNano())

	this.renderWindow = sf.NewRenderWindow(sf.VideoMode{width, height, 32}, "Pong (GoSFML2)", sf.StyleDefault, sf.DefaultContextSettings())

	// Load the sounds used in the game
	buffer, _ := sf.NewSoundBufferFromFile("resources/ball.wav")
	this.ballSound = sf.NewSound(buffer)

	// Create the left paddle
	this.leftPaddle, _ = sf.NewRectangleShape()
	this.leftPaddle.SetSize(sf.Vector2f{this.paddleSize.X - 3, this.paddleSize.Y - 3})
	this.leftPaddle.SetOutlineThickness(3)
	this.leftPaddle.SetOutlineColor(sf.ColorBlack())
	this.leftPaddle.SetFillColor(sf.Color{100, 100, 200, 255})
	this.leftPaddle.SetOrigin(sf.Vector2f{this.paddleSize.X / 2, this.paddleSize.Y / 2})

	// Create the right paddle
	this.rightPaddle, _ = sf.NewRectangleShape()
	this.rightPaddle.SetSize(sf.Vector2f{this.paddleSize.X - 3, this.paddleSize.Y - 3})
	this.rightPaddle.SetOutlineThickness(3)
	this.rightPaddle.SetOutlineColor(sf.ColorBlack())
	this.rightPaddle.SetFillColor(sf.Color{200, 100, 100, 255})
	this.rightPaddle.SetOrigin(sf.Vector2f{this.paddleSize.X / 2, this.paddleSize.Y / 2})

	// Create the ball
	this.ball, _ = sf.NewCircleShape()
	this.ball.SetRadius(ballRadius - 3)
	this.ball.SetOutlineThickness(3)
	this.ball.SetOutlineColor(sf.ColorBlack())
	this.ball.SetFillColor(sf.ColorWhite())
	this.ball.SetOrigin(sf.Vector2f{ballRadius / 2, ballRadius / 2})

	// Load the text font
	font, _ := sf.NewFontFromFile("resources/sansation.ttf")

	// Initialize the pause message
	pauseMessage, _ := sf.NewText(font)
	pauseMessage.SetCharacterSize(40)
	pauseMessage.SetPosition(sf.Vector2f{170, 150})
	pauseMessage.SetColor(sf.ColorWhite())
	pauseMessage.SetString("Welcome to SFML pong!\nPress space to start the game")

	for this.renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			//poll events
			this.renderWindow.DispatchEvents(this)

			//playing
			if this.isPlaying {
				deltaTime := time.Second / 60

				// Move the player's paddle
				if sf.KeyboardIsKeyPressed(sf.KeyUp) && this.leftPaddle.GetPosition().Y-this.paddleSize.Y/2 > 5 {
					this.leftPaddle.Move(sf.Vector2f{0, -paddleSpeed * float32(deltaTime.Seconds())})
				}

				if sf.KeyboardIsKeyPressed(sf.KeyDown) && this.leftPaddle.GetPosition().Y+this.paddleSize.Y/2 < height-5 {
					this.leftPaddle.Move(sf.Vector2f{0, paddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the computer's paddle
				if (this.rightPaddleSpeed < 0 && this.rightPaddle.GetPosition().Y-this.paddleSize.Y/2 > 5) ||
					(this.rightPaddleSpeed > 0 && this.rightPaddle.GetPosition().Y+this.paddleSize.Y/2 < height-5) {
					this.rightPaddle.Move(sf.Vector2f{0, this.rightPaddleSpeed * float32(deltaTime.Seconds())})
				}

				// Move the ball
				factor := ballSpeed * float32(deltaTime.Seconds())
				this.ball.Move(sf.Vector2f{float32(math.Cos(float64(this.ballAngle))) * factor, float32(math.Sin(float64(this.ballAngle))) * factor})

				// Check collisions between the ball and the screen
				if this.ball.GetPosition().X-ballRadius < 0 {
					this.isPlaying = false
					pauseMessage.SetString("You lost !\nPress space to restart or\nescape to exit")
				}

				if this.ball.GetPosition().X+ballRadius > width {
					this.isPlaying = false
					pauseMessage.SetString("You won !\nPress space to restart or\nescape to exit")
				}

				if this.ball.GetPosition().Y-ballRadius < 0 {
					this.ballAngle = -this.ballAngle
					this.ball.SetPosition(sf.Vector2f{this.ball.GetPosition().X, ballRadius + 0.1})
					this.ballSound.Play()
				}

				if this.ball.GetPosition().Y+ballRadius > height {
					this.ballAngle = -this.ballAngle
					this.ball.SetPosition(sf.Vector2f{this.ball.GetPosition().X, height - ballRadius - 0.1})
					this.ballSound.Play()
				}

				// Check the collisions between the ball and the paddles
				// Left Paddle
				if this.ball.GetPosition().X-ballRadius < this.leftPaddle.GetPosition().X+this.paddleSize.X/2 &&
					this.ball.GetPosition().X-ballRadius > this.leftPaddle.GetPosition().X &&
					this.ball.GetPosition().Y+ballRadius >= this.leftPaddle.GetPosition().Y-this.paddleSize.Y/2 &&
					this.ball.GetPosition().Y-ballRadius <= this.leftPaddle.GetPosition().Y+this.paddleSize.Y/2 {

					if this.ball.GetPosition().Y > this.leftPaddle.GetPosition().Y {
						this.ballAngle = math.Pi - this.ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						this.ballAngle = math.Pi - this.ballAngle - rand.Float32()*math.Pi*0.2
					}

					this.ball.SetPosition(sf.Vector2f{this.leftPaddle.GetPosition().X + ballRadius + this.paddleSize.X/2 + 0.1, this.ball.GetPosition().Y})
					this.ballSound.Play()
				}

				// Right Paddle
				if this.ball.GetPosition().X+ballRadius > this.rightPaddle.GetPosition().X-this.paddleSize.X/2 &&
					this.ball.GetPosition().X+ballRadius < this.rightPaddle.GetPosition().X &&
					this.ball.GetPosition().Y+ballRadius >= this.rightPaddle.GetPosition().Y-this.paddleSize.Y/2 &&
					this.ball.GetPosition().Y-ballRadius <= this.rightPaddle.GetPosition().Y+this.paddleSize.Y/2 {

					if this.ball.GetPosition().Y > this.rightPaddle.GetPosition().Y {
						this.ballAngle = math.Pi - this.ballAngle + rand.Float32()*math.Pi*0.2
					} else {
						this.ballAngle = math.Pi - this.ballAngle - rand.Float32()*math.Pi*0.2
					}

					this.ball.SetPosition(sf.Vector2f{this.rightPaddle.GetPosition().X - ballRadius - this.paddleSize.X/2 - 0.1, this.ball.GetPosition().Y})
					this.ballSound.Play()
				}
			}

			// Clear the window
			this.renderWindow.Clear(sf.Color{50, 200, 50, 0})

			if this.isPlaying {
				this.renderWindow.Draw(this.leftPaddle, sf.DefaultRenderStates())
				this.renderWindow.Draw(this.rightPaddle, sf.DefaultRenderStates())
				this.renderWindow.Draw(this.ball, sf.DefaultRenderStates())
			} else {
				this.renderWindow.Draw(pauseMessage, sf.DefaultRenderStates())
			}

			// Display things on screen
			this.renderWindow.Display()
		case <-AITicker.C:
			if this.ball.GetPosition().Y+ballRadius > this.rightPaddle.GetPosition().Y+this.paddleSize.Y/2 {
				this.rightPaddleSpeed = paddleSpeed
			} else if this.ball.GetPosition().Y-ballRadius < this.rightPaddle.GetPosition().Y-this.paddleSize.Y/2 {
				this.rightPaddleSpeed = -paddleSpeed
			} else {
				this.rightPaddleSpeed = 0
			}
		}
	}
}

func (this *Game) OnEvent(event sf.Event) {
	switch ev := event.(type) {
	case sf.EventKeyReleased:
		switch ev.Code {
		case sf.KeyEscape:
			this.renderWindow.Close()
		case sf.KeySpace:
			if !this.isPlaying {
				// (re)start the game
				this.isPlaying = true

				// reset position of the paddles and ball
				this.leftPaddle.SetPosition(sf.Vector2f{10 + this.paddleSize.X/2, height / 2})
				this.rightPaddle.SetPosition(sf.Vector2f{width - 10 - this.paddleSize.X/2, height / 2})
				this.ball.SetPosition(sf.Vector2f{width / 2, height / 2})

				// reset the ball angle
				for {
					// Make sure the ball initial angle is not too much vertical
					this.ballAngle = rand.Float32() * math.Pi * 2
					if math.Abs(math.Cos(float64(this.ballAngle))) > 0.7 {
						break
					}
				}
			}
		}
	case sf.EventClosed:
		this.renderWindow.Close()
	}
}

//main
func main() {
	game := Game{}
	game.play()
}
