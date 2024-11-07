package expression

import (
	"fmt"
)

func SumNaturalNumbers(n int) string {
	var sum = (n * (n + 1)) / 2

	return fmt.Sprintf("The sum of the first %d natural numbers is: %d\n", n, sum)
}
