package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when CRUD methods are introduced for Azure Virtual Machines, these tests can be extended
(see AWS S3 tests for reference).
*/

func TestKeyVaultSecretExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultSecretName := "fakeSecretName"
	_, err := KeyVaultSecretExistsE(testKeyVaultName, testKeyVaultSecretName)
	require.Error(t, err)
}

func TestKeyVaultKeyExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultKeyName := "fakeKeyName"
	_, err := KeyVaultKeyExistsE(testKeyVaultName, testKeyVaultKeyName)
	require.Error(t, err)
}

func TestKeyVaultCertificateExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultCertName := "fakeCertName"
	_, err := KeyVaultCertificateExistsE(testKeyVaultName, testKeyVaultCertName)
	require.Error(t, err)
}
