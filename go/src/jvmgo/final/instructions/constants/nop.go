package constants

import "jvmgo/final/instructions/base"
import "jvmgo/final/rtda"

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
