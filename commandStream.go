// Copyright (C) 2013 by krepa098. All rights reserved.
// Use of this source code is governed by a zlib-style
// license that can be found in the license.txt file.

package gosfml2

import (
	"runtime"
)

//Lecture:
//http://code.google.com/p/go-wiki/wiki/LockOSThread

/////////////////////////////////////
///		VARS
/////////////////////////////////////

var cstream = newCommandStream()

/////////////////////////////////////
///		STRUCTS
/////////////////////////////////////

//A command stream executes functions in a specific thread
type CommandStream struct {
	commands chan func()
}

/////////////////////////////////////
///		FUNCS
/////////////////////////////////////

//Creates a new command stream
func newCommandStream() CommandStream {
	cs := CommandStream{make(chan func(), 8)}

	go func() {
		runtime.LockOSThread()
		cs.Run()
	}()

	return cs
}

//Executes the pending commands
func (this CommandStream) Run() {
	for f := range this.commands {
		f()
	}
}

//Add a command to the stream
func (this CommandStream) Exec(f func()) {
	//this.ExecAndBlock(f)
	this.commands <- f
}

//Add a command to the stream and wait till the end of its execution
func (this CommandStream) ExecAndBlock(f func()) {
	done := make(chan bool, 1)
	this.commands <- func() {
		f()
		done <- true
	}
	<-done
}
