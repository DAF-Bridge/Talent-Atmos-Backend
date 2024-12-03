package domain

import (
	"time"

	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// ENUMS
//---------------------------------------------------------------------------

type Media string

const (
	// Media Enum
	MediaWebsite 	Media = "website"  
	MediaFacebook   Media = "facebook" 
	MediaIG         Media = "instagram"
	MediaTikTok     Media = "tiktok"
	MediaYoutube    Media = "youtube"
	MediaLinkedin   Media = "linkedin"
	MediaLine       Media = "line"    
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Organization struct {
	ID          			uint           			`gorm:"primaryKey;autoIncrement" json:"id"`
	Name        			string         			`gorm:"type:varchar(255);not null" json:"org_name"`
	Goal        			[]string       			`gorm:"type:text[];not null" json:"goal"`           // Detailed description of the organization's goal
	Expertise   			string         			`gorm:"type:varchar(255)" json:"expertise"`         // Organization's area of expertise
	Location    			string         			`gorm:"type:varchar(255)" json:"location"`          // General location
	Subdistrict 			string         			`gorm:"type:varchar(255)" json:"subdistrict"`       // Subdistrict name
	Province    			string         			`gorm:"type:varchar(255)" json:"province"`          // Province name
	PostalCode  			string         			`gorm:"type:varchar(20)" json:"postal_code"`        // Postal code, allowing for flexibility in format
	Latitude    			string         			`gorm:"type:varchar(50)" json:"latitude"`           // Geographic latitude (stored as string for precision)
	Longitude   			string         			`gorm:"type:varchar(50)" json:"longitude"`          // Geographic longitude (stored as string for precision)
	Email       			string         			`gorm:"type:varchar(255);unique" json:"org_email"`  // Email address (unique constraint)
	Phone       			string         			`gorm:"type:varchar(20)" json:"org_phone"`
	CreatedAt   			time.Time      			`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   			time.Time      			`gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   			gorm.DeletedAt 			`gorm:"index"`
	OrganizationContacts 	[]OrganizationContact 	`gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgOpenJobs 			[]OrgOpenJob 			`gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

type OrganizationContact struct {
	gorm.Model
	OrganizationID 	uint 		`json:"organization_id"`
	Media          	Media 		`gorm:"type:media;not null" json:"media"`
	MediaLink      	string 		`gorm:"type:varchar(255);not null" json:"media_link"`
}

type OrgOpenJob struct {
	gorm.Model
	OrganizationID  uint           `json:"organization_id"`
	Title       	string         `gorm:"type:varchar(255);not null" json:"title"`
	Scope       	string         `gorm:"type:varchar(255);not null" json:"scope"`
	Workplace   	string         `gorm:"type:varchar(255);not null" json:"workplace"`
	WorkType    	string         `gorm:"type:varchar(255);not null" json:"work_type"`
	Period      	string         `gorm:"type:varchar(255);not null" json:"period"`
	Description 	string         `gorm:"type:text" json:"description"`
	HoursPerDay 	string         `gorm:"type:varchar(255);not null" json:"hours_per_day"`
	Qualifications 	string         `gorm:"type:text" json:"qualifications"`
	Benefits 		string         `gorm:"type:text" json:"benefits"`
	Quantity    	int            `json:"quantity"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

// Organization
type OrganizationRepository interface {
	GetByID(id uint) (*Organization, error)
	GetAll() ([]Organization, error)
	Create(org *Organization) error
	// Update(org *Organization) error
	// Delete(id uint) error
}

type OrganizationService interface {
	GetOrgByID(id uint) (*Organization, error) 
	GetAllOrg() ([]Organization, error)
	CreateOrg(org *Organization) error
	// Update(org *Organization) error
	// Delete(id uint) error
}

//---------------------------------------------------------------------------
// OrgOpenJob
type OrgOpenJobRepository interface {
	GetByID(id uint) (*OrgOpenJob, error)
	GetAll() ([]OrgOpenJob, error)
	Create(org *OrgOpenJob) error
	// Update(org *OrgOpenJob) error
	// Delete(id uint) error
}

type OrgOpenJobService interface {
	GetOrgOpenJobByID(id uint) (*OrgOpenJob, error)
	GetAllOrgOpenJob() ([]OrgOpenJob, error)
	CreateOrgOpenJob(org *OrgOpenJob) error
	// Update(org *OrgOpenJob) error
	// Delete(id uint) error
}

//--------------------------------------------------------------------------
// OrganizationContact
type OrganizationContactRepository interface {
	Create(org *OrganizationContact) error
	// Update(org *OrganizationContact) error
	// Delete(id uint) error
}
