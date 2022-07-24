package iso20022

// Reference of a transfer and of a transfer confirmation.
type TransferReference6 struct {

	// Unique and unambiguous identifier for a group of individual transfers as assigned by the instructing party. This identifier links the individual transfers together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Unique and unambiguous identifier for a transfer instruction, as assigned by the instructing party.
	TransferReference *Max35Text `xml:"TrfRef,omitempty"`

	// Unique and unambiguous investor's identification of a transfer. This reference can typically be used in a hub scenario to give the reference of the transfer as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Unique and unambiguous identifier for a transfer execution, as assigned by a confirming party.
	TransferConfirmationReference *Max35Text `xml:"TrfConfRef,omitempty"`

	// Unambiguous identification of the transfer allocated by the counterparty.
	CounterpartyReference *AdditionalReference2 `xml:"CtrPtyRef,omitempty"`
}

func (t *TransferReference6) SetMasterReference(value string) {
	t.MasterReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetTransferReference(value string) {
	t.TransferReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetClientReference(value string) {
	t.ClientReference = (*Max35Text)(&value)
}

func (t *TransferReference6) SetTransferConfirmationReference(value string) {
	t.TransferConfirmationReference = (*Max35Text)(&value)
}

func (t *TransferReference6) AddCounterpartyReference() *AdditionalReference2 {
	t.CounterpartyReference = new(AdditionalReference2)
	return t.CounterpartyReference
}
