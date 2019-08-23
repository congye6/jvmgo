package constant_info

import (
	"jvmgo/classfile"
	"math"
)

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *classfile.ClassReader) {
	this.val = math.Float64frombits(reader.ReadUint64())
}
