package azure

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/require"
)

func TestKeyVaultSecretExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultSecretName := "fakeSecretName"
	_, err := azure.KeyVaultSecretExistsE(testKeyVaultName, testKeyVaultSecretName)
	require.Error(t, err)
}

func TestKeyVaultKeyExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultKeyName := "fakeKeyName"
	_, err := azure.KeyVaultKeyExistsE(testKeyVaultName, testKeyVaultKeyName)
	require.Error(t, err)
}

func TestKeyVaultCertificateExists(t *testing.T) {
	t.Parallel()

	testKeyVaultName := "fakeKeyVault"
	testKeyVaultCertName := "fakeCertName"
	_, err := azure.KeyVaultCertificateExistsE(testKeyVaultName, testKeyVaultCertName)
	require.Error(t, err)
}
