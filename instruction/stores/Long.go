package stores

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Store long into local variable
type LSTORE struct{ base.Index8 }

func (self *LSTORE) Execute(frame *stack.Frame) {
	lstore(frame, uint(self.Index))
}

type LSTORE_0 struct{ base.NoOperand }

func (self *LSTORE_0) Execute(frame *stack.Frame) {
	lstore(frame, 0)
}

type LSTORE_1 struct{ base.NoOperand }

func (self *LSTORE_1) Execute(frame *stack.Frame) {
	lstore(frame, 1)
}

type LSTORE_2 struct{ base.NoOperand }

func (self *LSTORE_2) Execute(frame *stack.Frame) {
	lstore(frame, 2)
}

type LSTORE_3 struct{ base.NoOperand }

func (self *LSTORE_3) Execute(frame *stack.Frame) {
	lstore(frame, 3)
}

func lstore(frame *stack.Frame, index uint) {
	val := frame.GetOperandStack().PopLong()
	frame.GetLocalVars().SetLong(index, val)
}
