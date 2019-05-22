package ch

import (
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

func init() {
}

func _ssci(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/nio/ch/ServerSocketChannelImpl", name, desc, method)
}
