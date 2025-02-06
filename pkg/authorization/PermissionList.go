package authorization

import (
	"sync"
)

// Define a private map and a sync.Once instance for lazy initialization

var once sync.Once

var allRole []string

// Read-only function to initialize permissionsList (called only once)
func initializePermissionsList() {
	allRole = []string{"moderator", "owner", "system_admin"}
}
