package main

import (
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

	input, err := os.ReadFile(filepath.Join(wd, "six_one/input"))
	if err != nil {
		return err
	}

	initial := strings.Split(string(input),",")
	countPerAge := make([]int, 9)
	for _, f := range initial {
		i, _ := strconv.Atoi(f)
		countPerAge[i]++
	}


	for i := 0; i < 256; i ++ {
		newCount := countPerAge[0]
		for j := 1; j < 9; j ++ {
			countPerAge[j-1] = countPerAge[j]
		}
		countPerAge[6] += newCount
		countPerAge[8] = newCount
	}

	var tot int
	for _, c := range countPerAge {
		tot += c
	}

	fmt.Println(tot)

	return nil
}

type fish int

func makeFish(s string) fish {
	i, _ := strconv.Atoi(s)
	return fish(i)
}

func (f *fish) advance() (fish, bool) {
	if *f == 0 {
		*f = 6
		return 8, true
	}
	*f--
	return 0, false
}