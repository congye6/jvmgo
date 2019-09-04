package constant_pool

import (
	"jvmgo/classfile/reader"
)

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *reader.ClassReader) {
	this.val = int64(reader.ReadUint64())
}
