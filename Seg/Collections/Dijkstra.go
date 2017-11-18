package Collections

//import "fmt"

//候选节点
type CandidateVertex struct {
	id         int
	prev int
	pathLength float64
}

const StartVertexFlag  = -10
const UndeterminedVertexFlag  = -1
func (g *Graph) FindShortestPath(from int, to int)  Path{
	resultPath := Path{}
	
	//表示每个顶点的最短路径的前一个顶点,初始赋值-1
	statusArray := make([]int,len(g.Vertexs)) 
	for i:=0;i<len(statusArray);i++{
		statusArray[i] = UndeterminedVertexFlag
	}
	statusArray[from] = StartVertexFlag
	//添加第一个顶点 起始点
	nextVertex := -1
	nextVertexWeight := -1.0
	for k,v := range g.Vertexs[from].Edges{
		if nextVertex == -1{
			nextVertex = k
			nextVertexWeight = v
			continue
		}
		if v < nextVertexWeight{
			nextVertex = k
			nextVertexWeight = v
		}
	}

	statusArray[nextVertex] = from
	
	for nextVertex != to{
		candidateVertexs := g.findCandidateVertexs(statusArray,from)
		//fmt.Println("len(candidateVertexs)",len(candidateVertexs))
		if len(candidateVertexs) == 0{
			break
		}
		choosedVertex := chooseCandidateVertex(candidateVertexs)
		nextVertex = choosedVertex.id
		//fmt.Println("nextVertex ",nextVertex)
		if nextVertex == -1{
			break
		}
		statusArray[nextVertex] = choosedVertex.prev
	}

	prev := to
	//fmt.Println("len(resultPath)",len(resultPath))
	for prev!=from {
		if prev == UndeterminedVertexFlag{
			resultPath = Path{}
			return resultPath
		}
		resultPath = append(resultPath,prev)
		prev = statusArray[prev]
	}
	resultPath = append(resultPath,from)
	prev = statusArray[from]
	//fmt.Println("len(resultPath)",len(resultPath))
	reverse(resultPath)
	return resultPath

}

//从候选顶点中找出最短节点
func chooseCandidateVertex(candidates []CandidateVertex)  CandidateVertex{
	final := -1
	finalPathLength := -1.0
	finalVertex := CandidateVertex{}
	//fmt.Println("chooseCandidateVertex")
	//fmt.Println("len(candidates)",len(candidates))
	for _, candidate := range candidates{
		if final == -1.0{
			final = candidate.id
			finalPathLength = candidate.pathLength
			finalVertex = candidate
			continue
		}
		if candidate.pathLength < finalPathLength{
			final = candidate.id
			finalPathLength = candidate.pathLength
			finalVertex = candidate
		}
	}
	return finalVertex
}

//根据已经确定节点找出候选节点
func (g *Graph) findCandidateVertexs(statusArray []int, from int) []CandidateVertex {
	candidateVertexs := make([]CandidateVertex,0)

	for i:=0;i<len(statusArray);i++ {

		//排除掉未确定节点
		if statusArray[i] == UndeterminedVertexFlag{
			continue
		}
		//从已确定节点指向的节点中选取
		for v,weight :=range g.Vertexs[i].Edges{
			//排除指向的节点是已确定节点
			if statusArray[v] != UndeterminedVertexFlag{
				continue
			}
			sharedePathLength := g.computePathLength(statusArray,i,from)
			//fmt.Println("sharedePathLength",sharedePathLength)

			if sharedePathLength == -1{
				continue
			}
			//添加到候选列表
			candidate :=CandidateVertex{}
			candidate.id=v
			candidate.prev=i
			candidate.pathLength = sharedePathLength+weight
			//fmt.Println("candidate.id",candidate.id)
			//fmt.Println("candidate.prev",candidate.prev)
			//fmt.Println("candidate.pathLength",candidate.pathLength)
			candidateVertexs = append(candidateVertexs,candidate)
		}
	}

	return candidateVertexs
}

//计算已经确定节点的最短路径长度
func  (g *Graph) computePathLength(statusArray []int, vertex int,from int)  float64{
	prev := statusArray[vertex]
	if vertex == from{
		//fmt.Println("warning")
		return 0
	}
	pathlength := g.Vertexs[prev].Edges[vertex]
	for  prev!=from {

		//如果出现-1，则认为路径断掉了。
		if prev == UndeterminedVertexFlag{
			return -1
		}

		pathlength += g.Vertexs[statusArray[prev]].Edges[prev]
		prev = statusArray[prev]
	}

	return pathlength
}

func reverse(s Path) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}