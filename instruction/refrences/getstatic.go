package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/method_area"
	"jvmgo/stack"
)

type GET_STATIC struct {
	base.Index16
}

func (this *GET_STATIC) Execute(frame *stack.Frame) {
	class := frame.GetMethod().GetClass()
	constantPool := class.GetConstantPool()
	field := constantPool.GetConstantInfo(this.Index).(method_area.ConstantFieldDefInfo)
	//field.link() 加载类时已解析，不用在运行时解析

	slotId := field.GetSlotId()
	descriptor := field.GetDescriptor()
	slots := class.GetStaticVars()
	operandStack := frame.GetOperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		operandStack.PushInt(slots.GetInt(slotId))
	case 'F':
		operandStack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		operandStack.PushLong(slots.GetLong(slotId))
	case 'D':
		operandStack.PushDouble(slots.GetDouble(slotId))
	case 'L':
		operandStack.PushRef(slots.GetRef(slotId))
	}
}
