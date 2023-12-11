package thir

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"unicode"
)

var (
	matrixMu sync.Mutex
	visited  [][]bool
)

var neighbors = [][]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func traverseLR(i int, j int, grid []string) int {
	visited[i][j] = true
	//Go towards the left
	l, r := j, j

	for l >= 0 && unicode.IsDigit(rune(grid[i][l])) {
		visited[i][l] = true
		l--
	}
	l++
	//Go towards the right
	for r < len(grid) && unicode.IsDigit(rune(grid[i][r])) {
		visited[i][r] = true
		r++
	}
	final := grid[i][l:r]
	//fmt.Println(final)

	res, err := strconv.Atoi(final)

	if err != nil {
		return 0
	}
	//fmt.Println("This is a new number")

	return res
}

func helper(grid []string) int {
	ans := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if unicode.IsDigit(rune(grid[i][j])) == false && grid[i][j] != '.' {
				for _, neighbor := range neighbors {
					ni, nj := i+neighbor[0], j+neighbor[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[0]) {
						if visited[ni][nj] == false && unicode.IsDigit(rune(grid[ni][nj])) {
							ans += traverseLR(ni, nj, grid)
						}
					}
				}
			}
		}
	}
	return ans
}

func solve() {
	matrixMu.Lock()
	defer matrixMu.Unlock()

	scanner := bufio.NewScanner(os.Stdin)
	//ans := 0

	var grid []string

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			fmt.Println("Input has be extracted")
			break
		}
		grid = append(grid, input)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	rows := len(grid)
	columns := len(grid[0])

	visited = make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, columns)
	}

	ans := helper(grid)
	fmt.Println(ans)
}

func main() {
	solve()
}
