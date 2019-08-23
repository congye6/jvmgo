package reader

import "encoding/binary"

// 封装class文件字节流的操作
type ClassReader struct {
	Data []byte
}

func (this *ClassReader) ReadUint8() uint8 {
	val := this.Data[0]
	this.Data = this.Data[1:]
	return val
}

func (this *ClassReader) ReadUint16() uint16 {
	val := binary.BigEndian.Uint16(this.Data)
	this.Data = this.Data[2:]
	return val
}

func (this *ClassReader) ReadUint32() uint32 {
	val := binary.BigEndian.Uint32(this.Data)
	this.Data = this.Data[4:]
	return val
}

func (this *ClassReader) ReadUint64() uint64 {
	val := binary.BigEndian.Uint64(this.Data)
	this.Data = this.Data[8:]
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
	bytes := this.Data[:n]
	this.Data = this.Data[n:]
	return bytes
}
