package PowerNLP

import "fmt"

//树节点 用Hash表存储<Character, Node>
type TrieNode struct {
	Num       int                //经过这个节点的单词数量
	Sons      map[rune]*TrieNode //所有子节点
	IsEnd     bool               //是否是某个单词的终点
	Character rune               //存储的值
}

func newTrieNode() *TrieNode {
	node := &TrieNode{}
	node.Num = 1
	node.Sons = make(map[rune]*TrieNode)
	node.IsEnd = false

	return node
}

type BaseTrie struct {
	Root *TrieNode
}

func NewBaseTrie() *BaseTrie {
	newtrie := &BaseTrie{}
	newtrie.Root = newTrieNode()
	return newtrie
}

//插入一个单词
func (this *BaseTrie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	//if this.Has(word) {
	//	return
	//}
	tmpNode := this.Root
	letters := []rune(word)
	for _, letter := range letters {

		//如果没有这个字母
		if _, ok := tmpNode.Sons[letter]; !ok {
			tmpNode.Sons[letter] = newTrieNode()
			tmpNode.Sons[letter].Character = letter
		} else {
			tmpNode.Sons[letter].Num++

		}
		tmpNode = tmpNode.Sons[letter]
	}
	tmpNode.IsEnd = true
}

//计算包含某个前缀的单词数量
func (this *BaseTrie) CountPrefix(prefix string) int {
	if len(prefix) == 0 {
		return -1
	}
	tmpNode := this.Root
	letters := []rune(prefix)
	for _, letter := range letters {

		//如果没有这个字母
		if _, ok := tmpNode.Sons[letter]; !ok {
			return 0
		} else {
			tmpNode = tmpNode.Sons[letter]
		}
	}
	return tmpNode.Num
}

//前序遍历节点
func PreTraverse(node *TrieNode) {
	if node != nil {
		fmt.Println(string(node.Character))
		for _, son := range node.Sons {
			PreTraverse(son)
		}
	}
}

func (this *BaseTrie) PreTraverseByPrefix(prefix string) {
	node := this.FindPrefix(prefix)
	if node == nil {
		return
	}
	PreTraverse(node)
}

//是否包含某个单词
func (this *BaseTrie) Has(word string) bool {
	if len(word) == 0 {
		return false
	}
	tmpNode := this.Root
	letters := []rune(word)
	for _, letter := range letters {

		//如果没有这个字母
		if _, ok := tmpNode.Sons[letter]; !ok {
			return false
		} else {
			tmpNode = tmpNode.Sons[letter]
		}
	}
	if tmpNode.IsEnd == true {
		return true
	}
	return false
}

//找到某个前缀的最终节点
func (this *BaseTrie) FindPrefix(prefix string) *TrieNode {
	if len(prefix) == 0 {
		return nil
	}
	tmpNode := this.Root
	letters := []rune(prefix)
	for _, letter := range letters {

		//如果没有这个字母
		if _, ok := tmpNode.Sons[letter]; !ok {
			return nil
		} else {
			tmpNode = tmpNode.Sons[letter]
		}
	}

	return tmpNode

}

//分词
func (this *BaseTrie) Segment(setence string) []string {
	chars := []rune(setence)
	results := make([]string, 0)
	//tmpChar := ""
	tmpHead := 0         //扫描全部字符过程中位于单词头部的指针位置
	tmpNode := this.Root //在词典里顺延的节点指针
	length := len(chars)
	for index := 0; index < length; index++ {

		//如果有这个字符，就继续顺下去（要求扫描指针未到尾部）
		// 防止尾部的单词是个完整词，同时也是一个另一个词的前缀

		if _, ok := tmpNode.Sons[chars[index]]; ok {
			tmpNode = tmpNode.Sons[chars[index]]
			if index != length-1 {
				continue
			} else {
				// 这个地方后移是为了接下来截取单词能正确，因
				// 为如果发现了匹配的单词，用到列表切片
				// 用了slice[start:end]。如果不是因为是最后一个，
				// index所在的位置都是单词之后

				index++
			}
		}

		//匹配到了单词（词典里面没有的字，或者是词典里面的词匹配完成了）

		//第一个字词典里面就没有
		if tmpHead == index {
			word := string(chars[tmpHead])

			results = append(results, word)

			//将头指针移到下一个字符
			tmpHead = index + 1
			// 将指针节点移到树根
			tmpNode = this.Root
		} else if tmpNode.IsEnd { //词典里面的词结束了
			word := string(chars[tmpHead:index])

			results = append(results, word)

			//将头指针移到现在的位置
			tmpHead = index
			index-- // index往回退一个，因为这一次的字符和之前的不连，不退就错过了一个字符
			// 将节点指针移到树根
			tmpNode = this.Root
		} else {
			//匹配到词典某个词的非词前缀（此前缀本身不是词）。
			//这一段只是词典里某个词的前缀 割出第一个字，扫描指针重新到刚刚的字后面
			word := string(chars[tmpHead])
			results = append(results, word)

			tmpHead++           //单词头指针移动一格
			index = tmpHead - 1 //扫描指针回到前面 这个地方减一是因为这次循环结束会加1
			tmpNode = this.Root
		}

	}

	return results
}
