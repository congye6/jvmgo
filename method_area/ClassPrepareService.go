package method_area

import "jvmgo/slot"

func prepare(class *Class) {
	calcInsatnceFieldId(class)
	calcStaticFieldId(class)
	initStaticVars(class)
}

// 为每个实例字段分配slotId，并统计slot的个数
// 父类的slot也需要统计
// 静态变量用另外的Slots
func calcInsatnceFieldId(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instancesSlotCount //父类的字段已占用instancesSlotCount个slot
	}
	for _, field := range class.fields {
		if field.isStatic() { //静态变量用另外一个slots
			continue
		}
		field.slotId = slotId
		slotId++
		if field.isLongOrDouble() {
			slotId++
		}
	}
	class.instancesSlotCount = slotId
}

// 为每个静态变量分配slotId，并统计slot个数
func calcStaticFieldId(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.staticSlotCount
	}
	for _, field := range class.fields {
		if !field.isStatic() {
			continue
		}
		field.slotId = slotId
		slotId++
		if field.isLongOrDouble() {
			slotId++
		}
	}
	class.staticSlotCount = slotId
}

func initStaticVars(class *Class) {
	class.staticVars = slot.NewSlots(class.staticSlotCount) // 因为go语言本身就会给slot分配默认值，所以默认值不需要再分配
	for _, field := range class.fields {
		if field.isStatic() && field.isFinal() { // static final的变量会有初始值，分配在const value常量中
			initStaticFinalVars(class, field)
		}
	}
}

func initStaticFinalVars(class *Class, field *Field) {
	staticVars := class.staticVars
	poolIndex := field.constValueIndex
	constPool := class.constantPool
	descriptor := field.descriptor
	slotId := field.slotId

	if descriptor == "Z" || descriptor == "B" || descriptor == "C" || descriptor == "S" || descriptor == "I" {
		staticVars.SetInt(slotId, constPool.GetInteger(poolIndex))
		return
	}

	if descriptor == "J" {
		staticVars.SetLong(slotId, constPool.GetLong(poolIndex))
		return
	}

	if descriptor == "D" {
		staticVars.SetDouble(slotId, constPool.GetDouble(poolIndex))
		return
	}

	if descriptor == "Ljava/lang/String" {
		// TODO 常量池暂不支持string
	}

}
