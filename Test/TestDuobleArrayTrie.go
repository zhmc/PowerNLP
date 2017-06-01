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

	//fmt.Println("插入 ab")
	//t.Insert("ab")
	//fmt.Println("存在 a?", t.Contains("a"))
	//fmt.Println("存在 ab?", t.Contains("ab"))
	//fmt.Println("存在 abc?", t.Contains("abc"))

	//fmt.Println("插入 a")
	//t.Insert("a")
	//fmt.Println("存在 a?", t.Contains("a"))
	//fmt.Println("存在 ab?", t.Contains("ab"))
	//fmt.Println("存在 abc?", t.Contains("abc"))

	//fmt.Println("插入 黑娃")
	//t.Insert("黑娃")
	//fmt.Println("存在 黑?", t.Contains("黑"))
	//fmt.Println("存在 黑娃?", t.Contains("黑娃"))
	//fmt.Println("存在 黑娃娃?", t.Contains("黑娃娃"))

	//fmt.Println("插入 黑娃娃")
	//t.Insert("黑娃娃")
	//fmt.Println("存在 黑?", t.Contains("黑"))
	//fmt.Println("存在 黑娃?", t.Contains("黑娃"))
	//fmt.Println("存在 黑娃娃?", t.Contains("黑娃娃"))
	//fmt.Println("存在 黑娃娃娃?", t.Contains("黑娃娃娃"))

	//fmt.Println("插入 黑大陆")
	//t.Insert("黑大陆")
	//fmt.Println("存在 黑?", t.Contains("黑"))
	//fmt.Println("存在 黑娃?", t.Contains("黑娃"))
	//fmt.Println("存在 黑娃娃?", t.Contains("黑娃娃"))
	//fmt.Println("存在 黑大?", t.Contains("黑大"))
	//fmt.Println("存在 黑大陆?", t.Contains("黑大陆"))
	//fmt.Println("存在 黑大陆陆?", t.Contains("黑大陆陆"))


	//fmt.Println("插入 黑")
	//t.Insert("黑")
	//fmt.Println("存在 黑？", t.Contains("黑"))
	//fmt.Println("存在 黑娃？", t.Contains("黑娃"))
	//fmt.Println("存在 黑娃娃?", t.Contains("黑娃娃"))
	//fmt.Println("存在 黑娃娃娃?", t.Contains("黑娃娃娃"))
	//fmt.Println("存在 黑大?", t.Contains("黑大"))
	//fmt.Println("存在 黑大陆?", t.Contains("黑大陆"))
	//fmt.Println("存在 黑色?", t.Contains("黑色"))


	fmt.Println("插入 bachelor")
	t.Insert("bachelor")
	fmt.Println("插入 jar")
	t.Insert("jar")
	fmt.Println("插入 badge")
	t.Insert("badge")
	fmt.Println("插入 baby")
	t.Insert("baby")
	fmt.Println("存在 bachelor?", t.Contains("bachelor"))
	fmt.Println("存在 jar?", t.Contains("jar"))
	fmt.Println("存在 badge?", t.Contains("badge"))
	fmt.Println("存在 baby?", t.Contains("baby"))

}
