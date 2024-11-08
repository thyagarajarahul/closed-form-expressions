package expression

import (
	"closed-form-expressions/math"
	"fmt"
)

func SumOddNumbers(n int64) string {
	var sum = math.Square(n)

	return fmt.Sprintf("The sum of the first %d odd numbers is: %d\n", n, int(sum))
}
