package bytesutil

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

func StringToBytes(data string) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}
func BytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}
func Uint64ToBytes(n uint64) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}

func Ufloat32ToBytes(n float32) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}
func Ufloat64ToBytes(n float64) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}

func Uint32ToBytes(n uint32) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}

func Uint16ToBytes(n uint16) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}

func Uint8ToBytes(n uint8) []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, n)
	return byteBuffer.Bytes()
}
