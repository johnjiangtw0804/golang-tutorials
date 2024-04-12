package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev_sum_1 := -2
	prev_sum_2 := -1
	return func() int {
		if prev_sum_1 == -2 && prev_sum_2 == -1 {
			prev_sum_1 = -1
			prev_sum_2 = 0
			return 0
		} else if prev_sum_1 == -1 && prev_sum_2 == 0 {
			prev_sum_1 = 0
			prev_sum_2 = 1
			return 1
		}
		prev_sum_1, prev_sum_2 = prev_sum_2, prev_sum_1+prev_sum_2
		return prev_sum_2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
