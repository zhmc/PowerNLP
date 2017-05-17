package main

import (
	"fmt"
	"github.com/zhmc/PowerNLP/Seg"
)

func main() {
	length := len([]rune("你好,hello"))
	fmt.Println(length)
	fmt.Println(len("你好,hello"))

	Seg.DeafaultSegment()
}
