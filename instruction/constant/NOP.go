package constant

import (
	"jvmgo/instruction/base"
	"jvmgo/stack"
)

// Do nothing
type NOP struct{ base.NoOperand }

func (self *NOP) Execute(frame *stack.Frame) {
	// really do nothing
}