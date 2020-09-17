package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoveryServicesVaultName(t *testing.T) {
	exists := RecoveryServicesVaultExists("", "", "")
	assert.False(t, exists, "vault does not exist")
}

func TestRecoveryServicesVaultBackupPolicyList(t *testing.T) {
	_, err := GetRecoveryServicesVaultBackupPolicyListE("", "", "")
	require.Error(t, err, "Backup policy list not faulted")
}

func TestRecoveryServicesVaultBackupProtectedVMList(t *testing.T) {
	_, err := GetRecoveryServicesVaultBackupProtectedVMListE("", "", "", "")
	require.Error(t, err, "Backup policy protected vm list not faulted")
}
