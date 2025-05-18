package graph

func DFS(graph *Graph, startVertex *Vertex, val int) *Vertex {
	visited := make(map[*Vertex]bool)
	return dfs(graph, startVertex, val, visited)
}

func dfs(graph *Graph, vertex *Vertex, val int, visited map[*Vertex]bool) *Vertex {
	visited[vertex] = true
	for _, v := range graph.Edges[vertex] {
		// It could be some other operation.
		if v.Val == val {
			return v
		}
		if !visited[v] {
			dfs(graph, v, val, visited)
		}
	}
	return nil
}

func BFS(graph *Graph, startVertex *Vertex, val int) *Vertex {
	visited := make(map[*Vertex]bool)
	queue := []*Vertex{startVertex}
	var vertex *Vertex
	for len(queue) > 0 {
		vertex, queue = queue[0], queue[1:]
		visited[vertex] = true

		// It could be some other operation.
		if vertex.Val == val {
			return vertex
		}
		for _, v := range graph.Edges[vertex] {
			if !visited[v] {
				queue = append(queue, v)
			}
		}
	}
	return nil
}
