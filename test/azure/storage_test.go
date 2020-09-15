// +build azure

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/stretchr/testify/require"
)

func TestStorageAccountExists(t *testing.T) {
	_, err := azure.StorageAccountExistsE("", "", "")
	require.Error(t, err)
}

func TestStorageBlobContainerExists(t *testing.T) {
	_, err := azure.StorageBlobContainerExistsE("", "", "", "")
	require.Error(t, err)
}

func TestStorageBlobContainerPublicAccess(t *testing.T) {
	_, err := azure.GetStorageBlobContainerPublicAccessE("", "", "", "")
	require.Error(t, err)
}

func TestGetStorageAccountKind(t *testing.T) {
	_, err := azure.GetStorageAccountKindE("", "", "")
	require.Error(t, err)
}

func TestGetStorageAccountSkuTier(t *testing.T) {
	_, err := azure.GetStorageAccountSkuTierE("", "", "")
	require.Error(t, err)
}

func TestGetStorageDNSString(t *testing.T) {
	_, err := azure.GetStorageDNSStringE("", "", "")
	require.Error(t, err)
}
