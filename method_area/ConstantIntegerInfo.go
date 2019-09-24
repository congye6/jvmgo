package method_area

// Integer常量，包括int，bool，char，byte，short
type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) GetVal() int32 {
	return this.val
}
