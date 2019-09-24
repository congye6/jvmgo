package stack

import (
	"jvmgo/method_area"
	"jvmgo/slot"
)

type Frame struct {
	next         *Frame              // 栈中下一个元素
	method       *method_area.Method // 方法的直接引用
	localVars    slot.Slots          // 局部变量表
	operandStack *OperandStack       // 操作数栈
	thread       *Thread
}

func NewFrame(thread *Thread, next *Frame, method *method_area.Method) *Frame {
	return &Frame{
		next:      next,
		method:    method,
		localVars: slot.NewSlots(uint(method.GetMaxLocals())),
		operandStack: &OperandStack{
			top:   0,
			slots: slot.NewSlots(uint(method.GetMaxStack())),
		},
		thread: thread,
	}
}

func (this *Frame) GetOperandStack() *OperandStack {
	return this.operandStack
}

func (this *Frame) GetLocalVars() slot.Slots {
	return this.localVars
}

func (this *Frame) GetMethod() *method_area.Method {
	return this.method
}

func (this *Frame) CurrentThread() *Thread {
	return this.thread
}

func (this *Frame) GetConstantPool() *method_area.ConstantPool {
	return this.method.GetClass().GetConstantPool()
}
