package main

import (
	"aisd-sortowanie/sort/data"
	"aisd-sortowanie/sort/heap"
	// "aisd-sortowanie/sort/insert"
	"aisd-sortowanie/sort/quick"
	"encoding/json"
	"fmt"
	"time"
)

type Result struct {
	Time       float64
	Name       string
	IsVVariant bool
}

type Result1 struct {
	// time       float64
	Name string
	// isVVariant bool
}

type Response1 struct {
	Page int
}

func main() {
	var sorts = make(map[string]func([]int) []int)
	sorts["quick"] = quick.Quick
	sorts["heap"] = heap.Heap
	// sorts["insert"] = insert.Insert

	// make(map[string]T)
	var testsCount int = 5 //settings
	var testsValues []int = []int{10, 100, 1000, 10000, 100000}

	// var results list.List
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
