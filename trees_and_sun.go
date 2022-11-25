package bluebird

func GrabTrees(arr []int, n int) int {
	result := 1

	nArr := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > nArr || arr[i] == nArr {
			result++
			nArr = arr[i]
		}
	}

	return result
}
