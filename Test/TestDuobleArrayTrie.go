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
	t.Insert("ab")
	fmt.Println(t.Contain("ab"))
	fmt.Println(t.Contain("abc"))
	t.Insert("黑娃")
	fmt.Println("存在 黑吗",t.Contain("黑"))
	fmt.Println("存在 黑娃吗",t.Contain("黑娃"))
	fmt.Println(t.Contain("黑娃娃"))
	t.Insert("黑大陆")
	fmt.Println("存在 黑娃吗",t.Contain("黑娃"))
	fmt.Println("存在 黑大陆吗",t.Contain("黑大陆"))

}
