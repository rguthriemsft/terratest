package k8s

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnmarshalJSONPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		jsonBlob    string
		jsonPath    string
		expectedOut interface{}
	}{
		{
			"boolField",
			`{"key": true}`,
			"{ .key }",
			[]bool{true},
		},
		{
			"nestedObject",
			`{"key": {"data": [1,2,3]}}`,
			"{ .key }",
			[]map[string][]int{
				map[string][]int{
					"data": []int{1, 2, 3},
				},
			},
		},
		{
			"nestedArray",
			`{"key": {"data": [1,2,3]}}`,
			"{ .key.data[*] }",
			[]int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		// capture range variable so that it doesn't update when the subtest goroutine swaps.
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			var output interface{}
			UnmarshalJSONPath(t, []byte(testCase.jsonBlob), testCase.jsonPath, &output)
			// NOTE: we have to do equality check on the marshalled json data to allow equality checks over dynamic
			// types in this table driven test.
			expectedOutJSON, err := json.Marshal(testCase.expectedOut)
			require.NoError(t, err)
			actualOutJSON, err := json.Marshal(output)
			require.NoError(t, err)
			assert.Equal(t, string(expectedOutJSON), string(actualOutJSON))
		})
	}
}
