// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when CRUD methods are introduced for Azure Virtual Machines, these tests can be extended
(see AWS S3 tests for reference).
*/

func TestLogAnalyticsWorkspace(t *testing.T) {
	t.Parallel()

	result := azure.LogAnalyticsWorkspaceExists("", "", "")
	assert.False(t, result)
}

func TestLogAnalyticsSku(t *testing.T) {
	t.Parallel()

	result := azure.GetLogAnalyticsWorkspaceSku("", "", "")
	assert.Equal(t, "", result)
}

func TestLogAnalyticsRetentionPeriodDays(t *testing.T) {
	t.Parallel()

	result := azure.GetLogAnalyticsWorkspaceRetentionPeriodDays("", "", "")
	assert.Equal(t, int32(-1), result)
}
