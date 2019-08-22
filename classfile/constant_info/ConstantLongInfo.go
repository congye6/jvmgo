package constant_info

import "jvmgo/classfile"

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *classfile.ClassReader) {
	this.val = int64(reader.ReadUint64())
}
