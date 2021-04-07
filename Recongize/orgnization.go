package Recongize
import (
	"fmt"
)

var nrf = 0
var ni = 1
var nic = 2
var nis = 3
var nit = 4
var m = 5

type OrganizationRecognition struct {
}

func NewOrganizationRecognition() (rcvr *OrganizationRecognition) {
	rcvr = &OrganizationRecognition{}
	return
}
func Recognition(pWordSegResult *List, wordNetOptimum *WordNet, wordNetAll *WordNet) (bool) {
	roleTagList := OrganizationRecognition.RoleTag(pWordSegResult, wordNetAll)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		for _, NTEnumItem := range roleTagList {
			sbLog.append('[')
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append(' ')
			sbLog.append(NTEnumItem)
			sbLog.append(']')
		}
		fmt.Printf("机构名角色观察：%s\n", sbLog.toString())
	}
	NTList := OrganizationRecognition.ViterbiExCompute(roleTagList)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		sbLog.append('[')
		for _, NT := range NTList {
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append('/')
			sbLog.append(NT)
			sbLog.append(" ,")
		}
		if sbLog.length() > 1 {
			sbLog.delete(sbLog.length()-2, sbLog.length())
		}
		sbLog.append(']')
		fmt.Printf("机构名角色标注：%s\n", sbLog.toString())
	}
	OrganizationDictionary.parsePattern(NTList, pWordSegResult, wordNetOptimum, wordNetAll)
	return true
}
func Recognition2(pWordSegResult *List, wordNetOptimum *WordNet, wordNetAll *WordNet) (bool) {
	roleTagList := OrganizationRecognition.RoleTag(pWordSegResult, wordNetAll)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		for _, NTEnumItem := range roleTagList {
			sbLog.append('[')
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append(' ')
			sbLog.append(NTEnumItem)
			sbLog.append(']')
		}
		fmt.Printf("机构名角色观察：%s\n", sbLog.toString())
	}
	NTList := OrganizationRecognition.ViterbiExCompute(roleTagList)
	if NLP.Config.DEBUG {
		sbLog := ""
		iterator := pWordSegResult.iterator()
		sbLog.append('[')
		for _, NT := range NTList {
			sbLog.append(<<unimp_obj.nm_*parser.GoMethodAccessVar>>)
			sbLog.append('/')
			sbLog.append(NT)
			sbLog.append(" ,")
		}
		if sbLog.length() > 1 {
			sbLog.delete(sbLog.length()-2, sbLog.length())
		}
		sbLog.append(']')
		fmt.Printf("机构名角色标注：%s\n", sbLog.toString())
	}
	OrganizationDictionary.parsePattern(NTList, pWordSegResult, wordNetOptimum, wordNetAll)
	return true
}
func RoleTag(vertexList *List, wordNetAll *WordNet) (*List) {
	var tagList []string
	for _, vertex := range vertexList {
		nature := vertex.guessNature()
		switch nature {
		case nrf:
			{
				if <<unimp_obj.nm_*parser.GoMethodAccessVar>> <= 1000 {
					tagList = append(tagList, NewEnumItem(NT.F, 1000))
				} else {
					break
				}
			}
			continue
			fallthrough
		case ni:
			{
				break
			}
			fallthrough
		case nic:
			{
				break
			}
			fallthrough
		case nis:
			{
				break
			}
			fallthrough
		case nit:
			{
				ntEnumItem := NewEnumItem(NT.K, 1000)
				ntEnumItem.addLabel(NT.D, 1000)
				tagList = append(tagList, ntEnumItem)
			}
			continue
			fallthrough
		case m:
			{
				ntEnumItem := NewEnumItem(NT.M, 1000)
				tagList = append(tagList, ntEnumItem)
			}
			continue
			fallthrough
		}
		NTEnumItem := OrganizationDictionary.dictionary.get(vertex.word)
		if NTEnumItem == nil {
			NTEnumItem = NewEnumItem(NT.Z, OrganizationDictionary.transformMatrixDictionary.getTotalFrequency(NT.Z))
		}
		tagList = append(tagList, NTEnumItem)
	}
	return tagList
}
func ViterbiExCompute(roleTagList *List) (*List) {
	return Viterbi.computeEnum(roleTagList, OrganizationDictionary.transformMatrixDictionary)
}
