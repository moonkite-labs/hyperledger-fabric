package db

import (
	"fmt"
	"gocert-gateway/models"
	"time"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

type PostgreIndividualService struct {
	baseDB *BaseDBService
}

// Initialise a new individual service with a baseDB instance
func NewIndividualService(baseDB *BaseDBService) *PostgreIndividualService {
	return &PostgreIndividualService{baseDB: baseDB}
}

//		Store an individual into the database.
//
//		Parameters:
//		  id specifies the name of the individual in the database.
//	   mspid the associated mspid of the individual
//		  publicKey the public key of the individual in bytes form.
//	   privateKey the private key of the individual in bytes form.
//
//		Returns:
//		  Error if any
func (p PostgreIndividualService) CreateIndividual(orgId string, cimPersonId string, cimOrganisationId string, country string, firstName string, lastName string, pictureUrl string, dateOfBirth time.Time, orgName string, wallet *models.Wallet) error {
	var err error

	i := models.Individual{
		OrgId:             orgId,
		CimPersonId:       cimPersonId,
		CimOrganisationId: cimOrganisationId,
		Country:           country,
		FirstName:         firstName,
		LastName:          lastName,
		PictureUrl:        pictureUrl,
		DateOfBirth:       datatypes.Date(dateOfBirth),
		OrgName:           orgName,
	}

	if wallet != nil {
		i.Wallet = *wallet
	}

	err = p.baseDB.DB.Debug().Model(&models.Individual{}).Create(&i).Error
	if err != nil {
		errors.Wrap(err, "Error saving individual to individual")
	}

	return err
}

//	 Get an individual by id from the database.
//
//		Parameters:
//		id the id of the individual in the database.
//
//		Returns:
//		The individual object and error if any.
func (p PostgreIndividualService) FindIndividualById(id uint64) (*models.Individual, error) {
	var err error

	i := models.Individual{}

	err = p.baseDB.DB.Debug().Model(&i).Preload("Wallet").First(&i, id).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error finding individual with id %d", id))
		return nil, err
	}

	return &i, nil
}

func (p PostgreIndividualService) FindIndividualByCimPersonId(cim_person_id string) (*models.Individual, error) {
	var err error

	i := models.Individual{}

	err = p.baseDB.DB.Debug().Model(&i).Preload("Wallet").Where("cim_person_id = ?", cim_person_id).First(&i).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error finding individual with cim_person_id %s", cim_person_id))
		return nil, err
	}

	return &i, nil
}

//	 Update an individual in the database.
//
//		Parameters:
//		id: the id of the individual in the database.
//
//		publicKey the public key of the individual in bytes form.
//	    privateKey the private key of the individual in bytes form.
//
//		Returns:
//		Error if any
func (p PostgreIndividualService) UpdateById(id uint64, options ...IndividualOption) error {
	var err error

	if !p.ExistsById(id) {
		return errors.New(fmt.Sprintf("Individual with id %d does not exist!", id))
	}

	i, err := p.FindIndividualById(id)

	if err != nil {
		return errors.Wrapf(err, "Failed to get individual of id %d", id)
	}

	for _, o := range options {
		o(i)
	}

	err = p.baseDB.DB.Debug().Save(&i).Error

	return err
}

//	 Check if an individual exists in the database.
//
//		Parameters:
//		id specifies the id of the individual in the database.
//
//		Returns:
//		True or false.
func (p PostgreIndividualService) ExistsById(id uint64) bool {

	i := models.Individual{}

	result := p.baseDB.DB.Debug().First(&i, id)

	if result.Error != nil {
		fmt.Errorf(errors.Wrap(result.Error, "Fail to check for individual existence").Error())
		return false
	}

	return result.RowsAffected > 0 // Return (count of records found > 0)
}

//	 Delete an individual from the database.
//
//		Parameters:
//		id specifies the id of the individual in the database.
//
//		Returns:
//		The individual object.
func (p PostgreIndividualService) DeleteById(id uint64) error {
	var err error

	i := models.Individual{}

	err = p.baseDB.DB.Debug().Delete(&i, id).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error deleting individual with id %d", id))
		return nil
	}

	return err
}
