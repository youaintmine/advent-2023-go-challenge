package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func helper(str string, myMap map[string]int) int {

	var first, last int
outerloop:
	for i, _ := range str {
		if unicode.IsDigit(rune(str[i])) {
			first = int(str[i] - '0')
			break
		} else {
			for key, value := range myMap {
				Keylength := min(len(str), i+len(key))
				substr := str[i:Keylength]
				if substr == key {
					first = value
					break outerloop
				}
			}
		}
	}

lastouterloop:
	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]
		if unicode.IsDigit(rune(char)) {
			last = int(char - '0')
			break
		} else {
			for key, value := range myMap {
				Keylength := min(len(str), i+len(key))
				substr := str[i:Keylength]
				if substr == key {
					last = value
					break lastouterloop
				}
			}
		}
	}
	ans := first*10 + last
	return ans
}

func solve() {
	scanner := bufio.NewScanner(os.Stdin)

	ans := 0

	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			fmt.Println("Input has be extracted")
			break
		}
		ans = ans + helper(input, myMap)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	fmt.Println("Answer is : ")
	fmt.Println(ans)
}

func main() {
	fmt.Println("Hell World")
	solve()
}
