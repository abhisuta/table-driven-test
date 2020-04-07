package calculator_test

import (
	"fmt"
	"testing"

	"github.com/abhisuta/table-driven-test/calculator"
	"github.com/stretchr/testify/require"
)

var testcases = []struct {
	name     string
	inputx   string
	inputy   string
	expected int64
	err      error
}{
	{"zero", "0", "0", 0, nil},
	{"positive numbers", "1", "2", 3, nil},
	{"negative numbers positive result", "-1", "2", 1, nil},
	{"negative numbers negative result", "-1", "-2", -3, nil},
	{"invalid x", "a", "2", 0, fmt.Errorf(`converting x failed: strconv.ParseInt: parsing "c": invalid syntax`)},
	{"invalid y", "1", "b", 0, fmt.Errorf(`converting y failed: strconv.ParseInt: parsing "b": invalid syntax`)},
}

func TestAdd(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := calculator.Add(tc.inputx, tc.inputy)

			if tc.err == nil {
				require.NoError(tt, err)
			} else {
				require.EqualError(tt, tc.err, err.Error())
			}
			require.Equal(tt, tc.expected, actual)
		})
	}
}
