package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsStraight(t *testing.T) {

	cases := []struct {
		message  string
		value    Cards
		expected bool
	}{
		{"Empty, should be minimum 5 cards", Cards{}, false},
		{"3 Cars, invalid hand, minimum 5", Cards{7, 3, 2}, false},
		{"Invalid hand, maximum 7 cards", Cards{7, 7, 12, 11, 3, 4, 14, 5}, false},
		{"5 cards valid, low values", Cards{2, 3, 4, 5, 6}, true},
		{"5 cards valid, high values", Cards{9, 10, 11, 12, 13}, true},
		{"'As' in the beginning, valid", Cards{14, 5, 4, 2, 3}, true},
		{"'As' in the end, valid", Cards{14, 13, 12, 10, 11}, true},
		{"7 cards, invalid", Cards{7, 7, 12, 11, 3, 4, 14}, false},
		{"Duplicated cards, invalid", Cards{7, 7, 12, 11, 3, 4, 14}, false},
		{"Duplicated cards, valid", Cards{6, 5, 5, 4, 3, 2, 2}, true},
		{"Special case, Q + K + AS + 2 + 3, invalid", Cards{12, 13, 14, 2, 3}, false},
	}

	for _, item := range cases {
		actual := IsStraight(item.value)
		assert.Equal(t, item.expected, actual, item.message)
	}
}
