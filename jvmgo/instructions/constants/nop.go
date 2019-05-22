package constants

import (
	"github.com/gottaBoy/jvm.go/jvmgo/instructions/base"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
)

// Do nothing
type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
