package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func helper(str *string) int {
	tries := strings.Split(*str, ";")
	//8 green, 6 blue, 20 red;
	//5 blue, 4 red, 13 green;
	//5 green, 1 red
	red, blue, green := 0, 0, 0

	for _, balls := range tries {
		ball := strings.Split(balls, ",")
		// 3 blue
		for _, b := range ball {
			parts := strings.Fields(b)
			number, err := strconv.Atoi(parts[0])
			if err != nil {
				return 0
			}
			switch parts[1] {
			case "red":
				if number > red {
					red = number
				}
			case "green":
				if number > green {
					green = number
				}
			case "blue":
				if number > blue {
					blue = number
				}

			}
		}
	}
	return red * blue * green
}

func getGameID(str *string) int {

	S := strings.Split(*str, " ")
	fmt.Println(S[1])
	num, err := strconv.Atoi(S[1])

	if err != nil {
		return 0
	}

	return num
}

func solve() {
	scanner := bufio.NewScanner(os.Stdin)
	ans := 0
	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			fmt.Println("Input has be extracted")
			break
		}
		game := strings.Split(input, ":")
		//gameID := getGameID(&game[0])

		ans += helper(&game[1])

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
