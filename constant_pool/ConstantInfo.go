package constant_pool

import (
	"jvmgo/classfile/reader"
)

// 所有常量数据类型的父类
type ConstantInfo interface {
	readInfo(reader *reader.ClassReader)
}

