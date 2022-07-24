package iso20022

// Details of the transaction.
type TransactionDetails3 struct {

	// Unique identification assigned to a trade. This is the reference generated by a firm or the reference allocated by the executing system if the trade was executed automatically.
	TradeReference *Max70Text `xml:"TradRef"`

	// Reference that links to other trades that are/will be sent, eg for straddles where put and call legs need to be reported together.
	AssociatedTradeReference []*Max70Text `xml:"AssoctdTradRef,omitempty"`

	// Identifies the execution venue. In the case of an exchange or a Multilateral Trading Facility (MTF), this should be identified using a MIC code. In the case of a systematic internaliser, place of trade should be identified using a BIC code.
	PlaceOfTrade *PlaceOfTradeIdentification2Choice `xml:"PlcOfTrad"`

	// Specifies the date/time on which the trade was executed.
	TradeDateTime *ISODateTime `xml:"TradDtTm"`

	// Provides details of the financial instrument for which the transaction report is being sent.
	FinancialInstrumentDetails *FinancialInstrument15 `xml:"FinInstrmDtls"`

	// Identifies whether the transaction was a buy or a sell from the perspective of the reporting firm.
	Side *OrderDriverCode `xml:"Sd"`

	// Identifies the regulator(s) to whom the transaction report must be sent.
	TransactionReportMarker []*PartyIdentification24Choice `xml:"TxRptMrkr,omitempty"`

	// Provides details of the counterparty.
	Counterparty *PartyIdentification11Choice `xml:"CtrPty"`

	// Provides details of the client.
	Client *PartyIdentification23 `xml:"Clnt,omitempty"`

	// Identifies the trading capacity of the firm reporting the transaction, eg  Agent or Principal.
	Capacity *TradingCapacity3Code `xml:"Cpcty"`

	// Specifies the currency and price at which the trade has been executed, excluding commission or accrued interest.
	ExecutedTradePrice *PriceRateOrAmountChoice `xml:"ExctdTradPric"`

	// Quantity of financial instrument executed by the trading party.
	ExecutedTradeQuantity *UnitOrFaceAmountChoice `xml:"ExctdTradQty"`

	// The total consideration or value.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Identifies the intended settlement date.
	SettlementDate *ISODateTime `xml:"SttlmDt,omitempty"`

	// Provides details of the person/organisation that has the power of attorney.
	ProxyHolder *PartyIdentification2Choice `xml:"PrxyHldr,omitempty"`

	// Additional domestic regulatory transaction information.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TransactionDetails3) SetTradeReference(value string) {
	t.TradeReference = (*Max70Text)(&value)
}

func (t *TransactionDetails3) AddAssociatedTradeReference(value string) {
	t.AssociatedTradeReference = append(t.AssociatedTradeReference, (*Max70Text)(&value))
}

func (t *TransactionDetails3) AddPlaceOfTrade() *PlaceOfTradeIdentification2Choice {
	t.PlaceOfTrade = new(PlaceOfTradeIdentification2Choice)
	return t.PlaceOfTrade
}

func (t *TransactionDetails3) SetTradeDateTime(value string) {
	t.TradeDateTime = (*ISODateTime)(&value)
}

func (t *TransactionDetails3) AddFinancialInstrumentDetails() *FinancialInstrument15 {
	t.FinancialInstrumentDetails = new(FinancialInstrument15)
	return t.FinancialInstrumentDetails
}

func (t *TransactionDetails3) SetSide(value string) {
	t.Side = (*OrderDriverCode)(&value)
}

func (t *TransactionDetails3) AddTransactionReportMarker() *PartyIdentification24Choice {
	newValue := new(PartyIdentification24Choice)
	t.TransactionReportMarker = append(t.TransactionReportMarker, newValue)
	return newValue
}

func (t *TransactionDetails3) AddCounterparty() *PartyIdentification11Choice {
	t.Counterparty = new(PartyIdentification11Choice)
	return t.Counterparty
}

func (t *TransactionDetails3) AddClient() *PartyIdentification23 {
	t.Client = new(PartyIdentification23)
	return t.Client
}

func (t *TransactionDetails3) SetCapacity(value string) {
	t.Capacity = (*TradingCapacity3Code)(&value)
}

func (t *TransactionDetails3) AddExecutedTradePrice() *PriceRateOrAmountChoice {
	t.ExecutedTradePrice = new(PriceRateOrAmountChoice)
	return t.ExecutedTradePrice
}

func (t *TransactionDetails3) AddExecutedTradeQuantity() *UnitOrFaceAmountChoice {
	t.ExecutedTradeQuantity = new(UnitOrFaceAmountChoice)
	return t.ExecutedTradeQuantity
}

func (t *TransactionDetails3) SetSettlementAmount(value, currency string) {
	t.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TransactionDetails3) SetSettlementDate(value string) {
	t.SettlementDate = (*ISODateTime)(&value)
}

func (t *TransactionDetails3) AddProxyHolder() *PartyIdentification2Choice {
	t.ProxyHolder = new(PartyIdentification2Choice)
	return t.ProxyHolder
}

func (t *TransactionDetails3) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}