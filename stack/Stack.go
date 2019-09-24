package stack

type Stack struct {
	top     *Frame // 栈顶的栈帧
	size    uint   // 当前大小
	maxSize uint   // 最多能放多少个元素
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("[ERROR] StackOverFlowException \n")
	}

	frame.next = this.top
	this.top = frame
	this.size++
}

func (this *Stack) pop() *Frame {
	result := this.top
	this.top = result.next
	result.next = nil
	this.size--
	return result
}

func (this *Stack) getTop() *Frame {
	return this.top
}
