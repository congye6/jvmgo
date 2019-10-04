package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

type NEW struct {
	base.Index16
}

func (this *NEW) Execute(frame *stack.Frame) {
	constans := frame.GetMethod().GetClass().GetConstantPool()
	class := constans.GetClass(this.Index)
	ref := class.NewObject()
	frame.GetOperandStack().PushRef(ref)
}
