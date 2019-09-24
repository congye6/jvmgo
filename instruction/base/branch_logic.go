package base

import "jvmgo/stack"

func BranchGoTo(frame *stack.Frame,offset uint){
	pc:=frame.CurrentThread().GetPC()
	pc=pc+offset
	frame.CurrentThread().SetPC(pc)
}
