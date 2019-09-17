package math

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
	"math"
)

// 求余
type DREM struct {
	base.NoOperand
}

func (this *DREM) Execute(frame *stack.Frame) {
	operandStack := frame.GetOperandStack()
	num1 := operandStack.PopDouble()
	num2 := operandStack.PopDouble()
	if num1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	operandStack.PushDouble(math.Mod(num2, num1))
}

// Remainder float
type FREM struct{ base.NoOperand }

func (self *FREM) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2))) // todo
	stack.PushFloat(result)
}

// Remainder int
type IREM struct{ base.NoOperand }

func (self *IREM) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 % v2
		stack.PushInt(result)
	}
}

// Remainder long
type LREM struct{ base.NoOperand }

func (self *LREM) Execute(frame *stack.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()

	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	} else {
		result := v1 % v2
		stack.PushLong(result)
	}
}
