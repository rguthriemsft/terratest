// +build azure azureslim,network

// NOTE: We use build tags to differentiate azure testing because we currently do not have azure access setup for
// CircleCI.

package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
The below tests are currently stubbed out, with the expectation that they will throw errors.
If/when methods can be mocked or Create/Delete APIs are added, these tests can be extended.
*/

func TestGetPublicIPAddressE(t *testing.T) {
	t.Parallel()

	paName := ""
	rgName := ""
	subID := ""

	_, err := GetPublicIPAddressE(t, paName, rgName, subID)

	require.Error(t, err)
}

func TestCheckPublicDNSNameAvailabilityE(t *testing.T) {
	t.Parallel()

	location := ""
	domain := ""
	subID := ""

	_, err := CheckPublicDNSNameAvailabilityE(t, location, domain, subID)

	require.Error(t, err)
}

func TestGetPublicAddressIPE(t *testing.T) {
	t.Parallel()

	paName := ""
	rgName := ""
	subID := ""

	_, err := GetPublicAddressIPE(t, paName, rgName, subID)

	require.Error(t, err)
}

func TestPublicAddressExistsE(t *testing.T) {
	t.Parallel()

	paName := ""
	rgName := ""
	subID := ""

	_, err := PublicAddressExistsE(t, paName, rgName, subID)

	require.Error(t, err)
}