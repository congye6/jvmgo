package constant_pool

import (
	"jvmgo/classfile/reader"
	"math"
)

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *reader.ClassReader) {
	this.val = math.Float32frombits(reader.ReadUint32())
}
