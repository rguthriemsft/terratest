package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete log analytics resources are added, these tests can be extended.
*/

func TestLogAnalyticsWorkspace(t *testing.T) {
	t.Parallel()

	result := LogAnalyticsWorkspaceExists(t, "fake", "", "")
	assert.False(t, result)
}

func TestLogAnalyticsSku(t *testing.T) {
	t.Parallel()

	result := GetLogAnalyticsWorkspaceSku(t, "fake", "", "")
	assert.Equal(t, "", result)
}

func TestLogAnalyticsRetentionPeriodDays(t *testing.T) {
	t.Parallel()

	result := GetLogAnalyticsWorkspaceRetentionPeriodDays(t, "fake", "", "")
	assert.Equal(t, int32(-1), result)
}
