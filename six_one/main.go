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

	var fishes []fish
	for _, s := range initial {
		fishes = append(fishes, makeFish(s))
	}

	for i := 0; i < 80; i ++ {
		var newFish []fish
		for i := range fishes {
			nf, spawn := fishes[i].advance()
			if spawn {
				newFish = append(newFish, nf)
			}
		}
		fishes = append(fishes, newFish...)
	}

	fmt.Println(len(fishes))

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