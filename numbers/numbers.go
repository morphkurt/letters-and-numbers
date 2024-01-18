package numbers

import (
	"container/heap"
	"math"
	"slices"
	"strconv"
)

type Node struct {
	value   int
	target  int
	path    string
	visited []int
}

type queue []Node

func (h queue) Len() int { return len(h) }
func (h queue) Less(i, j int) bool {
	if math.Abs(float64(h[i].value-h[i].target)) == math.Abs(float64(h[j].value-h[j].target)) {
		return len(h[i].visited) < len(h[j].visited)
	} else {
		return math.Abs(float64(h[i].value-h[i].target)) < math.Abs(float64(h[j].value-h[j].target))
	}
}
func (h queue) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *queue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *queue) Pop() interface{} {

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var op = []func(int, int) (int, rune){add, subtract, divide, multiply}

func Solve(c []int, s int) (int, string) {
	h := &queue{}
	heap.Init(h)
	closestValue := math.MaxInt
	finalValue := 0
	path := ""
	toBeVisited := nextNumIndex(c, []int{})
	for _, v := range toBeVisited {
		heap.Push(h, Node{value: c[v], target: s, path: strconv.Itoa(c[v]), visited: []int{v}})
	}
	for h.Len() > 0 {
		n := heap.Pop(h).(Node)
		if n.value == s {
			path = n.path
			finalValue = n.value
			break
		}
		if closestValue >= int(math.Abs(float64(n.value-s))) {
			closestValue = int(math.Abs(float64(n.value - s)))
			path = n.path
			finalValue = n.value
		}
		if len(n.visited) == len(c) {
			continue
		}
		toBeVisited := nextNumIndex(c, n.visited)
		for _, v := range toBeVisited {
			for _, o := range op {
				nextVal, p := o(n.value, c[v])
				visitedPath := n.path
				if nextVal != 0 {
					visitedPath += string(p) + strconv.Itoa(c[v])
				}
				visited := []int{v}
				visited = append(visited, n.visited...)
				heap.Push(h, Node{value: nextVal, target: s, path: visitedPath, visited: visited})
			}
		}
	}
	return finalValue, path
}

func nextNumIndex(c []int, visitedIndex []int) (nonVisited []int) {
	nonVisited = []int{}
	for i := 0; i < len(c); i++ {
		if !slices.Contains(visitedIndex, i) {
			nonVisited = append(nonVisited, i)
		}
	}
	return nonVisited
}

func add(a, b int) (int, rune) {
	return a + b, '+'
}

func subtract(a, b int) (int, rune) {
	return a - b, '-'
}

func multiply(a, b int) (int, rune) {
	return a * b, '*'
}

func divide(a, b int) (int, rune) {
	if a%b == 0 {
		return a / b, '/'
	}
	return 0, '/'
}
