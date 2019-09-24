package loads

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Load long from local variable
type LLOAD struct{ base.Index8 }

func (self *LLOAD) Execute(frame *stack.Frame) {
	lload(frame, uint(self.Index))
}

type LLOAD_0 struct{ base.NoOperand }

func (self *LLOAD_0) Execute(frame *stack.Frame) {
	lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperand }

func (self *LLOAD_1) Execute(frame *stack.Frame) {
	lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperand }

func (self *LLOAD_2) Execute(frame *stack.Frame) {
	lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperand }

func (self *LLOAD_3) Execute(frame *stack.Frame) {
	lload(frame, 3)
}

func lload(frame *stack.Frame, index uint) {
	val := frame.GetLocalVars().GetLong(index)
	frame.GetOperandStack().PushLong(val)
}
