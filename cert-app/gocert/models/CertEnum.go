package models

type Certification int
type Role string
type RecommendationLevel int
type Level int

type CopyType int

const (
	MASTER CopyType = iota
	RECIPIENT
)

type TemplateCategory int

const (
	BASE TemplateCategory = iota
	CUSTOM
)

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
