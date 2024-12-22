package service

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/golang-jwt/jwt/v5"
)

type OauthService struct {
	userRepo *repository.UserRepository
	profileRepo *repository.ProfileRepository
	jwtSecret string
}

func NewOauthService(userRepo *repository.UserRepository, profileRepo *repository.ProfileRepository ,jwtSecret string) *OauthService {
	return &OauthService{userRepo: userRepo, profileRepo: profileRepo ,jwtSecret: jwtSecret}
}

func (s * OauthService) AuthenticateUser(name, email, provider, providerID string) (string, error) {

	// Start a new transaction
	tx := s.userRepo.BeginTransaction()

	// Always defer rollback in case something goes wrong
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // Rollback in case of a panic
		}
	}()

	user := &domain.User{
		Name: 		name, 
		Email: 		email, 
		Provider: 	domain.Provider(provider), 
		ProviderID: providerID, 
	}

	// check if email is already taken
	if existedUser,err := s.userRepo.FindByEmail(email); err == nil {
		user.ID = existedUser.ID
		return s.generateJWT(user)
	}

	fname, lname := utils.SeparateName(name)

	profile := &domain.Profile{
		FirstName: fname, 
		LastName: lname, 
		Email: email, 
		Phone: "", 
	}
	
	// Start the transaction for creating the user and profile
	if err := s.userRepo.Create(user); err != nil {
		tx.Rollback() // Rollback if user creation fails
		return "", err
	}

	profile.UserID = user.ID

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

func (s *OauthService) generateJWT(user *domain.User) (string, error) {
	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	})
	return token.SignedString([]byte(s.jwtSecret))
}