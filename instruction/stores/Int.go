package stores

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

func istore(frame *stack.Frame, index uint) {
	val := frame.GetOperandStack().PopInt()
	frame.GetLocalVars().SetInt(index, val)
}

type ISTORE struct {
	base.Index8
}

func (this *ISTORE) Execute(frame *stack.Frame) {
	istore(frame, this.Index)
}

type ISTORE_0 struct{ base.NoOperand }

func (self *ISTORE_0) Execute(frame *stack.Frame) {
	istore(frame, 0)
}

type ISTORE_1 struct{ base.NoOperand }

func (self *ISTORE_1) Execute(frame *stack.Frame) {
	istore(frame, 1)
}

type ISTORE_2 struct{ base.NoOperand }

func (self *ISTORE_2) Execute(frame *stack.Frame) {
	istore(frame, 2)
}

type ISTORE_3 struct{ base.NoOperand }

func (self *ISTORE_3) Execute(frame *stack.Frame) {
	istore(frame, 3)
}
