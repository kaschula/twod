package slice

// todo make generic
// Range makes a range of numbers in assecnding order only
// if max is larger then min then max min are swopped
func Range(min, max int) []int {
	if min == max {
		return []int{min}
	}

	if max < min {
		min, max = max, min
	}

	r := []int{}
	for i := min; max+1 > i; i++ {
		r = append(r, i)
	}

	return r
}

func Reverse(is []int) []int {
	length := len(is)
	if length < 2 {
		return is
	}

	res := make([]int, length)
	lastIndex := length - 1

	for i, v := range is {
		res[lastIndex-i] = v
	}

	return res
}
