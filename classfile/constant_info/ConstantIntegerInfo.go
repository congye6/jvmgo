package constant_info

import (
	"jvmgo/classfile/reader"
)

// Integer常量，包括int，bool，char，byte，short
type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *reader.ClassReader){
	this.val = int32(reader.ReadUint32())
}
