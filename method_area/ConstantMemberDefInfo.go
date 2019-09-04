package method_area

import (
	"jvmgo/constant_pool"
)

// 类描述符和名称+描述符可以唯一确定一个字段或方法
type ConstantMemberDefInfo struct {
	constantPool *ConstantPool
	classIndex   uint16
	name         string
	desciptor    string
}

func (this *ConstantMemberDefInfo) copyInfo(constValue *constant_pool.ConstantMemberDefInfo) {
	this.classIndex = constValue.GetClassIndex()
	//this.classInfo = this.constantPool.getConstantInfo(constValue.GetClassIndex()).(*ConstantClassInfo)
	this.name, this.desciptor = constValue.GetNameAndType()

}
