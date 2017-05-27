package main

import (
	"github.com/zhmc/PowerNLP/Seg/Collections"
	"fmt"
)

func main() {
	t := Collections.NewDATrie()
	fmt.Println(Collections.EndRune)
	fmt.Println(string(Collections.EndRune))
	fmt.Println(len(t.Base))
	fmt.Println(len(t.RuneCodeMap))
	fmt.Println(t.GetRuneCode('d'))
	t.AppendToTailArray([]rune("hello"),3)
	fmt.Println(string(t.Tail[0]))
}
