package service

import (
	"errors"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/customerrors"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	profileRepo *repository.ProfileRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, profileRepo *repository.ProfileRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, profileRepo: profileRepo ,jwtSecret: jwtSecret}
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
		return "", customerrors.ErrEmailAlreadyRegistered
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedPasswordString := string(hashedPassword) // Convert []byte to string

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: &hashedPasswordString,
		Provider: domain.ProviderLocal,
	}

	// Create User
	if err := s.userRepo.Create(user); err != nil {
		tx.Rollback()
		return "", err
	}

	fname, lname := utils.SeparateName(name)

	profile := &domain.Profile{
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
		return "", err
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit().Error; err != nil {
		tx.Rollback() // Rollback if commit fails
		return "", err
	}

	// Generate JWT
	return s.generateJWT(user)
}

func (s *AuthService) LogIn(email, password string) (string, error) {
	// Find User
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	passwordStr := string(*user.Password)	// Convert *string to string

	// Check Password
	if err := bcrypt.CompareHashAndPassword([]byte(passwordStr), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	// fmt.Println(user.ID)
	// Generate JWT
	return s.generateJWT(user)

}

// Private Methods for local use
func (s *AuthService) generateJWT(user *domain.User) (string, error) {
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),  // 7 days
	})
	return token.SignedString([]byte(s.jwtSecret))
}
