package service

import (
	"context"

	"mime/multipart"

	dto "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

type UserService interface {
	CreateUser(user *models.User) error
	ListUsers() ([]dto.UserResponses, error)
	GetCurrentUserProfile(userId uuid.UUID) (*dto.ProfileResponses, error)
	UpdateUserPicture(ctx context.Context, userID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

type UserPreferenceService interface {
	CreateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) error
	GetUserPreference(userID uuid.UUID) (dto.UserPreferenceResponse, error)
	ListUserPreferences() (dto.UserPreferenceTrainingResponses, error)
	ListEventTrainingPreference() (dto.EventTrainingRespoonses, error)
	UpdateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) (dto.UserPreferenceResponse, error)
	DeleteUserPreference(userID uuid.UUID) error
}

type UserInteractService interface {
	IncrementUserInteractForEvent(userID uuid.UUID, eventID uint) error
	FindByUserID(userID uuid.UUID) ([]dto.UserInteractResponse, error)
	GetAll() ([]dto.UserInteractResponse, error)
	FindCategoryByIds(catIDs uint) ([]dto.UserInteractResponse, error)
}

type UserInteractEventService interface {
	IncrementUserInteractForEvent(userID uuid.UUID, eventID uint) error
	FindInteractedEventByUserID(userID uuid.UUID) (*dto.EventsAreInteractedByUserResponse, error)
	GetAll() ([]dto.UserInteractEventResponse, error)
	FindUserCategoriesStatsByUserID(userID uuid.UUID) (*dto.UserInteractCategoriesResponse, error)
	GetAllUserCategoriesStats() ([]dto.UserInteractCategoriesResponse, error)
	GetAllInteractedEventPerUser() ([]dto.EventsAreInteractedByUserResponse, error)
}

func convertToUserResponses(user *models.User) *dto.UserResponses {
	return &dto.UserResponses{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      string(user.Role),
		PicUrl:    user.PicUrl,
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}

// func convertToUserModel(user *dto.SignUpRequest) *models.User {
// 	return &models.User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: &user.Password,
// 		Role:     models.Role(user.Role),
// 		Provider: models.Provider(user.Provider),
// 	}
// }

func convertToProfileResponse(profile *models.Profile) *dto.ProfileResponses {
	user := convertToUserResponses(&profile.User)

	return &dto.ProfileResponses{
		ID:        user.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     user.Email,
		Phone:     profile.Phone,
		Language:  profile.Language,
		PicUrl:    profile.PicUrl,
		Role:      user.Role,
		UpdateAt:  profile.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}

func convertToUserInteractResponse(interact *models.UserInteract) *dto.UserInteractResponse {
	return &dto.UserInteractResponse{
		UserResponses: *convertToUserResponses(&interact.User),
		CategoryResponses: dto.CategoryResponses{
			Value: interact.Category.ID,
			Label: interact.Category.Name,
		},
		Count: interact.Count,
	}
}

func convertToUserInteractEventResponse(UserInteractEvent *models.UserInteractEvent) *dto.UserInteractEventResponse {

	event := UserInteractEvent.Event

	var Categories []dto.CategoryResponses
	for _, category := range event.Categories {
		Categories = append(Categories, dto.CategoryResponses{
			Value: category.ID,
			Label: category.Name,
		})
	}

	return &dto.UserInteractEventResponse{
		UserResponses:  *convertToUserResponses(&UserInteractEvent.User),
		EventResponses: ConvertToEventDocumentResponse(event),
		Count:          UserInteractEvent.Count,
	}

}

func convertToUserInteractCategoryResponse(user *models.User, events []models.Event) *dto.UserInteractCategoriesResponse {
	statCategories := make(map[uint]dto.CategoryWithCountResponse)

	for _, event := range events {
		for _, category := range event.Categories {
			if _, ok := statCategories[category.ID]; !ok {
				statCategories[category.ID] = dto.CategoryWithCountResponse{
					CategoryResponses: dto.CategoryResponses{
						Value: category.ID,
						Label: category.Name,
					},
					Amount: 0,
				}
			} else {
				oldAmount := statCategories[category.ID].Amount
				statCategories[category.ID] = dto.CategoryWithCountResponse{
					CategoryResponses: dto.CategoryResponses{
						Value: category.ID,
						Label: category.Name,
					},
					Amount: oldAmount + 1,
				}
			}

		}
	}

	var categories []dto.CategoryWithCountResponse
	for _, category := range statCategories {
		categories = append(categories, category)
	}

	return &dto.UserInteractCategoriesResponse{
		UserResponses: *convertToUserResponses(user),
		CategoryData:  categories,
		TotalEvents:   uint(len(events)),
	}

}

func convertToListUserInteractCategoryResponse(UserInteractEvent []models.UserInteractEvent) []dto.UserInteractCategoriesResponse {
	userInteractEvents := make(map[uuid.UUID][]models.Event)
	users := make(map[uuid.UUID]models.User)

	for _, interact := range UserInteractEvent {
		if _, ok := userInteractEvents[interact.UserID]; !ok {
			userInteractEvents[interact.UserID] = []models.Event{interact.Event}
		} else {
			userInteractEvents[interact.UserID] = append(userInteractEvents[interact.UserID], interact.Event)
		}
		users[interact.UserID] = interact.User
	}

	var userInteractCategoriesResponses []dto.UserInteractCategoriesResponse
	for userID, events := range userInteractEvents {
		user := users[userID]
		userInteractCategoriesResponses = append(userInteractCategoriesResponses, *convertToUserInteractCategoryResponse(&user, events))
	}

	return userInteractCategoriesResponses

}

func convertToEventsAreInteractedByUserResponse(user *models.User, userInteractEvents []models.UserInteractEvent) *dto.EventsAreInteractedByUserResponse {
	var Events []dto.EventWithCountResponses
	for _, userInteractEvent := range userInteractEvents {
		var Categories []dto.CategoryResponses
		event := userInteractEvent.Event
		for _, category := range event.Categories {
			Categories = append(Categories, dto.CategoryResponses{
				Value: category.ID,
				Label: category.Name,
			})
		}
		Events = append(Events, dto.EventWithCountResponses{
			EventResponse: ConvertToEventDocumentResponse(event),
			Count:         userInteractEvent.Count,
		})

	}
	return &dto.EventsAreInteractedByUserResponse{
		UserResponses:  *convertToUserResponses(user),
		EventResponses: Events,
	}
}

func convertToAllEventsAreInteractedByUserResponse(userInteractEvents []models.UserInteractEvent) []dto.EventsAreInteractedByUserResponse {
	userIDWithUserInteractEvents := make(map[uuid.UUID][]models.UserInteractEvent)
	users := make(map[uuid.UUID]models.User)

	for _, interact := range userInteractEvents {
		userID := interact.UserID
		if _, ok := userIDWithUserInteractEvents[userID]; !ok {
			userIDWithUserInteractEvents[interact.UserID] = []models.UserInteractEvent{interact}
		} else {
			userIDWithUserInteractEvents[interact.UserID] = append(userIDWithUserInteractEvents[interact.UserID], interact)
		}
		users[interact.UserID] = interact.User
	}

	var eventsAreInteractedByUserResponses []dto.EventsAreInteractedByUserResponse
	for userID, interactEvents := range userIDWithUserInteractEvents {
		user := users[userID]
		eventsAreInteractedByUserResponses = append(eventsAreInteractedByUserResponses, *convertToEventsAreInteractedByUserResponse(&user, interactEvents))
	}

	return eventsAreInteractedByUserResponses
}
