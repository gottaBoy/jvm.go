package awt

import (
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

func init() {
}

func _tk(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Toolkit", name, desc, method)
}
