package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Branch if reference comparison succeeds
type IF_ACMPEQ struct{ base.Branch }

func (self *IF_ACMPEQ) Execute(frame *stack.Frame) {
	if _acmp(frame) {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ACMPNE struct{ base.Branch }

func (self *IF_ACMPNE) Execute(frame *stack.Frame) {
	if !_acmp(frame) {
		base.BranchGoTo(frame, self.Offset)
	}
}

func _acmp(frame *stack.Frame) bool {
	stack := frame.GetOperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}
