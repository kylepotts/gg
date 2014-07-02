package gg
import (
	"container/list"
	"fmt"
)
type Vertex interface{
	ToString() string
}
type Edge struct {
	Vertex1 Vertex
	Vertex2 Vertex
}
type adjacenyList map[Vertex]* list.List
type Graph struct {
	edges adjacenyList
	edgeCount,vertexCount int
}

func(a adjacenyList)String() string {
	s := ""
	for vertex,vertList := range a{
		s+= vertex.ToString()
		for e:=vertList.Front(); e!= nil; e=e.Next(){
			s+=fmt.Sprintf("->%v",e.Value)
		}
		s+="\n"


	
	}
	return s


}

func (g *Graph)PrintAdjMatrix(){
	fmt.Print(g.edges)
}


func (g *Graph)String() string{
	return fmt.Sprintf("Graph edgeCount %d vertexCount %d", g.edgeCount, g.vertexCount)
}

func NewGraph() *Graph {
	return &Graph{edges:make(adjacenyList)}
}

func (g *Graph) addVertex(v Vertex) bool{
	if g.edges[v] == nil {
		g.edges[v] = list.New()
		g.vertexCount++
		return true
	}
	return false
}

func(g *Graph) AddEdge(e Edge){
	g.addVertex(e.Vertex1)
	g.addVertex(e.Vertex2)
	g.edges[e.Vertex1].PushBack(e.Vertex2)
	g.edges[e.Vertex2].PushBack(e.Vertex1)
}






