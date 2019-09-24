package loads

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Load double from local variable
type DLOAD struct{ base.Index8 }

func (self *DLOAD) Execute(frame *stack.Frame) {
	dload(frame, uint(self.Index))
}

type DLOAD_0 struct{ base.NoOperand }

func (self *DLOAD_0) Execute(frame *stack.Frame) {
	dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperand }

func (self *DLOAD_1) Execute(frame *stack.Frame) {
	dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperand }

func (self *DLOAD_2) Execute(frame *stack.Frame) {
	dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperand }

func (self *DLOAD_3) Execute(frame *stack.Frame) {
	dload(frame, 3)
}

func dload(frame *stack.Frame, index uint) {
	val := frame.GetLocalVars().GetDouble(index)
	frame.GetOperandStack().PushDouble(val)
}
