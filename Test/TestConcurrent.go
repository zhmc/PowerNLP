package main

import (
	"fmt"
	"github.com/zhmc/PowerNLP/Seg/SegList/DictSeg"
)

var wait = make(chan string, 10)

//测试字典分词的并发安全性
func main() {
	for i := 0; i < 10; i++ {
		go concurrent()
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-wait)
	}
}

func concurrent() {
	a := DictSeg.MapTrieSeg
	fmt.Println(a)
	wait <- "done"
}
