package data

import (
	"math/rand"
)

func GetArray(size int) []int {
	var arr []int
	for index := 0; index < size; index++ {
		arr[index] = rand.Intn(10000)
	}

	return arr
}

type Result struct {
	time       float64
	name       string
	isVVariant bool
}
