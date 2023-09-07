package contracts

import "time"

type CertificateType uint8

const (
	LICENSE_CERTIFICATE = iota
	QUALIFICATION_CERTIFICATE
	ATTENDANCE_LETTER
	WORK_ATTESTATION
	SKILLS
	ENROLLMENT
	TRANSCRIPT
	RECORD_OF_LEARNING
	INTERNSHIP
	MEMBERSHIP_TYPE
	PROOF_OF_INTERNSHIP
	PROOF_OF_ENROLLMENT
	REFERENCE_LETTER
	OPEN_BADGES
)

type Certificate struct {
	Id                  string              `json:"Id"`
	ClaimRequest        ClaimRequest        `json:"ClaimRequest"`
	Course              Course              `json:"Course"`
	CovidSelfTestEntity CovidSelfTestEntity `json:"CovidSelfTestEntity"`
	RecipientId         string              `json:"RecipientId"`
	ContestantName      string              `json:"ContestantName"`
	VoterName           string              `json:"VoterName"`
	EventName           string              `json:"EventName"`
	RecipientEmail      string              `json:"RecipientEmail"`
	// first name and last name
	RecipientName       string                  `json:"RecipientName"`
	IndividualPersonId  string                  `json:"IndividualPersonId"`
	IndividualPublicKey string                  `json:"IndividualPublicKey"`
	CertificateType     CertificateType         `json:"CertificateType"`
	TemplateName        string                  `json:"TemplateName"`
	CertificateText     string                  `json:"CertificateText"`
	CertificateHtml     string                  `json:"CertificateHtml"`
	CertFields          CertificateInnerFields  `json:"CertFields"`
	Attachments         []CertificateAttachment `json:"Attachments"`
	CustomTemplateUrl   string                  `json:"CustomTemplateUrl"`
	IssuerId            string                  `json:"IssuerId"`
	IssuerName          string                  `json:"IssuerName"`
	IssuedDate          time.Time               `json:"IssuedDate"`
	ExpiryDate          time.Time               `json:"ExpiryDate"`
	CreatedDate         time.Time               `json:"CreatedDate"`
	IssuedPersonId      string                  `json:"IssuedPersonId"`
	IssuerPublicKey     string                  `json:"IssuerPublicKey"`
	Uid                 string                  `json:"Uid"`
	Description         string                  `json:"Description"`
	CertificateName     string                  `json:"CertificateName"`
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
	SourceOfPdf     string   `json:"SourceOfPdf"`
	IsNotified      bool     `json:"IsNotified"`
	Status          string   `json:"Status"`
	CorrelationId   string   `json:"CorrelationId"`
	TemplateJSON    string   `json:"TemplateJSON"`
}

func (cert *Certificate) 
