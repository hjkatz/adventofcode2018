package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func load_input() (input []string) {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return
}

func count_chars(str string) (int, int) {
	seen := map[rune]int{}

	for _, char := range str {
		if _, ok := seen[char]; !ok {
			seen[char] = 0
		}

		seen[char] += 1
	}

	has_2 := 0
	has_3 := 0
	for _, value := range seen {
		if value == 2 {
			has_2 = 1
		}

		if value == 3 {
			has_3 = 1
		}
	}

	return has_2, has_3
}

func checksum(input []string) int {
	count_2 := 0
	count_3 := 0

	for _, id := range input {
		a, b := count_chars(id)

		count_2 += a
		count_3 += b
	}

	return count_2 * count_3
}

func main() {
	input := load_input()

	fmt.Printf("%s\n", checksum(input))
}
