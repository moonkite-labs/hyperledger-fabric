package db

import (
	"fabric-gateway/models"
	"fmt"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreWalletService struct {
	DB *gorm.DB
}

func (p *PostgreWalletService) Connect(host string, user string, password string, dbname string, port string) error {
	var err error

	// TODO: sslmode
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s charset=utf8 parseTime=True", host, port, user, dbname, password)
	Dbdriver := postgres.Open(dsn)

	p.DB, err = gorm.Open(Dbdriver, &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		return fmt.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("Connected to the %s database", Dbdriver)
	return nil
}

//		Put an identity into the wallet.
//
//		Parameters:
//		  label specifies the name of the identity in the wallet.
//	   mspid the associated mspid of the identity
//		  publicKey the public key of the identity in bytes form.
//	   privateKey the private key of the identity in bytes form.
//
//		Returns:
//		  Error if any
func (p PostgreWalletService) Put(label string, mspid string, publicKey []byte, privateKey []byte) error {
	var err error

	i := models.Wallet{
		Label:      label,
		MSPID:      mspid,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	err = p.DB.Debug().Model(&models.Wallet{}).Create(&i).Error
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
func (p PostgreWalletService) Get(label string) (*models.Wallet, error) {
	var err error

	i := models.Wallet{}

	err = p.DB.Debug().Where("label = ?", label).First(&i).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error finding identity with label %s", label))
		return nil, err
	}

	return &i, nil
}

//		 Update an identity in the wallet.
//
//			Parameters:
//			label specifies the name of the identity in the wallet.
//	     mspid the associated mspid of the identity
//			publicKey the public key of the identity in bytes form.
//		    privateKey the private key of the identity in bytes form.
//
//			Returns:
//			Error if any
func (p PostgreWalletService) Update(label string, mspid string, publicKey []byte, privateKey []byte) error {
	var err error

	if !p.Exists(label) {
		return errors.New(fmt.Sprintf("Identity with label %s does not exist!", label))
	}

	i, err := p.Get(label)

	if err != nil {
		return errors.Wrapf(err, "Failed to get label %s", label)
	}

	if mspid != "" {
		i.MSPID = mspid
	}

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
func (p PostgreWalletService) Exists(label string) bool {

	i := models.Wallet{}

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
func (p PostgreWalletService) Delete(label string) error {
	var err error

	i := models.Wallet{}

	err = p.DB.Debug().Where("label = ?", label).First(&i).Delete(&i).Error

	if err != nil {
		errors.Wrap(err, fmt.Sprintf("Error deleting identity with label %s", label))
		return nil
	}

	return err
}