package stack

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// 第一个槽出栈
type POP struct {
	base.NoOperand
}

func (this *POP) Execute(frame *stack.Frame){
	frame.GetOperandStack().PopSlot()
}

// 两字节如double，long出栈
type POP2 struct {
	base.NoOperand
}

func (this *POP2) Execute(frame *stack.Frame){
	frame.GetOperandStack().PopSlot()
	frame.GetOperandStack().PopSlot()
}