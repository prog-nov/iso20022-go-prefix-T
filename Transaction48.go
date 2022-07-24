package iso20022

// Identifies the details of the transaction.
type Transaction48 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Reference assigned to the trade by the investor or the trading party. This reference will be used throughout the trade life cycle to access/update the trade details.
	TradeIdentification []*RestrictedFINXMax16Text `xml:"TradId,omitempty"`

	// Collective reference identifying a set of messages.
	PoolIdentification *RestrictedFINXMax16Text `xml:"PoolId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the account servicer to unambiguously identify a corporate action event.
	CorporateActionEventIdentification *RestrictedFINXMax16Text `xml:"CorpActnEvtId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the triparty agent's point of view.
	TripartyAgentCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"TrptyAgtCollTxId,omitempty"`

	// Unique reference identifying the triparty collateral management transaction from the client's point of view.
	ClientTripartyCollateralTransactionIdentification *RestrictedFINXMax16Text `xml:"ClntTrptyCollTxId,omitempty"`

	// Unique identification assigned to the instruction by the client.
	ClientCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"ClntCollInstrId,omitempty"`

	// Unique identification assigned to the instruction by the triparty agent.
	TripartyCollateralInstructionIdentification *RestrictedFINXMax16Text `xml:"TrptyCollInstrId,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *TransactionDetails90 `xml:"TxDtls,omitempty"`
}

func (t *Transaction48) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) AddTradeIdentification(value string) {
	t.TradeIdentification = append(t.TradeIdentification, (*RestrictedFINXMax16Text)(&value))
}

func (t *Transaction48) SetPoolIdentification(value string) {
	t.PoolIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetCorporateActionEventIdentification(value string) {
	t.CorporateActionEventIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetTripartyAgentCollateralTransactionIdentification(value string) {
	t.TripartyAgentCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetClientTripartyCollateralTransactionIdentification(value string) {
	t.ClientTripartyCollateralTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetClientCollateralInstructionIdentification(value string) {
	t.ClientCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) SetTripartyCollateralInstructionIdentification(value string) {
	t.TripartyCollateralInstructionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *Transaction48) AddTransactionDetails() *TransactionDetails90 {
	t.TransactionDetails = new(TransactionDetails90)
	return t.TransactionDetails
}
