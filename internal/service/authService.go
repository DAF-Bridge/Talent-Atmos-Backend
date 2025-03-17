package service

import (
	"fmt"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo    repository.UserRepository
	profileRepo *repository.ProfileRepository
	jwtSecret   string
}

func NewAuthService(userRepo repository.UserRepository, profileRepo *repository.ProfileRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, profileRepo: profileRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) SignUp(name, email, password, phone string) (string, error) {

	// Begin Transaction
	tx := s.userRepo.BeginTransaction()

	// Always defer rollback in case something goes wrong
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback in case of a panic
		}
	}()

	// check if email is already taken
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		logs.Error("Email already registered")
		return "", errs.NewConflictError("email already registered")
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logs.Error("Failed to hash password")
		return "", errs.NewUnexpectedError()
	}

	hashedPasswordString := string(hashedPassword) // Convert []byte to string

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: &hashedPasswordString,
		Provider: models.ProviderLocal,
	}

	// Create User
	if err := s.userRepo.Create(user); err != nil {
		tx.Rollback()
		logs.Error("Failed to create user")
		return "", errs.NewConflictError(err.Error())
	}

	fname, lname := utils.SeparateName(name)

	profile := &models.Profile{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Phone:     phone,
		PicUrl:    "",
		UserID:    user.ID,
	}

	// Create the profile
	if err := s.profileRepo.Create(profile); err != nil {
		tx.Rollback() // Rollback if profile creation fails
		logs.Error("Failed to create profile")
		return "", errs.NewUnexpectedError()
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // Rollback if commit fails
		logs.Error("Failed to commit create user transaction")
		return "", errs.NewUnexpectedError()
	}

	// Generate JWT
	userJWT, err := s.generateJWT(user)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to generate JWT: %v", err))
		return "", errs.NewUnexpectedError()
	}

	return userJWT, nil
}

func (s *AuthService) LogIn(email, password string) (string, error) {
	// Find User
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to find user: %v", err))
		return "", errs.NewUnauthorizedError("invalid email or password")
	}

	// Check if user is not login with local Username and Password
	if user.Provider != models.ProviderLocal || user.Password == nil {
		logs.Error("User is not registered with local username and password")
		return "", errs.NewForbiddenError("User is not registered with Username and Password. Please log in using the other method.")
	}

	passwordStr := *user.Password // Convert *string to string

	// Check Password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordStr), []byte(password)); err != nil {
		logs.Error("Invalid email or password")
		return "", errs.NewUnauthorizedError("invalid email or password")
	}

	// fmt.Println(user.ID)
	// Generate JWT
	userJWT, err := s.generateJWT(user)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to generate JWT: %v", err))
		return "", errs.NewUnexpectedError()
	}

	return userJWT, nil
}

// Private Methods for local use
func (s *AuthService) generateJWT(user *models.User) (string, error) {
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	})
	return token.SignedString([]byte(s.jwtSecret))
}
