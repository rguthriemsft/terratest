package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
