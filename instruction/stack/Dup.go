package stack

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// 复制栈顶元素
type DUP struct {
	base.NoOperand
}

func (this *DUP) Execute(frame *stack.Frame){
	operandStack:=frame.GetOperandStack()
	slot := operandStack.PopSlot()
	operandStack.PushSlot(slot)
	operandStack.PushSlot(slot)
}

// Duplicate the top operand stack value and insert two values down
type DUP_X1 struct{ base.NoOperand }

func (self *DUP_X1) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top operand stack value and insert two or three values down
type DUP_X2 struct{ base.NoOperand }

func (self *DUP_X2) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values
type DUP2 struct{ base.NoOperand }

func (self *DUP2) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type DUP2_X1 struct{ base.NoOperand }

func (self *DUP2_X1) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type DUP2_X2 struct{ base.NoOperand }

func (self *DUP2_X2) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	val3 := stack.PopSlot()
	val4 := stack.PopSlot()
	stack.PushSlot(val2)
	stack.PushSlot(val1)
	stack.PushSlot(val4)
	stack.PushSlot(val3)
	stack.PushSlot(val2)
	stack.PushSlot(val1)
}
