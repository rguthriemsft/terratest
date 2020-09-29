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

func TestStorageAccountExists(t *testing.T) {
	_, err := StorageAccountExistsE("", "", "")
	require.Error(t, err)
}

func TestStorageBlobContainerExists(t *testing.T) {
	_, err := StorageBlobContainerExistsE("", "", "", "")
	require.Error(t, err)
}

func TestStorageBlobContainerPublicAccess(t *testing.T) {
	_, err := GetStorageBlobContainerPublicAccessE("", "", "", "")
	require.Error(t, err)
}

func TestGetStorageAccountKind(t *testing.T) {
	_, err := GetStorageAccountKindE("", "", "")
	require.Error(t, err)
}

func TestGetStorageAccountSkuTier(t *testing.T) {
	_, err := GetStorageAccountSkuTierE("", "", "")
	require.Error(t, err)
}

func TestGetStorageDNSString(t *testing.T) {
	_, err := GetStorageDNSStringE("", "", "")
	require.Error(t, err)
}
