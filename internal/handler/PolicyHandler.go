package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

type PolicyHandler struct {
	policyRoleService service.EmployeeManagementService
}

func NewPolicyHandler(policyRoleService service.EmployeeManagementService) *PolicyHandler {
	return &PolicyHandler{policyRoleService: policyRoleService}
}

func (p *PolicyHandler) AddPolicyForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization

	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Policy form Json Body
	var policy = struct {
		Role string `json:"role"`
		repository.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.AddPolicyForRoleInDomain(policy.Role, orgID, repository.Policy{Resource: policy.Resource, Action: policy.Action})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) AddPoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Policy form Json Body
	var policies = struct {
		Role     string              `json:"role"`
		Policies []repository.Policy `json:"policies"`
	}{}
	if err := c.BodyParser(&policies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.AddPoliciesForRoleInDomain(policies.Role, orgID, policies.Policies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) DeletePolicyForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization

	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Policy form Json Body
	var policy = struct {
		Role string `json:"role"`
		repository.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.DeletePolicyForRoleInDomain(policy.Role, orgID, repository.Policy{Resource: policy.Resource, Action: policy.Action})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) DeletePoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Policy form Json Body
	var policies = struct {
		Role     string              `json:"role"`
		Policies []repository.Policy `json:"policies"`
	}{}
	if err := c.BodyParser(&policies); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := p.policyRoleService.DeletePoliciesForRoleInDomain(policies.Role, orgID, policies.Policies)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (p *PolicyHandler) GetPoliciesForRoleInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Role form Json Body
	var role = struct {
		Role string `json:"role"`
	}{}
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	policies, err := p.policyRoleService.GetPoliciesForRoleInDomain(role.Role, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"policies": policies})
}

func (p *PolicyHandler) GetRolesForPolicyInDomain(c *fiber.Ctx) error {
	// Access the Organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Policy form Json Body
	var policy = struct {
		repository.Policy
	}{}
	if err := c.BodyParser(&policy); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	roles, err := p.policyRoleService.GetRolesForPolicyInDomain(orgID, policy.Resource, policy.Action)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"roles": roles})
}
