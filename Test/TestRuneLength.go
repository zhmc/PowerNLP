package main

import (
	"fmt"
	//"github.com/zhmc/PowerNLP/Seg"
)

func main() {
	length := len([]rune("你好,hello"))
	fmt.Println(length)
	fmt.Println(len("你好,hello"))

	//Seg.DeafaultSegment()
	runes := []rune("我是传奇mike82MIKE")
	for _,v := range runes{
		fmt.Println(int(v))
	}

}
