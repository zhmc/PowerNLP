package Collections

type Vertex struct {
	Id    int
	Edges map[int]float64 //指向别的顶点的边和权重
	//Character rune //顶点的值
	//BestPrev *Vertex

}

type Graph struct {
	Vertexes []Vertex
}

type Path []int

func NewGraph(vertexNum int) *Graph {

	graph := &Graph{}
	graph.Vertexes = make([]Vertex, vertexNum)
	for i := 0; i < vertexNum; i++ {
		graph.Vertexes[i].Id = i
		graph.Vertexes[i].Edges = make(map[int]float64)
	}
	return graph
}

func (g *Graph) Connect(from int, to int, weight float64) {
	g.Vertexes[from].Edges[to] = weight
}
