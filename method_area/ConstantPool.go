package method_area

import (
	"jvmgo/constant_pool"
)

type ConstantPool struct {
	constantValues *constant_pool.ConstantPool
	constants      []ConstantInfo
}

func NewConstantPool(constantValues *constant_pool.ConstantPool) *ConstantPool {
	return &ConstantPool{
		constantValues: constantValues,
	}
}

func (this *ConstantPool) Init() {
	pool := this.constantValues.GetPool()
	count := len(pool)
	this.constants = make([]ConstantInfo, count)
	for i := uint16(1); i < uint16(count); i++ {
		constValue := pool[i]
		switch constValue.(type) {
		case *constant_pool.ConstantIntegerInfo:
			this.constants[i] = &ConstantIntegerInfo{val: this.constantValues.GetInteger(i)}
		case *constant_pool.ConstantFloatInfo:
			this.constants[i] = &ConstantFloatInfo{val: this.constantValues.GetFloat(i)}
		case *constant_pool.ConstantLongInfo:
			this.constants[i] = &ConstantLongInfo{val: this.constantValues.GetLong(i)}
			i++
		case *constant_pool.ConstantDoubleInfo:
			this.constants[i] = &ConstantDoubleInfo{val: this.constantValues.GetDouble(i)}
			i++
		case *constant_pool.ConstantClassInfo:
			this.constants[i] = &ConstantClassInfo{name: this.constantValues.GetClassName(i)}
		case *constant_pool.ConstantFieldDefInfo:
			fieldInfo := &ConstantFieldDefInfo{}
			fieldInfo.constantPool = this
			fieldInfo.copyInfo(&constValue.(*constant_pool.ConstantFieldDefInfo).ConstantMemberDefInfo)
			this.constants[i] = fieldInfo
		case *constant_pool.ConstantMethodDefInfo:
			methodInfo := &ConstantMethodDefInfo{}
			methodInfo.constantPool = this
			methodInfo.copyInfo(&constValue.(*constant_pool.ConstantMethodDefInfo).ConstantMemberDefInfo)
			this.constants[i] = methodInfo
		case *constant_pool.ConstantInterfaceMethodDefInfo:
			interfaceMethod := &ConstantInterfaceMethodDefInfo{}
			interfaceMethod.copyInfo(&constValue.(*constant_pool.ConstantInterfaceMethodDefInfo).ConstantMemberDefInfo)
		}
	}
}

func (this *ConstantPool) GetConstantValues() *constant_pool.ConstantPool {
	return this.constantValues
}

func (this *ConstantPool) GetConstantInfo(index uint16) ConstantInfo {
	if constantInfo := this.constants[index]; constantInfo != nil {
		return constantInfo
	}
	panic("invalid constant pool index")
}

// 读取字符串字面量
func (this *ConstantPool) GetUtf8(index uint16) string {
	utf8Info := this.GetConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (this *ConstantPool) GetInteger(index uint16) int32 {
	intInfo := this.GetConstantInfo(index).(*ConstantIntegerInfo)
	return intInfo.val
}

func (this *ConstantPool) GetLong(index uint16) int64 {
	longInfo := this.GetConstantInfo(index).(*ConstantLongInfo)
	return longInfo.val
}

func (this *ConstantPool) GetDouble(index uint16) float64 {
	doubleInfo := this.GetConstantInfo(index).(*ConstantDoubleInfo)
	return doubleInfo.val
}

func (this *ConstantPool) GetClass(index uint16) *Class {
	classInfo := this.GetConstantInfo(index).(*ConstantClassInfo)
	return classInfo.linkedClass
}

// 获取名称和描述符
func (this *ConstantPool) GetNameAndType(index uint16) (string, string) {
	nameTypeInfo := this.constants[index].(*ConstantNameAndTypeInfo)
	return nameTypeInfo.name, nameTypeInfo.descriptor
}

// 获取类描述符
func (this *ConstantPool) GetClassName(index uint16) string {
	classInfo := this.constants[index].(*ConstantClassInfo)
	return classInfo.name
}
