package service

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type OauthService struct {
	userRepo *repository.UserRepository
	jwtSecret string
}

func NewOauthService(userRepo *repository.UserRepository, jwtSecret string) *OauthService {
	return &OauthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s * OauthService) AuthenticateUser(name, email, provider, providerID string) (string, error) {
	user := &domain.User{
		Name: 		name, 
		Email: 		email, 
		Provider: 	provider, 
		ProviderID: providerID, 
	}

	// check if email is already taken
	if _,err := s.userRepo.FindByEmail(email); err == nil {
		return s.generateJWT(user)
	}
	
	// Create User if not exists
	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}
	
	// Generate JWT
	return s.generateJWT(user)
}

func (s *OauthService) generateJWT(user *domain.User) (string, error) {
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}