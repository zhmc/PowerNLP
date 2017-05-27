package main

import (
	"github.com/zhmc/PowerNLP/Seg/Collections"
	"fmt"
)

func main() {
	t := Collections.NewDATrie()
	fmt.Println(t.EndRune)
	fmt.Println(string(t.EndRune))
	fmt.Println(len(t.Base))
	fmt.Println(len(t.RuneCodeMap))
	fmt.Println(t.GetRuneCode('d'))
	fmt.Println(t.Tail[0])

}
