package classfile

import (
	"jvmgo/classfile/attribute_info"
	"jvmgo/classfile/reader"
	"jvmgo/constant_pool"
)

type MemberInfo struct {
	constantPool    *constant_pool.ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []attribute_info.AttributeInfo
}

type FieldInfo struct {
	MemberInfo
}

type MethodInfo struct {
	MemberInfo
}

func (this *MemberInfo) GetName() string {
	return this.constantPool.GetUtf8(this.nameIndex)
}

func (this *MemberInfo) GetDescriptor() string {
	return this.constantPool.GetUtf8(this.descriptorIndex)
}

func (this *MemberInfo) GetAttributes() []attribute_info.AttributeInfo {
	return this.attributes
}

func (this *MemberInfo) GetAccessFlag() uint16 {
	return this.accessFlags
}

func readFields(reader *reader.ClassReader, pool *constant_pool.ConstantPool) []*FieldInfo {
	fieldCount := reader.ReadUint16()
	fields := make([]*FieldInfo, fieldCount)
	for i := range fields {
		fields[i] = &FieldInfo{*readMember(reader, pool)}
	}
	return fields
}

func (this *FieldInfo) GetConstAttribute() *attribute_info.ConstValueAttribute {
	for _, attribute := range this.attributes {
		if attribute.GetName() == attribute_info.ATTR_CONST_VALUE {
			return attribute.(*attribute_info.ConstValueAttribute)
		}
	}
	return nil
}

func readMethods(reader *reader.ClassReader, pool *constant_pool.ConstantPool) []*MethodInfo {
	methodCount := reader.ReadUint16()
	methods := make([]*MethodInfo, methodCount)
	for i := range methods {
		methods[i] = &MethodInfo{*readMember(reader, pool)}
	}
	return methods
}

func readMember(reader *reader.ClassReader, pool *constant_pool.ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    pool,
		accessFlags:     reader.ReadUint16(),
		nameIndex:       reader.ReadUint16(),
		descriptorIndex: reader.ReadUint16(),
		attributes:      attribute_info.ReadAttributes(reader, pool),
	}
}

func (this *MethodInfo) GetCodeAttribute() *attribute_info.CodeAttribute {
	for _, attribute := range this.attributes {
		if attribute.GetName() == attribute_info.ATTR_CODE {
			return attribute.(*attribute_info.CodeAttribute)
		}
	}
	return nil
}
