package stores

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Store float into local variable
type FSTORE struct{ base.Index8 }

func (self *FSTORE) Execute(frame *stack.Frame) {
	fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ base.NoOperand }

func (self *FSTORE_0) Execute(frame *stack.Frame) {
	fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperand }

func (self *FSTORE_1) Execute(frame *stack.Frame) {
	fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperand }

func (self *FSTORE_2) Execute(frame *stack.Frame) {
	fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperand }

func (self *FSTORE_3) Execute(frame *stack.Frame) {
	fstore(frame, 3)
}

func fstore(frame *stack.Frame, index uint) {
	val := frame.GetOperandStack().PopFloat()
	frame.GetLocalVars().SetFloat(index, val)
}
