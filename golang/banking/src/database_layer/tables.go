package database_layer

import (
	"time"

	"github.com/shopspring/decimal"
)

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

type User struct {
	Id          int64
	Password    string
	LastLogin   time.Time
	IsSuperuser bool
	FirstName   string
	LastName    string
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
	Phone       string
	Email       string
}

type Client struct {
	ClientId            int64 `gorm:"primaryKey"`
	RegistrationAddress string
	ResidentialAddress  string
	ClientType          ClientType `gorm:"type:client_type"`
	Ogrn                string
	Inn                 string
	Kpp                 string
	UserId              int64
	Cards               []Card
}

type Card struct {
	CardId   string `gorm:"primaryKey"`
	Balance  decimal.Decimal
	ValidTo  time.Time
	CvcCode  string
	CardType CardType `gorm:"type:card_type"`
	Currency string
	ClientId int64
	Client   Client `gorm:"foreignKey:ClientId"`
}

type Contract struct {
	ContractId      string       `gorm:"primaryKey"`
	ContractType    ContractType `gorm:"type:contract_type"`
	ConclusionDate  time.Time
	ContractContent string
	ClientId        int64
	Client          Client `gorm:"foreignKey:ClientId"`
}

type Credit struct {
	CreditID       string `gorm:"primaryKey"`
	Summ           decimal.Decimal
	Percent        decimal.Decimal
	ConclusionDate time.Time
	EndDate        time.Time
	ContractId     int64
	Contract       Contract `gorm:"foreignKey:ContractId"`
}

type PaymentSchedule struct {
	PaymentScheduleID      int64 `gorm:"primaryKey"`
	TotalSum               decimal.Decimal
	Currency               string
	Commission             decimal.Decimal
	RepaymentOfInterestSum decimal.Decimal
	SumRepaymentLoanPart   decimal.Decimal
	DateBegin              time.Time
	DateEnd                time.Time
	ContractID             int64
	Contract               Contract `gorm:"foreignKey:ContractId"`
}

type Transaction struct {
	TransactionID   int64           `gorm:"primaryKey"`
	TransactionType TransactionType `gorm:"type:transaction_type"`
	CardFromId      int64           `gorm:"foreignKey:card_from_id`
	CardToId        int64           `gorm:"foreignKey:card_to_id`
	CardId          int64           `gorm:"foreignKey:card_id"`
}

func (User) TableName() string {
	return "user"
}

func (Client) TableName() string {
	return "client"
}

func (Card) TableName() string {
	return "card"
}

func (Contract) TableName() string {
	return "contract"
}

func (Credit) TableName() string {
	return "credit"
}

func (PaymentSchedule) TableName() string {
	return "payment_schedule"
}

func (Transaction) TableName() string {
	return "transaction"
}
