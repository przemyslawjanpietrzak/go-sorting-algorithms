package quick

import "math/rand"

func Quick(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivotIndex := rand.Int() % len(arr)

	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	Quick(arr[:left])
	Quick(arr[left+1:])

	return arr
}
