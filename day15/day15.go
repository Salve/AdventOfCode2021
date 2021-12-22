package day15

import (
	"container/heap"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Printf("\n---- Day 15 ----\n")
	part1()
	part2()
}

func part1() {
	r := dijkstra(inputs(1), point{0, 0}, point{99, 99})
	fmt.Printf("Part 1 answer: %d\n\n", r)
}

func part2() {
	r := dijkstra(inputs(5), point{0, 0}, point{499, 499})
	fmt.Printf("Part 2 answer: %d\n\n", r)
}

func adjacent(c map[point]int, p point) []point {
	o := make([]point, 0, 4)
	for _, a := range []point{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}} {
		if _, ok := c[a]; !ok {
			continue
		}
		o = append(o, a)
	}
	return o
}

func dijkstra(c map[point]int, start, end point) int {
	n := len(c)
	dist := make(map[point]int, n)
	visited := make(map[point]struct{}, n)
	items := make(map[point]*Item, n)
	q := make(PriorityQueue, n)
	dist[start] = 0
	visited[start] = struct{}{}
	i := 0
	for p := range c {
		if p != start {
			dist[p] = math.MaxInt
		}
		item := &Item{
			p:        p,
			priority: dist[p],
			index:    i,
		}
		items[p] = item
		q[i] = item
		i++
	}
	heap.Init(&q)

	for len(q) > 0 {
		u := heap.Pop(&q).(*Item).p
		visited[u] = struct{}{}
		for _, v := range adjacent(c, u) {
			if _, ok := visited[v]; ok {
				continue
			}
			alt := dist[u] + c[v]
			if alt < dist[v] {
				dist[v] = alt
				q.update(items[v], v, alt)
			}
		}
	}
	return dist[end]
}

type point struct{ x, y int }

type Item struct {
	p        point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, p point, priority int) {
	item.p = p
	item.priority = priority
	heap.Fix(pq, item.index)
}

func inputs(tile int) map[point]int {
	o := make(map[point]int, 100*100*tile*tile)
	for tx := 0; tx < tile; tx++ {
		for ty := 0; ty < tile; ty++ {
			for x, line := range strings.Fields(input) {
				for y, s := range line {
					o[point{100*tx + x, 100*ty + y}] = 1 + (int(s-'0')+tx+ty-1)%9
				}
			}
		}
	}
	return o
}
