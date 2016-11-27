package algorithm

import "testing"
import "fmt"

func TestAdjacencyList(t *testing.T) {
	first, next := AdjacencyList()

	fmt.Println("顶点 0 的出边")
	TraverseEdge(0, first, next)
}
