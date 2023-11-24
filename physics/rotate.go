package physics

//todo, this belongs in a math package
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	Signed | Unsigned
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Numbers interface {
	Integer | Float
}

// rotate a value back round to allow max or 0 never to be reached
// could improve by adding a min
// todo, check this works
func CycleValue[T Numbers](max, value T) T {
	if value > max {
		return value - max
	}

	if value < 0 {
		return value + max
	}

	return value
}

func Cycle360Float64(value float64) float64 {
	return CycleValue[float64](359, value)
}

// indexCycler doesn't work with numbers twice the size of the length, not test with decrements
func indexCycler(sliceLength int, index int) int {
	lastIndex := sliceLength - 1

	if index > lastIndex {
		return index - sliceLength
	}

	// add decremeanste

	return index
}
