package internal

import (
	"strconv"
)

func GenerateFizzBuzz(int1, int2, limit uint, str1, str2 string) []string {
	result := make([]string, limit)

	for i := uint(1); i <= limit; i++ {
		div1 := i%int1 == 0
		div2 := i%int2 == 0

		switch {
		case div1 && div2:
			result[i-1] = str1 + str2
		case div1:
			result[i-1] = str1
		case div2:
			result[i-1] = str2
		default:
			result[i-1] = strconv.FormatUint(uint64(i), 10)
		}
	}
	return result
}
