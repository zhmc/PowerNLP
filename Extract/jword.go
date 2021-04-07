package Extract
import (
	"fmt"
	"sort"
)


const NKeyword = 10
const d = 0.85f
const max_iter = 200
const min_diff = 0.001f

type TextRankKeyword struct {
}

func NewTextRankKeyword() (rcvr *TextRankKeyword) {
	rcvr = &TextRankKeyword{}
	return
}
func (rcvr *TextRankKeyword) GetKeyword(title string, content string) (string) {
	termList := NLP.segment(fmt.Sprintf("%v%v", title, content))
	wordList := NewArrayList()
	for _, t := range termList {
		if rcvr.ShouldInclude(t) {
			wordList = append(wordList, t.word)
		}
	}
	words := NewHashMap()
	que := NewLinkedList()
	for _, w := range wordList {
		if !words.containsKey(w) {
			words.put(w, NewHashSet())
		}
		que.offer(w)
		if que.size() > 5 {
			que.poll()
		}
		for _, w1 := range que {
			for _, w2 := range que {
				if w1.equals(w2) {
					continue
				}
				words.get(w1).add(w2)
				words.get(w2).add(w1)
			}
		}
	}
	score := NewHashMap()
	for i := 0; i < max_iter; i++ {
		m := NewHashMap()
		max_diff := 0
		for _, entry := range words.entrySet() {
			key := entry.getKey()
			value := entry.getValue()
			m.put(key, 1-d)
			for _, other := range value {
				size := words.get(other).size()
				if key.equals(other) || size == 0 {
					continue
				}
				m.put(key, m.get(key)+d/size*<<unimp_expr[*grammar.JConditionalExpr]>>)
			}
			max_diff = Math.max(max_diff, Math.abs(m.get(key)-<<unimp_expr[*grammar.JConditionalExpr]>>))
		}
		score = m
		if max_diff <= min_diff {
			break
		}
	}
	entryList := NewArrayList(score.entrySet())
	Collections.sort(entryList, NewAnonymous_Comparator_0())
	result := ""
	for i := 0; i < NKeyword; i++ {
		result += entryList[i].getKey() + '#'
	}
	return result
}
// func main() {
// 	content := "程序员(英文Programmer)是从事程序开发、维护的专业人员。一般将程序员分为程序设计人员和程序编码人员，但两者的界限并不非常清楚，特别是在中国。软件从业人员分为初级程序员、高级程序员、系统分析员和项目经理四大类。"
// 	fmt.Println(NewTextRankKeyword().getKeyword("", content))
// }
func (rcvr *TextRankKeyword) ShouldInclude(term *Term) (bool) {
	return rcvr.ShouldInclude(term)
}
