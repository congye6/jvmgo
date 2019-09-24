package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

type CHECK_CAST struct {
	base.Index16
}

func (this *CHECK_CAST) Execute(frame *stack.Frame) {
	operandStack := frame.GetOperandStack()
	ref := operandStack.PopRef()
	operandStack.PushRef(ref)

	class := frame.GetMethod().GetClass()
	typeClass := class.GetConstantPool().GetClass(this.Index)

	if ref.IsInstanceOf(typeClass) {
		return
	}

	panic("java.lang.ClassCastException")
}
