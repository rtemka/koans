package graph

type Graph struct {
	Edges map[*Vertex][]*Vertex
}

type Vertex struct {
	Val int
}
