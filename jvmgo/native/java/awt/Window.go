package awt

import (
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

func _window(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/awt/Window", name, desc, method)
}
