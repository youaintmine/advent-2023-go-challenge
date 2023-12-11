package base

import (
	"bufio"
	"fmt"
	"os"
)

func helper(grid []string) int {

}

func solve() {

	scanner := bufio.NewScanner(os.Stdin)
	//ans := 0

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			fmt.Println("Input has be extracted")
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}

func main() {
	solve()
}
