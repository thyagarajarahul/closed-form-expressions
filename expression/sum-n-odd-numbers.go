package expression

import (
	"fmt"
	"math"
)

func SumOddNumbers(n int) string {
	var sum = math.Pow(float64(n), 2)

	return fmt.Sprintf("The sum of the first %d odd numbers is: %d\n", n, int(sum))
}
