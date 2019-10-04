package refrences

import (
	"jvmgo/instruction/base"
	"jvmgo/method_area"
	"jvmgo/stack"
)

type PUT_FIELD struct {
	base.Index16
}

func (this *PUT_FIELD) Execute(frame *stack.Frame) {
	class := frame.GetMethod().GetClass()
	field := class.GetConstantPool().GetConstantInfo(this.Index).(method_area.ConstantFieldDefInfo)

	operandStack := frame.GetOperandStack()
	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := operandStack.PopInt()
		ref := operandStack.PopRef()
		fields := ref.GetFields()
		fields.SetInt(slotId, val)
	case 'F':
		val := operandStack.PopFloat()
		ref := operandStack.PopRef()
		fields := ref.GetFields()
		fields.SetFloat(slotId, val)
	case 'J':
		val := operandStack.PopLong()
		ref := operandStack.PopRef()
		fields := ref.GetFields()
		fields.SetLong(slotId, val)
	case 'D':
		val := operandStack.PopDouble()
		ref := operandStack.PopRef()
		fields := ref.GetFields()
		fields.SetDouble(slotId, val)
	case 'L':
		val := operandStack.PopRef()
		ref := operandStack.PopRef()
		fields := ref.GetFields()
		fields.SetRef(slotId, val)
	}

}
