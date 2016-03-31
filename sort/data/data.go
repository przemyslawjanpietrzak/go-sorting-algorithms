package data

import (
	"math/rand"
)

func GetArray() [10000]int {
	var arr [10000]int
	for index := 0; index < 10000; index++ {
		arr[index] = rand.Intn(10000)
	}

	return arr
}

type Result struct {
	time       float64
	name       string
	isVVariant bool
}

var Results []Result
