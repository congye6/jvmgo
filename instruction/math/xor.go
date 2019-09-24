package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Boolean XOR int
type IXOR struct{ base.NoOperand }

func (self *IXOR) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long
type LXOR struct{ base.NoOperand }

func (self *LXOR) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
