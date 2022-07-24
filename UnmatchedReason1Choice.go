package iso20022

// Choice of format for the settlement transaction unmatched reason.
type UnmatchedReason1Choice struct {

	// Specifies the reason why the instruction has an unmatched status.
	Code *UnmatchedReason2Code `xml:"Cd"`

	// Specifies the reason why the instruction has an unmatched status.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (u *UnmatchedReason1Choice) SetCode(value string) {
	u.Code = (*UnmatchedReason2Code)(&value)
}

func (u *UnmatchedReason1Choice) AddProprietary() *GenericIdentification20 {
	u.Proprietary = new(GenericIdentification20)
	return u.Proprietary
}
