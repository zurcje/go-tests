package ordering

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdering(t *testing.T) {
	numbers := []int{3, 2, 7, 4, 1, 6, 5, 8, 9}
	result := Ordering(numbers)

	expectedResult := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}
