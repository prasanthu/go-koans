package go_koans

func divideFourBy(i int) int {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	return 4 / i
}

const __divisor__ = 2

func aboutPanics() {
	//assert(__delete_me__) // panics are exceptional errors at runtime
	n := divideFourBy(__divisor__)
	assert(n == 2) // panics are exceptional errors at runtime
}
