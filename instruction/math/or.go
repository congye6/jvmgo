package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Boolean OR int
type IOR struct{ base.NoOperand }

func (self *IOR) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

// Boolean OR long
type LOR struct{ base.NoOperand }

func (self *LOR) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
