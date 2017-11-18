package Collections

type Vertex struct {
	Id int
	Edges  map[int]float64   //指向别的顶点的边和权重
	//Character rune //顶点的值
	//BestPrev *Vertex

}

type Graph struct {
	Vertexs []Vertex
}

type Path []int

func NewGraph(vertexnum int)  *Graph{

	graph := &Graph{}
	graph.Vertexs = make([]Vertex,vertexnum)
	for i:=0;i<vertexnum;i++ {
		graph.Vertexs[i].Id = i
		graph.Vertexs[i].Edges = make(map[int]float64)
	}
	return graph
}

func (this *Graph) Connect(from int, to int, weight float64)  {
	this.Vertexs[from].Edges[to] = weight
}

