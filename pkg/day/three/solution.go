package three 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "sort"
	// "strconv"
	// "strings"
)

func Solution() {
	file, err := os.Open("pkg/day/three/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
