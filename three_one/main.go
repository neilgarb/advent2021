package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	f, err := os.Open(filepath.Join(wd, "three_one/input"))
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var tot int
	var ones []int
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if ones == nil {
			ones = make([]int, len(line))
		}
		for i, c := range line {
			if c == '1' {
				ones[i] ++
			}
		}
		tot++
	}

	var gamma, epsilon int
	for i := 0; i < len(ones); i ++ {
		if ones[i] >= tot/2 {
			gamma += 1<<(len(ones)-i-1)
		} else {
			epsilon += 1<<(len(ones)-i-1)
		}
	}

	fmt.Println(gamma, epsilon, gamma*epsilon)

	return scanner.Err()
}
