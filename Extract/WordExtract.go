package Extract

import (
	"math"
)

//关键词提取
func ex(wordlist []string){
	nKeyword := 10

    //阻尼系数（ＤａｍｐｉｎｇＦａｃｔｏｒ），一般取值为0.85
    d := 0.85

    //最大迭代次数
    max_iter := 200;
	//最小变化阈值
	min_diff := 0.001;

	wordAndNeighbor :=map[string]map[string]bool{}
	wordWindow := []string{}
	//记录每个单词窗口内邻居
	for _,word := range wordlist{
		if _,ok:= wordAndNeighbor[word];!ok{
			wordAndNeighbor[word] = make(map[string]bool)
		}
		if len(wordWindow) ==5{
			wordWindow = wordWindow[1:]
		}
		wordWindow =append(wordWindow,word)

		for _, neighbor1 :=range wordWindow{
			for _, neighbor2 :=range wordWindow{
				if neighbor1 == neighbor2{
					continue
				}
				wordAndNeighbor[neighbor1][neighbor2] = true
				wordAndNeighbor[neighbor2][neighbor1] = true

			}

		}
	}

	score := make(map[string]float64)
	for i:= 1; i< max_iter; i++{
		tempScore := make(map[string]float64)
		max_diff := .0
		for word,neighbors:=range wordAndNeighbor{
			tempScore[word] = 1 - d

			for neighbor := range neighbors{
				size := len(wordAndNeighbor[neighbor])
				if word == neighbor || size==0{continue}
				//m.put(key, m.get(key) + d / size * (score.get(other) == null ? 0 : score.get(other)));
				//如果该邻居有权重，更新
				if _, ok:=score[neighbor];ok{
					tempScore[word]= tempScore[word]+ d / float64(size) * score[neighbor]
				}
				//max_diff = Math.max(max_diff, Math.abs(m.get(key) - (score.get(key) == null ? 0 : score.get(key))));
				if _, ok:=score[word];ok{
					max_diff = math.Max(max_diff, math.Abs(tempScore[word]) - score[word])
				} else {
					max_diff = math.Max(max_diff, math.Abs(tempScore[word]) - 0)

				}

			}
		}
		score =tempScore
		//如果最大权重变化小于设定的一个小数
		if max_diff<min_diff{
			break
		}
	}
}
