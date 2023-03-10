package base

import "fmt"
import "strings"
import "jvmgo/final/rtda"
import "jvmgo/final/rtda/heap"

// InvokeMethod 这里的InvokeMethod是其它四条函数调用指令普适的部分，所以使用其提取出来
//	调用指令总结
/*  静态方法的调用，生成的是invokestatic指令。
    接口方法的调用，生成的是invokeinterface指令。
    其他的方法，一般生成invokevirtual指令，尽管final方法不可能被继承覆盖重写，但还是生成invokevirtual指令。
	调用构造器是生成invokespecial指令，此外super关键字方法调用，也生成invokespecial指令。
	最复杂的当属于invokedynamic指令了，其目的是实现动态类型
*/
func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	//_logInvoke(callerFrame.Thread().StackDepth(), method)
	//创建一个新的栈帧
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	//将n个变量从调用者的操作数栈中弹出并放进被调用方法的局部变量表中
	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		for i := argSlotCount - 1; i >= 0; i-- {
			//不需要对long和double类型做特殊处理，因为在slot中处理过了已经
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i), slot)
		}
	}
}

func _logInvoke(stackSize uint, method *heap.Method) {
	space := strings.Repeat(" ", int(stackSize))
	className := method.Class().Name()
	methodName := method.Name()
	fmt.Printf("[method]%v %v.%v()\n", space, className, methodName)
}
