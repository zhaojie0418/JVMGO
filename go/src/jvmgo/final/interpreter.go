package main

import "fmt"
import "jvmgo/final/instructions"
import "jvmgo/final/instructions/base"
import "jvmgo/final/rtda"

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

//recover 仅在延迟函数 defer 中有效，在正常的执行过程中，
//调用 recover 会返回 nil 并且没有其他任何效果，
//如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行
//recover相当于try-catch，用于捕获异常
func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		//如果还有待执行的栈帧则继续执行
		if thread.IsStackEmpty() {
			break
		}
	}
}

//用于在方法执行过程中打印指令信息
func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

//用于打印java虚拟机栈信息
func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		lineNum := method.GetLineNumber(frame.NextPC())
		fmt.Printf(">> line:%4d pc:%4d %v.%v%v \n",
			lineNum, frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
