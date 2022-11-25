package bluebird

import "sort"

type Arr struct {
	value int
	index int
}

type ByValue []Arr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func MinSwapArray(arr []int) int {
	nArr := make([]Arr, 0, len(arr))

	for i, v := range arr {
		nArr = append(nArr, Arr{v, i})
	}

	sort.Sort(ByValue(nArr))

	idx := make([]int, 0, len(arr))
	for _, ar := range nArr {
		idx = append(idx, ar.index)
	}

	var result int
	for i := 0; i < len(arr); i++ {
		if i == idx[i] {
			continue
		}

		arr[i], arr[idx[i]] = arr[idx[i]], arr[i]
		idx[i], idx[arr[idx[i]]-1] = i, idx[i]
		result++
	}

	return result
}
