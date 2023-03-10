package references

import "jvmgo/final/instructions/base"
import "jvmgo/final/rtda"
import "jvmgo/final/rtda/heap"

// Determine if object is of given type
//由于判断对象是否为某个类的实例，该指令需要两个参数
//第一个参数是给定的要判断类对应的类符号引用
//第二个参数是需要判断的对象（从操作数栈中弹出）

type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	//将结果压回操作数栈
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
