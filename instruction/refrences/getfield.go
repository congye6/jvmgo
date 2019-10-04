package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/method_area"
	"jvmgo/stack"
)

type GET_FIELD struct {
	base.Index16
}

func (this *GET_FIELD) Execute(frame *stack.Frame) {
	class := frame.GetMethod().GetClass()
	field := class.GetConstantPool().GetConstantInfo(this.Index).(method_area.ConstantFieldDefInfo)

	slotId := field.GetSlotId()
	operandStack := frame.GetOperandStack()
	descriptor := field.GetDescriptor()

	ref := operandStack.PopRef()
	fields := ref.GetFields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		operandStack.PushInt(fields.GetInt(slotId))
	case 'F':
		operandStack.PushFloat(fields.GetFloat(slotId))
	case 'J':
		operandStack.PushLong(fields.GetLong(slotId))
	case 'D':
		operandStack.PushDouble(fields.GetDouble(slotId))
	case 'L', '[':
		operandStack.PushRef(fields.GetRef(slotId))
	}
}
