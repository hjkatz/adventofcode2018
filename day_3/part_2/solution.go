package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	RE_CLAIM = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
)

type Claim struct {
	id     string
	x      int
	y      int
	width  int
	height int
}

type Fabric = [][]map[string]bool

func NewClaim(claim_str string) *Claim {
	matches := RE_CLAIM.FindStringSubmatch(claim_str)

	// runtime panics are OK for this small problem
	id := matches[1]
	x, _ := strconv.Atoi(matches[2])
	y, _ := strconv.Atoi(matches[3])
	width, _ := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])

	return &Claim{
		id:     id,
		x:      x,
		y:      y,
		width:  width,
		height: height,
	}
}

func load_input() (input []*Claim) {
	// strings := []string{
	// 	"#1 @ 1,3: 4x4",
	// 	"#2 @ 3,1: 4x4",
	// 	"#3 @ 5,5: 2x2",
	// }

	// for _, s := range strings {
	// 	input = append(input, NewClaim(s))
	// }

	// return

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, NewClaim(scanner.Text()))
	}

	return
}

func print_fabric(fabric *Fabric) {
	for _, row := range *fabric {
		for _, inch := range row {
			if len(inch) == 0 {
				fmt.Printf(".")
			} else if len(inch) == 1 {
				for k, _ := range inch {
					fmt.Printf("%s", k)
					break
				}
			} else {
				fmt.Printf("X")
			}
		}
		fmt.Printf("\n")
	}
}

func (claim *Claim) mark(fabric *Fabric) {
	for _, row := range (*fabric)[claim.y : claim.y+claim.height] {
		for _, inch := range row[claim.x : claim.x+claim.width] {
			inch[claim.id] = true
		}
	}
}

func find_unique_claim(fabric *Fabric) string {
	invalidated_ids := map[string]bool{}
	possible_ids := map[string]bool{}

	for _, row := range *fabric {
		for _, inch := range row {
			for id, _ := range inch {
				if len(inch) > 1 {
					invalidated_ids[id] = true
					delete(possible_ids, id)
				} else {
					if _, ok := invalidated_ids[id]; ok {
						delete(possible_ids, id)
					} else {
						possible_ids[id] = true
					}
				}
			}
		}
	}

	// return the first id, there can only be 1!
	for id, _ := range possible_ids {
		return id
	}

	return ""
}

func main() {
	input := load_input()

	fabric := make(Fabric, 1000)
	for i, _ := range fabric {
		fabric[i] = make([]map[string]bool, 1000)
		for j, _ := range fabric[i] {
			fabric[i][j] = map[string]bool{}
		}
	}

	for _, claim := range input {
		claim.mark(&fabric)
	}

	fmt.Printf("%s\n", find_unique_claim(&fabric))
}
