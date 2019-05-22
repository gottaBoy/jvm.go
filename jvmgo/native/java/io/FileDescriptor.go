package io

import (
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

func init() {
}

func _fd(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/io/FileDescriptor", name, desc, method)
}
