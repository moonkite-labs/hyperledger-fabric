package service

import (
	"encoding/json"
	"gocert-gateway/db"
	"gocert-gateway/models"
	"gocert-gateway/utils"
	"net/http"

	"github.com/google/uuid"
)

// CertificateService is put here to avoid cyclic import
// Services always import from db but CertificateService
// needs to import ContractService which is located here
type CertificateService struct {
	baseDB          *db.BaseDBService
	indiService     db.PostgreIndividualService
	walletService   db.PostgreWalletService
	contractService ContractService
}

func NewCertificateService(baseDB *db.BaseDBService, indiService db.PostgreIndividualService, walletService db.PostgreWalletService, contractService ContractService) *CertificateService {
	cs := &CertificateService{}
	cs.baseDB = baseDB
	cs.indiService = indiService
	cs.walletService = walletService
	cs.contractService = contractService
	return cs
}

func (cs CertificateService) autoSignUpIssueCertificate(w http.ResponseWriter, r *http.Request) {
	_ = uuid.New().String()

	// body, err := io.ReadAll(r.Body)

	// if err != nil {

	// 	w.WriteHeader(http.StatusUnprocessableEntity)
	// 	errstruct := struct {
	// 		Error string `json:"error"`
	// 	}{
	// 		Error: err.Error(),
	// 	}
	// 	errmsg := json.NewEncoder(w).Encode(errstruct)
	// 	if err != nil {
	// 		fmt.Fprintf(w, "%s", errmsg.Error())
	// 	}
	// 	return
	// }

	// // TODO CER-260 :  SEQ-#8.1
	// requester, err := cs.indiService.FindIndividualByCimPersonId(body.IndividualCimPersonId)

	// // TODO CER-260 :  SEQ-#8.2
	// org := orgService.findOrgByCimOrgId(request.getIssuerCimOrgId())

	// // TODO CER-260 :  SEQ-#8.3
	// crypto := cryptoService.findCryptoDetailsByOrgId(org.getId())

	// // TODO CER-260 :  SEQ-#8.5
	// user := userService.findUserByCimPersonId(request.getIssuedCimPersonId())
	// if StringUtils.isBlank(request.getIndividualPublickKey()) {
	// 	request.setIndividualPublickKey(
	// 		service.getCryptoDetailsByCimPersonId(requester.getCimPersonId()))
	// }

	// // TODO CER-260 :  SEQ-#8.6
	// var course string
	// if request.getCourseId() != nil {
	// 	course = courseService.getCourse(request.getCourseId())
	// }
	// expirationDate = request.getExpirationDate()
	// TimeWatcher.recordMethod(id, "SEQ-#8.6", stopwatch.stop())

	// randomNumber := UUID.randomNumber()
	// certificateName := this.getCertificateName(request)

	// // generate badge to ceritificate
	// certificate = &Certificate{}
	// certificate.IssuedDate = time.Now()
	// certificate.CreatedDate = time.Now()

	// if request.getCertificateType() == CertificateType.LICENSE_CERTIFICATE {
	// 	certificate.setExpiryDate(calculateExpiryDate(certificate.getIssuedDate()))
	// } else if request.getCertificateType() == QUALIFICATION_CERTIFICATE && course != nil {
	// 	certificateName = course.getCourse_name()
	// 	request.setSpecialization(course.getCourse_description())
	// 	request.setCourseName(course.getCourse_name())
	// 	// start : change made for CER-17
	// 	request.setStartDate(
	// 		CertificateHelper.getDateIfNullOrNotEqual(request, course).get("startDate"))
	// 	request.setEndDate(
	// 		CertificateHelper.getDateIfNullOrNotEqual(request, course).get("endDate"))
	// 	expirationDate =
	// 		CertificateHelper.getDateIfNullOrNotEqual(request, course).get("expirationDate")

	// 	// end : change made for CER-17
	// } else if request.getCertificateType() == CertificateType.OPEN_BADGES && course != null {
	// 	certificateName = course.getCourse_name()
	// 	var courseStartDate, courseEndDate time.Time
	// 	if course.getCourseStartDate() != nil {
	// 		courseStartDate = course.getCourseStartDate()
	// 	}
	// 	if course.getCourseEndDate() != nil {
	// 		courseEndDate = course.getCourseEndDate()
	// 	}
	// 	request.setStartDate(courseStartDate)
	// 	request.setEndDate(courseEndDate)
	// 	if Objects.isNull(expirationDate) {
	// 		expirationDate =
	// 			badgeService.getBadgeTemplateById(request.getBadgeImage()).getExpiryDate()
	// 	}
	// }

	// templateName := request.getTemplateName()
	// if templateName == "" {
	// 	templateName = org.getTemplates().get(string(request.getCertificateType()))
	// }

	// if request.getCertificateType() == CertificateType.OPEN_BADGES {
	// 	templateName = request.getBadgeImage()
	// }

	// TODO CER-260 :  SEQ-#8.7
}

func (cs *CertificateService) IssueCertificate(requesterCimPersonId string, recipientCimPersonId string, issuerCert *models.Certificate, rCert *models.Certificate) error {
	chaincodeName := "certificate-manager"
	contractName := "CertificateContract"

	requester, err := cs.indiService.FindIndividualByCimPersonId(recipientCimPersonId)

	if err != nil {
		return err
	}

	requesterCrypto := requester.Wallet

	signer, err := requesterCrypto.ToSign()

	if err != nil {
		return err
	}

	recipient, err := cs.indiService.FindIndividualByCimPersonId(recipientCimPersonId)

	if err != nil {
		return err
	}

	recipientPublicKey := recipient.Wallet.PublicKey

	if err != nil {
		return err
	}

	// Serialise issuerCert, TODO: May need to change for a better encoder
	issuerCertByte, err := json.Marshal(issuerCert)

	if err != nil {
		return err
	}

	// Calculate issuerCertHash
	issuerCertHash, err := utils.SHASumHex(issuerCertByte, 512)

	if err != nil {
		return err
	}

	// Serialise issuerCert, TODO: May need to change for a better encoder
	recipientCertByte, err := json.Marshal(rCert)

	if err != nil {
		return err
	}

	// Calculate issuerCertHash
	recipientCertHash, err := utils.SHASumHex(recipientCertByte, 512)

	if err != nil {
		return err
	}

	_, err = cs.contractService.IssueCertificate(chaincodeName, contractName, requesterCimPersonId, signer, issuerCert.Id, issuerCertHash, recipientCertHash, recipientCimPersonId, recipientPublicKey, rCert.Id)

	return err
}
