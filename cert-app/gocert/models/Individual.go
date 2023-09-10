package models

import (
	"time"

	"gorm.io/datatypes"
)

type Individual struct {
	Id                uint64 `gorm:"primaryKey;auto_increment"`
	OrgId             string
	CimPersonId       string
	CimOrganisationId string
	Country           string
	FirstName         string
	LastName          string
	PictureUrl        string
	DateOfBirth       datatypes.Date
	OrgName           string
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	Wallet            Wallet    `gorm:"foreignKey:Label;references:cim_person_id"`
}

// type Individual struct {
// 	Id                string
// 	LinkedInProfileId string
// 	Country           string
// 	CimPersonId       string
// 	CimOrganisationId string
// 	FirstName         string
// 	LastName          string
// 	Email             string
// 	PictureUrl        string
// 	AvailableForWork  bool
// 	Status            string
// 	AccountStatus     string // TODO: AccountStatus struct
// 	Title             string
// 	IsSocialLogin     bool
// 	Phone             string // TODO: Phone struct
// 	CompanyInfo       string // TODO: CompanyInfo struct
// 	ProfileUpdated    bool
// 	DateOfBirth       time.Time
// 	Gender            string
// 	OrgName           string
// 	Channel           string
// 	OrgId             string
// 	CreatedAt         time.Time
// }

// TableName overrides the table name used by Individual to `individuals`
func (Individual) TableName() string {
	return "individual"
}
