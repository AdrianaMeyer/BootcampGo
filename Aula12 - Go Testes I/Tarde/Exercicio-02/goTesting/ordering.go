package goTesting

func Ordering(array []int) []int {
	
	for i := 1; i < len(array); i++ {
		atual := array[i]
		j := i - 1
		for j >= 0 && array[j] > atual {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = atual
	}

	return array
}