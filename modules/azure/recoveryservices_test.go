package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when CRUD methods are introduced for Azure Virtual Machines, these tests can be extended
(see AWS S3 tests for reference).
*/

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
