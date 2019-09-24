package stack

import (
	"jvmgo/method_area"
	"jvmgo/slot"
)

type OperandStack struct {
	top   uint //栈顶下标+1
	slots slot.Slots
}

func (this *OperandStack) PushInt(num int32) {
	this.slots.SetInt(this.top, num)
	this.top++
}

func (this *OperandStack) PopInt() int32 {
	this.top--
	return this.slots.GetInt(this.top)
}

func (this *OperandStack) PushLong(num int64) {
	this.slots.SetLong(this.top, num)
	this.top += 2
}

func (this *OperandStack) PopLong() int64 {
	this.top -= 2
	return this.slots.GetLong(this.top)
}

func (this *OperandStack) PushDouble(num float64) {
	this.slots.SetDouble(this.top, num)
	this.top += 2
}

func (this *OperandStack) PopDouble() float64 {
	this.top -= 2
	return this.slots.GetDouble(this.top)
}

func (this *OperandStack) PushFloat(num float32) {
	this.slots.SetFloat(this.top, num)
	this.top++
}

func (this *OperandStack) PopFloat() float32 {
	this.top--
	return this.slots.GetFloat(this.top)
}

func (this *OperandStack) PushRef(ref *method_area.Object) {
	this.slots.SetRef(this.top, ref)
	this.top++
}

func (this *OperandStack) PopRef() *method_area.Object {
	this.top--
	return this.slots.GetRef(this.top)
}

func (this *OperandStack) PushSlot(slot *slot.Slot) {
	this.slots.SetSlot(this.top, slot)
	this.top++
}

func (this *OperandStack) PopSlot() *slot.Slot {
	this.top--
	return this.slots.GetSlot(this.top)
}
