package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateFizzBuzz(t *testing.T) {
	t.Run("primes 2 and 3", func(t *testing.T) {
		str1, str2 := "Fizz", "Buzz"
		result := GenerateFizzBuzz(2, 3, 6, str1, str2)

		assert.Len(t, result, 6)
		assert.Equal(t, []string{
			"1",         // 1
			str1,        // 2
			str2,        // 3
			str1,        // 4
			"5",         // 5
			str1 + str2, // 6
		}, result)
	})

	t.Run("non-primes 2 and 4", func(t *testing.T) {
		str1, str2 := "X", "Y"
		result := GenerateFizzBuzz(2, 4, 8, str1, str2)

		assert.Len(t, result, 8)
		assert.Equal(t, []string{
			"1",         // 1
			str1,        // 2
			"3",         // 3
			str1 + str2, // 4
			"5",         // 5
			str1,        // 6
			"7",         // 7
			str1 + str2, // 8
		}, result)
	})

	t.Run("limit of 1", func(t *testing.T) {
		str1, str2 := "A", "B"
		result := GenerateFizzBuzz(2, 3, 1, str1, str2)

		assert.Len(t, result, 1)
		assert.Equal(t, "1", result[0])
	})

}
