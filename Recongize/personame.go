package Recongize
import (
	"fmt"
)

var nr = 0
var nnt = 1
var nic = 2
var nis = 3
var nit = 4
var m = 5


type PersonRecognition struct {
}

func NewPersonRecognition() (rcvr *PersonRecognition) {
	rcvr = &PersonRecognition{}
	return
}
func Recognition(pWordSegResult *List, wordNetOptimum *WordNet, wordNetAll *WordNet) (bool) {
	roleTagList := PersonRecognition.RoleObserve(pWordSegResult)
	if NLP.Config.DEBUG {
		sbLog := NewStringBuilder()
		iterator := pWordSegResult.iterator()
		for _, nrEnumItem := range roleTagList {
			sbLog.append('[')
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append(' ')
			sbLog.append(nrEnumItem)
			sbLog.append(']')
		}
		fmt.Printf("人名角色观察：%s\n", sbLog.toString())
	}
	nrList := PersonRecognition.ViterbiComputeSimply(roleTagList)
	if NLP.Config.DEBUG {
		sbLog := NewStringBuilder()
		iterator := pWordSegResult.iterator()
		sbLog.append('[')
		for _, nr := range nrList {
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append('/')
			sbLog.append(nr)
			sbLog.append(" ,")
		}
		if sbLog.length() > 1 {
			sbLog.delete(sbLog.length()-2, sbLog.length())
		}
		sbLog.append(']')
		fmt.Printf("人名角色标注：%s\n", sbLog.toString())
	}
	PersonDictionary.parsePattern(nrList, pWordSegResult, wordNetOptimum, wordNetAll)
	return true
}
func Recognition2(segResult *List, wordNetOptimum *WordNet, wordNetAll *WordNet) {
	sbName := NewStringBuilder()
	appendTimes := 0
	listIterator := segResult.listIterator()
	listIterator.next()
	line := 1
	activeLine := 1
	for listIterator.hasNext() {
		vertex := listIterator.next()
		if appendTimes > 0 {
			if vertex.guessNature() == Nature.nrf || TranslatedPersonDictionary.containsKey(vertex.realWord) {
				sbName.append(vertex.realWord)
				appendTimes++
			} else {
				if appendTimes > 1 {
					if NLP.Config.DEBUG {
						fmt.Println(fmt.Sprintf("%v%v", "音译人名识别出：", sbName.toString()))
					}
					wordNetOptimum.insert(activeLine, NewVertex(Predefine.TAG_PEOPLE, sbName.toString(), NewCoreDictionary.Attribute(Nature.nrf), WORD_ID), wordNetAll)
				}
				sbName.setLength(0)
				appendTimes = 0
			}
		} else {
			if vertex.guessNature() == Nature.nrf {
				sbName.append(vertex.realWord)
				appendTimes++
				activeLine = line
			}
		}
		line += vertex.realWord.length()
	}
}
func RoleObserve(wordSegResult *List) (*List) {
	tagList := NewLinkedList()
	iterator := wordSegResult.iterator()
	iterator.next()
	tagList = append(tagList, NewEnumItem(NR.A, NR.K))
	for iterator.hasNext() {
		vertex := iterator.next()
		nrEnumItem := PersonDictionary.dictionary.get(vertex.realWord)
		if nrEnumItem == nil {
			switch vertex.guessNature() {
			case nr:
				{
					if <<unimp_obj.nm_*parser.GoMethodAccessVar>> <= 1000 && vertex.realWord.length() == 2 {
						nrEnumItem = NewEnumItem(NR.X, NR.G)
					} else {
						nrEnumItem = NewEnumItem(NR.A, PersonDictionary.transformMatrixDictionary.getTotalFrequency(NR.A))
					}
				}
			case nnt:
				{
					nrEnumItem = NewEnumItem(NR.G, NR.K)
				}
			default:
				{
					nrEnumItem = NewEnumItem(NR.A, PersonDictionary.transformMatrixDictionary.getTotalFrequency(NR.A))
				}
			}
		}
		tagList = append(tagList, nrEnumItem)
	}
	return tagList
}
func ViterbiCompute(roleTagList *List) (*List) {
	return Viterbi.computeEnum(roleTagList, PersonDictionary.transformMatrixDictionary)
}
func ViterbiComputeSimply(roleTagList *List) (*List) {
	return Viterbi.computeEnumSimply(roleTagList, PersonDictionary.transformMatrixDictionary)
}
