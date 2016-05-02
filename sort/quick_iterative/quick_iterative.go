package quick_iterative

func NewStack() *Stack {
	return &Stack{}
}

type Stack struct {
	nodes []int
	count int
}

func (s *Stack) Push(n int) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

func (s *Stack) Pop() int {
	s.count--
	return s.nodes[s.count]
}

func IterativeQsort(arr []int) []int {

	var stack = NewStack()
	stack.Push(0)
	stack.Push(len(arr))
	for stack.count != 0 {
		var end int = stack.Pop()
		var start int = stack.Pop()
		if end-start < 2 {
			continue
		}
		var p = start + ((end - start) / 2)
		p = partition(arr, p, start, end)

		stack.Push(p + 1)
		stack.Push(end)

		stack.Push(start)
		stack.Push(p)

	}
	return stack.nodes
}

func partition(arr []int, p int, start int, end int) int {
	var l int = start
	var h int = end - 2
	var piv int = arr[p]
	swap(arr, p, end-1)

	for l < h {
		if arr[l] < piv {
			l++
		} else if arr[h] >= piv {
			h--
		} else {
			swap(arr, l, h)
		}
	}
	var idx int = h
	if arr[h] < piv {
		idx++
	}
	swap(arr, end-1, idx)
	return idx
}

func swap(arr []int, i int, j int) {
	var temp int = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
