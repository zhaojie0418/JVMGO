package references

import "jvmgo/final/instructions/base"
import "jvmgo/final/rtda"
import "jvmgo/final/rtda/heap"

// Check whether object is of given type
//相比于instanceof，CHECK_CAST并不会把结果压栈而是直接如果判断失败就抛异常

type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	//取出来再给压回去，保证其操作数栈不改变
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
