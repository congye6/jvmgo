package extended

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Branch if reference is null
type IFNULL struct{ base.Branch }

func (self *IFNULL) Execute(frame *stack.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref == nil {
		base.BranchGoTo(frame, self.Offset)
	}
}

// Branch if reference not null
type IFNONNULL struct{ base.Branch }

func (self *IFNONNULL) Execute(frame *stack.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref != nil {
		base.BranchGoTo(frame, self.Offset)
	}
}
