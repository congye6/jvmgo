package control

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Return void from method
type RETURN struct{ base.NoOperand }

func (self *RETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	thread.PopFrame()
}

// Return reference from method
type ARETURN struct{ base.NoOperand }

func (self *ARETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	ref := currentFrame.GetOperandStack().PopRef()
	invokerFrame.GetOperandStack().PushRef(ref)
}

// Return double from method
type DRETURN struct{ base.NoOperand }

func (self *DRETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.GetOperandStack().PopDouble()
	invokerFrame.GetOperandStack().PushDouble(val)
}

// Return float from method
type FRETURN struct{ base.NoOperand }

func (self *FRETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.GetOperandStack().PopFloat()
	invokerFrame.GetOperandStack().PushFloat(val)
}

// Return int from method
type IRETURN struct{ base.NoOperand }

func (self *IRETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.GetOperandStack().PopInt()
	invokerFrame.GetOperandStack().PushInt(val)
}

// Return double from method
type LRETURN struct{ base.NoOperand }

func (self *LRETURN) Execute(frame *stack.Frame) {
	thread := frame.CurrentThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	val := currentFrame.GetOperandStack().PopLong()
	invokerFrame.GetOperandStack().PushLong(val)
}
