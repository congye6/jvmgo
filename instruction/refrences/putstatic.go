package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/method_area"
	"jvmgo/stack"
)

type PUT_STATIC struct {
	base.Index16
}

func (this *PUT_STATIC) Execute(frame *stack.Frame) {
	method := frame.GetMethod()
	class := method.GetClass()
	constantPool := class.GetConstantPool()
	field := constantPool.GetConstantInfo(this.Index).(method_area.ConstantFieldDefInfo)
	descriptor := field.GetDescriptor()

	slotId := field.GetSlotId()
	slots := class.GetStaticVars()
	operandStack := frame.GetOperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, operandStack.PopInt())
	case 'F':
		slots.SetFloat(slotId, operandStack.PopFloat())
	case 'J':
		slots.SetLong(slotId, operandStack.PopLong())
	case 'D':
		slots.SetDouble(slotId, operandStack.PopDouble())
	case 'L':
		slots.SetRef(slotId, operandStack.PopRef())
	}
}
