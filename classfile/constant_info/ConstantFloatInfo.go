package constant_info

import (
	"jvmgo/classfile"
	"math"
)

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *classfile.ClassReader) {
	this.val = math.Float32frombits(reader.ReadUint32())
}
