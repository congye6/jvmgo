package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Divide double
type DDIV struct{ base.NoOperand }

func (self *DDIV) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

// Divide float
type FDIV struct{ base.NoOperand }

func (self *FDIV) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

// Divide int
type IDIV struct{ base.NoOperand }

func (self *IDIV) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushInt(result)
	}
}

// Divide long
type LDIV struct{ base.NoOperand }

func (self *LDIV) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 / v2
		stack.PushLong(result)
	}
}
