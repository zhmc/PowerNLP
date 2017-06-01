package Collections

import "fmt"

//双数组Trie树
type DATrie struct {
	Base         []int        //base数组
	Check        []int        //check数组
	Tail         [][]rune     // 存放尾串的数组
	tailPosition int          // 现在尾串的位置
	RuneCodeMap  map[rune]int //<字符，code码>hash表

}

//标记结束的字符
const EndRune = '#'

//初始化双数组Tire
func NewDATrie() *DATrie {
	newDATrie := DATrie{}
	newDATrie.Base = make([]int, 1024)
	//Base数组0位置不用，1是根节点
	newDATrie.Base[1] = 1
	newDATrie.Check = make([]int, 1024)
	newDATrie.Tail = make([][]rune, 1) //tail数组第一个必须空，因为base中存负值，索引为0的用不上
	newDATrie.tailPosition = 0
	newDATrie.RuneCodeMap = make(map[rune]int)
	newDATrie.RuneCodeMap[EndRune] = 1
	for i := 0; i < 26; i++ {
		//+1是因为code从1开始
		newDATrie.RuneCodeMap[rune('a'+i)] = len(newDATrie.RuneCodeMap) + 1
	}

	return &newDATrie
}

//将双数组扩充一定长度
func (this *DATrie) extendBaseCheck(addSize int) {
	new := make([]int, addSize)
	this.Base = append(this.Base, new[:]...)
	this.Check = append(this.Check, new[:]...)
}

//获得字符的code码
func (this *DATrie) GetRuneCode(_rune rune) int {
	if _, ok := this.RuneCodeMap[_rune]; !ok {
		this.RuneCodeMap[_rune] = len(this.RuneCodeMap) + 1
	}
	return this.RuneCodeMap[_rune]
}

//寻找到新的base值，能够满足按照转移得到的子节点的位置都没有被占用
func (this *DATrie) x_check(checklist []int) int {

	//从1开始寻找新的base值
	for i := 1; ; i++ {
		stopflag := true

		//遍历所有子节点的转移字符（到达子节点的code）
		for _, inputCode := range checklist {
			//新的子节点位置
			newSonNodeIndex := i + inputCode
			//如果这个位置已经被占据，退出
			if this.Base[newSonNodeIndex] != 0 || this.Check[newSonNodeIndex] != 0 {
				stopflag = false
				break
			}
			//新的子节点位置已经超过原数组大小了
			if newSonNodeIndex > len(this.Base) {
				this.extendBaseCheck(newSonNodeIndex - len(this.Base) + 1)
			}
		}

		//遍历所有子节点的转移字符结束，发现可以满足要求
		if stopflag {
			return i
		}
		return 0
	}
}

//找出某个节点的所有子节点
func (this *DATrie) getChildList(fatherIndex int) []int {
	childList := make([]int, 0)
	//遍历所有转移字符，看看这个节点是否有这一条边
	for i := 1; i < len(this.RuneCodeMap); i++ {
		maybeSonIndex := this.Base[fatherIndex] + i
		if maybeSonIndex > len(this.Base) {
			break
		}
		if this.Check[maybeSonIndex] == fatherIndex {
			childList = append(childList, i)
		}
	}
	return childList
}

//将一个字符串的尾串添加到TAIL数组中, 返回位置
func (this *DATrie) AppendToTailArray(runes []rune, position int) int {
	tailRunes := runes[position:]
	this.Tail = append(this.Tail, tailRunes)
	return len(this.Tail) - 1
}

//添加单词 最核心部分
func (this *DATrie) Insert(word string) {
	wordRunes := []rune(word)
	wordRunes = append(wordRunes, EndRune)

	prePosition := 1        //之前位置
	var currentPosition int //现在位置-走一个单词的路径过程中当前字符在base数组中的索引位置

	//index用于取尾串
	for index, char := range wordRunes {
		//获取该字符连接的子节点的位置
		currentPosition = this.Base[prePosition] + this.GetRuneCode(char)

		//扩充长度
		if currentPosition > len(this.Base)-1 {
			this.extendBaseCheck(currentPosition - len(this.Base) + 1)
		}
		//该子节点未被占用
		if this.Base[currentPosition] == 0 && this.Check[currentPosition] == 0 {
			//先插入ba#, bc#，又插入b#（新的单词是旧的子串
			if char == EndRune{
				this.Base[currentPosition] = 0
				this.Check[currentPosition] = prePosition
				return //结束了
			}

			this.AppendToTailArray(wordRunes, index+1) //index要不要加1
			//尾串添加到tail数组中的位置为len(this.Tail)-1
			this.Base[currentPosition] = -(len(this.Tail) - 1)
			this.Check[currentPosition] = prePosition
			return //结束了
		}

		//该节点已经被占用
		//如果可以正常转移 未发生冲突
		if this.Base[currentPosition] > 0 && this.Check[currentPosition] == prePosition {
			prePosition = currentPosition
			continue
		}
		//发生冲突

		//冲突 1：遇到 Base[cur_p]小于0的，即遇到一个被压缩存到Tail中的字符串
		if this.Base[currentPosition] < 0 && this.Check[currentPosition] == prePosition {
			tailIndex := -this.Base[currentPosition]
			//发生冲突的单词（树的路径）的尾串完全一样，就停止了
			if string(this.Tail[tailIndex]) == string(wordRunes[index+1]) {
				return
			}
			//尾串不一样。取出共同前缀，存入Base数组，独立区分尾串存入Tail
			//相同的字符
			if this.Tail[tailIndex][0] == wordRunes[index+1] {
				tailHeadCode := this.GetRuneCode(wordRunes[index+1])
				newBaseValue := this.x_check([]int{tailHeadCode})
				//换上新的base值，从负值到正值（有一个子节点）
				this.Base[currentPosition] = newBaseValue

				//改变tail数组中存放的。去掉第一个
				this.Tail[tailIndex] = this.Tail[tailIndex][1:]
				//这条边到达的子节点在Base数组中位置是newBaseValue+tailHeadCode
				this.Base[newBaseValue+tailHeadCode] = -tailIndex
				this.Check[newBaseValue+tailHeadCode] = currentPosition
				prePosition = currentPosition
				continue

			} else { //不同的字符 可能有一个为结束符
				fmt.Println("string(this.Tail[tailIndex][0])",string(this.Tail[tailIndex][0]))
				tailHeadCode := this.GetRuneCode(this.Tail[tailIndex][0])
				nextCharCode := this.GetRuneCode(wordRunes[index+1])
				newBaseValue := this.x_check([]int{tailHeadCode, nextCharCode})
				//换上新的base值，从负值到正值（有两个子节点）
				this.Base[currentPosition] = newBaseValue
				this.Check[newBaseValue+tailHeadCode] = currentPosition
				this.Check[newBaseValue+nextCharCode] = currentPosition

				//Tail 为END_FLAG 的情况
				if this.Tail[tailIndex][0] == EndRune {
					fmt.Println("this.Tail[tailIndex][0] == EndRune")
					this.Base[newBaseValue+tailHeadCode] = 0
				} else {
					//改变tail数组中存放的。去掉第一个
					this.Tail[tailIndex] = this.Tail[tailIndex][1:]
					//这条边到达的子节点在Base数组中位置是newBaseValue+tailHeadCode
					this.Base[newBaseValue+tailHeadCode] = -tailIndex
				}

				//word转化的rune数组中剩余的部分存入tail
				//为什么要用index+2？因为index+1已经作为到达子节点的边存在双数组中了
				if wordRunes[index+1] == EndRune {
					fmt.Println("wordRunes[index+1] == EndRune")
					this.Base[newBaseValue+nextCharCode] = 0
				} else {
					newTailIndex := this.AppendToTailArray(wordRunes, index+2)
					this.Base[newBaseValue+nextCharCode] = -newTailIndex
				}
				return

			}

		}

		//TODO

		//冲突2：当前结点已经被占用，需要调整pre的base
		//这里也就是整个DATrie最复杂的地方了
		if this.Check[currentPosition] != prePosition {

		}

	}
}

//确认是否存在某个单词
func (this *DATrie) Contains(word string) bool {
	exist := false
	chars := []rune(word)
	chars = append(chars, EndRune)
	prePosition := 1
	currentPosition := 0
	for index, char := range chars {
		currentPosition = this.Base[prePosition] + this.GetRuneCode(char)
		//等于0，根本没有或者是结束符转移到了这里
		if this.Base[currentPosition] == 0 {
			fmt.Println("this.Base[currentPosition] == 0")
			if this.Check[currentPosition] == prePosition && index==len(chars)-1{
				fmt.Println("this.Check[currentPosition] == prePosition && index==len(chars)-1")
				return true
			}
			return false
		} else if this.Base[currentPosition] > 0 {
			//大于0，继续转移
			if this.Check[currentPosition] != prePosition {
				return false
			}
			prePosition = currentPosition
		} else {
			//小于0，去比较尾串

			fmt.Println("this.Base[currentPosition]", this.Base[currentPosition])
			fmt.Println("string(char) ",string(char))
			fmt.Println("string(this.Tail[-this.Base[currentPosition]])",string(this.Tail[-this.Base[currentPosition]]))
			fmt.Println("string(chars[index+1:])",string(chars[index+1:]))

			if string(this.Tail[-this.Base[currentPosition]]) == string(chars[index+1:]) {
				return true
			} else {return false}

		}

	}
	return exist
}
