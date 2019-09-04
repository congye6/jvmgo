package method_area

// 字段
type ConstantFieldDefInfo struct {
	ConstantMemberDefInfo
	field      *Field //字段直接引用
	fieldClass *Class //属于哪个类的字段
}

// 字段常量的解析，在类加载阶段完成
// 字段属于哪个类直接根据符号引用（即类名）查找，没有多态的动态链接
func (this *ConstantFieldDefInfo) link(originClass *Class) *Field {
	if this.field == nil {
		classInfo := this.constantPool.getConstantInfo(this.classIndex).(*ConstantClassInfo)
		this.fieldClass = classInfo.link(originClass) //先解析字段所在的类
		this.field = this.fieldClass.searchField(this.name)
		// TODO 可访问性检查
	}
	return this.field
}
