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

func RemoveIntSliceElement(intSlice []int, index int) []int {
	intLine := make([]int, 0, len(intSlice)-1)
	intLine = append(intLine, intSlice[:index]...)
	intLine = append(intLine, intSlice[index+1:]...)
	//return append(, ...)
	return intLine
}
