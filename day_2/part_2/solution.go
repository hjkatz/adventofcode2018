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

func string_differ(a, b string) (int, string) {
	differ := 0
	common := []byte{}

	for i := 0; i < len(a); i += 1 {
		if a[i] == b[i] {
			common = append(common, a[i])
		} else {
			differ += 1
		}
	}

	return differ, string(common)
}

func find_common_id(input []string) string {
	for idx, id := range input {
		for _, compare := range input[idx+1:] {
			if amount, common := string_differ(id, compare); amount == 1 {
				return common
			}
		}
	}

	return ""
}

func main() {
	input := load_input()

	fmt.Printf("%s\n", find_common_id(input))
}
