package two

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "sort"
	"strconv"
	"strings"
)

func Solution() {
	file, err := os.Open("pkg/day/two/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		slice := []int{}
		textNums := strings.Split(text, " ")
		for _, v := range textNums {
			num, err := strconv.Atoi(v)
			if err != nil {
				// Handle the error
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Converted number:", num)
				slice = append(slice, num)
			}
		}

		safe := reportIsSafe(slice)
		fmt.Println("Safe?: ", safe)
		if safe == true {
			count += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// part 1 answer: 660
	fmt.Println(count)
}

func reportIsSafe(slice []int) bool {
	safe := true
	previousValue := 0
	direction := ""
	initialDirection := ""

	for i, v := range slice {
		if i == 0 {
			previousValue = v
			continue
		}

		if i == 1 {
			if v > previousValue {
				initialDirection = "inc"
			} else {
				initialDirection = "dec"
			}
		}

		if v > previousValue {
			direction = "inc"
		} else {
			direction = "dec"
		}

		if direction != initialDirection {
			safe = false
			fmt.Println("Direction changed, ", initialDirection, direction)
			break
		}

		diff := difference(previousValue, v)
		if diff > 3 || diff < 1 {
			safe = false
			fmt.Println("Difference ", diff)
			break
		}

		previousValue = v
	}
	return safe
}

func difference(num1 int, num2 int) int {
	return abs(num1 - num2)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
