package knothash

import "fmt"

func KnotHash(alen int, lengths []int) []int {
	arr := initArray(alen)
	var start, end, skip int
	for _, l := range lengths {
		end = start + l
		for i := 0; i < (end-start)/2; i++ {
			t := arr[(start+i)%len(arr)]
			arr[(start+i)%len(arr)] = arr[(end-i-1)%len(arr)]
			arr[(end-i-1)%len(arr)] = t
		}
		start = end + skip
		skip++
		fmt.Println(l, "=>", arr)
	}
	return arr
}

func initArray(alen int) []int {
	arr := make([]int, alen)
	for i := 0; i < alen; i++ {
		arr[i] = i
	}
	return arr
}

func SparseHash(in string) []int {
	b := []byte(in)

	return nil
}
