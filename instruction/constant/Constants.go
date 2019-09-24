package constant

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// 把隐含在操作码中的常量压入操作数栈

type ACONST_NULL struct {
	base.NoOperand
}

func (this *ACONST_NULL) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushRef(nil)
}

type DCONST_0 struct {
	base.NoOperand
}

func (this *DCONST_0) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushDouble(0.0)
}

type DCONST_1 struct{ base.NoOperand }

func (this *DCONST_1) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushDouble(1.0)
}

// Push float
type FCONST_0 struct{ base.NoOperand }

func (self *FCONST_0) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushFloat(0.0)
}

type FCONST_1 struct{ base.NoOperand }

func (self *FCONST_1) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushFloat(1.0)
}

type FCONST_2 struct{ base.NoOperand }

func (self *FCONST_2) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushFloat(2.0)
}

// Push int constant
type ICONST_M1 struct{ base.NoOperand }

func (self *ICONST_M1) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(-1)
}

type ICONST_0 struct{ base.NoOperand }

func (self *ICONST_0) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(0)
}

type ICONST_1 struct{ base.NoOperand }

func (self *ICONST_1) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(1)
}

type ICONST_2 struct{ base.NoOperand }

func (self *ICONST_2) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(2)
}

type ICONST_3 struct{ base.NoOperand }

func (self *ICONST_3) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(3)
}

type ICONST_4 struct{ base.NoOperand }

func (self *ICONST_4) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(4)
}

type ICONST_5 struct{ base.NoOperand }

func (self *ICONST_5) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushInt(5)
}

// Push long constant
type LCONST_0 struct{ base.NoOperand }

func (self *LCONST_0) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushLong(0)
}

type LCONST_1 struct{ base.NoOperand }

func (self *LCONST_1) Execute(frame *stack.Frame) {
	frame.GetOperandStack().PushLong(1)
}
