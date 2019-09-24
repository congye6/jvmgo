package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Compare float
type FCMPG struct{ base.NoOperand }

func (self *FCMPG) Execute(frame *stack.Frame) {
	_fcmp(frame, true)
}

type FCMPL struct{ base.NoOperand }

func (self *FCMPL) Execute(frame *stack.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *stack.Frame, gFlag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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
