package constant_pool

import (
	"jvmgo/classfile/reader"
)

// 字段名，描述符等字符串常量
// 使用MUTF8，如果包含null或补充字符将和utf8不兼容
type ConstantUtf8Info struct {
	str string
}

func (this *ConstantUtf8Info) readInfo(reader *reader.ClassReader) {
	length := uint32(reader.ReadUint16())
	bytes := reader.ReadBytes(length)
	this.str = string(bytes)
}
