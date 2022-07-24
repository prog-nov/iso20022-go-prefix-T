package iso20022

// Transaction totals during the reconciliation period, for a certain type of transaction.
type TransactionTotals2 struct {

	// Identifier assigned by the merchant identifying a set of POI terminals performing some categories of transactions.
	POIGroupIdentification *Max35Text `xml:"POIGrpId,omitempty"`

	// Category of cards related the acceptance processing rules defined by the acquirer.
	CardProductProfile *Exact4NumericText `xml:"CardPdctPrfl,omitempty"`

	// Currency associated with the transaction totals.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Identification of the type of transaction.
	Type *TypeTransactionTotals2Code `xml:"Tp"`

	// Total number of transactions during a reconciliation period.
	TotalNumber *Max35NumericText `xml:"TtlNb"`

	// Total amount of a collection of transactions.
	CumulativeAmount *ImpliedCurrencyAndAmount `xml:"CmltvAmt"`
}

func (t *TransactionTotals2) SetPOIGroupIdentification(value string) {
	t.POIGroupIdentification = (*Max35Text)(&value)
}

func (t *TransactionTotals2) SetCardProductProfile(value string) {
	t.CardProductProfile = (*Exact4NumericText)(&value)
}

func (t *TransactionTotals2) SetCurrency(value string) {
	t.Currency = (*CurrencyCode)(&value)
}

func (t *TransactionTotals2) SetType(value string) {
	t.Type = (*TypeTransactionTotals2Code)(&value)
}

func (t *TransactionTotals2) SetTotalNumber(value string) {
	t.TotalNumber = (*Max35NumericText)(&value)
}

func (t *TransactionTotals2) SetCumulativeAmount(value, currency string) {
	t.CumulativeAmount = NewImpliedCurrencyAndAmount(value, currency)
}
