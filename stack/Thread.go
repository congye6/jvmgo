package stack

type Thread struct {
	pc    uint   //pc计数器
	stack *Stack //栈
}

func NewThread(maxSize uint) *Thread {
	return &Thread{
		pc:    0,
		stack: newStack(maxSize),
	}
}

func (this *Thread) SetPC(pc uint) {
	this.pc = pc
}

func (this *Thread) GetPC() uint {
	return this.pc
}

func (this *Thread) Push(frame *Frame) {
	this.stack.push(frame)
}

func (this *Thread) Pop() *Frame {
	return this.stack.pop()
}

func (this *Thread) CurrentFrame() *Frame {
	return this.stack.getTop()
}

func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}

func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}
