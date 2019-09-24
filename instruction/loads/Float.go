package loads

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Load float from local variable
type FLOAD struct{ base.Index8 }

func (self *FLOAD) Execute(frame *stack.Frame) {
	fload(frame, uint(self.Index))
}

type FLOAD_0 struct{ base.NoOperand }

func (self *FLOAD_0) Execute(frame *stack.Frame) {
	fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperand }

func (self *FLOAD_1) Execute(frame *stack.Frame) {
	fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperand }

func (self *FLOAD_2) Execute(frame *stack.Frame) {
	fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperand }

func (self *FLOAD_3) Execute(frame *stack.Frame) {
	fload(frame, 3)
}

func fload(frame *stack.Frame, index uint) {
	val := frame.GetLocalVars().GetFloat(index)
	frame.GetOperandStack().PushFloat(val)
}
