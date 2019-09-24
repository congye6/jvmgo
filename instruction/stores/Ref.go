package stores

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Store reference into local variable
type ASTORE struct{ base.Index8 }

func (self *ASTORE) Execute(frame *stack.Frame) {
	astore(frame, uint(self.Index))
}

type ASTORE_0 struct{ base.NoOperand }

func (self *ASTORE_0) Execute(frame *stack.Frame) {
	astore(frame, 0)
}

type ASTORE_1 struct{ base.NoOperand }

func (self *ASTORE_1) Execute(frame *stack.Frame) {
	astore(frame, 1)
}

type ASTORE_2 struct{ base.NoOperand }

func (self *ASTORE_2) Execute(frame *stack.Frame) {
	astore(frame, 2)
}

type ASTORE_3 struct{ base.NoOperand }

func (self *ASTORE_3) Execute(frame *stack.Frame) {
	astore(frame, 3)
}

func astore(frame *stack.Frame, index uint) {
	ref := frame.GetOperandStack().PopRef()
	frame.GetLocalVars().SetRef(index, ref)
}
