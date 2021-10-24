package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SelectionSort(t *testing.T) {
	tt := []struct {
		name      string
		data      []int
		expectedR []int
	}{
		{
			name:      "should return valid result",
			data:      []int{9, 28, 54, 1, 6, 7, 2},
			expectedR: []int{1, 2, 6, 7, 9, 28, 54},
		},
		{
			name:      "shoudl return valid result2",
			data:      []int{1, 2, 3, 4, 5},
			expectedR: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(*testing.T) {
			r := SelctionSort(tc.data)
			assert.Equal(t, tc.expectedR, r)
		})
	}
}
