package insert

import (
	"fmt"
	"time"
	"aisd/sort/data"
)

func insert(arr []int) []int{

	var i int
	currentTime := time.Now()
	// arr := [5]int{1, 21, 3, 4, 5}
	arr := data.GetArray()
	arrQuick := data.GetArray()
	arrQuick = quick.sort(arrQuick)
	for i=2; i<len(arr); i++ {
		 	key := arr[i]
			j := i-1
			for (j>0 && arr[j] > key) {
				arr[j+1] = arr[ j ]
				arr[j] = key
				j = j - 1
		}
	}
	fmt.Printf("%v", arrQuick)
	fmt.Printf(time.Since(currentTime).String())

	return arr
}
