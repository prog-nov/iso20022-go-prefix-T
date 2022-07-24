package iso20022

// Result of the verifications performed by the issuer to deliver or decline the authorisation.
type TransactionVerificationResult1 struct {

	// Result of an e-commerce authentication process.
	ElectronicCommerceAuthenticationResult *Max500Text `xml:"ElctrncComrcAuthntcnRslt,omitempty"`

	// Result of the printed card security code (CSC) validation.
	CSCResult *CSCResult1Code `xml:"CSCRslt,omitempty"`

	// Result of the cardholder verification address checks on the street number and the postal code.
	CardholderAddressVerificationResult *CardholderAddressVerificationResult1Code `xml:"CrdhldrAdrVrfctnRslt,omitempty"`

	// Product code for which the authorisation was declined.
	DeclinedProductCode []*Max70Text `xml:"DclndPdctCd,omitempty"`
}

func (t *TransactionVerificationResult1) SetElectronicCommerceAuthenticationResult(value string) {
	t.ElectronicCommerceAuthenticationResult = (*Max500Text)(&value)
}

func (t *TransactionVerificationResult1) SetCSCResult(value string) {
	t.CSCResult = (*CSCResult1Code)(&value)
}

func (t *TransactionVerificationResult1) SetCardholderAddressVerificationResult(value string) {
	t.CardholderAddressVerificationResult = (*CardholderAddressVerificationResult1Code)(&value)
}

func (t *TransactionVerificationResult1) AddDeclinedProductCode(value string) {
	t.DeclinedProductCode = append(t.DeclinedProductCode, (*Max70Text)(&value))
}
