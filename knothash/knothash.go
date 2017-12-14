package knothash

import "bytes"
import "fmt"

// KnotHash returns a result of a single round of knot hash function
func KnotHash(arrlen int, lengths []byte) []byte {
	arr := initArray(arrlen)
	return knotHash(arr, lengths, 1)
}

func initArray(arrlen int) []byte {
	arr := make([]byte, arrlen)
	for i := 0; i < arrlen; i++ {
		arr[i] = byte(i)
	}
	return arr
}

func knotHash(arr []byte, lengths []byte, rounds int) []byte {
	var start, end, skip int
	for i := 0; i < rounds; i++ {
		for _, l := range lengths {
			end = start + int(l)
			for i := 0; i < (end-start)/2; i++ {
				t := arr[(start+i)%len(arr)]
				arr[(start+i)%len(arr)] = arr[(end-i-1)%len(arr)]
				arr[(end-i-1)%len(arr)] = t
			}
			start = end + skip
			skip++
		}
	}
	return arr
}

// Hash calculates a sparse hash value running knot hash function 64 times
func Hash(arrlen int, in string) string {
	lengths := []byte(in)
	lengths = append(lengths, 17, 31, 73, 47, 23)

	arr := initArray(arrlen)
	sh := knotHash(arr, lengths, 64)

	var buff []byte
	for i := 0; i < 16; i++ {
		var xor byte
		for _, v := range sh[i*16 : (i+1)*16] {
			xor = xor ^ v
		}
		buff = append(buff, xor)
	}

	var buffer bytes.Buffer
	for _, v := range buff {
		buffer.WriteString(fmt.Sprintf("%02x", v))
	}

	return buffer.String()
}
