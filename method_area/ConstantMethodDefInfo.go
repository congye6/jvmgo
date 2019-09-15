package method_area

import "fmt"

// 方法
type ConstantMethodDefInfo struct {
	ConstantMemberDefInfo
	methodClass *Class  //方法所在的类
	method      *Method //方法的直接引用
}

// 先根据静态类型（即引用类型）查找，如果可以静态链接就赋值，否则不赋值，等待运行时动态链接
func (this *ConstantMethodDefInfo) staticLink(originClass *Class) *Method {
	classInfo := this.constantPool.getConstantInfo(this.classIndex).(*ConstantClassInfo)
	findClass := classInfo.link(originClass)
	if findClass == nil {
		fmt.Printf("[WARNING] method class:%s not found \n", classInfo.name)
		return nil
	}
	findMethod := findClass.searchMethod(this.name, this.desciptor)

	if findMethod == nil {
		return nil
	}

	if !findMethod.isStaticLink() { // 不可以静态链接
		return nil
	}

	this.method = findMethod
	this.methodClass = findClass
	return this.method
}
