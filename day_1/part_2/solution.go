package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	solution := 0
	series := map[int]bool{}

	frequencies := []int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		frequencies = append(frequencies, i)
	}

	for {
		for _, reading := range frequencies {
			if _, ok := series[solution]; ok {
				// duplicate found
				fmt.Printf("%s\n", solution)
				os.Exit(0)
			} else {
				series[solution] = true
			}

			solution += reading
		}
	}
}
