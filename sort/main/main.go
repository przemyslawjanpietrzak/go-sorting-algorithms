package main

import (
	"aisd-sortowanie/sort/data"
	"aisd-sortowanie/sort/heap"
	"aisd-sortowanie/sort/insert"
	"aisd-sortowanie/sort/quick"
	"aisd-sortowanie/sort/quick_iterative"
	"encoding/json"
	"fmt"
	"time"
)

type Result struct {
	Time       float64
	Name       string
	IsVVariant bool
}

func main() {
	var sorts = make(map[string]func([]int) []int)
	sorts["quick"] = quick.Quick
	sorts["heap"] = heap.Heap
	sorts["quick_interative"] = quick_iterative.IterativeQsort
	sorts["insert"] = insert.Insert

	var testsCount int = 5
	var testsValues []int = []int{10, 100, 1000, 10000, 100000}

	var results = make([]Result, len(testsValues)*len(sorts))

	var arr []int
	var index int
	for key := range sorts {
		for _, testValue := range testsValues {
			var timeValue float64
			for i := 0; i < testsCount; i++ {

				arr = data.GetArray(testValue)
				start := time.Now()
				sorts[key](arr)

				elapsed := time.Since(start)
				timeValue += elapsed.Seconds()
			}

			results[index] = Result{
				Time:       timeValue / float64(testsCount),
				Name:       key,
				IsVVariant: false,
			}

			index++
		}
	}

	responseJSon, _ := json.Marshal(results)
	fmt.Println(string(responseJSon))

}
