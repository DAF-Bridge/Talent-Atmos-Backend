package authorization

import (
	"sync"
)

// Define a private map and a sync.Once instance for lazy initialization
var normalPermissionsList map[string][]string
var RoleManagerPermissionsList map[string][]string
var OrganizationAdminPermissionsList map[string][]string
var SystemAdminPermissionsList map[string][]string
var once sync.Once

var specialRole []string

// Read-only function to initialize permissionsList (called only once)
func initializePermissionsList() {
	normalPermissionsList = map[string][]string{
		"Employees": {"read", "add", "remove", "edit"},
		"Events":    {"read", "create", "delete", "edit"},
		"Job":       {"read", "create", "delete", "edit"},
	}
	SystemAdminPermissionsList = map[string][]string{
		"Domain": {"read", "create", "delete", "edit"},
	}
	RoleManagerPermissionsList = map[string][]string{
		"Role": {"read", "create", "delete", "edit"},
	}

	specialRole = []string{"Role Manager", "System Admin", "Organization Admin"}
}

// GetPermissions returns a copy of permissions for a specific resource
func GetPermissions(key string) ([]string, bool) {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	permissions, exists := normalPermissionsList[key]
	if !exists {
		return nil, false
	}
	// Return a copy of the permissions slice to ensure immutability
	copyPermissions := make([]string, len(permissions))
	copy(copyPermissions, permissions)
	return copyPermissions, true
}

// GetAllPermissions returns a copy of the entire permissions list
func GetAllPermissions() map[string][]string {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	// Create a new map to return a copy
	copyPermissionsList := make(map[string][]string)
	for key, permissions := range normalPermissionsList {
		// Copy each slice of permissions
		copyPermissions := make([]string, len(permissions))
		copy(copyPermissions, permissions)
		copyPermissionsList[key] = copyPermissions
	}
	return copyPermissionsList
}

// GetAllResources returns a slice of all resource names (keys)
func GetAllResources() []string {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	resources := make([]string, 0, len(normalPermissionsList))
	for resource := range normalPermissionsList {
		resources = append(resources, resource)
	}
	return resources
}

// CheckAction checks if a specific action is allowed for a resource
func CheckAction(resource string, action string) bool {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	permissions, exists := GetPermissions(resource)
	if !exists {
		return false
	}
	for _, permission := range permissions {
		if permission == action {
			return true
		}
	}
	return false
}

// GetActions returns the actions (permissions) for a resource
func GetActions(resource string) []string {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	permissions, exists := GetPermissions(resource)
	if !exists {
		return nil
	}
	return permissions
}

func GetSpecialRole() []string {
	return specialRole
}
