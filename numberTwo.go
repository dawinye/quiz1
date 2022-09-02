package main

import (
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

//https://gobyexample.com/sorting-by-functions
type byLength []Int

func (s byLength) Less(i, j int) bool {
	return s[i] < s[j]
}
func main() {
	var intArg, _ = strconv.Atoi(os.Args[1])

	array := [1000]int{}

	var slice []int
	for i := 0; i < intArg; {
		slice = append(slice, rand.Intn(1000))
	}
	startSortTime := time.Now()
	sort.Ints(slice)

	for i := 0; i < intArg; {
		slice[i] = rand.Intn(1000)
	}
	sort.Stable(array)
}
