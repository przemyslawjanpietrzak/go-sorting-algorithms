package heap

// import "fmt"

func siftDown(arr []int, lo, hi, first int) {
  root := lo
  for {
  	child := 2*root + 1
  	if child >= hi {
  		break
  	}
  	if child+1 < hi && arr[first+child] < arr[first+child+1] {
  		child++
  	}
  	if arr[first+root] >=arr[first+child] {
  		return
  	}
  	arr[first+root], arr[first+child] = arr[first+child], arr[first+root]
  	root = child
  }
}

func heapSort(arr []int, a, b int) {
	first := a
	lo := 0
	hi := b - a
	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		arr[first], arr[first+i] = arr[first+i], arr[first]
		siftDown(arr, lo, i, first)
	}
}
