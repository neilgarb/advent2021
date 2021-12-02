package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	if err := do(); err != nil {
		panic(err)
	}
}

func do() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	f, err := os.Open(filepath.Join(wd, "two_one/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var horizontal, depth int
	for scanner.Scan() {
		horizontal, depth = applyCommand(scanner.Text(), horizontal, depth)
		fmt.Println(horizontal, depth, horizontal*depth)
	}


	return scanner.Err()
}

func applyCommand(text string, horizontal, depth int) (int, int) {
	parts := strings.SplitN(text, " ", 2)
	if len(parts) < 2 {
		return horizontal, depth
	}
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return horizontal, depth
	}
	switch parts[0] {
	case "up":
		depth -= val
	case "down":
		depth += val
	case "forward":
		horizontal += val
	}
	return horizontal, depth
}
