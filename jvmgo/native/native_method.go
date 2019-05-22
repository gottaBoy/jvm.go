package native

import (
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
	//"github.com/gottaBoy/jvm.go/jvmgo/rtda/heap"
)

type NativeMethod func(frame *rtda.Frame)
