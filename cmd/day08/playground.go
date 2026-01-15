package main

import (
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y, z int64
}

type Pair struct {
	p1, p2 int
	distSq int64
}

type UnionFind struct {
	parent []int
	size   []int
}

func parseInput(input string) []Point {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		x, _ := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		y, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		z, _ := strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 64)
		points = append(points, Point{x, y, z})
	}
	return points
}

func distSq(p1, p2 Point) int64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)
	if rootI != rootJ {
		if uf.size[rootI] < uf.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
	}
}

func solve(points []Point, connections int) int64 {
	n := len(points)
	pairs := make([]Pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, Pair{i, j, distSq(points[i], points[j])})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].distSq != pairs[j].distSq {
			return pairs[i].distSq < pairs[j].distSq
		}
		if pairs[i].p1 != pairs[j].p1 {
			return pairs[i].p1 < pairs[j].p1
		}
		return pairs[i].p2 < pairs[j].p2
	})

	if connections > len(pairs) {
		connections = len(pairs)
	}

	uf := NewUnionFind(n)
	for i := 0; i < connections; i++ {
		uf.Union(pairs[i].p1, pairs[i].p2)
	}

	sizes := make([]int, 0)
	for i := 0; i < n; i++ {
		if uf.parent[i] == i {
			sizes = append(sizes, uf.size[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	var result int64 = 1
	for i := 0; i < 3 && i < len(sizes); i++ {
		result *= int64(sizes[i])
	}
	return result
}

func solvePart2(points []Point) int64 {
	n := len(points)
	pairs := make([]Pair, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, Pair{i, j, distSq(points[i], points[j])})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].distSq != pairs[j].distSq {
			return pairs[i].distSq < pairs[j].distSq
		}
		if pairs[i].p1 != pairs[j].p1 {
			return pairs[i].p1 < pairs[j].p1
		}
		return pairs[i].p2 < pairs[j].p2
	})

	uf := NewUnionFind(n)
	numCircuits := n
	var lastXProduct int64

	for _, p := range pairs {
		root1 := uf.Find(p.p1)
		root2 := uf.Find(p.p2)
		if root1 != root2 {
			uf.Union(p.p1, p.p2)
			numCircuits--
			if numCircuits == 1 {
				lastXProduct = points[p.p1].x * points[p.p2].x
				break
			}
		}
	}

	return lastXProduct
}
