package authorization

// Define a private map and a sync.Once instance for lazy initialization

var allRole []string
var permissionsList [][]string

// Read-only function to initialize permissionsList (called only once)
func init() {
	allRole = []string{"moderator", "owner", "system_admin"}
	moderatorPermissionsMap := map[string][]string{
		"Event":               {"read"},
		"Organization":        {"delete", "update"},
		"OrganizationContact": {"delete", "update", "create"},
		"OrganizationOpenJob": {"delete", "update", "create"},
	}
	moderatorPermissionsList := createCasbinPermissionsList("moderator", moderatorPermissionsMap)
	permissionsList = append(permissionsList, moderatorPermissionsList...)

	ownerPermissionsMap := map[string][]string{
		"Organization": {"create"},
		"Role":         {"remove", "edit", "invite", "read"},
	}
	mergeMapSlice(ownerPermissionsMap, moderatorPermissionsMap)
	ownerPermissionsList := createCasbinPermissionsList("owner", ownerPermissionsMap)
	permissionsList = append(permissionsList, ownerPermissionsList...)
}

func GetPermissionsList() [][]string {
	return permissionsList
}

func GetAllRole() []string {
	return allRole
}

func createCasbinPermissionsList(role string, policy map[string][]string) [][]string {
	CasbinPermissionsList := make([][]string, 0)
	for key, value := range policy {
		for _, v := range value {
			CasbinPermissionsList = append(CasbinPermissionsList, []string{"p", role, key, v, "allow"})
		}
	}
	return CasbinPermissionsList
}

func mergeMapSlice[T comparable, U comparable](dist, src map[T][]U) map[T][]U {
	for k, v := range src {
		dist[k] = append(dist[k], v...)
	}

	return dist

}
