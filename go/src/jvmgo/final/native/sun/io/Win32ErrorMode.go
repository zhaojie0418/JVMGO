package io

import "jvmgo/final/native"
import "jvmgo/final/rtda"

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
