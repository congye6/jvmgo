package method_area

// 类描述符，包括类，超类，接口都用这个表示
type ConstantClassInfo struct {
	name         string // 类名
	linkedClass  *Class // 类直接引用
}

// 类的解析
func (this *ConstantClassInfo) link(originClass *Class) *Class {
	if this.linkedClass == nil {
		classLoader := originClass.loader
		this.linkedClass = classLoader.LoadClass(this.name)
		// TODO 访问权限检查
	}
	return this.linkedClass
}
