package classfile

import "encoding/binary"

// 封装class文件字节流的操作
type ClassReader struct {
	data []byte
}

func (this *ClassReader) ReadUint8() uint8 {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

func (this *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]
	return val
}

func (this *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]
	return val
}

func (this *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]
	return val
}

// 读取n个uint16，长度由第一个uint16决定
func (this *ClassReader) ReadUint16s() []uint16 {
	n := this.ReadUint16()
	result := make([]uint16, n)
	for i := uint16(0); i < n; i++ {
		result[i] = this.ReadUint16()
	}
	return result
}

func (this *ClassReader) ReadBytes(n uint32) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}
