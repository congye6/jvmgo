package stores

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Store double into local variable
type DSTORE struct{ base.Index8 }

func (self *DSTORE) Execute(frame *stack.Frame) {
	dstore(frame, uint(self.Index))
}

type DSTORE_0 struct{ base.NoOperand }

func (self *DSTORE_0) Execute(frame *stack.Frame) {
	dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperand }

func (self *DSTORE_1) Execute(frame *stack.Frame) {
	dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperand }

func (self *DSTORE_2) Execute(frame *stack.Frame) {
	dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperand }

func (self *DSTORE_3) Execute(frame *stack.Frame) {
	dstore(frame, 3)
}

func dstore(frame *stack.Frame, index uint) {
	val := frame.GetOperandStack().PopDouble()
	frame.GetLocalVars().SetDouble(index, val)
}
