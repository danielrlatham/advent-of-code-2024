package one

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("pkg/day/one/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	left_slice := []int{}
	right_slice := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		words := strings.Split(text, "   ")
		left_num, err := strconv.Atoi(words[0])
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Converted number:", left_num)
		}
		right_num, err := strconv.Atoi(words[1])
		if err != nil {
			// Handle the error
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Converted number:", right_num)
		}
		left_slice = append(left_slice, left_num)
		right_slice = append(right_slice, right_num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(left_slice, func(i, j int) bool {
		return left_slice[i] < left_slice[j]
	})

	sort.Slice(right_slice, func(i, j int) bool {
		return right_slice[i] < right_slice[j]
	})

	total_distance := 0
	for i, value := range left_slice {
		val := value - right_slice[i]
		total_distance += abs(val)
	}

	// part 1 answer: 1970720
	fmt.Println(total_distance)

	score_map := make(map[int]int)

	for _, lvalue := range left_slice {
		total := 0
		for _, rvalue := range right_slice {
			if lvalue == rvalue {
				total += lvalue
			}
		}
		if _, ok := score_map[lvalue]; ok {
			score_map[lvalue] = score_map[lvalue] + total
		} else {
			score_map[lvalue] = total
		}
	}

	total_score := 0
	for _, value := range score_map {
		total_score += value
	}
	// part 2 answer: 17191599
	fmt.Println(total_score)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
