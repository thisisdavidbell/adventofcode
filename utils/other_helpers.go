package utils

func IntAbs(x int) (abs int) {
	if x < 0 {
		return -x
	}
	return x
}

func IgnoreError(interf interface{}, err error) int {
	if err != nil {
		panic(err)
	}
	return interf.(int)
}
