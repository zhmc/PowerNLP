package Recongize
import (
	"fmt"
)

type PlaceRecognition struct {
}

func NewPlaceRecognition() (rcvr *PlaceRecognition) {
	rcvr = &PlaceRecognition{}
	return
}
func Recognition(pWordSegResult *List, wordNetOptimum *WordNet, wordNetAll *WordNet) (bool) {
	roleTagList := PlaceRecognition.RoleTag(pWordSegResult, wordNetAll)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		for _, NSEnumItem := range roleTagList {
			sbLog.append('[')
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append(' ')
			sbLog.append(NSEnumItem)
			sbLog.append(']')
		}
		fmt.Printf("地名角色观察：%s\n", sbLog.toString())
	}
	NSList := PlaceRecognition.ViterbiExCompute(roleTagList)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		sbLog.append('[')
		for _, NS := range NSList {
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append('/')
			sbLog.append(NS)
			sbLog.append(" ,")
		}
		if sbLog.length() > 1 {
			sbLog.delete(sbLog.length()-2, sbLog.length())
		}
		sbLog.append(']')
		fmt.Printf("地名角色标注：%s\n", sbLog.toString())
	}
	PlaceDictionary.parsePattern(NSList, pWordSegResult, wordNetOptimum, wordNetAll)
	return true
}
func insert(listIterator *ListIterator, tagList *List, wordNetAll *WordNet, line int, ns *NS) {
	vertex := wordNetAll.getFirst(line)
	listIterator.add(vertex)
	tagList = append(tagList, NewEnumItem(ns, 1000))
}
func RoleTag(vertexList *List, wordNetAll *WordNet) (*List) {
	var tagList []string
	listIterator := vertexList.listIterator()
	for listIterator.hasNext() {
		vertex := listIterator.next()
		if Nature.ns == vertex.getNature() && <<unimp_obj.nm_*parser.GoMethodAccessVar>> <= 1000 {
			if vertex.realWord.length() < 3 {
				tagList = append(tagList, NewEnumItem(NS.H, NS.G))
			} else {
				tagList = append(tagList, NewEnumItem(NS.G))
			}
			continue
		}
		var NSEn *EnumItem
		umItem = PlaceDictionary.dictionary.get(vertex.word)
		if NSEnumItem == nil {
			NSEnumItem = NewEnumItem(NS.Z, PlaceDictionary.transformMatrixDictionary.getTotalFrequency(NS.Z))
		}
		tagList = append(tagList, NSEnumItem)
	}
	return tagList
}
func ViterbiExCompute(roleTagList *List) (*List) {
	return Viterbi.computeEnum(roleTagList, PlaceDictionary.transformMatrixDictionary)
}
