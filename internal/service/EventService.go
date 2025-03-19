package service

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/search"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/sync"
	"github.com/opensearch-project/opensearch-go"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

const numberOfEvent uint = 12

// EventService is a service that provides operations on events.
type eventService struct {
	eventRepo repository.EventRepository
	DB        *gorm.DB
	OS        *opensearch.Client
	S3        *infrastructure.S3Uploader
}

//--------------------------------------------//

func NewEventService(eventRepo repository.EventRepository, db *gorm.DB, os *opensearch.Client, s3 *infrastructure.S3Uploader) EventService {
	return eventService{
		eventRepo: eventRepo,
		DB:        db,
		OS:        os,
		S3:        s3}
}

func (s eventService) CountEventByOrgID(orgID uint) (int64, error) {
	count, err := s.eventRepo.CountsByOrgID(orgID)

	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}

	return count, nil
}

func (s eventService) SyncEvents() error {
	return sync.SyncEventsToOpenSearch(s.DB, s.OS)
}

func (s eventService) SearchEvents(query dto.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error) {
	eventsRes, err := search.SearchEvents(s.OS, query, page, Offset)
	if err != nil {
		if len(eventsRes.Events) == 0 {
			return dto.SearchEventResponse{}, errs.NewNotFoundError("No search results found")
		}

		return dto.SearchEventResponse{}, errs.NewUnexpectedError()
	}
	return eventsRes, nil
}

func (s eventService) NewEvent(orgID uint, req dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error {
	categoryIDs := make([]uint, 0)
	for _, category := range req.Categories {
		categoryIDs = append(categoryIDs, category.Value)
	}

	categories, err := s.eventRepo.FindCategoryByIds(categoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("categories not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	contacts := make([]models.ContactChannel, 0)
	for _, contact := range req.ContactChannels {
		contacts = append(contacts, models.ContactChannel{
			Media:     models.Media(contact.Media),
			MediaLink: contact.MediaLink,
		})
	}

	event := requestConvertToEvent(orgID, req, categories, contacts)

	err = s.eventRepo.Create(orgID, &event)
	if err != nil {
		logs.Error(err)

		return errs.NewUnexpectedError()
	}

	// Upload image to S3
	if file != nil {
		picURL, err := s.S3.UploadEventPictureFile(ctx, file, fileHeader, orgID, event.ID)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}

		// Update event record in database
		err = s.eventRepo.UpdateEventPicture(orgID, event.ID, picURL)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errs.NewNotFoundError("event not found")
			}

			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	return nil
}

func (s eventService) GetAllEvents() ([]dto.EventResponses, error) {
	events, err := s.eventRepo.GetAll()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	EventResponses := make([]dto.EventResponses, 0)
	for _, event := range events {
		eventResponse := ConvertToEventResponse(event)
		EventResponses = append(EventResponses, eventResponse)
	}

	return EventResponses, nil
}

func (s eventService) GetAllEventsByOrgID(orgID uint) ([]dto.EventResponses, error) {
	events, err := s.eventRepo.GetAllByOrgID(orgID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	EventResponses := make([]dto.EventResponses, 0)
	for _, event := range events {
		eventResponse := ConvertToEventResponse(event)
		EventResponses = append(EventResponses, eventResponse)
	}

	return EventResponses, nil
}

func (s eventService) GetEventByID(eventID uint) (*dto.EventResponses, error) {
	event, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*event)
	return &eventResponse, nil
}

func (s eventService) GetEventByIDwithOrgID(orgID uint, eventID uint) (*dto.EventResponses, error) {
	event, err := s.eventRepo.GetByIDwithOrgID(orgID, eventID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*event)

	return &eventResponse, nil
}

func (s eventService) ListAllCategories() (*dto.CategoryListResponse, error) {
	categories, err := s.eventRepo.GetAllCategories()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("No categories found")
		}

		return nil, err
	}

	var responses dto.CategoryListResponse
	for _, category := range categories {
		responses.Categories = append(responses.Categories, dto.CategoryResponses{
			Value: category.ID,
			Label: category.Name,
		})
	}

	return &responses, nil
}

func (s eventService) GetEventPaginate(page uint) ([]dto.EventDocumentDTOResponse, error) {
	events, err := s.eventRepo.GetPaginate(page, numberOfEvent)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	EventResponses := make([]dto.EventDocumentDTOResponse, 0)
	for _, event := range events {
		eventResponse := ConvertToEventDocumentResponse(event)
		EventResponses = append(EventResponses, eventResponse)
	}

	return EventResponses, nil
}

func (s eventService) GetFirst() (*dto.EventResponses, error) {
	event, err := s.eventRepo.GetFirst()

	if err != nil {

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*event)

	return &eventResponse, nil
}

func (s eventService) CountEvent() (int64, error) {
	count, err := s.eventRepo.Count()

	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}

	return count, nil
}

func (s eventService) UpdateEvent(orgID uint, eventID uint, req dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.EventResponses, error) {

	existingEvent, err := s.eventRepo.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var categoryIDs []uint
	for _, category := range req.Categories {
		categoryIDs = append(categoryIDs, category.Value)
	}

	categories := make([]models.Category, 0)
	if len(categoryIDs) > 0 {
		categories, err = s.eventRepo.FindCategoryByIds(categoryIDs)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errs.NewNotFoundError("categories not found")
			}

			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	}

	var contacts []models.ContactChannel
	for _, contact := range req.ContactChannels {
		contacts = append(contacts, models.ContactChannel{
			EventID:   existingEvent.ID,
			Media:     models.Media(contact.Media),
			MediaLink: contact.MediaLink,
		})
	}

	// Convert request to Event
	event := requestConvertToEvent(orgID, req, categories, contacts)
	event.ID = eventID

	if file != nil {
		picURL, err := s.S3.UploadEventPictureFile(ctx, file, fileHeader, orgID, eventID)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
		event.PicUrl = picURL
	} else {
		event.PicUrl = existingEvent.PicUrl
	}

	// Update event record in database
	updateEvent, err := s.eventRepo.Update(orgID, event.ID, &event)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*updateEvent)

	return &eventResponse, nil
}

func (s eventService) UploadEventPicture(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID uint, eventID uint) (string, error) {
	picURL, err := s.S3.UploadEventPictureFile(ctx, file, fileHeader, orgID, eventID)

	if err != nil {
		logs.Error(err)
		return "", errs.NewUnexpectedError()
	}

	return picURL, nil
}

func (s eventService) DeleteEvent(orgID uint, eventID uint) error {
	err := s.eventRepo.Delete(orgID, eventID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}
