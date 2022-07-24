package iso20022

// Provides transaction type and identification information.
type TransactionIdentifications38 struct {

	// Unambiguous identification of the transaction as known by the account owner (or the instructing party managing the account).
	AccountOwnerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctOwnrTxId"`

	// Unambiguous identification of the transaction as known by the account servicer.
	AccountServicerTransactionIdentification *RestrictedFINXMax16Text `xml:"AcctSvcrTxId,omitempty"`

	// Identification of a transaction assigned by a market infrastructure other than a central securities depository, for example, Target2-Securities.
	MarketInfrastructureTransactionIdentification *RestrictedFINXMax16Text `xml:"MktInfrstrctrTxId,omitempty"`

	// Identification of the transaction assigned by the processor of the instruction other than the account owner the account servicer and the market infrastructure.
	ProcessorTransactionIdentification *RestrictedFINXMax16Text `xml:"PrcrTxId,omitempty"`

	// Unique reference agreed upon by the two trade counterparties to identify the trade.
	CommonIdentification *RestrictedFINXMax16Text `xml:"CmonId,omitempty"`

	// Identification assigned by the netting service provider to identify the net transaction resulting from the netting process.
	NettingServiceProviderIdentification *RestrictedFINXMax16Text `xml:"NetgSvcPrvdrId,omitempty"`
}

func (t *TransactionIdentifications38) SetAccountOwnerTransactionIdentification(value string) {
	t.AccountOwnerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionIdentifications38) SetAccountServicerTransactionIdentification(value string) {
	t.AccountServicerTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionIdentifications38) SetMarketInfrastructureTransactionIdentification(value string) {
	t.MarketInfrastructureTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionIdentifications38) SetProcessorTransactionIdentification(value string) {
	t.ProcessorTransactionIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionIdentifications38) SetCommonIdentification(value string) {
	t.CommonIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (t *TransactionIdentifications38) SetNettingServiceProviderIdentification(value string) {
	t.NettingServiceProviderIdentification = (*RestrictedFINXMax16Text)(&value)
}
