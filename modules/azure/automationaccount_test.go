// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete network resources are added, these tests can be extended.
*/

func TestGetAutomationAccountClientE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	client, err := GetAutomationAccountClientE(subscriptionID)

	require.NoError(t, err)
	assert.NotEmpty(t, *client)
}

func TestGetAutomationAccountE(t *testing.T) {
	t.Parallel()

	automationAccountName := ""
	resourceGroupName := ""
	subscriptionID := ""

	_, err := GetAutomationAccountE(t, automationAccountName, resourceGroupName, subscriptionID)

	require.NoError(t, err)
}
