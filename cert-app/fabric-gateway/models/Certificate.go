package models

import "time"

type Certificate struct {
	Id             string `gorm:"primaryKey" json:"Id"`
	Course         string `json:"Course"`
	RecipientId    string `json:"RecipientId"`
	EventName      string `json:"EventName"`
	RecipientEmail string `json:"RecipientEmail"`
	// first name and last name
	RecipientName string `json:"RecipientName"`
	// RecipientId
	IndividualPersonId  string          `json:"IndividualPersonId"`
	IndividualPublicKey []byte          `json:"IndividualPublicKey"`
	CertificateType     CertificateType `json:"CertificateType"`
	TemplateName        string          `json:"TemplateName"`
	CustomTemplateUrl   string          `json:"CustomTemplateUrl"`
	// Organisation who issued the certificate
	IssuerId    string    `json:"IssuerId"`
	IssuerName  string    `json:"IssuerName"`
	IssuedDate  time.Time `json:"IssuedDate"`
	ExpiryDate  time.Time `json:"ExpiryDate"`
	CreatedDate time.Time `json:"CreatedDate"`
	// Organisation admin
	IssuedPersonId  string `json:"IssuedPersonId"`
	IssuerPublicKey []byte `json:"IssuerPublicKey"`
	Uid             string `json:"Uid"`
	Description     string `json:"Description"`
	CertificateName string `json:"CertificateName"`
	// OrgnaisationId or IndividualId
	OwnerId    string `json:"OwnerId"`
	OwnerName  string `json:"OwnerName"`
	OwnerEmail string `json:"OwnerEmail"`
	// Certificate Id
	SharedFrom string `json:"SharedFrom"`
	// Certificate Id which is Issued by orgnaisation
	MasterId        string   `json:"MasterId"`
	TypeOfCopy      CopyType `json:"TypeOfCopy"`
	IsValid         bool     `json:"IsValid"`
	RevokedReason   string   `json:"RevokedReason"`
	BadgeUrl        string   `json:"BadgeUrl"`
	TemplateType    string   `json:"TemplateType"`
	CertificateDesc string   `json:"CertificateDesc"`
	Status          string   `json:"Status"`
	TemplateJSON    string   `json:"TemplateJSON" gorm:"column:template_json"`
}
