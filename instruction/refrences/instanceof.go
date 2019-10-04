package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

type INSTANCE_OF struct {
	base.Index16
}

func (this *INSTANCE_OF) Execute(frame *stack.Frame) {
	operandStack := frame.GetOperandStack()
	ref := operandStack.PopRef()
	if ref == nil {
		operandStack.PushInt(0)
		return
	}

	currentClass := frame.GetMethod().GetClass()
	instanceClass := currentClass.GetConstantPool().GetClass(this.Index)

	if ref.IsInstanceOf(instanceClass) {
		operandStack.PushInt(1)
	} else {
		operandStack.PushInt(0)
	}
}
