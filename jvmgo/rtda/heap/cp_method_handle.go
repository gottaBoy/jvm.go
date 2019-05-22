package heap

import (
	cf "github.com/gottaBoy/jvm.go/jvmgo/classfile"
)

type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func newConstantMethodHandle(mhInfo cf.ConstantMethodHandleInfo) *ConstantMethodHandle {
	return &ConstantMethodHandle{
		referenceKind:  mhInfo.ReferenceKind,
		referenceIndex: mhInfo.ReferenceIndex,
	}
}
