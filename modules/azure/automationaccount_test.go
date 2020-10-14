// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAutomationAccountClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}
