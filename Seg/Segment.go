package Seg

import (
	"github.com/zhmc/PowerNLP/Seg/Collections"
	"github.com/zhmc/PowerNLP/Seg/SegList/DictSeg"
)

//分词器
type Tokenizer interface {
	Segment(sentence string) []string //分词
}

func NewMapTrieSeg() *Collections.MapTrie {
	return DictSeg.MapTrieSeg
}

//默认分词器
func DefaultSegment() Tokenizer {

	return NewMapTrieSeg()
}
