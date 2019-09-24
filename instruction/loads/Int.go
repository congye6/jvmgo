package loads

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

func iload(frame *stack.Frame, index uint) {
	val := frame.GetLocalVars().GetInt(index)
	frame.GetOperandStack().PushInt(val)
}

type ILOAD struct {
	base.Index8
}

func (this *ILOAD) Execute(frame *stack.Frame) {
	iload(frame, this.Index)
}

type ILOAD_0 struct{ base.NoOperand }

func (self *ILOAD_0) Execute(frame *stack.Frame) {
	iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperand }

func (self *ILOAD_1) Execute(frame *stack.Frame) {
	iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperand }

func (self *ILOAD_2) Execute(frame *stack.Frame) {
	iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperand }

func (self *ILOAD_3) Execute(frame *stack.Frame) {
	iload(frame, 3)
}

