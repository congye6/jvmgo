package loads

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Load reference from local variable
type ALOAD struct{ base.Index8 }

func (self *ALOAD) Execute(frame *stack.Frame) {
	aload(frame, uint(self.Index))
}

type ALOAD_0 struct{ base.NoOperand }

func (self *ALOAD_0) Execute(frame *stack.Frame) {
	aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperand }

func (self *ALOAD_1) Execute(frame *stack.Frame) {
	aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperand }

func (self *ALOAD_2) Execute(frame *stack.Frame) {
	aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperand }

func (self *ALOAD_3) Execute(frame *stack.Frame) {
	aload(frame, 3)
}

func aload(frame *stack.Frame, index uint) {
	ref := frame.GetLocalVars().GetRef(index)
	frame.GetOperandStack().PushRef(ref)
}
