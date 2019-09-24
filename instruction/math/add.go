package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Add double
type DADD struct{ base.NoOperand }

func (self *DADD) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// Add float
type FADD struct{ base.NoOperand }

func (self *FADD) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// Add int
type IADD struct{ base.NoOperand }

func (self *IADD) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// Add long
type LADD struct{ base.NoOperand }

func (self *LADD) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
