package Collections

type Dict struct {
	Root           *MapTrieNode
	Words          []*MapTrieNode //所有单词的指针列表，方便遍历
	TotalFrequency int            //总词频
}

func NewDict() *Dict {
	NewDict := &Dict{}
	NewDict.Root = newMapTrieNode()
	return NewDict
}

//树节点 用Hash表存储<Character, Node>
type MapTrieNode struct {
	Sons      map[rune]*MapTrieNode //所有子节点
	IsEnd     bool                  //是否是某个单词的终点
	Character rune                  //存储的值
	Frequency int                   //该词频
	Weight    float64               //-log（该词频/总词频）= log（总词频/该词频）=log(总词频）-log（该词频）  用作图中边的权重
}

func newMapTrieNode() *MapTrieNode {
	node := &MapTrieNode{}
	node.Sons = make(map[rune]*MapTrieNode)
	node.IsEnd = false
	return node
}

//加载字典文件
func (d *Dict) loadDictionaryFile(filePath string) {

}

//增加一个词
func (d *Dict) Insert(word string,frequency int) {
	if len(word) == 0 {
		return
	}

	tmpNode := d.Root
	letters := []rune(word)
	for _, letter := range letters {

		//如果没有这个字母
		if _, ok := tmpNode.Sons[letter]; !ok {
			tmpNode.Sons[letter] = newMapTrieNode()
			tmpNode.Sons[letter].Character = letter
		}
		tmpNode = tmpNode.Sons[letter]
	}
	tmpNode.IsEnd = true
	tmpNode.Frequency = frequency
	d.TotalFrequency += frequency
	d.Words=append(d.Words, tmpNode)
}

//计算所有词的权重


//查询某个词是否存在，给出权值

//和一段文本比对，给出词典中有的某个字开头的所有词
func (d *Dict) findMatchWord(words []rune,from int)  {

}

