package service

import (
	"errors"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repository.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s *AuthService) SignUp(name , email, password string) (string, error) {
	// check if email is already taken
	if _,err := s.userRepo.FindByEmail(email); err == nil {
		return "", errors.New("email already taken")
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Create User
	user := &domain.User{
		Name: name, 
		Email: email, 
		Password: string(hashedPassword), 
		Provider: "local", 
	}

	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	// fmt.Println(user.ID)
	// Generate JWT
	return s.generateJWT(user)
}

func (s *AuthService) LogIn(email, password string) (string, error) {
	// Find User
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Check Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}