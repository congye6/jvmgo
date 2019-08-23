package constant_info

import (
	"jvmgo/classfile/reader"
)

// 类描述符，包括类，超类，接口都用这个表示
type ConstantClassInfo struct {
	nameIndex    uint16 // 符号引用的索引 ？
}

func (this *ConstantClassInfo) readInfo(reader *reader.ClassReader){
	this.nameIndex=reader.ReadUint16()
}