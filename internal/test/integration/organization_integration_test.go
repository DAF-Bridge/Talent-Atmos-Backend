//go:build integration

package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/test"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func generateMockJWT(userID uuid.UUID, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   "talentsatmos@gmail.com",
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func TestOrganizationHandlerIntegrationService(t *testing.T) {
	// ARRANGE
	userID := uuid.New()
	jwtSecret := "testSecret"
	mockToken, err := generateMockJWT(userID, jwtSecret)
	assert.NoError(t, err)

	t.Run("TestGetOrganizationByItsID", func(t *testing.T) {
		// ARRANGE
		organizationID := 1
		expected := models.Organization{
			Model:     gorm.Model{ID: 1, UpdatedAt: time.Now()},
			Name:      "Talents Atmos",
			Email:     "talentsatmos@gmail.com",
			Phone:     "+66876428591",
			PicUrl:    "https://talentsatmos.com",
			HeadLine:  "We are the best",
			Specialty: "We are the best",
			Address:   "Chiang Mai University",
			Province:  "Chiang Mai",
			Country:   "TH",
			Latitude:  18.7953,
			Longitude: 98.9523,
			Industries: []*models.Industry{
				{
					Model:    gorm.Model{ID: 18},
					Industry: "Cybersecurity & Data Privacy",
				},
			},
			OrganizationContacts: []models.OrganizationContact{
				{
					Model:          gorm.Model{ID: 1},
					OrganizationID: 1,
					Media:          models.Media("facebook"),
					MediaLink:      "https://facebook.com",
				},
			},
		}

		// Integration interface
		organizationRepo := repository.NewOrganizationRepositoryMock()
		casbinRoleRepository := repository.NewCasbinRoleRepository(initializers.Enforcer)
		organizationService := service.NewOrganizationService(organizationRepo, casbinRoleRepository, initializers.S3)
		organizationHandler := handler.NewOrganizationHandler(organizationService)

		app := fiber.New()
		app.Get("/orgs/get/:id", middleware.AuthMiddleware("testSecret"), organizationHandler.GetOrganizationByID)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/orgs/get/%v", organizationID), nil)
		req.AddCookie(&http.Cookie{
			Name:     "authToken",
			Value:    mockToken,
			Expires:  time.Now().Add(time.Hour),
			HttpOnly: true,
			Path:     "/",
		})

		resp, err := app.Test(req)

		// ACT
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// ASSERT
		if assert.Equal(t, fiber.StatusOK, resp.StatusCode) {
			body, _ := io.ReadAll(resp.Body)

			var actual dto.OrganizationResponse
			expectedResponse := service.ConvertToOrgResponse(expected)

			err := json.Unmarshal(body, &actual)
			assert.NoError(t, err)
			assert.Equal(t, expectedResponse, actual)
		}
	})

	t.Run("TestGetJobByID", func(t *testing.T) {
		// ARRANGE
		organizationID := 1
		jobID := 1

		expected := models.OrgOpenJob{
			Model:          gorm.Model{ID: 1, UpdatedAt: time.Now()},
			Title:          "Software Engineer",
			PicUrl:         "https://talentsatmos.com",
			Scope:          "Software Development",
			Location:       "Chiang Mai",
			Organization:   models.Organization{Name: "Talents Atmos"},
			Workplace:      models.Workplace("remote"),
			WorkType:       models.WorkType("fulltime"),
			CareerStage:    models.CareerStage("entrylevel"),
			Period:         "1 year",
			Description:    "This is a description",
			HoursPerDay:    "8 hours",
			Qualifications: "Bachelor's degree in Computer Science",
			Benefits:       "Health insurance",
			Quantity:       1,
			Salary:         30000,
			Status:         "published",
			Categories: []models.Category{
				{
					Model: gorm.Model{ID: 12},
					Name:  "social",
				},
			},
		}

		// Integration interface
		jobRepo := repository.NewOrgOpenJobRepositoryMock()
		casbinRoleRepository := repository.NewCasbinRoleRepository(initializers.Enforcer)
		jobSrv := service.NewOrgOpenJobService(jobRepo, casbinRoleRepository, test.DB_TEST, initializers.ESClient, initializers.S3)
		jobHandler := handler.NewOrgOpenJobHandler(jobSrv)

		app := fiber.New()
		app.Get("/orgs/:orgID/jobs/get/:id", middleware.AuthMiddleware("testSecret"), jobHandler.GetOrgOpenJobByID)

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/orgs/%v/jobs/get/%v", organizationID, jobID), nil)
		req.AddCookie(&http.Cookie{
			Name:     "authToken",
			Value:    mockToken,
			Expires:  time.Now().Add(time.Hour),
			HttpOnly: true,
		})

		resp, err := app.Test(req)

		// ACT
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		// ASSERT
		if assert.Equal(t, fiber.StatusOK, resp.StatusCode) {
			body, _ := io.ReadAll(resp.Body)

			var actual dto.JobResponses
			expectedResponse := service.ConvertToJobResponse(expected)

			err := json.Unmarshal(body, &actual)
			assert.NoError(t, err)
			assert.Equal(t, expectedResponse, actual)
		}

	})
}
