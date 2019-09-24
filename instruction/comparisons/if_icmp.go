package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Branch if int comparison succeeds
type IF_ICMPEQ struct{ base.Branch }

func (self *IF_ICMPEQ) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ICMPNE struct{ base.Branch }

func (self *IF_ICMPNE) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ICMPLT struct{ base.Branch }

func (self *IF_ICMPLT) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ICMPLE struct{ base.Branch }

func (self *IF_ICMPLE) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ICMPGT struct{ base.Branch }

func (self *IF_ICMPGT) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

type IF_ICMPGE struct{ base.Branch }

func (self *IF_ICMPGE) Execute(frame *stack.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.BranchGoTo(frame, self.Offset)
	}
}

func _icmpPop(frame *stack.Frame) (val1, val2 int32) {
	stack := frame.GetOperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
