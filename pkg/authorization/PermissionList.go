package authorization

import (
	"sync"
)

// Define a private map and a sync.Once instance for lazy initialization
var permissionsList map[string][]string
var once sync.Once

// Read-only function to initialize permissionsList (called only once)
func initializePermissionsList() {
	permissionsList = map[string][]string{
		"Role":      {"read", "create", "delete", "edit"},
		"Employees": {"read", "add", "remove", "edit"},
		"Events":    {"read", "create", "delete", "edit"},
		"Job":       {"read", "create", "delete", "edit"},
	}
}

// GetPermissions returns a copy of permissions for a specific resource
func GetPermissions(key string) ([]string, bool) {
	// Ensure the map is initialized
	once.Do(initializePermissionsList)

	permissions, exists := permissionsList[key]
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
	for key, permissions := range permissionsList {
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

	resources := make([]string, 0, len(permissionsList))
	for resource := range permissionsList {
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
