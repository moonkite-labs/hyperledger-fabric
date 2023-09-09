package db

import (
	"gocert-gateway/models"
	"time"

	"gorm.io/datatypes"
)

type IndividualOption func(*models.Individual)

// Options to be used for updating fields of an individual object.

func WithOrgId(orgId string) IndividualOption {
	return func(i *models.Individual) {
		i.OrgId = orgId
	}
}

func WithCimPersonId(cimPersonId string) IndividualOption {
	return func(i *models.Individual) {
		i.CimPersonId = cimPersonId
	}
}

func WithCimOrganisationId(cimOrganisationId string) IndividualOption {
	return func(i *models.Individual) {
		i.CimOrganisationId = cimOrganisationId
	}
}

func WithCountry(country string) IndividualOption {
	return func(i *models.Individual) {
		i.Country = country
	}
}

func WithFirstName(firstName string) IndividualOption {
	return func(i *models.Individual) {
		i.FirstName = firstName
	}
}

func WithLastName(lastName string) IndividualOption {
	return func(i *models.Individual) {
		i.LastName = lastName
	}
}

func WithPictureUrl(pictureUrl string) IndividualOption {
	return func(i *models.Individual) {
		i.PictureUrl = pictureUrl
	}
}

func WithDateOfBirth(dateOfBirth time.Time) IndividualOption {
	return func(i *models.Individual) {
		i.DateOfBirth = datatypes.Date(dateOfBirth)
	}
}

func WithOrgName(orgName string) IndividualOption {
	return func(i *models.Individual) {
		i.OrgName = orgName
	}
}

func WithWallet(wallet models.Wallet) IndividualOption {
	return func(i *models.Individual) {
		i.Wallet = wallet
	}
}
