package PowerNLP

import "github.com/zhmc/PowerNLP/Seg"

//默认分词方法
func Segment(sentence string) []string {
	return Seg.DeafaultSegment().Segment(sentence)
}
