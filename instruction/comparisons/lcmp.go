package comparisons

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

type LCMP struct {
	base.NoOperand
}

func (this *LCMP) Execute(frame *stack.Frame) {
	operandStack := frame.GetOperandStack()
	v2 := operandStack.PopLong()
	v1 := operandStack.PopLong()
	if v1 > v2 {
		operandStack.PushInt(1)
	} else if v1 == v2 {
		operandStack.PushInt(0)
	} else {
		operandStack.PushInt(-1)
	}
}
