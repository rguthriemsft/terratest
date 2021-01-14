package azure

import (
	"context"
	"fmt"
	"testing"

	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	kvmng "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/stretchr/testify/require"
)

// KeyVaultSecretExists indicates whether a key vault secret exists; otherwise false
func KeyVaultSecretExists(t *testing.T, keyVaultName string, secretName string) bool {
	result, err := KeyVaultSecretExistsE(keyVaultName, secretName)
	require.NoError(t, err)
	return result
}

// KeyVaultKeyExists indicates whether a key vault key exists; otherwise false.  This function would fail the test if there is an error.
func KeyVaultKeyExists(t *testing.T, keyVaultName string, keyName string) bool {
	result, err := KeyVaultKeyExistsE(keyVaultName, keyName)
	require.NoError(t, err)
	return result
}

// KeyVaultCertificateExists indicates whether a key vault certificate exists; otherwise false.
// This function would fail the test if there is an error.
func KeyVaultCertificateExists(t *testing.T, keyVaultName string, certificateName string) bool {
	result, err := KeyVaultCertificateExistsE(keyVaultName, certificateName)
	require.NoError(t, err)
	return result
}

// KeyVaultCertificateExistsE indicates whether a certificate exists in key vault; otherwise false.
// This function would fail the test if there is an error.
func KeyVaultCertificateExistsE(keyVaultName, certificateName string) (bool, error) {
	keyVaultSuffix, err := GetKeyVaultURISuffixE()
	if err != nil {
		return false, err
	}
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetCertificateVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		certificateName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// KeyVaultKeyExistsE indicates whether a key exists in the key vault; otherwise false.
// This function would fail the test if there is an error.
func KeyVaultKeyExistsE(keyVaultName, keyName string) (bool, error) {
	keyVaultSuffix, err := GetKeyVaultURISuffixE()
	if err != nil {
		return false, err
	}
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetKeyVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		keyName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// KeyVaultSecretExistsE indicates whether a secret exists in the key vault; otherwise false.
// This function would fail the test if there is an error.
func KeyVaultSecretExistsE(keyVaultName, secretName string) (bool, error) {
	client, err := GetKeyVaultClientE()
	if err != nil {
		return false, err
	}
	keyVaultSuffix, err := GetKeyVaultURISuffixE()
	if err != nil {
		return false, err
	}
	var maxVersionsCount int32 = 1
	versions, err := client.GetSecretVersions(context.Background(),
		fmt.Sprintf("https://%s.%s", keyVaultName, keyVaultSuffix),
		secretName,
		&maxVersionsCount)
	if err != nil {
		return false, err
	}
	items := versions.Values()

	if len(items) > 0 {
		return true, nil
	}
	return false, nil
}

// GetKeyVaultClientE creates a KeyVault client.
// This function would fail the test if there is an error.
func GetKeyVaultClientE() (*keyvault.BaseClient, error) {
	kvClient := keyvault.New()
	authorizer, err := NewKeyVaultAuthorizerE()
	if err != nil {
		return nil, err
	}
	kvClient.Authorizer = *authorizer
	return &kvClient, nil
}

// NewKeyVaultAuthorizerE will return dataplan Authorizer for KeyVault.
// This function would fail the test if there is an error.
func NewKeyVaultAuthorizerE() (*autorest.Authorizer, error) {
	authorizer, err := kvauth.NewAuthorizerFromCLI()
	return &authorizer, err
}

// GetKeyVault is a helper function that gets the keyvault management object.
// This function would fail the test if there is an error.
func GetKeyVault(t *testing.T, resGroupName string, keyVaultName string, subscriptionID string) *kvmng.Vault {
	keyVault, err := GetKeyVaultE(t, resGroupName, keyVaultName, subscriptionID)
	require.NoError(t, err)

	return keyVault
}

// GetKeyVaultE is a helper function that gets the keyvault management object.
// This function would fail the test if there is an error.
func GetKeyVaultE(t *testing.T, resGroupName string, keyVaultName string, subscriptionID string) (*kvmng.Vault, error) {
	// Create akey vault management client
	vaultClient, err := GetKeyVaultManagementClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	//Get the corresponding server client
	keyVault, err := vaultClient.Get(context.Background(), resGroupName, keyVaultName)
	if err != nil {
		return nil, err
	}

	//Return keyvault
	return &keyVault, nil
}

// GetKeyVaultManagementClientE is a helper function that will setup a key vault management client
// This function would fail the test if there is an error.
func GetKeyVaultManagementClientE(subscriptionID string) (*kvmng.VaultsClient, error) {
	// Validate Azure subscription ID
	subscriptionID, err := getTargetAzureSubscription(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create a keyvault management client
	vaultClient := kvmng.NewVaultsClient(subscriptionID)

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}

	// Attach authorizer to the client
	vaultClient.Authorizer = *authorizer

	return &vaultClient, nil
}
