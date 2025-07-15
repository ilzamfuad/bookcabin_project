package utils

import "math/rand"

func RandomizeSlice(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

	return slice
}
