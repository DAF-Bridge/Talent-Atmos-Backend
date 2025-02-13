package dto

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type RoleResponse struct {
	Role                 string                    `json:"role" example:"owner"`
	OrganizationResponse OrganizationShortResponse `json:"organization" example:"{id: 1, name: 'DAF Bridge'}"`
	UserResponse         UserResponses             `json:"user" example:"{id: 1, username: 'DAF Bridge'}"`
}

func BuildRoleResponse(role models.RoleInOrganization) RoleResponse {
	return RoleResponse{
		Role:                 role.Role,
		OrganizationResponse: BuildOrganizationShortResponse(role.Organization),
		UserResponse:         BuildUserResponses(role.User),
	}

}

type UserWithRoleResponse struct {
	Role         string        `json:"role" example:"owner"`
	UserResponse UserResponses `json:"user" example:"{id: 1, username: 'DAF Bridge'}"`
}

func BuildUserWithRoleResponse(role models.RoleInOrganization) UserWithRoleResponse {
	return UserWithRoleResponse{
		Role:         role.Role,
		UserResponse: BuildUserResponses(role.User),
	}
}

type ListUserWithRoleInOrganizationResponse struct {
	Users []UserWithRoleResponse `json:"users_with_role" example:"[{id: 1, username: 'DAF Bridge', role: 'owner'}]"`
}

func BuildListUserWithRoleInOrganizationResponse(roles []models.RoleInOrganization) ListUserWithRoleInOrganizationResponse {
	var usersWithRole []UserWithRoleResponse
	for _, role := range roles {
		usersWithRole = append(usersWithRole, BuildUserWithRoleResponse(role))
	}
	return ListUserWithRoleInOrganizationResponse{
		Users: usersWithRole,
	}
}
