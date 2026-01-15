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
	const threeDSpace = 3
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([]Point, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ",")
		if len(parts) != threeDSpace {
			continue
		}
		x, _ := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		y, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		z, _ := strconv.ParseInt(strings.TrimSpace(parts[2]), 10, 64)
		points = append(points, Point{x, y, z})
	}
	return points
}

func squareDistance(p1, p2 Point) int64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z
	return dx*dx + dy*dy + dz*dz
}

func NewUnionFind(elements int) *UnionFind {
	parent := make([]int, elements)
	size := make([]int, elements)
	for i := 0; i < elements; i++ {
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

func connectJunctionBoxes(points []Point, connections int) int64 {
	sizePoints, pairs := createAllPossibleUniquePairs(points)

	sortAsc(pairs)

	if connections > len(pairs) {
		connections = len(pairs)
	}

	uf := NewUnionFind(sizePoints)
	for i := 0; i < connections; i++ {
		uf.Union(pairs[i].p1, pairs[i].p2)
	}

	largestCircuits := getLargestCircuits(sizePoints, uf)

	var result int64 = 1
	const amountLargestJB = 3
	for i := 0; i < amountLargestJB && i < len(largestCircuits); i++ {
		result *= int64(largestCircuits[i])
	}
	return result
}

func getLargestCircuits(sizePoints int, uf *UnionFind) []int {
	largestCircuits := make([]int, 0)
	for i := 0; i < sizePoints; i++ {
		if uf.parent[i] == i {
			largestCircuits = append(largestCircuits, uf.size[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(largestCircuits)))
	return largestCircuits
}

func sortAsc(pairs []Pair) {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].distSq != pairs[j].distSq {
			return pairs[i].distSq < pairs[j].distSq
		}
		if pairs[i].p1 != pairs[j].p1 {
			return pairs[i].p1 < pairs[j].p1
		}
		return pairs[i].p2 < pairs[j].p2
	})
}

func createAllPossibleUniquePairs(points []Point) (int, []Pair) {
	sizePoints := len(points)
	optimizedCap := sizePoints * (sizePoints - 1) / 2
	pairs := make([]Pair, 0, optimizedCap)
	for i := 0; i < sizePoints; i++ {
		for j := i + 1; j < sizePoints; j++ {
			pairs = append(pairs, Pair{i, j, squareDistance(points[i], points[j])})
		}
	}
	return sizePoints, pairs
}

func multiplyXCoordLastTwoJB(points []Point) int64 {
	sizePoints, pairs := createAllPossibleUniquePairs(points)

	sortAsc(pairs)

	uf := NewUnionFind(sizePoints)
	numCircuits := sizePoints
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
