package insert

func swap(arrayzor []int, i, j int) {
	tmp := arrayzor[j]
	arrayzor[j] = arrayzor[i]
	arrayzor[i] = tmp
}

func Insert(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			swap(array, j, j-1)
		}
	}
	return array
}
