package conversions

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Convert float to double
type F2D struct{ base.NoOperand }

func (self *F2D) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

// Convert float to int
type F2I struct{ base.NoOperand }

func (self *F2I) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

// Convert float to long
type F2L struct{ base.NoOperand }

func (self *F2L) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
