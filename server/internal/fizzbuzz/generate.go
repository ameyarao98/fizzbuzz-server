package fizzbuzz

import (
	"strconv"
)

func GenerateFizzBuzz(int1, int2, limit uint, str1, str2 string) []string {
	lcm := (int1 * int2) / gcd(int1, int2)

	result := make([]string, limit)

	for i := uint(1); i <= limit; i++ {
		switch {
		case i%lcm == 0:
			result[i-1] = str1 + str2
		case i%int1 == 0:
			result[i-1] = str1
		case i%int2 == 0:
			result[i-1] = str2
		default:
			result[i-1] = strconv.FormatUint(uint64(i), 10)
		}
	}

	return result
}
