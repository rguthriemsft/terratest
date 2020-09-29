package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when CRUD methods are introduced for Azure Virtual Machines, these tests can be extended
(see AWS S3 tests for reference).
*/

func TestLogAnalyticsWorkspace(t *testing.T) {
	t.Parallel()

	result := LogAnalyticsWorkspaceExists("", "", "")
	assert.False(t, result)
}

func TestLogAnalyticsSku(t *testing.T) {
	t.Parallel()

	result := GetLogAnalyticsWorkspaceSku("", "", "")
	assert.Equal(t, "", result)
}

func TestLogAnalyticsRetentionPeriodDays(t *testing.T) {
	t.Parallel()

	result := GetLogAnalyticsWorkspaceRetentionPeriodDays("", "", "")
	assert.Equal(t, int32(-1), result)
}
