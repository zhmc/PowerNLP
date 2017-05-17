package main

import (
	"fmt"
	"github.com/zhmc/PowerNLP/Seg"
	BaseTrie "github.com/zhmc/PowerNLP/Seg/Collections"
)

func main() {
	tree := BaseTrie.NewMapTrie()
	tree.Insert("word")
	tree.Insert("wor")
	tree.Insert("wx")
	tree.Insert("abastract")
	tree.Insert("中国人")
	tree.Insert("国足")
	//tree.Insert("中国")
	//Trie.PreTraverse(tree.Root)
	//fmt.Println(tree.CountPrefix("wordf"))
	//fmt.Println(tree.Has("word"))
	//fmt.Println(tree.CountPrefix("wor"))
	//fmt.Println(tree.CountPrefix("wo"))
	//fmt.Println(tree.CountPrefix("w"))
	//fmt.Println(tree.CountPrefix("ab"))
	//fmt.Println(tree.CountPrefix("中"))
	//fmt.Println(tree.Has("ab"))
	//fmt.Println(tree.Has("中国人"))
	//fmt.Println(tree.Has("中国"))

	r := tree.Segment("大中国word中国人wxabwo大中国足abastract")
	for _, v := range r {
		fmt.Println(v)
	}
	Seg.DeafaultSegment().Segment("")
}
