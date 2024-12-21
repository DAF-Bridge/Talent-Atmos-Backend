package handler

import (
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type PolicyHandler struct {
	policyRoleService *service.PolicyRoleService
}

func NewPolicyHandler(policyService *service.PolicyRoleService) *PolicyHandler {
	return &PolicyHandler{policyRoleService: policyService}
}

func (p *PolicyHandler) AddPolicyForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Policy form Json Body
	var policy = struct {
		Role string `json:"role"`
		domain.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.AddPolicyForRoleInDomain(policy.Role, fmt.Sprint(orgID), policy.Obj, policy.Action)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) AddPoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Policy form Json Body
	var policies = struct {
		Role     string          `json:"role"`
		Policies []domain.Policy `json:"policies"`
	}{}
	if err := c.BodyParser(&policies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.AddPoliciesForRoleInDomain(policies.Role, fmt.Sprint(orgID), policies.Policies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) DeletePolicyForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Policy form Json Body
	var policy = struct {
		Role string `json:"role"`
		domain.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.DeletePolicyForRoleInDomain(policy.Obj, fmt.Sprint(orgID), policy.Action, policy.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) DeletePoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Policy form Json Body
	var policies = struct {
		Role     string          `json:"role"`
		Policies []domain.Policy `json:"policies"`
	}{}
	if err := c.BodyParser(&policies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.DeletePoliciesForRoleInDomain(policies.Role, fmt.Sprint(orgID), policies.Policies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) GetPoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Role form Json Body
	var role = struct {
		Role string `json:"role"`
	}{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	policies, err := p.policyRoleService.GetPoliciesForRoleInDomain(role.Role, fmt.Sprint(orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"policies": policies})
}

func (p *PolicyHandler) GetRolesForPolicyInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	// Policy form Json Body
	var policy = struct {
		domain.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	roles, err := p.policyRoleService.GetRolesForPolicyInDomain(fmt.Sprint(orgID), policy.Obj, policy.Action)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"roles": roles})
}
