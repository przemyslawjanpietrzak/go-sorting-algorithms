package main

import "fmt"

func main() {
	var sorts map[string]func([]int)
  var arr = []int{12,3,4,5}

  for key, sort := sorts {
    sort(arr)
  }

	fmt.Println("Hello, 世界")
}
