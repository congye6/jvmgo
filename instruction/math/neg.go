package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Negate double
type DNEG struct{ base.NoOperand }

func (self *DNEG) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// Negate float
type FNEG struct{ base.NoOperand }

func (self *FNEG) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// Negate int
type INEG struct{ base.NoOperand }

func (self *INEG) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// Negate long
type LNEG struct{ base.NoOperand }

func (self *LNEG) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
