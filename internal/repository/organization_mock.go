package repository

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type organizationRepositoryMock struct {
	org *models.Organization
}

type organizationContactRepositoryMock struct {
	orgContact *models.OrganizationContact
}

type orgOpenJobRepositoryMock struct {
	job *models.OrgOpenJob
}

func (r orgOpenJobRepositoryMock) CountsByOrgID(orgID uint) (int64, error) {
	return 0, nil
}

func NewOrganizationRepositoryMock() OrganizationRepository {
	org := &models.Organization{
		Model:     gorm.Model{ID: 1, UpdatedAt: time.Now()},
		Email:     "talentsatmos@gmail.com",
		Phone:     "+66876428591",
		Name:      "Talents Atmos",
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
		OrgOpenJobs: []models.OrgOpenJob{
			{
				Model:  gorm.Model{ID: 1},
				Title:  "Software Engineer",
				PicUrl: "https://talentsatmos.com",
				Categories: []models.Category{
					{
						Model: gorm.Model{ID: 12},
						Name:  "social",
					},
				},
			},
		},
	}

	return &organizationRepositoryMock{org: org}
}

func NewOrganizationContactRepositoryMock(orgContact *models.OrganizationContact) OrganizationContactRepository {
	return &organizationContactRepositoryMock{orgContact: orgContact}
}

func NewOrgOpenJobRepositoryMock() OrgOpenJobRepository {
	newJob := &models.OrgOpenJob{
		Model:          gorm.Model{ID: 1, UpdatedAt: time.Now()},
		OrganizationID: 1,
		Organization:   models.Organization{Name: "Talents Atmos"},
		Title:          "Software Engineer",
		PicUrl:         "https://talentsatmos.com",
		Scope:          "Software Development",
		Location:       "Chiang Mai",
		Prerequisite:   pq.StringArray{"Great at problem solving", "Reliable"},
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

	return &orgOpenJobRepositoryMock{job: newJob}
}

// ----------------------------------------------
//
//	OrganizationRepository
//
// ----------------------------------------------
func (r *organizationRepositoryMock) CreateOrganization(userID uuid.UUID, org *models.Organization) (*models.Organization, error) {
	if org == nil {
		return nil, errs.NewNotFoundError("organization not found")
	}
	// Simulate auto increment ID
	if org.ID == 0 {
		org.ID = r.org.ID + 1
	}
	r.org = org

	return r.org, nil
}

func (r organizationRepositoryMock) FindIndustryByIds(industryIDs []uint) ([]models.Industry, error) {
	idSet := make(map[uint]struct{})
	for _, id := range industryIDs {
		idSet[id] = struct{}{}
	}

	var foundIndustries []models.Industry
	for _, industryPtr := range r.org.Industries {
		if _, ok := idSet[industryPtr.ID]; ok {
			foundIndustries = append(foundIndustries, *industryPtr)
		}
	}

	return foundIndustries, nil
}

func (r organizationRepositoryMock) GetAllIndustries() ([]models.Industry, error) {
	return nil, nil
}

func (r organizationRepositoryMock) GetByOrgID(id uint) (*models.Organization, error) {
	if r.org.ID == id {
		return r.org, nil
	}
	return nil, nil
}

func (r organizationRepositoryMock) GetAllOrganizations() ([]models.Organization, error) {
	return nil, nil
}

func (r organizationRepositoryMock) GetOrganizations(userID uuid.UUID) ([]models.Organization, error) {
	return nil, nil
}

func (r organizationRepositoryMock) GetOrgsPaginate(page uint, size uint) ([]models.Organization, error) {
	return nil, nil
}

func (r organizationRepositoryMock) UpdateOrganization(org *models.Organization) (*models.Organization, error) {
	return nil, nil
}

func (r organizationRepositoryMock) UpdateOrganizationPicture(id uint, picURL string) error {
	return nil
}

func (r organizationRepositoryMock) UpdateOrganizationBackgroundPicture(id uint, picURL string) error {
	return nil
}

func (r organizationRepositoryMock) DeleteOrganization(org uint) error {
	return nil
}

// ----------------------------------------------
// 			OrganizationContactRepository
// ----------------------------------------------

func (r organizationContactRepositoryMock) Create(orgID uint, org *models.OrganizationContact) error {
	return nil
}

func (r organizationContactRepositoryMock) GetByID(orgID uint, id uint) (*models.OrganizationContact, error) {
	if r.orgContact.ID == id {
		return r.orgContact, nil
	}
	return nil, nil
}

func (r organizationContactRepositoryMock) GetAllByOrgID(orgID uint) ([]models.OrganizationContact, error) {
	return nil, nil
}

func (r organizationContactRepositoryMock) Update(org *models.OrganizationContact) (*models.OrganizationContact, error) {
	return nil, nil
}

func (r organizationContactRepositoryMock) Delete(orgID uint, id uint) error {
	return nil
}

// ----------------------------------------------
//
//	OrgOpenJobRepository
//
// ----------------------------------------------
func (r orgOpenJobRepositoryMock) CreateJob(orgID uint, job *models.OrgOpenJob) error {
	return nil
}

func (r orgOpenJobRepositoryMock) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {

	return nil, nil
}

func (r orgOpenJobRepositoryMock) GetJobByID(orgID uint, jobID uint) (*models.OrgOpenJob, error) {
	if r.job.ID == jobID {
		return r.job, nil
	}
	return nil, nil
}

func (r orgOpenJobRepositoryMock) GetAllJobs() ([]models.OrgOpenJob, error) {
	return nil, nil
}

func (r orgOpenJobRepositoryMock) GetAllJobsByOrgID(OrgId uint) ([]models.OrgOpenJob, error) {
	return nil, nil
}

func (r orgOpenJobRepositoryMock) GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error) {
	return nil, nil
}

func (r orgOpenJobRepositoryMock) UpdateJob(job *models.OrgOpenJob) (*models.OrgOpenJob, error) {
	return nil, nil
}

func (r orgOpenJobRepositoryMock) UpdateJobPicture(orgID uint, jobID uint, picURL string) error {
	return nil
}

func (r orgOpenJobRepositoryMock) DeleteJob(orgID uint, jobID uint) error {
	return nil
}
