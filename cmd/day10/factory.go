package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const EPSILON = 1e-9

type Machine struct {
	Target         []bool
	Buttons        [][]int
	JoltageTargets []int
}

func parseInput(input string) []Machine {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var machines []Machine

	reTarget := regexp.MustCompile(`\[([.#]+)\]`)
	reButtons := regexp.MustCompile(`\(([\d,]+)\)`)
	reJoltage := regexp.MustCompile(`\{([\d,]+)\}`)

	for _, line := range lines {
		if line == "" {
			continue
		}
		targetMatch := reTarget.FindStringSubmatch(line)
		targetStr := targetMatch[1]
		target := make([]bool, len(targetStr))
		for i, char := range targetStr {
			if char == '#' {
				target[i] = true
			}
		}

		buttonMatches := reButtons.FindAllStringSubmatch(line, -1)
		buttons := make([][]int, len(buttonMatches))
		for i, match := range buttonMatches {
			nums := strings.Split(match[1], ",")
			for _, numStr := range nums {
				numStr = strings.TrimSpace(numStr)
				num, _ := strconv.Atoi(numStr)
				buttons[i] = append(buttons[i], num)
			}
		}

		joltageMatch := reJoltage.FindStringSubmatch(line)
		var joltageTargets []int
		if len(joltageMatch) > 1 {
			nums := strings.Split(joltageMatch[1], ",")
			for _, numStr := range nums {
				numStr = strings.TrimSpace(numStr)
				num, _ := strconv.Atoi(numStr)
				joltageTargets = append(joltageTargets, num)
			}
		}

		machines = append(machines, Machine{
			Target:         target,
			Buttons:        buttons,
			JoltageTargets: joltageTargets,
		})
	}

	return machines
}

// findMinPresses solves Part 1 using Breadth-First Search (BFS)
func findMinPresses(machine Machine) int {
	type state struct {
		lights int
		dist   int
	}

	targetLights := 0
	for i, b := range machine.Target {
		if b {
			targetLights |= (1 << i)
		}
	}

	queue := []state{{0, 0}}
	seen := map[int]bool{0: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.lights == targetLights {
			return curr.dist
		}

		for _, btn := range machine.Buttons {
			nextLights := curr.lights
			for _, bit := range btn {
				if bit < 32 { // safety check for bitmask
					nextLights ^= (1 << bit)
				}
			}
			if !seen[nextLights] {
				seen[nextLights] = true
				queue = append(queue, state{nextLights, curr.dist + 1})
			}
		}
	}
	return -1
}

func countBtnPress(input string) int {
	machines := parseInput(input)
	total := 0
	for _, m := range machines {
		res := findMinPresses(m)
		if res != -1 {
			total += res
		}
	}
	return total
}

// --- PART 2 LOGIC (Gaussian Elimination + DFS) ---

type Matrix struct {
	data         [][]float64
	rows, cols   int
	dependents   []int
	independents []int
}

func newMatrix(m Machine) *Matrix {
	rows := len(m.JoltageTargets)
	cols := len(m.Buttons)
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols+1)
	}

	for c, button := range m.Buttons {
		for _, r := range button {
			if r < rows {
				data[r][c] = 1.0
			}
		}
	}

	for r, val := range m.JoltageTargets {
		data[r][cols] = float64(val)
	}

	matrix := &Matrix{data: data, rows: rows, cols: cols}
	matrix.gaussianElimination()
	return matrix
}

func (m *Matrix) gaussianElimination() {
	pivot := 0
	col := 0

	for pivot < m.rows && col < m.cols {
		bestRow := pivot
		maxVal := math.Abs(m.data[pivot][col])

		for r := pivot + 1; r < m.rows; r++ {
			if val := math.Abs(m.data[r][col]); val > maxVal {
				maxVal = val
				bestRow = r
			}
		}

		if maxVal < EPSILON {
			m.independents = append(m.independents, col)
			col++
			continue
		}

		m.data[pivot], m.data[bestRow] = m.data[bestRow], m.data[pivot]
		m.dependents = append(m.dependents, col)

		pv := m.data[pivot][col]
		for j := col; j <= m.cols; j++ {
			m.data[pivot][j] /= pv
		}

		for r := 0; r < m.rows; r++ {
			if r != pivot {
				factor := m.data[r][col]
				if math.Abs(factor) > EPSILON {
					for j := col; j <= m.cols; j++ {
						m.data[r][j] -= factor * m.data[pivot][j]
					}
				}
			}
		}
		pivot++
		col++
	}

	for j := col; j < m.cols; j++ {
		m.independents = append(m.independents, j)
	}
}

func (m *Matrix) valid(values []int) (int, bool) {
	total := 0
	for _, v := range values {
		total += v
	}

	for row := 0; row < len(m.dependents); row++ {
		val := m.data[row][m.cols]
		for i, colIdx := range m.independents {
			val -= m.data[row][colIdx] * float64(values[i])
		}

		if val < -EPSILON {
			return 0, false
		}
		rounded := math.Round(val)
		if math.Abs(val-rounded) > EPSILON {
			return 0, false
		}
		total += int(rounded)
	}
	return total, true
}

func dfs(matrix *Matrix, idx int, values []int, minPresses *int, maxVal int) {
	if idx == len(matrix.independents) {
		if total, ok := matrix.valid(values); ok {
			if total < *minPresses {
				*minPresses = total
			}
		}
		return
	}

	currentSum := 0
	for i := 0; i < idx; i++ {
		currentSum += values[i]
	}

	for val := 0; val < maxVal; val++ {
		if currentSum+val >= *minPresses {
			break
		}
		values[idx] = val
		dfs(matrix, idx+1, values, minPresses, maxVal)
	}
}

func findMinJoltagePresses(machine Machine) int {
	matrix := newMatrix(machine)
	maxJolt := 0
	for _, j := range machine.JoltageTargets {
		if j > maxJolt {
			maxJolt = j
		}
	}

	min := math.MaxInt
	values := make([]int, len(matrix.independents))
	dfs(matrix, 0, values, &min, maxJolt+1)
	if min == math.MaxInt {
		return -1
	}
	return min
}

func countJoltagePresses(input string) int {
	machines := parseInput(input)
	results := make(chan int, len(machines))
	var wg sync.WaitGroup

	for _, m := range machines {
		wg.Add(1)
		go func(machine Machine) {
			defer wg.Done()
			res := findMinJoltagePresses(machine)
			if res != -1 {
				results <- res
			} else {
				results <- 0
			}
		}(m)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	total := 0
	for res := range results {
		total += res
	}
	return total
}
