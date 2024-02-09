package utils

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func AbsDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}

	return x - y
}

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func UintMin(x, y uint) uint {
	if x < y {
		return x
	}
	return y
}

func UintMax(x, y uint) uint {
	if x > y {
		return x
	}
	return y
}
