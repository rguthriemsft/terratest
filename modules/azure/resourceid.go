package azure

import "github.com/gruntwork-io/terratest/modules/collections"

// GetNameFromResourceID gets the Name from an Azure Resource ID
func GetNameFromResourceID(resourceID string) string {
	id, err := collections.GetSliceLastValueE(resourceID, "/")
	if err != nil {
		return ""
	}
	return id
}
