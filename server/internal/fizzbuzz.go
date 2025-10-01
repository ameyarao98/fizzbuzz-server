package internal

import (
	"errors"
	"strconv"
)

func GenerateFizzBuzz(int1, int2, limit int, str1, str2 string) ([]string, error) {
	// Validation
	if int1 <= 0 {
		return nil, errors.New("int1 must be greater than zero")
	}
	if int2 <= 0 {
		return nil, errors.New("int2 must be greater than zero")
	}
	if limit <= 0 {
		return nil, errors.New("limit must be greater than zero")
	}
	if str1 == "" {
		return nil, errors.New("str1 cannot be empty")
	}
	if str2 == "" {
		return nil, errors.New("str2 cannot be empty")
	}
	lcm := (int1 * int2) / gcd(int1, int2)

	result := make([]string, limit)

	for i := 1; i <= limit; i++ {
		switch {
		case i%lcm == 0:
			result[i-1] = str1 + str2
		case i%int1 == 0:
			result[i-1] = str1
		case i%int2 == 0:
			result[i-1] = str2
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result, nil
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
