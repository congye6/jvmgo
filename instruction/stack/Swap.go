package stack

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

type SWAP struct {
	base.NoOperand
}

func (this *SWAP) Execute(frame *stack.Frame){
	operandStack:=frame.GetOperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot2)
}
