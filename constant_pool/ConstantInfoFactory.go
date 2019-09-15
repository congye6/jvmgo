package constant_pool

// 不同常量类型的tag
const (
	INFO_CLASS                = 7
	INFO_FIELD_DEF            = 9
	INFO_METHOD_DEF           = 10
	INFO_INTERFACE_METHOD_DEF = 11
	INFO_STRING               = 8
	INFO_INTEGER              = 3
	INFO_FLOAT                = 4
	INFO_LONG                 = 5
	INFO_DOUBLE               = 6
	INFO_NAME_AND_TYPE        = 12
	INFO_UTF8                 = 1
	INFO_METHOD_HANDLE        = 15
	INFO_METHOD_TYPE          = 16
	INFO_INVOKE_DYNAMIC       = 18
)

// 创建constant info
type ConstantInfoFactory struct {
}

func (this *ConstantInfoFactory) newConstantInfo(tag uint8, pool *ConstantPool) ConstantInfo {
	if tag == INFO_INTEGER {
		return &ConstantIntegerInfo{}
	}

	if tag == INFO_CLASS {
		return &ConstantClassInfo{constantPool: pool}
	}

	if tag == INFO_DOUBLE {
		return &ConstantDoubleInfo{}
	}

	if tag == INFO_FIELD_DEF {
		return &ConstantFieldDefInfo{ConstantMemberDefInfo{constantPool: pool}}
	}

	if tag == INFO_FLOAT {
		return &ConstantFloatInfo{}
	}

	if tag == INFO_INTERFACE_METHOD_DEF {
		return &ConstantInterfaceMethodDefInfo{ConstantMemberDefInfo{constantPool: pool}}
	}

	if tag == INFO_LONG {
		return &ConstantLongInfo{}
	}

	if tag == INFO_METHOD_DEF {
		return &ConstantMethodDefInfo{ConstantMemberDefInfo{constantPool: pool}}
	}

	if tag == INFO_NAME_AND_TYPE {
		return &ConstantNameAndTypeInfo{constantPool: pool}
	}

	if tag == INFO_STRING {
		return &ConstantStringInfo{}
	}

	if tag == INFO_UTF8 {
		return &ConstantUtf8Info{}
	}
	return nil
}
