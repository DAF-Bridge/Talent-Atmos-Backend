package dto

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

type EditRoleRequest struct {
	Role   string    `json:"role" example:"owner" validate:"required,oneof=owner moderator"`
	UserID uuid.UUID `json:"user_id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910" validate:"required,uuid4"`
}

type RemoveMemberRequest struct {
	UserID uuid.UUID `json:"user_id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910" validate:"required,uuid4"`
}

type RoleResponse struct {
	Role                 string                    `json:"role" example:"owner"`
	OrganizationResponse OrganizationShortResponse `json:"organization" example:"{id: 1, name: 'DAF Bridge' ,picUrl: 'https://example.com/image.jpg',email:'example@gmail.com' ,phone:'0812345678'}"`
	UserResponse         UserResponses             `json:"user" example:"{id: 48a18dd9-48c3-45a5-b4f3-e8d7a60e2910, username: 'DAF Bridge' , email: 'test@example.com' , picUrl: 'https://example.com/image.jpg' , role : 'User' ,updatedAt: '2025-01-24T13:22:10.532645Z'}"`
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
	UserResponse UserResponses `json:"user" example:"{id: 48a18dd9-48c3-45a5-b4f3-e8d7a60e2910, username: 'DAF Bridge' , email: 'test@example.com' , picUrl: 'https://example.com/image.jpg' , role : 'User' ,updatedAt: '2025-01-24T13:22:10.532645Z'}"`
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
