package one

import (
	"fmt"
	"io"
	"log"
	"os"
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

	b, err := io.ReadAll(file)
  text := string(b)
	fmt.Print(text)
}
