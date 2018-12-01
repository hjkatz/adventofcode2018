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

		solution += i
	}

	fmt.Printf("%s\n", solution)
}
