package main

import (
	"fmt"

	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"unsafe"
)

// https://geektutu.com/post/hpg-struct-alignment.html
func main() {
	var a string
	a = "abcdefghijklmnopabcdefghijklmnop"
	// unsafe.Sizeof 返回的是这个类型所占用的字节数，因此无论 a 字符串有多长，返回都是 16
	//unsafe.Sizeof(wrapperspb.Int32Value{})=48，分为一个空 proto 生成的结构体 40 字节，
	//protoimpl.MessageState{} 占 8 字节，protoimpl.SizeCache 占 4 个字节，这时占 12 个字节
	//protoimpl.UnknownFields{} 占 24 字节，由于 nknownFields{} 实际为 []byte，任何类型的数组按该类型的大小看待，
	//又三个成员最大的是 8 字节，所以按 8 字节对齐，因此 protoimpl.UnknownFields{} 按对齐字节的整数位置开始，16+24=40
	// 同理， int32 占 4 字节，按 8 字节对齐，因此 unsafe.Sizeof(wrapperspb.Int32Value{})=48
	fmt.Println(len(a), unsafe.Sizeof(a), unsafe.Sizeof(wrapperspb.Int32Value{}),
		unsafe.Sizeof(&wrapperspb.Int32Value{}), unsafe.Sizeof(protoimpl.MessageState{}),
		unsafe.Sizeof(protoimpl.UnknownFields{}), unsafe.Alignof(wrapperspb.Int32Value{}))
}
