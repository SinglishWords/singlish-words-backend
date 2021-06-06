package util

import (
	"math/rand"
)

// ShufflePickN
// Randomly Pick (at most) n unique values from the given slice
// Time complexity: O(n), where n is the parameter n
// Space complexity: O(1), inplace modify the slices
//
// Important note:
// This function will shuffle the slice.
// And it won't totally shuffle the slice.
// It only pick n random entries, so at most shuffle n*2 entries.
func ShufflePickN(in interface{}, n int) ([]interface{}, bool) {
	arr, success := TakeSliceArg(in)
	if !success {
		return nil, success
	}

	length := len(arr)
	if n > length {
		n = length
	}

	for n > 0 {
		r := rand.Intn(length)
		arr[r], arr[length-1] = arr[length-1], arr[r]
		length--
		n--
	}

	return arr[length:], true
}
