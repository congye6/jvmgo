package constant_info

import "jvmgo/classfile"

// Integer常量，包括int，bool，char，byte，short
type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *classfile.ClassReader){
	this.val = int32(reader.ReadUint32())
}
