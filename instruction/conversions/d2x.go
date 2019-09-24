package conversions

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Convert double to float
type D2F struct{ base.NoOperand }

func (self *D2F) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

// Convert double to int
type D2I struct{ base.NoOperand }

func (self *D2I) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

// Convert double to long
type D2L struct{ base.NoOperand }

func (self *D2L) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
