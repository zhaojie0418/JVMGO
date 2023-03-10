package reserved

import "jvmgo/final/instructions/base"
import "jvmgo/final/rtda"
import "jvmgo/final/native"
import _ "jvmgo/final/native/java/io"
import _ "jvmgo/final/native/java/lang"
import _ "jvmgo/final/native/java/security"
import _ "jvmgo/final/native/java/util/concurrent/atomic"
import _ "jvmgo/final/native/sun/io"
import _ "jvmgo/final/native/sun/misc"
import _ "jvmgo/final/native/sun/reflect"

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
