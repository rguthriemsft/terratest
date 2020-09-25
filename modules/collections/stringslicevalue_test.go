package collections

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetSliceLastValue(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName       string
		sliceSource    string
		sliceSeperator string
		expectedReturn string
		expectedError  bool
	}{
		{"longSlice", "this/is/a/long/slash/separated/string/success", "/", "success", false},
		{"shortendSlice", "this/is/a/long/slash/separated", "/", "separated", false},
		{"dashSlice", "this-is-a-long-dash-separated-string-success", "-", "success", false},
		{"seperatorNotPresent", "this-is-a-long-dash-separated-string-success", "/", "", true},
		{"sourceNoSeperator", "noslicepresent", "/", "", true},
		{"emptyStrings", "", "", "", true},
	}

	for _, tc := range testCases {
		testFor := tc //necessary range capture

		t.Run(testFor.testName, func(t *testing.T) {
			actualReturn, err := GetSliceLastValueE(testFor.sliceSource, testFor.sliceSeperator)
			switch testFor.expectedError {
			case true:
				require.Error(t, err)
			case false:
				require.NoError(t, err)
			}
			assert.Equal(t, testFor.expectedReturn, actualReturn)
		})
	}
}

func TestGetSliceIndexValue(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		sliceIndex     int
		expectedReturn string
		expectedError  bool
	}{
		{-1, "", true},
		{0, "this", false},
		{4, "slash", false},
		{7, "success", false},
		{10, "", true},
	}

	sliceSource := "this/is/a/long/slash/separated/string/success"
	sliceSeperator := "/"

	for _, tc := range testCases {
		testFor := tc //necessary range capture

		t.Run(fmt.Sprintf("Index_%v", testFor.sliceIndex), func(t *testing.T) {
			actualReturn, err := GetSliceIndexValueE(sliceSource, sliceSeperator, testFor.sliceIndex)
			switch testFor.expectedError {
			case true:
				require.Error(t, err)
			case false:
				require.NoError(t, err)
			}
			assert.Equal(t, testFor.expectedReturn, actualReturn)
		})
	}
}
