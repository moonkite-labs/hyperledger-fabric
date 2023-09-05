package service

import (
	"fabric-gateway/models"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type WalletService interface {
	Put(label string, publicKey []byte, privateKey []byte) error
	Get(label string) models.Identity
	Update(label string) error
	Delete(label string) error
}

type PostgreWalletService struct {
	DB *gorm.DB
}

func (p *PostgreWalletService) Connect(host string, user string, password string, dbname string, port string) error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	Dbdriver := postgres.Open(dsn)

	p.DB, err = gorm.Open(Dbdriver, &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		return fmt.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Connected to the %s database", Dbdriver)
	return nil
}

//	 Put an identity into the wallet.
//
//		Parameters:
//		label specifies the name of the identity in the wallet.
//		publicKey the public key of the identity in bytes form.
//	 privateKey the private key of the identity in bytes form.
//
//		Returns:
//		Error if any
func (p *PostgreWalletService) Put(label string, publicKey []byte, privateKey []byte) error {
	var err error

	i := models.Identity{
		Label:      label,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	err = p.DB.Debug().Model(&models.Identity{}).Create(&i).Error
	if err != nil {
		errors.Wrap(err, "Error saving identity to wallet")
	}

	return err
}

//	 Get an identity from the wallet.
//
//		Parameters:
//		label specifies the name of the identity in the wallet.
//
//		Returns:
//		The identity object.
func (p *PostgreWalletService) Get(label string) *models.Identity {
	var err error

	i := models.Identity{}

	err = p.DB.Debug().Where("label = ?", label).First(&i).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error finding identity with label %s", label))
		return nil
	}

	return &i
}

//	 Update an identity in the wallet.
//
//		Parameters:
//		label specifies the name of the identity in the wallet.
//		publicKey the public key of the identity in bytes form.
//	 privateKey the private key of the identity in bytes form.
//
//		Returns:
//		Error if any
func (p *PostgreWalletService) Update(label string, publicKey []byte, privateKey []byte) error {
	var err error

	if !p.Exists(label) {
		return errors.New(fmt.Sprintf("Identity with label %s does not exist!", label))
	}

	i := p.Get(label)

	if publicKey != nil {
		i.PublicKey = publicKey
	}

	if privateKey != nil {
		i.PrivateKey = privateKey
	}

	err = p.DB.Debug().Save(&i).Error

	return err
}

//	 Check if an identity exists in the wallet.
//
//		Parameters:
//		label specifies the name of the identity in the wallet.
//
//		Returns:
//		True or false.
func (p *PostgreWalletService) Exists(label string) bool {

	i := models.Identity{}

	result := p.DB.Debug().Where("label = ?", label).First(&i)

	if result.Error != nil {
		fmt.Errorf(errors.Wrap(result.Error, "Fail to check for identity existence").Error())
	}

	return result.RowsAffected > 0 // Return (count of records found > 0)
}

//	 Delete an identity from the wallet.
//
//		Parameters:
//		label specifies the name of the identity in the wallet.
//
//		Returns:
//		The identity object.
func (p *PostgreWalletService) Delete(label string) error {
	var err error

	i := models.Identity{}

	err = p.DB.Debug().Where("label = ?", label).First(&i).Delete(&i).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error deleting identity with label %s", label))
		return nil
	}

	return err
}
