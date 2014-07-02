package main

import (
	"toj/gg"
	"fmt"
)

type intVertex struct {
	value int

}

func (v intVertex) ToString() string{
	return fmt.Sprintf("%d",v.value)
}

func main(){
	fmt.Println("Hello")
	g := gg.NewGraph()
	fmt.Println(g)
	v1:= intVertex{1}
	v2:= intVertex{2}
	v3:= intVertex{3}
	e := gg.Edge{v1,v2}
	e2:= gg.Edge{v1,v3}
	g.AddEdge(e)
	g.AddEdge(e2)
	g.PrintAdjMatrix()


}