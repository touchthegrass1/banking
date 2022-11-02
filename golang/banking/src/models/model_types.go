package models

type CardType string
type ClientType string
type ContractType string
type TransactionType string

const (
	debit  CardType = "debit"
	credit CardType = "credit"

	jp         ClientType = "jp"
	individual ClientType = "individual"
	ie         ClientType = "ie"

	loanAgreement                     ContractType = "loan agreement"
	bankAccountAgreement              ContractType = "bank account agreement"
	settlementAndCashServiceAgreement ContractType = "settlement and cash service agreement"

	deposit  TransactionType = "deposit"
	withdraw TransactionType = "withdraw"
	transfer TransactionType = "transfer"
)
