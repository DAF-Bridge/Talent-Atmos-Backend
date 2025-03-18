package handler

import (
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary Create a new user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 201 {object} models.User
// @Failure 400 {object} fiber.Map "Bad request - Invalid user object"
// @Failure 500 {object} fiber.Map "Internal server error - Failed to create user"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// @Summary List all users
// @Description List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} fiber.Map "Internal server error - Internal Server Error"
// @Router /users [get]
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

// @Summary Get current user profile
// @Description Get current user profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Profile
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/me [get]
func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
	userData, ok := c.Locals("user").(jwt.MapClaims)
	// fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

	if !ok {
		logs.Error("Failed to get user data")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// Access the user_id
	userID, ok := userData["user_id"].(string) // JSON numbers are parsed as string

	if !ok {
		logs.Error("Failed to get user_id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id to uuid"})
	}
	// println(userID)

	// Convert user_id to uint
	currentUserID, err := uuid.Parse(userID)
	if err != nil {
		// println(err.Error())
		logs.Error("Failed to parse user_id")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	currentUserProfile, err := h.service.GetCurrentUserProfile(currentUserID)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to get current user profile: %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(currentUserProfile)
}

// @Summary Upload profile picture
// @Description Upload profile picture for a user
// @Tags Users
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "User ID"
// @Param image formData file true "Profile picture file"
// @Success 200 {object} dto.UploadResponse "Uploaded image URL"
// @Failure 400 {object} fiber.Map "Bad request - Invalid file/user ID"
// @Failure 500 {object} fiber.Map "Internal server error - Failed to update profile picture"
// @Router /users/{id}/upload-profile [post]
func (h *UserHandler) UploadProfilePicture(c *fiber.Ctx) error {
	claims, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	fileHeader, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid file"})
	}

	src, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer src.Close()

	// Upload to S3 & Update DB
	picURL, err := h.service.UpdateUserPicture(c.Context(), userID, src, fileHeader)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to update profile picture: %v", err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update profile picture"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"picUrl": picURL})
}

type UserPreferenceHandler struct {
	service service.UserPreferenceService
}

func NewUserPreferenceHandler(service service.UserPreferenceService) *UserPreferenceHandler {
	return &UserPreferenceHandler{service: service}
}

func (h *UserPreferenceHandler) CreateUserPreference(c *fiber.Ctx) error {
	var userPreference dto.UserPreferenceRequest
	if err := utils.ParseJSONAndValidate(c, &userPreference); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	parsedUserID, err := uuid.Parse(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.service.CreateUserPreference(parsedUserID, userPreference); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User preference created successfully"})
}

func (h *UserPreferenceHandler) GetUserPreferenceByUserID(c *fiber.Ctx) error {
	user, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	parsedUserID, err := uuid.Parse(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	userPreference, err := h.service.GetUserPreference(parsedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(userPreference)
}

func (h *UserPreferenceHandler) UpdateUserPreference(c *fiber.Ctx) error {
	var userPreference dto.UserPreferenceRequest
	if err := utils.ParseJSONAndValidate(c, &userPreference); err != nil {
		return err
	}

	user, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	parsedUserID, err := uuid.Parse(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	updatedUserPreference, err := h.service.UpdateUserPreference(parsedUserID, userPreference)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(updatedUserPreference)
}

func (h *UserPreferenceHandler) DeleteUserPreference(c *fiber.Ctx) error {
	user, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	parsedUserID, err := uuid.Parse(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	if err := h.service.DeleteUserPreference(parsedUserID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User preference deleted successfully"})
}
