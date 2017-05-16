package main

import (
	"fmt"
	BaseTrie "github.com/PowerNLP/BaseTrie"
)

func main() {
	tree := BaseTrie.NewBaseTrie()
	tree.Insert("word")
	tree.Insert("wor")
	tree.Insert("wx")
	tree.Insert("abastract")
	tree.Insert("中国人")
	//tree.Insert("中国")
	//BaseTrie.PreTraverse(tree.Root)
	fmt.Println(tree.CountPrefix("wordf"))
	fmt.Println(tree.Has("word"))
	fmt.Println(tree.CountPrefix("wor"))
	fmt.Println(tree.CountPrefix("wo"))
	fmt.Println(tree.CountPrefix("w"))
	fmt.Println(tree.CountPrefix("ab"))
	fmt.Println(tree.CountPrefix("中"))
	fmt.Println(tree.Has("ab"))
	fmt.Println(tree.Has("中国人"))
	fmt.Println(tree.Has("中国"))

	r := tree.Segment("大中国word中国人wxabwoabastract")
	for _, v := range r {
		fmt.Println(v)
	}

}
