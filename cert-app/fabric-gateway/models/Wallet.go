package models

import "time"

type Tabler interface {
	TableName() string
}

type Identity struct {
	ID         uint64 `gorm:"primaryKey;auto_increment"`
	Label      string
	PublicKey  []byte
	PrivateKey []byte
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

// TableName overrides the table name used by Wallet to `wallet`
func (Identity) TableName() string {
	return "wallet"
}
