package constant_pool

import (
	"jvmgo/classfile/reader"
	"math"
)

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *reader.ClassReader) {
	this.val = math.Float64frombits(reader.ReadUint64())
}
