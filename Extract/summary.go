package Extract
import (
	"fmt"
	"sort"
)

// 阻尼系数（ＤａｍｐｉｎｇＦａｃｔｏｒ），一般取值为0.85
const d = 0.85f
const max_iter = 200
const min_diff = 0.001f

type TextRankSummary struct {
	D          int   //文档句子的个数
	docs       *List //拆分为[句子[单词]]形式的文档
	top        *TreeMap //排序后的最终结果 score <-> index
	weight     []float64 //句子和其他句子的相关程度
	weight_sum []float64 //该句子和其他句子相关程度之和
	vertex     []float64 //迭代之后收敛的权重
	bm25       *BM25 //BM25相似度
}

func NewTextRankSummary(docs *List) (rcvr *TextRankSummary) {
	rcvr = &TextRankSummary{}
	rcvr.docs = docs
	rcvr.bm25 = NewBM25(docs)
	rcvr.D = len(docs)
	rcvr.weight = make([]float64, rcvr.D, rcvr.D)
	rcvr.weight_sum = make([]float64, rcvr.D)
	rcvr.vertex = make([]float64, rcvr.D)
	rcvr.top = NewTreeMap(Collections.reverseOrder())
	rcvr.solve()
	return
}
func NewTextRankSummary2(docs *List) (rcvr *TextRankSummary) {
	rcvr = &TextRankSummary{}
	rcvr.docs = docs
	rcvr.bm25 = NewBM25(docs)
	rcvr.D = len(docs)
	rcvr.weight = make([]float64, rcvr.D, rcvr.D)
	rcvr.weight_sum = make([]float64, rcvr.D)
	rcvr.vertex = make([]float64, rcvr.D)
	rcvr.top = NewTreeMap(Collections.reverseOrder())
	rcvr.solve()
	return
}
func (rcvr *TextRankSummary) GetTopSentence(size int) (<<array>>) {
	values := rcvr.top.values()
	size = Math.min(size, values.size())
	indexArray := make([]int, size)
	it := values.iterator()
	for i := 0; i < size; i++ {
		indexArray[i] = it.next()
	}
	return indexArray
}
func (rcvr *TextRankSummary) GetTopSentence2(size int) (<<array>>) {
	values := rcvr.top.values()
	size = Math.min(size, values.size())
	indexArray := make([]int, size)
	it := values.iterator()
	for i := 0; i < size; i++ {
		indexArray[i] = it.next()
	}
	return indexArray
}
func GetTopSentenceList(document string, size int) (*List) {
	sentenceList := TextRankSummary.spiltSentence(document)
	var docs []doc
	for _, sentence := range sentenceList {
		termList := NLP.segment(sentence)
		var wordList []doc
		for _, term := range termList {
			if TextRankSummary.ShouldInclude(term) {
				wordList = append(wordList, term.word)
			}
		}
		docs = append(docs, wordList)
	}
	textRankSummary := NewTextRankSummary(docs)
	topSentence := textRankSummary.getTopSentence(size)
	var resultList []doc
	for _, i := range topSentence {
		resultList = append(resultList, sentenceList[i])
	}
	return resultList
}
// func main() {
// 	document := "算法可大致分为基本算法、数据结构的算法、数论算法、计算几何的算法、图的算法、动态规划以及数值分析、加密算法、排序算法、检索算法、随机化算法、并行算法、厄米变形模型、随机森林算法。\n算法可以宽泛的分为三类，\n一，有限的确定性算法，这类算法在有限的一段时间内终止。他们可能要花很长时间来执行指定的任务，但仍将在一定的时间内终止。这类算法得出的结果常取决于输入值。\n二，有限的非确定算法，这类算法在有限的时间内终止。然而，对于一个（或一些）给定的数值，算法的结果并不是唯一的或确定的。\n三，无限的算法，是那些由于没有定义终止定义条件，或定义的条件无法由输入的数据满足而不终止运行的算法。通常，无限算法的产生是由于未能确定的定义终止条件。"
// 	fmt.Println(TextRankSummary.GetTopSentenceList(document, 3))
// }
func ShouldInclude(term *Term) (bool) {
	return TextRankSummary.ShouldInclude(term)
}
func (rcvr *TextRankSummary) solve() {
	cnt := 0
	for _, sentence := range rcvr.docs {
		scores := rcvr.bm25.simAll(sentence)
		weight[cnt] = scores
		weight_sum[cnt] = TextRankSummary.sum(scores) - scores[cnt]
		vertex[cnt] = 1.0
		cnt++
	}
	for _ := 0; _ < rcvr.max_iter; _++ {
		m := make([]float64, rcvr.D)
		max_diff := 0
		for i := 0; i < rcvr.D; i++ {
			m[i] = 1 - rcvr.d
			for j := 0; j < rcvr.D; j++ {
				if j == i || weight_sum[j] == 0 {
					continue
				}
				m[i] += rcvr.d * weight[j][i] / weight_sum[j] * vertex[j]
			}
			diff := Math.abs(m[i] - vertex[i])
			if diff > max_diff {
				max_diff = diff
			}
		}
		rcvr.vertex = m
		if max_diff <= rcvr.min_diff {
			break
		}
	}
	for i := 0; i < rcvr.D; i++ {
		rcvr.top.put(vertex[i], i)
	}
}
func spiltSentence(document string) (*List) {
	var sentences []string
	if document == nil {
		return sentences
	}
	for _, line := range document.split("[\r\n]") {
		line = line.trim()
		if line.length() == 0 {
			continue
		}
		for _, sent := range line.split("[，,。:：“”？?！!；;]") {
			sent = sent.trim()
			if sent.length() == 0 {
				continue
			}
			sentences = append(sentences, sent)
		}
	}
	return sentences
}
func spiltSentence2(document string) (*List) {
	var sentences []string
	if document == nil {
		return sentences
	}
	for _, line := range document.split("[\r\n]") {
		line = line.trim()
		if line.length() == 0 {
			continue
		}
		for _, sent := range line.split("[，,。:：“”？?！!；;]") {
			sent = sent.trim()
			if sent.length() == 0 {
				continue
			}
			sentences = append(sentences, sent)
		}
	}
	return sentences
}
func sum(array []float64) (float64) {
	total := 0
	for _, v := range array {
		total += v
	}
	return total
}
