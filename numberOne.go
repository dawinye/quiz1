package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var intArg, _ = strconv.Atoi(os.Args[1])
	//var arrayX [intArg + 1]int
	startSlice := time.Now()
	var sliceX []int
	elapseSlice := time.Since(startSlice)
	//https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
	fmt.Printf("Slice took %s to initialize", elapseSlice)

	startMap := time.Now()
	var mapX = make(map[int]int)
	elapseMap := time.Since(startMap)
	fmt.Printf("Map took %s tp initialize", elapseMap)

	startInitMap := time.Now()
	for i := 0; i < intArg+1; i++ {
		mapX[i] = i
	}
	endInitMap := time.Since(startInitMap)
	fmt.Printf("Map took %s to initialize", endInitMap)

	startInitSlice := time.Now()
	for i := 0; i < intArg+1; i++ {
		sliceX = append(sliceX, i)
	}
	endInitSlice := time.Since(startInitSlice)
	fmt.Printf("Map took %s to initialize", endInitSlice)
}
