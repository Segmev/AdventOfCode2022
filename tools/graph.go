package tools

type GraphNode[T any] struct {
	Id     string
	Parent *GraphNode[T]
	Nodes  map[string]*GraphNode[T]
	Value  T
}

func (g GraphNode[T]) Contains(key string) bool {
	_, ok := g.Nodes[key]
	return ok
}
