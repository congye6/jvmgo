package attribute_info

import (
	"jvmgo/classfile/reader"
	"jvmgo/constant_pool"
)

const (
	ATTR_CODE        = "Code"
	ATTR_CONST_VALUE = "ConstantValue"
	ATTR_EXCEPTIONS  = "Exceptions"
	ATTR_SOURCE_FILE = "SourceFile"
)

type AttributeInfo interface {
	readInfo(reader *reader.ClassReader)
	GetName() string
}

func ReadAttributes(reader *reader.ClassReader, pool *constant_pool.ConstantPool) []AttributeInfo {
	attributeCount := reader.ReadUint16()
	attributes := make([]AttributeInfo, attributeCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, pool)
	}
	return attributes
}

func readAttribute(classReader *reader.ClassReader, pool *constant_pool.ConstantPool) AttributeInfo {
	nameIndex := classReader.ReadUint16()
	name := pool.GetUtf8(nameIndex)
	length := classReader.ReadUint32()
	attribute := newAttributeInfo(name, length, pool)
	attribute.readInfo(classReader)
	return attribute
}

func newAttributeInfo(name string, length uint32, pool *constant_pool.ConstantPool) AttributeInfo {
	if name == ATTR_CODE {
		return &CodeAttribute{constantPool: pool}
	}

	if name == ATTR_CONST_VALUE {
		return &ConstValueAttribute{}
	}

	if name == ATTR_EXCEPTIONS {
		return &ExceptionAttribute{}
	}

	if name == ATTR_SOURCE_FILE {
		return &SourceFileAttribute{constantPool: pool}
	}

	return &UnparsedAttribute{name: name, length: length}
}
