package constant_pool

import (
	"jvmgo/classfile/reader"
)

// java中的字符串类型，本身不存字符串，存指向具体utf8常量的索引
type ConstantStringInfo struct {
	stringIndex  uint16 //字符串在常量池的索引
}

func (this *ConstantStringInfo) readInfo(reader *reader.ClassReader) {
	this.stringIndex = reader.ReadUint16()
}
