package lang

import "runtime"
import "jvmgo/final/native"
import "jvmgo/final/rtda"

const jlRuntime = "java/lang/Runtime"

func init() {
	native.Register(jlRuntime, "availableProcessors", "()I", availableProcessors)
}

// public native int availableProcessors();
// ()I
func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()

	stack := frame.OperandStack()
	stack.PushInt(int32(numCPU))
}
