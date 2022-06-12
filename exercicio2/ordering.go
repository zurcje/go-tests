package ordering

import "sort"

func Ordering(numbers []int) []int {
	sort.Ints(numbers)
	return numbers

}
