package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Compare double
type DCMPG struct{ base.NoOperand }

func (self *DCMPG) Execute(frame *stack.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperand }

func (self *DCMPL) Execute(frame *stack.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *stack.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
