package main

import (
	"aisd-sortowanie/sort/data"
	"aisd-sortowanie/sort/quick"
	"fmt"
	"time"
)

type Result struct {
	time       float64
	name       string
	isVVariant bool
}

func main() {
	var sorts map[string]func([]int) []int
	sorts["quick"] = quick.Quick

	var testsCount int = 5 //settings
	var testsValues []int = []int{10, 100, 1000, 10000, 100000, 1000000}

	// var results list.List
	var results = make([]Result, len(testsValues)*len(sorts))

	var arr = []int{12, 3, 4, 5}
	var index int
	for key := range sorts {
		for testValue := range testsValues {

			var timeValue float64
			for i := 0; i < testsCount; i++ {
				arr = data.GetArray(testValue)
				start := time.Now()
				sorts[key](arr)
				elapsed := time.Since(start)
				timeValue += elapsed.Seconds()
			}

			results[index] = Result{
				time:       timeValue / float64(testsCount),
				name:       key,
				isVVariant: false,
			}

			index++
		}
	}

	fmt.Println("Hello, 世界")
}
