package stack

import (
	"github.com/gottaBoy/jvm.go/jvmgo/instructions/base"
	"github.com/gottaBoy/jvm.go/jvmgo/rtda"
)

// Swap the top two operand stack values
type SWAP struct{ base.NoOperandsInstruction }

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopSlot()
	val2 := stack.PopSlot()
	stack.PushSlot(val1)
	stack.PushSlot(val2)
}
