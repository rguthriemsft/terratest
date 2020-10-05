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

	_, err := LogAnalyticsWorkspaceExistsE("fake", "", "")
	assert.Error(t, err, "Workspace")
}

func TestLogAnalyticsSku(t *testing.T) {
	t.Parallel()

	_, err := GetLogAnalyticsWorkspaceSkuE("fake", "", "")
	assert.Error(t, err, "Sku")
}

func TestLogAnalyticsRetentionPeriodDays(t *testing.T) {
	t.Parallel()

	_, err := GetLogAnalyticsWorkspaceRetentionPeriodDaysE("fake", "", "")
	assert.Error(t, err, "RetentionPeriod")
}
