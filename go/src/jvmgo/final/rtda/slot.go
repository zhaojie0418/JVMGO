package rtda

import "jvmgo/final/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
