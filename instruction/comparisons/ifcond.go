package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Branch if int comparison with zero succeeds
type IFEQ struct{ base.Branch }

func (self *IFEQ) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val == 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IFNE struct{ base.Branch }

func (self *IFNE) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val != 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IFLT struct{ base.Branch }

func (self *IFLT) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val < 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IFLE struct{ base.Branch }

func (self *IFLE) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val <= 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IFGT struct{ base.Branch }

func (self *IFGT) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val > 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IFGE struct{ base.Branch }

func (self *IFGE) Execute(frame *stack.Frame) {
	val := frame.GetOperandStack().PopInt()
	if val >= 0 {
		base.BranchGoTo(frame, self.Offset)
	}
}
