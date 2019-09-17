package conversions

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Convert long to double
type L2D struct{ base.NoOperand }

func (self *L2D) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

// Convert long to float
type L2F struct{ base.NoOperand }

func (self *L2F) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

// Convert long to int
type L2I struct{ base.NoOperand }

func (self *L2I) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
