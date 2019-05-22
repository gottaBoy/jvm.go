package reserved

import (
	"github.com/gottaBoy/jvm.go/jvmgo/instructions/base"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
)

// Invoke native method
type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method().NativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
