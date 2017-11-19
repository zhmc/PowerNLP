package main

import (
	"github.com/zhmc/PowerNLP/Seg/Collections"

	"fmt"
)

func main()  {
	g := Collections.NewGraph(7)
	g.Connect(1,2,1)
	g.Connect(1,3,2)
	g.Connect(1,4,4)
	g.Connect(2,5,5)
	g.Connect(3,5,3)
	g.Connect(3,4,2.1)
	g.Connect(4,5,0.5)
	g.Connect(5,6,2)
	g.Connect(4,6,4)

	fmt.Println(g.Vertexs[4].Edges[6])

	fmt.Println(g.FindShortestPath(1,6))
	fmt.Println(g.FindShortestPath(3,6))
	fmt.Println(g.FindShortestPath(2,4))
	fmt.Println(g.FindShortestPath(1,4))
}

