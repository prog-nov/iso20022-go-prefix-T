package iso20022

// Parameters applied to the settlement of a security transfer.
type Transfer5 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Date at which the instructing party places the transfer instruction.
	TransferDate *DateFormat1Choice `xml:"TrfDt,omitempty"`

	// Total quantity of securities to be settled.
	TotalUnitsNumber *FinancialInstrumentQuantity1 `xml:"TtlUnitsNb"`

	// Indicates whether the transfer results in a change of beneficial owner.
	OwnAccountTransferIndicator *YesNoIndicator `xml:"OwnAcctTrfInd,omitempty"`

	// Additional specific settlement information for non-regulated traded funds.
	NonStandardSettlementInformation *Max350Text `xml:"NonStdSttlmInf,omitempty"`
}

func (t *Transfer5) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *Transfer5) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *Transfer5) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *Transfer5) AddTransferDate() *DateFormat1Choice {
	t.TransferDate = new(DateFormat1Choice)
	return t.TransferDate
}

func (t *Transfer5) AddTotalUnitsNumber() *FinancialInstrumentQuantity1 {
	t.TotalUnitsNumber = new(FinancialInstrumentQuantity1)
	return t.TotalUnitsNumber
}

func (t *Transfer5) SetOwnAccountTransferIndicator(value string) {
	t.OwnAccountTransferIndicator = (*YesNoIndicator)(&value)
}

func (t *Transfer5) SetNonStandardSettlementInformation(value string) {
	t.NonStandardSettlementInformation = (*Max350Text)(&value)
}
