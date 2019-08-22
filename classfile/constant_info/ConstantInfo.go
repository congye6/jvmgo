package constant_info

import "jvmgo/classfile"

// 所有常量数据类型的父类
type ConstantInfo interface {
	readInfo(reader *classfile.ClassReader)
}

