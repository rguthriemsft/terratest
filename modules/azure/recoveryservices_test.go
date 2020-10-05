package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods to create and delete recovery services resources are added, these tests can be extended.
*/

func TestRecoveryServicesVaultName(t *testing.T) {
	_, err := GetRecoveryServicesVaultE("", "", "")
	require.Error(t, err, "vault")
}

func TestRecoveryServicesVaultExists(t *testing.T) {
	_, err := RecoveryServicesVaultExistsE("", "", "")
	require.Error(t, err, "vault exists")
}

func TestRecoveryServicesVaultBackupPolicyList(t *testing.T) {
	_, err := GetRecoveryServicesVaultBackupPolicyListE("", "", "")
	require.Error(t, err, "Backup policy list not faulted")
}

func TestRecoveryServicesVaultBackupProtectedVMList(t *testing.T) {
	_, err := GetRecoveryServicesVaultBackupProtectedVMListE("", "", "", "")
	require.Error(t, err, "Backup policy protected vm list not faulted")
}
