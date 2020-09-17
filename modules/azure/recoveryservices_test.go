package azure

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoveryServicesVaultName(t *testing.T) {
	exists := azure.RecoveryServicesVaultExists("", "", "")
	assert.False(t, exists, "vault does not exist")
}

func TestRecoveryServicesVaultBackupPolicyList(t *testing.T) {
	_, err := azure.GetRecoveryServicesVaultBackupPolicyListE("", "", "")
	require.Error(t, err, "Backup policy list not faulted")
}

func TestRecoveryServicesVaultBackupProtectedVMList(t *testing.T) {
	_, err := azure.GetRecoveryServicesVaultBackupProtectedVMListE("", "", "", "")
	require.Error(t, err, "Backup policy protected vm list not faulted")
}
