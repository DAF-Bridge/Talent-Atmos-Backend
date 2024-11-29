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
	jwtSecret string
}

func NewOauthService(userRepo *repository.UserRepository, jwtSecret string) *OauthService {
	return &OauthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (s * OauthService) AuthenticateUser(name, email, provider, providerID, picUrl string) (string, error) {

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
		PicUrl: picUrl, 
	}
	
	// Start the transaction for creating the user and profile
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback() // Rollback if user creation fails
		return "", err
	}

	profile.UserID = user.ID

	// Create the profile
	if err := tx.Create(profile).Error; err != nil {
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
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(s.jwtSecret))
}