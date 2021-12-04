package utils

// CopyIntStringMap -  copy map
func CopyIntStringMap(origMap map[int]string) (newMap map[int]string) {
	// Copy from the original map to the target map

	newMap = make(map[int]string)
	for key, value := range origMap {
		newMap[key] = value
	}
	return
}
