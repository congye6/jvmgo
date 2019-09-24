package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Boolean AND int
type IAND struct{ base.NoOperand }

func (self *IAND) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

// Boolean AND long
type LAND struct{ base.NoOperand }

func (self *LAND) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}
